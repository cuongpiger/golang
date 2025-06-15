package api

import (
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestHelloHandler(t *testing.T) {
    // Create a request
    req, err := http.NewRequest("GET", "/hello", nil)
    assert.NoError(t, err, "creating request should not fail")

    // Create a response recorder
    rr := httptest.NewRecorder()

    // Call the handler
    handler := http.HandlerFunc(HelloHandler)
    handler.ServeHTTP(rr, req)

    // Check the status code
    assert.Equal(t, http.StatusOK, rr.Code, "handler should return 200 OK")

    // Check the content type
    assert.Equal(t, "application/json", rr.Header().Get("Content-Type"),
                "content type should be application/json")

    // Check the response body
    var response Response
    err = json.Unmarshal(rr.Body.Bytes(), &response)
    assert.NoError(t, err, "unmarshaling response should not fail")
    assert.Equal(t, "Hello, World!", response.Message, "message should match")
}