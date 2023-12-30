package serrors

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	err1 := New("test", "foo", "bar")
	err2 := Format("%w: %s", err1, "message")("baz", 2)
	err3 := fmt.Errorf("%w: %s", err2, "msg")

	assert.Equal(t, []any{"baz", 2, "foo", "bar"}, Attributes(err3))
	assert.Equal(t, "test: message: msg", err3.Error())
	assert.True(t, errors.Is(err3, err1))
}
