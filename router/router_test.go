package router

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNew_shouldSuccess(t *testing.T) {
	app := New()

	assert.NotNil(t, app)
}
