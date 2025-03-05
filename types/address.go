package types

import (
	"encoding/hex"
	"fmt"
)

type Address [20]uint8

func (a Address) String() string {
	//return fmt.Sprintf("0x%02x%02x%02x", a[0], a[1], a[2])
	return hex.EncodeToString(a.ToSlice())
}

func (a Address) ToSlice() []byte {
	return a[:]
}
func AddressFromBytes(b []byte) Address {
	if len(b) != 20 {
		msg := fmt.Sprintf("Address must be 20 bytes long, got %d", len(b))
		panic(msg)

	}
	return Address(b)
}
