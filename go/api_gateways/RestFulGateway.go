package apigateways

import (
	"github.com/saichler/my.simple/go/api_gateways/model"
	"github.com/saichler/my.simple/go/utils/logs"
	"net/http"
	"os"
	"strconv"
	"time"
)

type RestFulGateway struct {
	webServer *http.Server
	config    *model.RestFulGatewayConfig
	endpoints []string
}

func NewRestFulGateway(config *model.RestFulGatewayConfig) (*RestFulGateway, error) {
	rs := &RestFulGateway{}
	rs.config = config
	rs.endpoints = make([]string, 0)

	if rs.config.Crt != "" {
		_, err := os.Open(rs.config.Crt)
		if err != nil {
			return rs, err
		}
	}
	return rs, nil
}

func (gateway *RestFulGateway) Start() error {
	var err error
	gateway.webServer = &http.Server{
		Addr:    gateway.config.Host + ":" + strconv.Itoa(int(gateway.config.Port)),
		Handler: http.DefaultServeMux,
	}
	if gateway.config.Crt != "" && gateway.config.Key != "" {
		logs.Info("Starting https service")
		err = gateway.webServer.ListenAndServeTLS(gateway.config.Crt, gateway.config.Key)
	} else {
		logs.Info("Starting http service")
		err = gateway.webServer.ListenAndServe()
	}
	return err
}

func (gateway *RestFulGateway) Stop() {
	gateway.webServer.Shutdown(gateway)
}

func (gateway *RestFulGateway) AddEndpoint(endPoint string) {
	http.DefaultServeMux.HandleFunc(endPoint, gateway.forward)
}

func (gateway *RestFulGateway) forward(writer http.ResponseWriter, request *http.Request) {

}

func (gateway *RestFulGateway) Deadline() (deadline time.Time, ok bool) {
	return time.Now(), true
}

func (gateway *RestFulGateway) Done() <-chan struct{} {
	return nil
}

func (gateway *RestFulGateway) Err() error {
	return nil
}

func (gateway *RestFulGateway) Value(key any) any {
	return nil
}
