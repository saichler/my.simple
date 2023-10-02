package common

import (
	"errors"
	"google.golang.org/protobuf/proto"
	"net"
	"time"
)

func Send(pb proto.Message, port Port) error {
	data, err := proto.Marshal(pb)
	if err != nil {
		return err
	}
	return port.Send(data)
}

func Write(data []byte, conn net.Conn) error {
	if conn == nil {
		return errors.New("No Connection Available")
	}
	_, e := conn.Write(Long2Bytes(int64(len(data))))
	if e != nil {
		return e
	}
	_, e = conn.Write(data)
	return e
}

func Read(conn net.Conn) ([]byte, error) {
	sizebytes, err := ReadSize(8, conn)
	if sizebytes == nil || err != nil {
		return nil, err
	}
	size := Bytes2Long(sizebytes)
	if size > MAX_DATA_SIZE {
		return nil, errors.New("Max Size Exceeded!")
	}
	data, err := ReadSize(int(size), conn)
	return data, err
}

func ReadSize(size int, conn net.Conn) ([]byte, error) {
	data := make([]byte, size)
	n, e := conn.Read(data)
	if e != nil {
		return nil, errors.New("Failed to read data size:" + e.Error())
	}

	if n < size {
		if n == 0 {
			time.Sleep(time.Second)
		}
		data = data[0:n]
		left, e := ReadSize(size-n, conn)
		if e != nil {
			return nil, errors.New("Failed to read packet size:" + e.Error())
		}
		data = append(data, left...)
	}
	return data, nil
}

func Bytes2Long(data []byte) int64 {
	v1 := int64(data[0]) << 56
	v2 := int64(data[1]) << 48
	v3 := int64(data[2]) << 40
	v4 := int64(data[3]) << 32
	v5 := int64(data[4]) << 24
	v6 := int64(data[5]) << 16
	v7 := int64(data[6]) << 8
	v8 := int64(data[7])
	return v1 + v2 + v3 + v4 + v5 + v6 + v7 + v8
}

func Long2Bytes(s int64) []byte {
	size := make([]byte, 8)
	size[7] = byte(s)
	size[6] = byte(s >> 8)
	size[5] = byte(s >> 16)
	size[4] = byte(s >> 24)
	size[3] = byte(s >> 32)
	size[2] = byte(s >> 40)
	size[1] = byte(s >> 48)
	size[0] = byte(s >> 56)
	return size
}
