package driver

import (
	"github.com/edgexfoundry/go-mod-core-contracts/v2/models"
	"time"
)

// Custom clients should implement this interface
type Client interface {
	Ping() bool
	Close()
	GetValue(param *GetParam) (interface{}, error)
	SetValue(param *SetParam) error
}

type templateClient struct {
	// todo add custom configuration if any
}

// NewClient Create the client via ConnectionInfo
func NewClient(info *ConnectionInfo) (Client, error) {
	// todo
	return &templateClient{}, nil
}

// Ping This method is used to check if the client is still connected
func (c *templateClient) Ping() bool {
	// todo
	return true
}

// Close This method is used to close connection
func (c *templateClient) Close() {
	// todo
	return
}

// GetValue This method is used to send `GET` command and return value
func (c *templateClient) GetValue(param *GetParam) (interface{}, error) {
	// todo
	return nil, nil
}

// SetValue This method is used to send `SET` command
func (c *templateClient) SetValue(param *SetParam) error {
	// todo
	return nil
}

// NewClientFromCache
func NewClientFromCache(deviceId string, protocols map[string]models.ProtocolProperties) (Client, error) {
	client, ok := driver.clientMap.Load(deviceId)
	if !ok || client == nil || !client.Ping() {
		// 创建连接信息
		info, err := CreateConnectionInfo(protocols)
		if err != nil {
			return nil, err
		}

		// 创建客户端
		client, err = NewClient(info)
		if err != nil {
			return nil, err
		}
		driver.clientMap.Store(deviceId, client)
	} else {
		// 命中的client增加寿命
		driver.clientMap.Grow(deviceId, 3*time.Minute)
	}

	return client, nil
}
