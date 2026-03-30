package router

import (
	"errors"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestNew_shouldSuccess(t *testing.T) {
	app := New()

	assert.NotNil(t, app)
}

func TestNew_shouldPanicWhenSetTrustedProxiesFails(t *testing.T) {
	orig := setTrustedProxies
	setTrustedProxies = func(app *gin.Engine) error {
		return errors.New("proxy error")
	}
	defer func() { setTrustedProxies = orig }()

	assert.Panics(t, func() { New() })
}
