package driver

import (
	"fmt"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/models"
)

type ServiceConfig struct {
	CustomConfig CustomConfig
}

func (sw *ServiceConfig) UpdateFromRaw(rawConfig interface{}) bool {
	configuration, ok := rawConfig.(*ServiceConfig)
	if !ok {
		return false //errors.New("unable to cast raw config to type 'ServiceConfig'")
	}

	*sw = *configuration

	return true
}

// ConnectionInfo The necessary information to create a client connection
type ConnectionInfo struct {
	// todo
}

func CreateConnectionInfo(protocols map[string]models.ProtocolProperties) (info *ConnectionInfo, err error) {
	info = new(ConnectionInfo)
	protocol, ok := protocols[Protocol]
	if !ok {
		return nil, fmt.Errorf("unable to load config, '%s' not exist", Protocol)
	}
	err = Load(protocol, info)

	return
}

// CustomConfig Driver custom configuration
type CustomConfig struct {
	// todo your rename this struct
}
