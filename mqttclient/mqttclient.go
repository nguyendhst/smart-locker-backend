// Package mqttclient is a wrapper for the Eclipse Paho MQTT Go client library with custom connection options to AdafruitIO
package mqttclient

import (
	"crypto/tls"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

const (
	host        = "io.adafruit.com"
	port        = "1883"
	port_secure = "8883"
	keepalive   = 60 * time.Second
)

var (
	// QoS
	QOS_0 = byte(0)
	QOS_1 = byte(1)
	QOS_2 = byte(2)

	// Feeds
	FeedTopics = []string{}
)

type (
	Client struct {
		username     string
		service_host string
		service_port string
		client       mqtt.Client
	}
)

func NewClient(username string, authKey string, secure bool) *Client {

	var _port string
	opts := mqtt.NewClientOptions()
	if secure {
		opts.AddBroker("ssl://" + host + ":" + port_secure)

		opts.SetTLSConfig(&tls.Config{
			InsecureSkipVerify: true,
			ClientAuth:         tls.NoClientCert,
		})

		_port = port_secure
	} else {
		opts.AddBroker("tcp://" + host + ":" + port)
		_port = port
	}
	opts.SetUsername(username)
	opts.SetPassword(authKey)
	//opts.SetClientID("smart-locker")
	opts.SetCleanSession(true)
	opts.SetKeepAlive(keepalive)
	opts.SetResumeSubs(false)

	return &Client{
		username:     username,
		service_host: host,
		service_port: _port,
		client:       mqtt.NewClient(opts),
	}
}

func (c *Client) Connect() error {
	token := c.client.Connect()
	if token.Wait() && token.Error() != nil {
		return token.Error()
	}
	return nil
}

func (c *Client) Disconnect() {
	c.client.Disconnect(250)
}

func (c *Client) Subscribe(topic string, qos byte, callback mqtt.MessageHandler) error {
	token := c.client.Subscribe(topic, qos, callback)
	if token.Wait() && token.Error() != nil {
		return token.Error()
	}
	return nil
}

func (c *Client) IsConnected() bool {
	return c.client.IsConnected()
}
