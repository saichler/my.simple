package protobuf_object

import (
	"math"
)

type Float32 struct{}

func (_type *Float32) Add(any interface{}) ([]byte, int) {
	f := any.(float32)
	i := math.Float32bits(f)

	bytes := make([]byte, 4)
	bytes[3] = byte(i)
	bytes[2] = byte(i >> 8)
	bytes[1] = byte(i >> 16)
	bytes[0] = byte(i >> 24)

	return bytes, 4
}

func (_type *Float32) Get(data []byte, location int) (interface{}, int) {
	var result uint32
	v1 := uint32(data[location]) << 24
	v2 := uint32(data[location+1]) << 16
	v3 := uint32(data[location+2]) << 8
	v4 := uint32(data[location+3])
	result = v1 + v2 + v3 + v4
	return math.Float32frombits(result), 4
}
