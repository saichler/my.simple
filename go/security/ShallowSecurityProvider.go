package security

import (
	"errors"
	"github.com/saichler/my.simple/go/common"
	"net"
	"strconv"
)
import "github.com/saichler/my.security/go/sec_common"

type ShallowSecurityProvider struct {
	secret string
	key    string
}

func NewShallowSecurityProvider(key, secret string) *ShallowSecurityProvider {
	sp := &ShallowSecurityProvider{}
	sp.key = key
	sp.secret = secret
	sec_common.MySecurityProvider = sp
	return sp
}

func (sp *ShallowSecurityProvider) CanDial(host string, port uint32, salts ...interface{}) (net.Conn, error) {
	return net.Dial("tcp", host+":"+strconv.Itoa(int(port)))
}

func (sp *ShallowSecurityProvider) CanAccept(conn net.Conn) error {
	return nil
}

func (sp *ShallowSecurityProvider) ValidateConnection(conn net.Conn, uuid string, salts ...interface{}) (string, error) {
	err := sp.writeEncrypted(conn, sp.secret)
	if err != nil {
		conn.Close()
		return "", err
	}

	secret, err := sp.readEncrypted(conn)
	if err != nil {
		conn.Close()
		return "", err
	}

	if sp.secret != secret {
		conn.Close()
		return "", errors.New("incorrect Secret/Key, aborting connection")
	}

	err = sp.writeEncrypted(conn, uuid)
	if err != nil {
		conn.Close()
		return "", err
	}

	zside, err := sp.readEncrypted(conn)
	if err != nil {
		conn.Close()
		return "", err
	}

	return zside, nil
}

func (sp *ShallowSecurityProvider) Encrypt(data []byte) (string, error) {
	return common.Encrypt(data, sp.key)
}

func (sp *ShallowSecurityProvider) Decrypt(data string) ([]byte, error) {
	return common.Decrypt(data, sp.key)
}
