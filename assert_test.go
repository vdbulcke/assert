package assert_test

import (
	"errors"
	"testing"

	"github.com/vdbulcke/assert"
)

func TestErrorWithColor(t *testing.T) {

	err := errors.New("an error")
	assert.IsTTY = true
	assert.NoErr(err, assert.DefaultMode, "foo", "var", []string{"hello", "world"})
}
