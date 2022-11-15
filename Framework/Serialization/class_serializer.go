package Serialization

import (
	"errors"
	"github.com/kmsar/laravel-go/Framework/Contracts/IContainer"
	"github.com/kmsar/laravel-go/Framework/Contracts/ISerialize"
	"github.com/kmsar/laravel-go/Framework/Contracts/Support"
	"github.com/kmsar/laravel-go/Framework/Support/class"

	"github.com/kmsar/laravel-go/Framework/Serialization/serializers"
	"reflect"
	"sync"
)

type Class struct {
	Class   string `json:"c"`
	Payload string `json:"p"`
}

func NewClassSerializer(container IContainer.IContainer) ISerialize.ClassSerializer {
	return &Serializer{
		container:  container,
		classes:    sync.Map{},
		serializer: serializers.Json{},
	}
}

type Serializer struct {
	container  IContainer.IContainer
	classes    sync.Map
	serializer ISerialize.Serializer
}

func (this *Serializer) Register(class Support.Class) {
	this.classes.Store(class.ClassName(), class)
}

func (this *Serializer) Serialize(instance interface{}) string {
	return this.serializer.Serialize(Class{
		Class:   class.Make(instance).ClassName(),
		Payload: this.serializer.Serialize(instance),
	})
}

func (this *Serializer) Parse(serialized string) (interface{}, error) {
	var c Class
	if err := this.serializer.UnSerialize(serialized, &c); err != nil {
		return nil, err
	}

	classItem, exists := this.classes.Load(c.Class)
	if !exists {
		return nil, errors.New("unregistered class")
	}

	targetClass := classItem.(Support.Class)

	instance := reflect.New(targetClass.GetType()).Interface()

	if err := this.serializer.UnSerialize(c.Payload, instance); err != nil {
		return nil, err
	}

	if component, isComponent := instance.(IContainer.Component); isComponent {
		component.Construct(this.container)
		return component, nil
	}

	return instance, nil
}
