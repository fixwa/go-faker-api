package fakery

import (
	"encoding/binary"
	"github.com/bxcodec/faker/v3"
	"reflect"
	"time"
)

func Generator() {
	faker.AddProvider("customIdFaker", func(v reflect.Value) (interface{}, error) {
		var b [12]byte
		binary.BigEndian.PutUint32(b[0:4], uint32(time.Now().Unix()))
		return b, nil
	})
}
