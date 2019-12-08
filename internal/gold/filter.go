package gold

import (
	"bytes"
	"encoding/json"
	"regexp"
)

var (
	regexBcryptHash = regexp.MustCompile(`\$2[ayb]\$.{56}`)
	regexUUID       = regexp.MustCompile(`[0123456789abcdef-]{36}`)
)

// A Filter is a function that takes some content ([]byte) and returns a
// modified version of that content.
//
// This is useful for doing some work on the output before saving the content in
// a file.
type Filter func([]byte) []byte

// FilterFormatJSON formats the given JSON payload.
//
// It panics if it cannot parse the payload (invalid JSON).
func FilterFormatJSON(src []byte) []byte {
	dst := &bytes.Buffer{}

	err := json.Indent(dst, src, "", "\t")
	if err != nil {
		panic(err)
	}

	return dst.Bytes()
}

// FilterBcryptHashes swaps bcrypt hashes with "_HASH_".
func FilterBcryptHashes(src []byte) []byte {
	return regexBcryptHash.ReplaceAll(src, []byte("_HASH_"))
}

// FilterUUIDs swaps UUIDs with "00000000-0000-0000-0000-000000000000".
func FilterUUIDs(src []byte) []byte {
	return regexUUID.ReplaceAll(src, []byte("00000000-0000-0000-0000-000000000000"))
}
