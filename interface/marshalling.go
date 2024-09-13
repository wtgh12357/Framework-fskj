package _interface

type MarshalHandler[T any] interface {
	Marshal(t T) ([]byte, error)
}

type UnMarshalHandler[T any] interface {
	Unmarshal(t T, bs []byte) error
}
