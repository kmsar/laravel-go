package Encription

import (
	"github.com/laravel-go-version/v2/pkg/Illuminate/Contracts/IEncryption"
	"github.com/laravel-go-version/v2/pkg/Illuminate/Support/Str"
)

type Factory struct {
	encryptors map[string]IEncryption.Encryptor
}

func (this *Factory) Encode(value string) string {
	return this.Driver("default").Encode(value)
}

func (this *Factory) Decode(payload string) (string, error) {
	return this.Driver("default").Decode(payload)
}

func (this *Factory) Extend(key string, encryptor IEncryption.Encryptor) {
	this.encryptors[key] = encryptor
}

func (this *Factory) Driver(key string) IEncryption.Encryptor {
	return this.encryptors[Str.StringOr(key, "default")]
}
