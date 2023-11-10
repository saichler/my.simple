package tests

import (
	"fmt"
	"github.com/saichler/my.simple/go/api_gateways"
	"github.com/saichler/my.simple/go/api_gateways/model"
	"github.com/saichler/my.simple/go/common"
	"testing"
	"time"
)

func TestRestfulServer(t *testing.T) {
	common.CreateDefaultTestCertificate()
	config := &model.RestFulGatewayConfig{
		Host: "127.0.0.1",
		Port: 8980,
		Crt:  "/tmp/test-crt.crt",
		Key:  "/tmp/test-crt.crtKey",
	}
	server, err := apigateways.NewRestFulGateway(config)
	if err != nil {
		t.Fail()
		fmt.Println(err)
		return
	}
	go server.Start()
	time.Sleep(time.Second * 5)
}
