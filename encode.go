package karigo

import (
	"errors"
	"fmt"

	"github.com/vmihailenco/msgpack"
)

// Encode ...
func Encode(v uint, e Entry) ([]byte, error) {
	var (
		enc []byte
		err error
	)

	switch v {
	case 0:
		// Version 0
		enc, err = msgpack.Marshal(e)
		if err != nil {
			return nil, fmt.Errorf("cannot encode: %s", err)
		}
	default:
		return nil, errors.New("unsupported version")
	}

	return enc, nil
}

// Decode ...
func Decode(v uint, raw []byte) (Entry, error) {
	if len(raw) == 0 {
		return nil, errors.New("no bytes")
	}

	e := Entry{}

	switch v {
	case 0:
		// Version 0
		err := msgpack.Unmarshal(raw, &e)
		if err != nil {
			return nil, fmt.Errorf("cannot decode: %s", err)
		}
	default:
		return nil, errors.New("unsupported version")
	}

	return e, nil
}
