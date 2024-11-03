package streamlogger

import "github.com/rabbitmq/rabbitmq-stream-go-client/pkg/stream"

type StreamLogger interface {
	Log(message string) error
}

type streamlogger struct {
	env *stream.Environment
}

func New(uri string, vhost string) (StreamLogger, error) {
	s := &streamlogger{}
	env, err := stream.NewEnvironment(
		stream.NewEnvironmentOptions().
			SetUri(uri).SetVHost(vhost))

	if err != nil {
		return nil, err
	}
	s.env = env
	return s, nil
}

func (s *streamlogger) Log(message string) error {
	return nil
}
