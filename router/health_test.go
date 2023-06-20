package router

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type HealthResponse struct {
	Message string `json:"message"`
}

func TestHealth(t *testing.T) {
	expected := HealthResponse{
		Message: "OK",
	}

	gin.SetMode(gin.TestMode)
	res := httptest.NewRecorder()
	ctx, r := gin.CreateTestContext(res)
	r.GET("/health", Health)

	req, err := http.NewRequestWithContext(ctx, "GET", "/health", nil)
	if err != nil {
		t.Errorf("got error: %s", err)
	}

	r.ServeHTTP(res, req)
	var response HealthResponse
	err = json.Unmarshal(res.Body.Bytes(), &response)

	assert.NoError(t, err, "failed to unmarshal response")
	assert.Equal(t, http.StatusOK, res.Code, "unexpected status code")
	assert.Equal(t, expected.Message, response.Message, "unexpected response message")
}
