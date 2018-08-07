package msgpack

import (
	"errors"
)

var defaultSerializer = SerializeAsMap

func Serialize(v interface{}) ([]byte, error) {
	return defaultSerializer(v)
}

func SerializeAsArray(v interface{}) ([]byte, error) {
	s := serializer{}
	s.Exec(v)
	return []byte{}, errors.New("not implement")
}

func SerializeAsMap(v interface{}) ([]byte, error) {
	return []byte{}, errors.New("not implement")
}
