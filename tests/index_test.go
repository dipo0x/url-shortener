package test

import (
    "encoding/json"
	"io"
	"net/http"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestIndexRoute(t *testing.T) {

	req, _ := http.NewRequest("GET", "/api/index", nil)

	res, err := testApp.Test(req, -1)

	assert.Nil(t, err)
	assert.Equal(t, 200, res.StatusCode)

	body, _ := io.ReadAll(res.Body)

    var response map[string]interface{}
    err = json.Unmarshal(body, &response)
    assert.Nil(t, err)

    assert.Equal(t, true, response["success"])
}
