package test

import (
    "bytes"
    "encoding/json"
    "net/http"
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestCreateURLSuccess(t *testing.T) {
    payload := map[string]string{
        "url":       "https://github.com/dipo0",
        "expiresAt": "24",
    }
    body, _ := json.Marshal(payload)

    req, _ := http.NewRequest("POST", "/api/url/create-url", bytes.NewReader(body))
    req.Header.Set("Content-Type", "application/json")

    res, err := testApp.Test(req, -1)

    assert.Nil(t, err)
    assert.Equal(t, 200, res.StatusCode)

    var response map[string]interface{}
    json.NewDecoder(res.Body).Decode(&response)

    assert.Equal(t, true, response["success"])
    assert.Equal(t, 200, int(response["status"].(float64)))
    assert.NotNil(t, response["data"])
}

func TestCreateURLFailure(t *testing.T) {
    payload := map[string]string{
        "url": "https://github.com/dipo0",
    }
    body, _ := json.Marshal(payload)

    req, _ := http.NewRequest("POST", "/api/url/create-url", bytes.NewReader(body))
    req.Header.Set("Content-Type", "application/json")

    res, err := testApp.Test(req, -1)

    assert.Nil(t, err)
    assert.Equal(t, 400, res.StatusCode)

    var response map[string]interface{}
    json.NewDecoder(res.Body).Decode(&response)

    assert.Equal(t, false, response["success"])
    assert.Equal(t, 400, int(response["status"].(float64)))
    assert.NotNil(t, response["error"])
}