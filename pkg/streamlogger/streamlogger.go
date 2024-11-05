package streamlogger

import (
	"github.com/rabbitmq/rabbitmq-stream-go-client/pkg/amqp"
	"github.com/rabbitmq/rabbitmq-stream-go-client/pkg/stream"
)

type StreamLogger interface {
	Log(message string) error
	Close() error
}

type streamlogger struct {
	env        *stream.Environment
	streamName string
	producer   *stream.Producer
}

func New(uri string, vhost string, streamName string) (StreamLogger, error) {
	env, err := stream.NewEnvironment(
		stream.NewEnvironmentOptions().
			SetUri(uri).SetVHost(vhost))

	if err != nil {
		return nil, err
	}
	err = env.DeclareStream(streamName,
		stream.NewStreamOptions().
			SetMaxLengthBytes(stream.ByteCapacity{}.GB(2)))

	if err != nil {
		return nil, err
	}

	producer, err := env.NewProducer(streamName, nil)

	if err != nil {
		return nil, err
	}

	s := &streamlogger{
		producer:   producer,
		env:        env,
		streamName: streamName,
	}

	return s, nil
}

func (s *streamlogger) Log(message string) error {
	amqp_message := amqp.NewMessage([]byte(message))
	return s.producer.Send(amqp_message)
}

func (s *streamlogger) Close() error {
	return s.env.Close()
}
