package query

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/mfcochauxlaberge/jsonapi"
)

// Encode ...
func Encode(v uint, ops []Op) ([]byte, error) {
	if len(ops) == 0 {
		return nil, errors.New("no ops")
	}

	var (
		enc []byte
		err error
	)

	switch v {
	case 0:
		// Version 0
		enc, err = encodeV0(ops)
		if err != nil {
			return nil, err
		}
	default:
		return nil, errors.New("unsupported version")
	}

	return enc, nil
}

// Decode ...
func Decode(v uint, data []byte) ([]Op, error) {
	if len(data) == 0 {
		return nil, errors.New("no bytes")
	}

	var (
		ops []Op
		err error
	)

	switch v {
	case 0:
		// Version 0
		ops, err = decodeV0(data)
		if err != nil {
			return nil, err
		}
	default:
		return nil, errors.New("unsupported version")
	}

	return ops, nil
}

func encodeV0(ops []Op) ([]byte, error) {
	enc := make([]json.RawMessage, 0)

	for _, op := range ops {
		m := map[string]interface{}{}

		key, err := json.Marshal(map[string]string{
			"set":   op.Key.Set,
			"id":    op.Key.ID,
			"field": op.Key.Field,
		})
		if err != nil {
			return nil, err
		}

		m["key"] = json.RawMessage(key)
		m["op"] = string(op.Op)
		m["value"] = op.Value

		if op.Value == nil {
			return nil, errors.New("op has nil value")
		}

		m["type"] = jsonapi.GetAttrTypeString(
			jsonapi.GetAttrType(fmt.Sprintf("%T", op.Value)),
		)

		data, err := json.Marshal(m)
		if err != nil {
			return nil, err
		}

		enc = append(enc, data)
	}

	res, err := json.Marshal(enc)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func decodeV0(data []byte) ([]Op, error) {
	ops := []Op{}

	raws := []json.RawMessage{}

	err := json.Unmarshal(data, &raws)
	if err != nil {
		return nil, err
	}

	for _, raw := range raws {
		op := Op{}
		m := map[string]json.RawMessage{}

		err = json.Unmarshal(raw, &m)
		if err != nil {
			return nil, err
		}

		// Key
		err = json.Unmarshal(m["key"], &op.Key)
		if err != nil {
			return nil, err
		}

		// Op
		opStr := ""

		err := json.Unmarshal(m["op"], &opStr)
		if err != nil {
			return nil, err
		}

		op.Op = NewOp(opStr)
		if op.Op == 0 {
			return nil, fmt.Errorf("unknown op %q", op.Op)
		}

		// Value
		typ := strings.Trim(string(m["type"]), `"`)

		t, n := jsonapi.GetAttrType(typ)
		attr := jsonapi.Attr{
			Name:     "",
			Type:     t,
			Nullable: n,
		}

		op.Value, err = attr.UnmarshalToType(m["value"])
		if err != nil {
			return nil, err
		}

		ops = append(ops, op)
	}

	return ops, nil
}
