package bin

import (
	"encoding/binary"
	"time"

	"github.com/mfcochauxlaberge/jsonapi"
)

// To takes a value and returns its binary representation.
//
// Two bytes are added at the beginning and store metadata. The first byte
// is the value's type, encoded as an integer taken from jsonapi's constants
// (for example, jsonapi.AttrTypeString which is 1).
//
// The second byte reports wether the value is null or not by storing 1 for null
// or 0 for not null.
//
// It panics if v is of unsupported type. See the implementation for a list of
// supported types.
func To(v interface{}) []byte {
	switch v2 := v.(type) {
	case string:
		data := make([]byte, 0, len(v2)+2)
		data = append(data, jsonapi.AttrTypeString)
		data = append(data, 0)
		data = append(data, v2...)

		return data
	case int:
		data := make([]byte, 9)
		_ = binary.PutVarint(data, int64(v2))

		return data
	case int8:
		return []byte{byte(v2)}
	case int16:
		data := make([]byte, 3)
		_ = binary.PutVarint(data, int64(v2))

		return data
	case int32:
		data := make([]byte, 5)
		_ = binary.PutVarint(data, int64(v2))

		return data
	case int64:
		data := make([]byte, 9)
		_ = binary.PutVarint(data, v2)

		return data
	case uint:
		data := make([]byte, 9)
		_ = binary.PutUvarint(data, uint64(v2))

		return data
	case uint8:
		return []byte{v2}
	case uint16:
		data := make([]byte, 3)
		_ = binary.PutUvarint(data, uint64(v2))

		return data
	case uint32:
		data := make([]byte, 5)
		_ = binary.PutUvarint(data, uint64(v2))

		return data
	case uint64:
		data := make([]byte, 9)
		_ = binary.PutUvarint(data, v2)

		return data
	case bool:
		if v2 {
			return []byte{jsonapi.AttrTypeBool, 1}
		}

		return []byte{jsonapi.AttrTypeBool, 0}
	case time.Time:
		// RFC3339Nano should always fit
		// in 37 bytes. An extra byte is
		// added for the type.
		data := make([]byte, 1, 38)
		data[0] = jsonapi.AttrTypeTime
		data = append(data, []byte(v2.Format(time.RFC3339Nano))...)

		return data
	case []byte:
		return v2
	}

	return nil
}
