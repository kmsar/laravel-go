package class

import (
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/IAuth"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/ICache"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/IContainer"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/IDatabase"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/IEvent"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/IExeption"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/IFilesystem"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/IFoundation"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/IHashing"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/IHttp"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/IRedis"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/Support"
)

var (
	Application  = Define(new(IFoundation.IApplication))
	Container    = Define(new(IContainer.IContainer))
	MagicalFunc  = Define(new(IContainer.MagicalFunc))
	Component    = Define(new(IContainer.Component))
	Json         = Define(new(Support.Json))
	HttpRequest  = Define(new(IHttp.IHttpRequest))
	HttpResponse = Define(new(IHttp.IHttpResponse))
	Fields       = Define(new(Support.FieldsProvider))
	Hash         = Define(new(IHashing.Hasher))
	Exception    = Define(new(IExeption.Exception))

	Auth  = Define(new(IAuth.Auth))
	Guard = Define(new(IAuth.Guard))

	Validatable    = Define(new(Support.Validatable))
	ShouldValidate = Define(new(Support.ShouldValidate))

	Redis      = Define(new(IRedis.RedisConnection))
	Cache      = Define(new(ICache.CacheStore))
	FileSystem = Define(new(IFilesystem.FileSystem))

	Event    = Define(new(IEvent.Event))
	Listener = Define(new(IEvent.EventListener))

	DB           = Define(new(IDatabase.DBConnection))
	SqlExecutor  = Define(new(IDatabase.SqlExecutor))
	QueryBuilder = Define(new(IDatabase.QueryBuilder))
	Model        = Define(new(IDatabase.Model))
)
