package gosimplefilelog_test

import (
	"gosimplefilelog"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOnlyGetMessage(t *testing.T) {
	l := gosimplefilelog.LogContext{}
	message := "log1"

	assert := assert.New(t)
	assert.Equal(l.GetMessage(message), message)
}

func TestOnlyStart(t *testing.T) {
	l := gosimplefilelog.LogContext{}
	input := "log1"
	expected := "-> " + input

	assert := assert.New(t)
	assert.Equal(l.Start(input), expected)
}
