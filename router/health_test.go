package router

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
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

	//Act
	r.ServeHTTP(res, req)

	var response HealthResponse
	err = json.Unmarshal(res.Body.Bytes(), &response)
	require.NoError(t, err, "failed to unmarshal response")

	//Assert
	require.Equal(t, http.StatusOK, res.Code, "unexpected status code")
	require.Equal(t, expected.Message, response.Message, "unexpected response message")
}
