package core

import "io"

type Transaction struct {
	Data []byte
}

func (tx *Transaction) EncodeBinary(reader io.Writer) error {
	return nil
}
func (tx *Transaction) DecodeBinary(reader io.Reader) error {
	return nil
}
