package class

import (
	"github.com/kmsar/laravel-go/Framework/Contracts/IAuth"
	"github.com/kmsar/laravel-go/Framework/Contracts/ICache"
	"github.com/kmsar/laravel-go/Framework/Contracts/IContainer"
	"github.com/kmsar/laravel-go/Framework/Contracts/IDatabase"
	"github.com/kmsar/laravel-go/Framework/Contracts/IEvent"
	"github.com/kmsar/laravel-go/Framework/Contracts/IExeption"
	"github.com/kmsar/laravel-go/Framework/Contracts/IFilesystem"
	"github.com/kmsar/laravel-go/Framework/Contracts/IFoundation"
	"github.com/kmsar/laravel-go/Framework/Contracts/IHashing"
	"github.com/kmsar/laravel-go/Framework/Contracts/IHttp"
	"github.com/kmsar/laravel-go/Framework/Contracts/IRedis"
	"github.com/kmsar/laravel-go/Framework/Contracts/Support"
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
