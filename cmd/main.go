package main

import (
	"github.com/edgexfoundry/device-sdk-go/v2"
	"github.com/edgexfoundry/device-sdk-go/v2/pkg/startup"
	"github.com/kainhuck/device-template-go/internal/driver"
	"log"
	"net/http"
)

const (
	serviceName string = "device-service-template"
)

func main() {
	go func() {
		log.Println(http.ListenAndServe("0.0.0.0:61618", nil))
	}()

	sd := driver.NewProtocolDriver()
	startup.Bootstrap(serviceName, device.Version, sd)
}
