// Package util provides some basic utilities like hashing and data storage.
package util

import (
	"fmt"
	"hash/fnv"
	"reflect"
	"bytes"
	"encoding/binary"
)

// Simple 2D vector struct
type Vector struct {
	X, Y int32
}

// Hashes the given value, first by reflecting the value (which is likely a
// struct) and converting each part into bytes which gets fed into the hash.
// Because reflection only gives access to exposed members, changes to unexposed
// members will not affect the hash.
func Hash32(i interface{}) uint32 {
	hashSum := fnv.New32()
	values := reflect.ValueOf(i)
	for i := 0; i < values.NumField(); i++ {
		val := values.Field(i).Interface()
		hashSum.Write(toBytes(val))
	}
	return hashSum.Sum32()
}

// Converts any datatype into a slice of bytes
func toBytes(i interface{}) []byte {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.LittleEndian, i)
	if err != nil {
		fmt.Printf("Error converting to byte slice: %v", err)
	}
	return buf.Bytes()
}
