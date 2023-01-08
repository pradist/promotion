package router

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestHealth(t *testing.T) {
	expected := gin.H{
		"message": "OK",
	}

	gin.SetMode(gin.TestMode)
	res := httptest.NewRecorder()
	ctx, r := gin.CreateTestContext(res)
	r.GET("/health", Health)

	req, err := http.NewRequestWithContext(ctx, "GET", "/health", nil)
	if err != nil {
		t.Errorf("got error: %s", err)
	}

	//Act
	r.ServeHTTP(res, req)

	var response map[string]string
	err = json.Unmarshal([]byte(res.Body.Bytes()), &response)
	if err != nil {
		t.Errorf("got error: %s", err)
	}
	value, exists := response["message"]

	//Assert
	assert.Equal(t, http.StatusOK, res.Code)
	assert.True(t, exists)
	assert.Equal(t, expected["message"], value)
}
