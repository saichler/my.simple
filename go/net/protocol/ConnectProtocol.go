package protocol

import (
	"errors"
	"github.com/saichler/my.simple/go/common"
	"github.com/saichler/my.simple/go/utils/security"
	"net"
	"strconv"
)

func ConnectToAndValidateSecretAndKey(host, secret, key string, port int) (net.Conn, error) {
	conn, err := net.Dial("tcp", host+":"+strconv.Itoa(port))
	if err != nil {
		return nil, err
	}
	data, err := security.Encode([]byte(secret), key)
	if err != nil {
		return nil, err
	}

	err = common.Write([]byte(data), conn)
	if err != nil {
		return nil, err
	}

	inData, err := common.Read(conn)
	if string(inData) != "OK" {
		return nil, errors.New("Failed to connect, incorrect Key/Secret")
	}
	return conn, err
}

func ExchangeUuid(uuid, key string, conn net.Conn) (string, error) {
	data, err := security.Encode([]byte(uuid), key)
	if err != nil {
		return "", err
	}

	err = common.Write([]byte(data), conn)
	if err != nil {
		return "", err
	}

	inData, err := common.Read(conn)
	if err != nil {
		return "", err
	}
	return string(inData), nil
}
