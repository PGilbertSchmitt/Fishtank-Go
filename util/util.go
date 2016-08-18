// Package util provides some basic utilities like hashing and data storage.
package util

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"hash/fnv"
	"reflect"
)

// Vertex is a simple 2D point
type Vertex struct {
	X, Y int32
}

// Hash32 provides a hash of the given struct, first by reflecting the struct
// value and converting each part into bytes which gets fed into the hash.
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
