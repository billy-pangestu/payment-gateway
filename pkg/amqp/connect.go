package amqp

import (
	"github.com/streadway/amqp"
)

// Connection ...
type Connection struct {
	URL string
}

// Connect ...
func (m Connection) Connect() (*amqp.Connection, *amqp.Channel, error) {
	var (
		err         error
		conn        *amqp.Connection
		amqpChannel *amqp.Channel
	)

	conn, err = amqp.Dial(m.URL)
	if err != nil {
		return conn, amqpChannel, err
	}

	amqpChannel, err = conn.Channel()

	return conn, amqpChannel, err
}
