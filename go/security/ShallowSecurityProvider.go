package security

import (
	"errors"
	"fmt"
	"github.com/saichler/my.security/go/sec"
	"github.com/saichler/my.simple/go/common"
	"net"
	"strconv"
)

type ShallowSecurityProvider struct {
	secret string
	key    string
}

func NewShallowSecurityProvider(key, secret string) *ShallowSecurityProvider {
	sp := &ShallowSecurityProvider{}
	fmt.Println(key)
	sp.key = key
	sp.secret = secret
	return sp
}

func (sp *ShallowSecurityProvider) CanDial(host string, port uint32, salts ...interface{}) (net.Conn, error) {
	return net.Dial("tcp", host+":"+strconv.Itoa(int(port)))
}

func (sp *ShallowSecurityProvider) CanAccept(conn net.Conn, salts ...interface{}) error {
	return nil
}

func (sp *ShallowSecurityProvider) ValidateConnection(conn net.Conn, uuid string, salts ...interface{}) (string, error) {
	err := sec.WriteEncrypted(conn, []byte(sp.secret), salts...)
	if err != nil {
		conn.Close()
		return "", err
	}

	secret, err := sec.ReadEncrypted(conn, salts...)
	if err != nil {
		conn.Close()
		return "", err
	}

	if sp.secret != secret {
		conn.Close()
		fmt.Println(sp.secret, ":", secret)
		return "", errors.New("incorrect Secret/Key, aborting connection")
	}

	err = sec.WriteEncrypted(conn, []byte(uuid), salts...)
	if err != nil {
		conn.Close()
		return "", err
	}

	zside, err := sec.ReadEncrypted(conn, salts...)
	if err != nil {
		conn.Close()
		return "", err
	}

	return zside, nil
}

func (sp *ShallowSecurityProvider) Encrypt(data []byte, salts ...interface{}) (string, error) {
	return common.Encrypt(data, sp.key)
}

func (sp *ShallowSecurityProvider) Decrypt(data string, salts ...interface{}) ([]byte, error) {
	return common.Decrypt(data, sp.key)
}

func (sp *ShallowSecurityProvider) CanDo(action sec.Action, endpoint string, token string, salts ...interface{}) error {
	return nil
}
func (sp *ShallowSecurityProvider) CanView(typ string, attrName string, token string, salts ...interface{}) error {
	return nil
}
