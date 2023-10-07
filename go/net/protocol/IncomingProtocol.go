package protocol

import (
	"errors"
	"github.com/saichler/my.simple/go/common"
	"github.com/saichler/my.simple/go/utils/security"
	"net"
)

func Incoming(conn net.Conn, key, secret, uuid string) (string, error) {
	initData, err := common.Read(conn)
	if err != nil {
		conn.Close()
		return "", err
	}

	data, err := security.Decode(string(initData), key)
	if err != nil {
		conn.Close()
		return "", err
	}

	if string(data) != secret {
		conn.Close()
		return "", errors.New("Incorrect Secret/Key, aborting connection")
	}

	err = common.Write([]byte("OK"), conn)
	if err != nil {
		conn.Close()
		return "", errors.New("Failed to write response")
	}

	initData, err = common.Read(conn)
	if err != nil {
		conn.Close()
		return "", err
	}

	data, err = security.Decode(string(initData), key)
	if err != nil {
		conn.Close()
		return "", err
	}

	portUuid := string(data)

	err = common.Write([]byte(uuid), conn)
	if err != nil {
		conn.Close()
		return "", errors.New("Failed to write response")
	}
	return portUuid, nil
}
