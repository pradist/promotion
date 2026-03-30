package router

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"runtime/debug"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type VersionResponse struct {
	Version string `json:"version"`
}

func TestVersion(t *testing.T) {
	expected := VersionResponse{
		Version: AppVersion(),
	}

	gin.SetMode(gin.TestMode)
	res := httptest.NewRecorder()
	ctx, r := gin.CreateTestContext(res)
	r.GET("/version", Version)

	req, err := http.NewRequestWithContext(ctx, "GET", "/version", nil)
	if err != nil {
		t.Errorf("got error: %s", err)
	}

	r.ServeHTTP(res, req)
	var response VersionResponse
	err = json.Unmarshal(res.Body.Bytes(), &response)

	assert.NoError(t, err, "failed to unmarshal response")
	assert.Equal(t, http.StatusOK, res.Code, "unexpected status code")
	assert.Equal(t, expected.Version, response.Version, "unexpected version")
}

func TestAppVersion_returnUnknownWhenReadBuildInfoFails(t *testing.T) {
	orig := readBuildInfo
	readBuildInfo = func() (*debug.BuildInfo, bool) {
		return nil, false
	}
	defer func() { readBuildInfo = orig }()

	assert.Equal(t, "unknown", AppVersion())
}
