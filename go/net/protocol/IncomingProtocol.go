package protocol

import (
	"errors"
	"github.com/saichler/my.security/go/sec_common"
	"github.com/saichler/my.simple/go/common"
	"net"
)

func Incoming(conn net.Conn, uuid string) (string, error) {
	initData, err := common.Read(conn)
	if err != nil {
		conn.Close()
		return "", err
	}

	data, err := sec_common.MySecurityProvider.Decrypt(string(initData))
	if err != nil {
		conn.Close()
		return "", err
	}

	if !sec_common.MySecurityProvider.IsSecret(string(data)) {
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

	data, err = sec_common.MySecurityProvider.Decrypt(string(initData))
	if err != nil {
		conn.Close()
		return "", err
	}

	portUuid := string(data)

	// @TODO - need to encrypt this as well

	err = common.Write([]byte(uuid), conn)
	if err != nil {
		conn.Close()
		return "", errors.New("Failed to write response")
	}
	return portUuid, nil
}
