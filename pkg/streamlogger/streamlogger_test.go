package streamlogger

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewStreamLogger(t *testing.T) {
	testURI := "rabbitmq-stream://elasticWMS:elasticWMS@sof-srv-dev-rabbit-01.sting.sf:5552"
	testVHost := "test/ElasticWMS"
	s, err := New(testURI, testVHost)

	assert.Nil(t, err)
	assert.NotNil(t, s)
	assert.Implements(t, (*StreamLogger)(nil), s)

}
