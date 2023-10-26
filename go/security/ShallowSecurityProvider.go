package security

import "github.com/saichler/my.simple/go/common"
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

func (sp *ShallowSecurityProvider) CanConnectFrom(protocol, source string, port uint32) bool {
	return true
}

func (sp *ShallowSecurityProvider) CanConnectTo(protocol, source string, port uint32) bool {
	return true
}

func (sp *ShallowSecurityProvider) Encrypt(data []byte) (string, error) {
	return common.Encrypt(data, sp.key)
}

func (sp *ShallowSecurityProvider) Decrypt(data string) ([]byte, error) {
	return common.Decrypt(data, sp.key)
}

func (sp *ShallowSecurityProvider) EncryptedSecret() (string, error) {
	return common.Encrypt([]byte(sp.secret), sp.key)
}

func (sp *ShallowSecurityProvider) IsSecret(secret string) bool {
	return sp.secret == secret
}
