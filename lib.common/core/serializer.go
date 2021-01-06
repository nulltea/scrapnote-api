package core

type Serializer interface {
	Encode(data interface{}) ([]byte, error)
	DecodeInto(input []byte, target interface{}) error
}