package karigo_test

import (
	"bytes"
	"testing"
	"time"

	. "github.com/mfcochauxlaberge/karigo"

	"github.com/stretchr/testify/assert"
)

func TestOpEncode(t *testing.T) {
	assert := assert.New(t)

	now, _ := time.Parse(time.RFC3339Nano, "2013-06-24T22:03:34.8276Z")

	tests := []struct {
		name    string
		ops     []Op
		encoded []byte
		err     bool
	}{
		{
			name: "nil ops",
			ops:  nil,
			err:  true,
		}, {
			name: "no ops",
			ops:  []Op{},
			err:  true,
		}, {
			name: "1 op",
			ops: []Op{
				{
					Key:   Key{Set: "set", ID: "id", Field: "field"},
					Op:    OpSet,
					Value: "string value",
				},
			},
			encoded: []byte(`[
				{"key":{"field":"field","id":"id","set":"set"},
					"op":"=","type":"string","value":"string value"}
			]`),
		}, {
			name: "one op per type",
			ops: []Op{
				{
					Key:   Key{Set: "set", ID: "id", Field: "string"},
					Op:    OpSet,
					Value: "string",
				}, {
					Key:   Key{Set: "set", ID: "id", Field: "integer"},
					Op:    OpSet,
					Value: 1,
				}, {
					Key:   Key{Set: "set", ID: "id", Field: "integer8"},
					Op:    OpSet,
					Value: int8(2),
				}, {
					Key:   Key{Set: "set", ID: "id", Field: "integer16"},
					Op:    OpSet,
					Value: int16(3),
				}, {
					Key:   Key{Set: "set", ID: "id", Field: "integer32"},
					Op:    OpSet,
					Value: int32(4),
				}, {
					Key:   Key{Set: "set", ID: "id", Field: "integer64"},
					Op:    OpSet,
					Value: int64(5),
				}, {
					Key:   Key{Set: "set", ID: "id", Field: "uinteger"},
					Op:    OpSet,
					Value: 1,
				}, {
					Key:   Key{Set: "set", ID: "id", Field: "uinteger8"},
					Op:    OpSet,
					Value: uint8(2),
				}, {
					Key:   Key{Set: "set", ID: "id", Field: "uinteger16"},
					Op:    OpSet,
					Value: uint16(3),
				}, {
					Key:   Key{Set: "set", ID: "id", Field: "uinteger32"},
					Op:    OpSet,
					Value: uint32(4),
				}, {
					Key:   Key{Set: "set", ID: "id", Field: "uinteger64"},
					Op:    OpSet,
					Value: uint64(5),
				}, {
					Key:   Key{Set: "set", ID: "id", Field: "boolean"},
					Op:    OpSet,
					Value: true,
				}, {
					Key:   Key{Set: "set", ID: "id", Field: "time"},
					Op:    OpSet,
					Value: now,
				}, {
					Key:   Key{Set: "set", ID: "id", Field: "bytes"},
					Op:    OpSet,
					Value: []byte{1, 2, 3},
				}, {
					Key:   Key{Set: "set", ID: "id", Field: "string"},
					Op:    OpSet,
					Value: ptr("string"),
				}, {
					Key:   Key{Set: "set", ID: "id", Field: "integer"},
					Op:    OpSet,
					Value: ptr(1),
				}, {
					Key:   Key{Set: "set", ID: "id", Field: "integer8"},
					Op:    OpSet,
					Value: ptr(int8(2)),
				}, {
					Key:   Key{Set: "set", ID: "id", Field: "integer16"},
					Op:    OpSet,
					Value: ptr(int16(3)),
				}, {
					Key:   Key{Set: "set", ID: "id", Field: "integer32"},
					Op:    OpSet,
					Value: ptr(int32(4)),
				}, {
					Key:   Key{Set: "set", ID: "id", Field: "integer64"},
					Op:    OpSet,
					Value: ptr(int64(5)),
				}, {
					Key:   Key{Set: "set", ID: "id", Field: "uinteger"},
					Op:    OpSet,
					Value: ptr(1),
				}, {
					Key:   Key{Set: "set", ID: "id", Field: "uinteger8"},
					Op:    OpSet,
					Value: ptr(uint8(2)),
				}, {
					Key:   Key{Set: "set", ID: "id", Field: "uinteger16"},
					Op:    OpSet,
					Value: ptr(uint16(3)),
				}, {
					Key:   Key{Set: "set", ID: "id", Field: "uinteger32"},
					Op:    OpSet,
					Value: ptr(uint32(4)),
				}, {
					Key:   Key{Set: "set", ID: "id", Field: "uinteger64"},
					Op:    OpSet,
					Value: ptr(uint64(5)),
				}, {
					Key:   Key{Set: "set", ID: "id", Field: "boolean"},
					Op:    OpSet,
					Value: ptr(true),
				}, {
					Key:   Key{Set: "set", ID: "id", Field: "time"},
					Op:    OpSet,
					Value: ptr(now),
				}, {
					Key:   Key{Set: "set", ID: "id", Field: "bytes"},
					Op:    OpSet,
					Value: ptr([]byte{1, 2, 3}),
				}, {
					Key:   Key{Set: "set", ID: "id", Field: "string"},
					Op:    OpSet,
					Value: nilptr("string"),
				}, {
					Key:   Key{Set: "set", ID: "id", Field: "integer"},
					Op:    OpSet,
					Value: nilptr("int"),
				}, {
					Key:   Key{Set: "set", ID: "id", Field: "integer8"},
					Op:    OpSet,
					Value: nilptr("int8"),
				}, {
					Key:   Key{Set: "set", ID: "id", Field: "integer16"},
					Op:    OpSet,
					Value: nilptr("int16"),
				}, {
					Key:   Key{Set: "set", ID: "id", Field: "integer32"},
					Op:    OpSet,
					Value: nilptr("int32"),
				}, {
					Key:   Key{Set: "set", ID: "id", Field: "integer64"},
					Op:    OpSet,
					Value: nilptr("int64"),
				}, {
					Key:   Key{Set: "set", ID: "id", Field: "uinteger"},
					Op:    OpSet,
					Value: nilptr("uint"),
				}, {
					Key:   Key{Set: "set", ID: "id", Field: "uinteger8"},
					Op:    OpSet,
					Value: nilptr("uint8"),
				}, {
					Key:   Key{Set: "set", ID: "id", Field: "uinteger16"},
					Op:    OpSet,
					Value: nilptr("uint16"),
				}, {
					Key:   Key{Set: "set", ID: "id", Field: "uinteger32"},
					Op:    OpSet,
					Value: nilptr("uint32"),
				}, {
					Key:   Key{Set: "set", ID: "id", Field: "uinteger64"},
					Op:    OpSet,
					Value: nilptr("uint64"),
				}, {
					Key:   Key{Set: "set", ID: "id", Field: "boolean"},
					Op:    OpSet,
					Value: nilptr("bool"),
				}, {
					Key:   Key{Set: "set", ID: "id", Field: "time"},
					Op:    OpSet,
					Value: nilptr("time.Time"),
				}, {
					Key:   Key{Set: "set", ID: "id", Field: "bytes"},
					Op:    OpSet,
					Value: nilptr("[]byte"),
				},
			},
			encoded: []byte(`[
				{"key":{"field":"string","id":"id","set":"set"},
					"op":"=","type":"string","value":"string"},
				{"key":{"field":"integer","id":"id","set":"set"},
					"op":"=","type":"int","value":1},
				{"key":{"field":"integer8","id":"id","set":"set"},
					"op":"=","type":"int8","value":2},
				{"key":{"field":"integer16","id":"id","set":"set"},
					"op":"=","type":"int16","value":3},
				{"key":{"field":"integer32","id":"id","set":"set"},
					"op":"=","type":"int32","value":4},
				{"key":{"field":"integer64","id":"id","set":"set"},
					"op":"=","type":"int64","value":5},
				{"key":{"field":"uinteger","id":"id","set":"set"},
					"op":"=","type":"int","value":1},
				{"key":{"field":"uinteger8","id":"id","set":"set"},
					"op":"=","type":"uint8","value":2},
				{"key":{"field":"uinteger16","id":"id","set":"set"},
					"op":"=","type":"uint16","value":3},
				{"key":{"field":"uinteger32","id":"id","set":"set"},
					"op":"=","type":"uint32","value":4},
				{"key":{"field":"uinteger64","id":"id","set":"set"},
					"op":"=","type":"uint64","value":5},
				{"key":{"field":"boolean","id":"id","set":"set"},
					"op":"=","type":"bool","value":true},
				{"key":{"field":"time","id":"id","set":"set"},
					"op":"=","type":"time","value":"2013-06-24T22:03:34.8276Z"},
				{"key":{"field":"bytes","id":"id","set":"set"},
					"op":"=","type":"bytes","value":"AQID"},
				{"key":{"field":"string","id":"id","set":"set"},
					"op":"=","type":"*string","value":"string"},
				{"key":{"field":"integer","id":"id","set":"set"},
					"op":"=","type":"*int","value":1},
				{"key":{"field":"integer8","id":"id","set":"set"},
					"op":"=","type":"*int8","value":2},
				{"key":{"field":"integer16","id":"id","set":"set"},
					"op":"=","type":"*int16","value":3},
				{"key":{"field":"integer32","id":"id","set":"set"},
					"op":"=","type":"*int32","value":4},
				{"key":{"field":"integer64","id":"id","set":"set"},
					"op":"=","type":"*int64","value":5},
				{"key":{"field":"uinteger","id":"id","set":"set"},
					"op":"=","type":"*int","value":1},
				{"key":{"field":"uinteger8","id":"id","set":"set"},
					"op":"=","type":"*uint8","value":2},
				{"key":{"field":"uinteger16","id":"id","set":"set"},
					"op":"=","type":"*uint16","value":3},
				{"key":{"field":"uinteger32","id":"id","set":"set"},
					"op":"=","type":"*uint32","value":4},
				{"key":{"field":"uinteger64","id":"id","set":"set"},
					"op":"=","type":"*uint64","value":5},
				{"key":{"field":"boolean","id":"id","set":"set"},
					"op":"=","type":"*bool","value":true},
				{"key":{"field":"time","id":"id","set":"set"},
					"op":"=","type":"*time","value":"2013-06-24T22:03:34.8276Z"},
				{"key":{"field":"bytes","id":"id","set":"set"},
					"op":"=","type":"*bytes","value":"AQID"},
				{"key":{"field":"string","id":"id","set":"set"},
					"op":"=","type":"*string","value":null},
				{"key":{"field":"integer","id":"id","set":"set"},
					"op":"=","type":"*int","value":null},
				{"key":{"field":"integer8","id":"id","set":"set"},
					"op":"=","type":"*int8","value":null},
				{"key":{"field":"integer16","id":"id","set":"set"},
					"op":"=","type":"*int16","value":null},
				{"key":{"field":"integer32","id":"id","set":"set"},
					"op":"=","type":"*int32","value":null},
				{"key":{"field":"integer64","id":"id","set":"set"},
					"op":"=","type":"*int64","value":null},
				{"key":{"field":"uinteger","id":"id","set":"set"},
					"op":"=","type":"*uint","value":null},
				{"key":{"field":"uinteger8","id":"id","set":"set"},
					"op":"=","type":"*uint8","value":null},
				{"key":{"field":"uinteger16","id":"id","set":"set"},
					"op":"=","type":"*uint16","value":null},
				{"key":{"field":"uinteger32","id":"id","set":"set"},
					"op":"=","type":"*uint32","value":null},
				{"key":{"field":"uinteger64","id":"id","set":"set"},
					"op":"=","type":"*uint64","value":null},
				{"key":{"field":"boolean","id":"id","set":"set"},
					"op":"=","type":"*bool","value":null},
				{"key":{"field":"time","id":"id","set":"set"},
					"op":"=","type":"*time","value":null},
				{"key":{"field":"bytes","id":"id","set":"set"},
					"op":"=","type":"*bytes","value":null}
			]`),
		},
	}

	for _, test := range tests {
		expected := bytes.ReplaceAll(test.encoded, []byte("\t"), []byte(""))
		expected = bytes.ReplaceAll(expected, []byte("\n"), []byte(""))

		data, err := Encode(0, test.ops)
		assert.Equal(expected, data, test.name)
		assert.Equal(test.err, err != nil, test.name)

		if !test.err {
			ops, err := Decode(0, data)
			assert.Equal(test.ops, ops, test.name)
			assert.Equal(test.err, err != nil, test.name)
		}
	}
}

func TestOpEncodeV0InvalidData(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		name string
		ops  []Op
	}{
		{
			name: "nil slice",
			ops: []Op{
				{
					Key: Key{
						Set:   "set",
						ID:    "id",
						Field: "field",
					},
					Op:    OpSet,
					Value: nil,
				},
			},
		}, {
			name: "nil slice",
			ops: []Op{
				{
					Key: Key{
						Set:   "set",
						ID:    "id",
						Field: "field",
					},
					Op: OpSet,
					Value: func() string {
						return "can't marshal a function"
					},
				},
			},
		},
	}

	for _, test := range tests {
		_, err := Encode(0, test.ops)
		assert.Error(err, test.name)
	}
}

func TestOpDecodeV0InvalidData(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		name    string
		encoded []byte
	}{
		{
			name:    "nil slice",
			encoded: nil,
		}, {
			name:    "zero bytes",
			encoded: []byte{},
		}, {
			name:    "invalid data",
			encoded: []byte(`{invalid_data}`),
		}, {
			name: "invalid op",
			encoded: []byte(`[
				{"key":{"field":"field","id":"id","set":"set"},
					"op":"!@#","type":"string","value":"string value"}
			]`),
		}, {
			name: "invalid op (integer)",
			encoded: []byte(`[
				{"key":{"field":"field","id":"id","set":"set"},
					"op":999,"type":"string","value":"string value"}
			]`),
		}, {
			name: "invalid op (integer)",
			encoded: []byte(`[
				{"key":{"field":["invalid","field"],"id":"id","set":"set"},
					"op":999,"type":"string","value":"string value"}
			]`),
		}, {
			name: "invalid op (integer)",
			encoded: []byte(`[
				{"key":{"field":"field","id":"id","set":"set"},
					"op":"=","type":"int","value":"won't fit in int"}
			]`),
		},
	}

	for _, test := range tests {
		_, err := Decode(0, test.encoded)
		assert.Error(err, test.name)
	}
}

func TestOpInvalidVersion(t *testing.T) {
	_, err := Encode(999, []Op{{}, {}})
	assert.EqualError(t, err, "unsupported version")

	_, err = Decode(999, []byte(`some bytes`))
	assert.EqualError(t, err, "unsupported version")
}

func ptr(v interface{}) interface{} {
	switch c := v.(type) {
	case string:
		return &c
	case int:
		return &c
	case int8:
		return &c
	case int16:
		return &c
	case int32:
		return &c
	case int64:
		return &c
	case uint:
		return &c
	case uint8:
		return &c
	case uint16:
		return &c
	case uint32:
		return &c
	case uint64:
		return &c
	case bool:
		return &c
	case time.Time:
		return &c
	case []byte:
		return &c
	default:
		return nil
	}
}

func nilptr(t string) interface{} {
	switch t {
	case "string":
		var p *string
		return p
	case "int":
		var p *int
		return p
	case "int8":
		var p *int8
		return p
	case "int16":
		var p *int16
		return p
	case "int32":
		var p *int32
		return p
	case "int64":
		var p *int64
		return p
	case "uint":
		var p *uint
		return p
	case "uint8":
		var p *uint8
		return p
	case "uint16":
		var p *uint16
		return p
	case "uint32":
		var p *uint32
		return p
	case "uint64":
		var p *uint64
		return p
	case "bool":
		var p *bool
		return p
	case "time.Time":
		var p *time.Time
		return p
	case "[]byte":
		var p *[]byte
		return p
	default:
		return nil
	}
}
