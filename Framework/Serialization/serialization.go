package Serialization

import (
	"github.com/kmsar/laravel-go/Framework/Contracts/ISerialize"
	"github.com/kmsar/laravel-go/Framework/Serialization/serializers"
)

func New() ISerialize.Serialization {
	return &Serialization{serializers: map[string]ISerialize.Serializer{
		"json": serializers.Json{},
		"gob":  serializers.Gob{},
		"xml":  serializers.Xml{},
	}}
}

type Serialization struct {
	serializers map[string]ISerialize.Serializer
}

func (s *Serialization) Method(name string) ISerialize.Serializer {
	return s.serializers[name]
}

func (s *Serialization) Extend(name string, serializer ISerialize.Serializer) {
	s.serializers[name] = serializer
}
