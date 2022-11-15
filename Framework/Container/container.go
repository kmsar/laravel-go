package Container

import (
	"errors"
	"fmt"
	"github.com/kmsar/laravel-go/Framework/Contracts/IContainer"
	"github.com/kmsar/laravel-go/Framework/Support/Exceptions"
	"github.com/kmsar/laravel-go/Framework/Support/Utils"
	"github.com/kmsar/laravel-go/Framework/Support/supports-master/utils"
	"reflect"
	"sync"
)

var (
	CallerTypeError = errors.New("CallerTypeError:invalid caller type")
)

type Container struct {
	//binds        map[string]IContainer.MagicalFunc
	transients      map[string]IContainer.MagicalFunc
	scopeds         map[string]IContainer.MagicalFunc //binds
	singletons      map[string]IContainer.MagicalFunc
	instances       sync.Map //store instance and singletone instance
	scopedInstances sync.Map //store scoped instance
	aliases         sync.Map
	argProviders    []func(key string, p reflect.Type, arguments ArgumentsTypeMap) any
	TypeRegister    *TypeRegister
}

func newInstanceProvider(provider any) IContainer.MagicalFunc {
	magicalFn := NewMagicalFunc(provider)
	if magicalFn.NumOut() != 1 {
		Exceptions.Throw(CallerTypeError)
	}
	return magicalFn
}

func New() IContainer.IContainer {
	container := &Container{}
	container.TypeRegister = &TypeRegister{}
	container.argProviders = []func(key string, p reflect.Type, arguments ArgumentsTypeMap) any{

		func(key string, _ reflect.Type, arguments ArgumentsTypeMap) any {
			return arguments.Pull(key)
		},

		func(key string, argType reflect.Type, arguments ArgumentsTypeMap) any {
			return arguments.FindConvertibleArg(key, argType)
		},

		func(key string, argType reflect.Type, arguments ArgumentsTypeMap) any {
			return container.GetByArguments(key, arguments)
		},

		func(key string, argType reflect.Type, arguments ArgumentsTypeMap) any {
			var (
				tempInstance any
				isPtr        = argType.Kind() == reflect.Ptr
			)
			if isPtr {
				tempInstance = reflect.New(argType.Elem()).Interface()
			} else {
				tempInstance = reflect.New(argType).Interface()
			}
			container.DIByArguments(tempInstance, arguments)
			if isPtr {
				return tempInstance
			}
			return reflect.ValueOf(tempInstance).Elem().Interface()
		},
	}
	container.Flush()
	return container
}
func (this *Container) Instance(key string, instance any) {
	this.instances.Store(this.GetKey(key), instance)
}

func (this *Container) bind(key string, provider any, lifeTime ServiceLifeTime, isLazy bool) {

	funcType := reflect.TypeOf(provider)
	if funcType.Kind() != reflect.Func {
		Exceptions.Throw("container: the resolver must be a function")
	}

	retCount := funcType.NumOut()
	if retCount == 0 || retCount > 2 {
		Exceptions.Throw("container: resolver function signature is invalid - it must return abstract, or abstract and error")
	}

	magicalFn := newInstanceProvider(provider)
	resolveType := funcType.Out(0)
	for _, arg := range magicalFn.Arguments() {
		if arg == resolveType {
			Exceptions.Throw("container: resolver function signature is invalid - depends on abstract it returns")
		}
	}

	if key == "" {
		key = this.TypeRegister.SetOrGet(resolveType)
	}

	this.Alias(key, Utils.GetTypeKey(magicalFn.Returns()[0]))

	switch lifeTime {
	case Singleton:
		this.singletons[this.GetKey(key)] = magicalFn
	case Scoped:
		this.scopeds[this.GetKey(key)] = magicalFn
	case Transient:
		this.transients[this.GetKey(key)] = magicalFn
	}
}

func (this *Container) Singleton(provider any) {
	this.bind("", provider, Singleton, true)
}
func (this *Container) Bind(provider any) {
	this.Scoped(provider)
}
func (this *Container) Scoped(provider any) {
	this.bind("", provider, Scoped, true)
}
func (this *Container) Transient(provider any) {
	this.bind("", provider, Transient, true)
}

func (this *Container) NamedSingleton(key string, provider any) {
	this.bind(key, provider, Singleton, true)
}
func (this *Container) NamedBind(key string, provider any) {
	this.NamedScoped(key, provider)
}
func (this *Container) NamedScoped(key string, provider any) {
	this.bind(key, provider, Scoped, true)
}
func (this *Container) NamedTransient(key string, provider any) {
	this.bind(key, provider, Transient, true)
}

func (this *Container) HasBound(key string) bool {
	key = this.GetKey(key)
	if _, existsTransient := this.transients[key]; existsTransient {
		return true
	}
	if _, existsScoped := this.scopeds[key]; existsScoped {
		return true
	}
	if _, existsSingleton := this.singletons[key]; existsSingleton {
		return true
	}
	if _, existsInstance := this.instances.Load(key); existsInstance {
		return true
	}
	return false
}

func (this *Container) Alias(key string, alias string) {
	this.aliases.Store(alias, key)
}

func (this *Container) GetKey(alias string) string {
	if value, existsAlias := this.aliases.Load(alias); existsAlias {
		return value.(string)
	}
	return alias
}

func (this *Container) FlushScopedInstances() {
	this.scopedInstances = sync.Map{}
}
func (this *Container) Flush() {

	this.singletons = make(map[string]IContainer.MagicalFunc, 0)
	this.scopeds = make(map[string]IContainer.MagicalFunc, 0)
	this.transients = make(map[string]IContainer.MagicalFunc, 0)
	this.aliases = sync.Map{}
	this.scopedInstances = sync.Map{}
	this.instances = sync.Map{}
}

func (this *Container) Get(key string, args ...any) any {
	key = this.GetKey(key)
	if tempInstance, existsInstance := this.instances.Load(key); existsInstance {
		return tempInstance
	}
	if singletonProvider, existsProvider := this.singletons[key]; existsProvider {
		value := this.Call(singletonProvider, args...)[0]
		this.instances.Store(key, value)
		return value
	}

	if tempInstance, existsInstance := this.scopedInstances.Load(key); existsInstance {
		return tempInstance
	}
	if scopedProvider, existsProvider := this.scopeds[key]; existsProvider {
		value := this.Call(scopedProvider, args...)[0]
		this.scopedInstances.Store(key, value)
		return value
	}

	if instanceProvider, existsProvider := this.transients[key]; existsProvider {
		return this.Call(instanceProvider, args...)[0]
	}
	return nil
}

func (this *Container) GetTypeName(resolvedType any) string {
	return this.TypeRegister.getName(reflect.TypeOf(resolvedType))
}

func (this *Container) GetByArguments(key string, arguments ArgumentsTypeMap) any {

	key = this.GetKey(key)
	if tempInstance, existsInstance := this.instances.Load(key); existsInstance {
		return tempInstance
	}
	if singletonProvider, existsProvider := this.singletons[key]; existsProvider {
		value := this.StaticCallByArguments(singletonProvider, arguments)[0]
		this.instances.Store(key, value)
		return value
	}

	if tempInstance, existsInstance := this.scopedInstances.Load(key); existsInstance {
		return tempInstance
	}
	if singletonProvider, existsProvider := this.scopeds[key]; existsProvider {
		value := this.StaticCallByArguments(singletonProvider, arguments)[0]
		this.scopedInstances.Store(key, value)
		return value
	}

	if instanceProvider, existsProvider := this.transients[key]; existsProvider {
		return this.StaticCallByArguments(instanceProvider, arguments)[0]
	}
	return nil
}

// StaticCall
func (this *Container) StaticCall(magicalFn IContainer.MagicalFunc, args ...any) []any {
	return this.StaticCallByArguments(magicalFn, NewArgumentsTypeMap(append(args, this)))
}

// StaticCallByArguments
func (this *Container) StaticCallByArguments(magicalFn IContainer.MagicalFunc, arguments ArgumentsTypeMap) []any {
	fnArgs := make([]reflect.Value, 0)

	for _, arg := range magicalFn.Arguments() {
		key := Utils.GetTypeKey(arg)
		fnArgs = append(fnArgs, reflect.ValueOf(this.findArg(key, arg, arguments)))
	}

	results := make([]any, 0)

	for _, result := range magicalFn.Call(fnArgs) {
		results = append(results, result.Interface())
	}

	return results
}

func (this *Container) Call(fn any, args ...any) []any {
	if magicalFn, isMagicalFunc := fn.(IContainer.MagicalFunc); isMagicalFunc {
		return this.StaticCall(magicalFn, args...)
	}
	return this.StaticCall(NewMagicalFunc(fn), args...)
}

func (this *Container) findArg(key string, p reflect.Type, arguments ArgumentsTypeMap) (result any) {
	for _, provider := range this.argProviders {
		if value := provider(key, p, arguments); value != nil {
			return value
		}
	}
	return
}

func (this *Container) DIByArguments(object any, arguments ArgumentsTypeMap) {
	if component, ok := object.(IContainer.Component); ok {
		component.Construct(this)
		return
	}

	objectValue := reflect.ValueOf(object)

	switch objectValue.Kind() {
	case reflect.Ptr:
		if objectValue.Elem().Kind() != reflect.Struct {
			Exceptions.Throw(errors.New("object is not struct"))
		}
		objectValue = objectValue.Elem()
	default:
		Exceptions.Throw(errors.New("object is not ptr"))
	}

	valueType := objectValue.Type()

	var (
		fieldNum  = objectValue.NumField()
		tempValue = reflect.New(valueType).Elem()
	)

	tempValue.Set(objectValue)

	for i := 0; i < fieldNum; i++ {
		var (
			field          = valueType.Field(i)
			key            = Utils.GetTypeKey(field.Type)
			fieldTags      = Utils.ParseStructTag(field.Tag)
			fieldValue     = tempValue.Field(i)
			fieldInterface any
		)

		if di, existsDiTag := fieldTags["di"]; existsDiTag {
			if len(di) > 0 {
				fieldInterface = this.Get(di[0])
			}
			if fieldInterface == nil {
				fieldInterface = this.findArg(key, field.Type, arguments)
			}
		}

		if fieldInterface != nil {
			fieldType := reflect.TypeOf(fieldInterface)
			if fieldType.ConvertibleTo(field.Type) {
				value := reflect.ValueOf(fieldInterface)
				if key != Utils.GetTypeKey(fieldType) {
					value = value.Convert(field.Type)
				}
				fieldValue.Set(value)
			} else {
				Exceptions.Throw(errors.New(fmt.Sprintf("Cannot inject %s, because the type does not match, the target type is %s, and the type of injection is %s", field.Name, field.Type.String(), fieldType.String())))
			}
		}
	}

	objectValue.Set(tempValue)

	return
}

// NamedResolve takes abstraction and its name and fills it with the related concrete.
func (this *Container) Resolve(abstraction any, args ...any) {
	receiverType := reflect.TypeOf(abstraction)
	if receiverType == nil {
		Exceptions.Throw("container: invalid abstraction")
	}

	if receiverType.Kind() == reflect.Ptr {
		elem := receiverType.Elem()

		key := utils.GetTypeKey(elem)
		instance := this.Get(key)
		reflect.ValueOf(abstraction).Elem().Set(reflect.ValueOf(instance))
		return
		//if concrete, exist := c[elem][name]; exist {
		//	if instance, err := concrete.make(c); err == nil {
		//		reflect.ValueOf(abstraction).Elem().Set(reflect.ValueOf(instance))
		//		return
		//	} else {
		//		Exceptions.Throw("container: no concrete found for:  " + err)
		//	}
		//}

		//Exceptions.Throw("container: no concrete found for: " + elem.String())
	}

	Exceptions.Throw("container: invalid abstraction")

}

func (this *Container) DI(object any, args ...any) {
	this.DIByArguments(object, NewArgumentsTypeMap(append(args, this)))
}
func (this *Container) Fill(object any, args ...any) {
	this.DIByArguments(object, NewArgumentsTypeMap(append(args, this)))
}
