package security

import (
	"github.com/saichler/my.simple/go/common"
	"net"
)

func (sp *ShallowSecurityProvider) writeEncrypted(conn net.Conn, str string) error {
	encData, err := common.Encrypt([]byte(str), sp.key)
	if err != nil {
		return err
	}
	err = common.Write([]byte(encData), conn)
	if err != nil {
		return err
	}
	return nil
}

func (sp *ShallowSecurityProvider) readEncrypted(conn net.Conn) (string, error) {
	inData, err := common.Read(conn)
	if err != nil {
		conn.Close()
		return "", err
	}

	decData, err := sp.Decrypt(string(inData))
	if err != nil {
		conn.Close()
		return "", err
	}
	return string(decData), nil
}
