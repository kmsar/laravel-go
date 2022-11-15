package serializers

import (
	"bytes"
	"encoding/gob"
)

type Gob struct {
}

func (j Gob) Serialize(i interface{}) string {
	var buf bytes.Buffer

	enc := gob.NewEncoder(&buf)

	if err := enc.Encode(i); err != nil {
		panic(err)
	}

	return buf.String()

}

func (j Gob) UnSerialize(s string, i interface{}) error {
	if err := gob.NewDecoder(bytes.NewBufferString(s)).Decode(i); err != nil {
		return err
	}
	return nil
}
