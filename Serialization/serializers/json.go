package serializers

import (
	"encoding/json"
)

type Json struct {
}

func (j Json) Serialize(i interface{}) string {
	var (
		buf []byte
		err error
	)

	if buf, err = json.Marshal(i); err != nil {
		panic(err)
	}

	return string(buf)
}

func (j Json) UnSerialize(s string, i interface{}) error {
	if err := json.Unmarshal([]byte(s), i); err != nil {
		return err
	}
	return nil
}
