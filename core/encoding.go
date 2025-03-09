package core

import "io"

type Encoder[T any] interface {
	Encode(w io.Writer, T any) error
}

type Decoder[T any] interface {
	Decode(r io.Reader) (T, error)
}
