package broker

import (
	"time"

	"github.com/segmentio/kafka-go"
)

type KafkaClient struct {
	Connection *kafka.Conn
}

func (c *KafkaClient) SendMessage(value string) error {
	err := c.Connection.SetWriteDeadline(time.Now().Add(10 * time.Minute))
	if err != nil {
		return err
	}
	_, err = c.Connection.Write([]byte(value))
	if err != nil {
		return err
	}
	return nil
}
