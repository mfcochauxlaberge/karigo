package bin

import (
	"encoding/binary"
	"errors"
	"time"

	"github.com/mfcochauxlaberge/jsonapi"
)

// From takes a slice of bytes and returns the value it represents.
//
// It reverses the operation done by To.
//
// It returns an error if the operation can't be done.
func From(data []byte) (interface{}, error) {
	if len(data) < 2 {
		return nil, errors.New("invalid binary")
	}

	typ := int(data[0])

	// A length of exactly two represents
	// a value with no data (only metada),
	// which means the value must be the
	// zero value of the type.
	if len(data) == 2 {
		nullable := data[1] == 1
		return jsonapi.GetZeroValue(typ, nullable), nil
	}

	switch typ {
	case jsonapi.AttrTypeString:
		return string(data[2:]), nil
	case jsonapi.AttrTypeInt:
		v, _ := binary.Varint(data[2:])
		return int(v), nil
	case jsonapi.AttrTypeInt8:
		v, _ := binary.Varint(data[2:])
		return int8(v), nil
	case jsonapi.AttrTypeInt16:
		v, _ := binary.Varint(data[2:])
		return int16(v), nil
	case jsonapi.AttrTypeInt32:
		v, _ := binary.Varint(data[2:])
		return int32(v), nil
	case jsonapi.AttrTypeInt64:
		v, _ := binary.Varint(data[2:])
		return v, nil
	case jsonapi.AttrTypeUint:
		v, _ := binary.Uvarint(data[2:])
		return uint(v), nil
	case jsonapi.AttrTypeUint8:
		v, _ := binary.Uvarint(data[2:])
		return uint8(v), nil
	case jsonapi.AttrTypeUint16:
		v, _ := binary.Uvarint(data[2:])
		return uint16(v), nil
	case jsonapi.AttrTypeUint32:
		v, _ := binary.Uvarint(data[2:])
		return uint32(v), nil
	case jsonapi.AttrTypeUint64:
		v, _ := binary.Uvarint(data[2:])
		return v, nil
	case jsonapi.AttrTypeBool:
		if data[3] == 0 {
			return false, nil
		}

		return true, nil
	case jsonapi.AttrTypeTime:
		t, err := time.Parse(time.RFC3339Nano, string(data[2:]))
		return t, err
	case jsonapi.AttrTypeBytes:
		return data[2:], nil
	default:
		return nil, errors.New("unsupported type")
	}
}
