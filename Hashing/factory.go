package Hashing

import (
	"fmt"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/IConfig"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/IHashing"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/Support"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Support/Field"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Support/Str"
)

type Factory struct {
	config  IConfig.Config
	hashes  map[string]IHashing.Hasher
	drivers map[string]IHashing.HasherProvider
}

func (f *Factory) Info(hashedValue string) Support.Fields {
	return f.Driver("default").Info(hashedValue)
}

func (f *Factory) Make(value string, options Support.Fields) string {
	return f.Driver("default").Make(value, options)
}

func (f *Factory) Check(value, hashedValue string, options Support.Fields) bool {
	return f.Driver("default").Check(value, hashedValue, options)
}

func (f Factory) getConfig(name string) Support.Fields {
	return f.config.GetFields(
		Str.IfString(name == "default", "hashing", fmt.Sprintf("hashing.hashes.%s", name)),
	)
}

func (f *Factory) Driver(name string) IHashing.Hasher {
	if hashed, existsHashed := f.hashes[name]; existsHashed {
		return hashed
	}

	config := f.getConfig(name)
	driver := Field.GetStringField(config, "driver", "bcrypt")
	driveProvider, existsProvider := f.drivers[driver]

	if !existsProvider {
		Logs.WithFields(nil).Fatal(fmt.Sprintf("：%s", driver))
	}

	f.hashes[name] = driveProvider(config)

	return f.hashes[name]
}

func (f *Factory) Extend(driver string, hashedProvider IHashing.HasherProvider) {
	f.drivers[driver] = hashedProvider
}
