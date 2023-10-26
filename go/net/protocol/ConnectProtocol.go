package protocol

import (
	"errors"
	"github.com/saichler/my.security/go/sec_common"
	"github.com/saichler/my.simple/go/common"
	"net"
	"strconv"
)

func ConnectToAndValidateSecretAndKey(host string, port int32) (net.Conn, error) {
	conn, err := net.Dial("tcp", host+":"+strconv.Itoa(int(port)))
	if err != nil {
		return nil, err
	}
	data, err := sec_common.MySecurityProvider.EncryptedSecret()
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

func ExchangeUuid(uuid string, conn net.Conn) (string, error) {
	data, err := sec_common.MySecurityProvider.Encrypt([]byte(uuid))
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
