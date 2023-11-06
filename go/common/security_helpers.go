package common

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"io"
	"io/ioutil"
	"math/big"
	mathrand "math/rand"
	"net"
	"os"
	"time"
)

var l = []rune("abcdefghijklmnopqrstuvwxyz0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ")

func GenerateAES256Key() string {
	mathrand.Seed(time.Now().UnixNano())
	key := make([]rune, 32)
	for i := range key {
		key[i] = l[mathrand.Intn(len(l))]
	}
	return string(key)
}

func Encrypt(dataToEncode []byte, key string) (string, error) {
	k := []byte(key)
	block, err := aes.NewCipher(k)
	if err != nil {
		return "", err
	}

	l := len(dataToEncode)
	cipherdata := make([]byte, aes.BlockSize+l)

	iv := cipherdata[:aes.BlockSize]
	_, err = io.ReadFull(rand.Reader, iv)
	if err != nil {
		return "", err
	}

	cfb := cipher.NewCFBEncrypter(block, iv)
	cfb.XORKeyStream(cipherdata[aes.BlockSize:], dataToEncode)
	return base64.StdEncoding.EncodeToString(cipherdata), nil
}

func Decrypt(stringToDecode, key string) ([]byte, error) {
	encData, err := base64.StdEncoding.DecodeString(stringToDecode)
	if err != nil {
		return nil, err
	}
	if len(encData) < aes.BlockSize {
		err = errors.New("encrypted data does not have an iv spec")
		return nil, err
	}
	k := []byte(key)
	block, err := aes.NewCipher(k)
	if err != nil {
		return nil, err
	}
	iv := encData[:aes.BlockSize]
	encData = encData[aes.BlockSize:]
	cfb := cipher.NewCFBDecrypter(block, iv)
	data := make([]byte, len(encData))
	cfb.XORKeyStream(data, encData)
	return data, nil
}

func CreateCA(filenamePrefix, org, country, county, city, street, zipcode, email string, years int) (*x509.Certificate, *rsa.PrivateKey, error) {
	_, e := os.Stat(filenamePrefix + ".ca")
	if e != nil {
		ca := &x509.Certificate{
			SerialNumber: big.NewInt(2019),
			Subject: pkix.Name{
				Organization:  []string{org},
				Country:       []string{country},
				Province:      []string{county},
				Locality:      []string{city},
				StreetAddress: []string{street},
				PostalCode:    []string{zipcode},
			},
			NotBefore:             time.Now(),
			NotAfter:              time.Now().AddDate(years, 0, 0),
			IsCA:                  true,
			ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
			KeyUsage:              x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
			BasicConstraintsValid: true,
			EmailAddresses:        []string{email},
		}

		caKey, err := rsa.GenerateKey(rand.Reader, 4096)
		if err != nil {
			return nil, nil, err
		}

		caData, err := x509.CreateCertificate(rand.Reader, ca, ca, &caKey.PublicKey, caKey)
		if err != nil {
			return nil, nil, err
		}

		caPEM := &bytes.Buffer{}
		pem.Encode(caPEM, &pem.Block{
			Type:  "CERTIFICATE",
			Bytes: caData,
		})

		err = ioutil.WriteFile(filenamePrefix+".ca", caPEM.Bytes(), 0777)
		if err != nil {
			return nil, nil, err
		}

		caKeyPEM := &bytes.Buffer{}
		err = pem.Encode(caKeyPEM, &pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: x509.MarshalPKCS1PrivateKey(caKey),
		})
		if err != nil {
			return nil, nil, err
		}
		err = ioutil.WriteFile(filenamePrefix+".caKey", caKeyPEM.Bytes(), 0777)
		return ca, caKey, err
	} else {
		return nil, nil, errors.New("Certificate Authority " + filenamePrefix + " already exists!")
	}
}

func CreateCrt(filenamePrefix, org, country, county, city, street, zipcode, email, ip, secret string, port int64, years int, ca *x509.Certificate, caKey *rsa.PrivateKey) error {
	_, e := os.Stat(filenamePrefix + ".crt")
	if e != nil {
		ipAddress := net.ParseIP(ip)
		crt := &x509.Certificate{
			SerialNumber: big.NewInt(port),
			Subject: pkix.Name{
				Organization:  []string{org},
				Country:       []string{country},
				Province:      []string{county},
				Locality:      []string{city},
				StreetAddress: []string{street},
				PostalCode:    []string{zipcode},
			},
			EmailAddresses: []string{email},
			IPAddresses:    []net.IP{ipAddress},
			NotBefore:      time.Now(),
			NotAfter:       time.Now().AddDate(years, 0, 0),
			SubjectKeyId:   []byte(secret),
			ExtKeyUsage:    []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
			KeyUsage:       x509.KeyUsageDigitalSignature,
		}

		key, err := rsa.GenerateKey(rand.Reader, 4096)
		if err != nil {
			return err
		}

		crtData, err := x509.CreateCertificate(rand.Reader, crt, ca, &key.PublicKey, caKey)
		if err != nil {
			return err
		}

		crtPEM := new(bytes.Buffer)
		pem.Encode(crtPEM, &pem.Block{
			Type:  "CERTIFICATE",
			Bytes: crtData,
		})

		err = ioutil.WriteFile(filenamePrefix+".crt", crtPEM.Bytes(), 0777)
		if err != nil {
			return err
		}

		keyPEM := new(bytes.Buffer)
		pem.Encode(keyPEM, &pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: x509.MarshalPKCS1PrivateKey(key),
		})
		err = ioutil.WriteFile(filenamePrefix+".crtKey", keyPEM.Bytes(), 0777)
		return err
	} else {
		return errors.New("Certificate " + filenamePrefix + " already exists!")
	}
}

func CreateDefaultTestCertificate() error {
	ca, caKey, err := CreateCA("/tmp/test-ca", "saichler", "USA", "Santa Clara",
		"San Jose", "", "", "saichler@gmail.com", 1)
	if err != nil {
		return err
	} else {
		err = CreateCrt("/tmp/test-crt", "saichler", "USA", "Santa Clara",
			"San Jose", "", "", "saichler@gmail.com", "127.0.0.1",
			"MySecret", int64(8980), 1, ca, caKey)
		return err
	}
}
