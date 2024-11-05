package streamlogger

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewStreamLogger(t *testing.T) {
	testURI := "rabbitmq-stream://elasticWMS:elasticWMS@sof-srv-dev-rabbit-01.sting.sf:5552"
	testVHost := "/"
	s, err := New(testURI, testVHost, "log.check")

	assert.Nil(t, err)
	assert.NotNil(t, s)
	assert.Implements(t, (*StreamLogger)(nil), s)

	defer s.Close()
	err = s.Log("test_message")

	assert.Nil(t, err)
}
