package http

import (
	"io/ioutil"
	"strings"
	"time"
	"uit_payment/entities/payment/delivery/response"
)

func NewHealth() *Health {
	return &Health{}
}

type Health struct {
	Handler
}

func (this *Health) Handle() {
	this.RenderSuccess(response.HealthResponse{getVersion(), time.Now().String()})
}

func getVersion() string {

	version, err := ioutil.ReadFile("./version")
	if err != nil {
		return ""
	}
	return strings.TrimRight(string(version), "\r\n")
}
