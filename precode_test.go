package main

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4
	req, _ := http.NewRequest("GET", "/cafe", nil)

	req.URL.RawQuery = prepareQuery(req, "5", "moscow")

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	// здесь нужно добавить необходимые проверки
	resp := responseRecorder.Body.String()
	body := strings.Split(resp, ",")
	assert.Len(t, body, totalCount)
}

func TestMainHandlerWhenUnsupportedCity(t *testing.T) {
	expectedError := "wrong city value"
	expectedCode := http.StatusBadRequest
	req, err := http.NewRequest("GET", "/cafe", nil)
	if err != nil {
		require.NoError(t, err)
	}

	req.URL.RawQuery = prepareQuery(req, "1", "ufa")

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)
	resCode := responseRecorder.Code
	require.Equal(t, expectedCode, resCode)
	resp := responseRecorder.Body.String()
	require.Equal(t, resp, expectedError)
}

func TestMainHandlerWhenQueryIsCorrect(t *testing.T) {
	expectedCode := http.StatusOK
	req, err := http.NewRequest("GET", "/cafe", nil)
	if err != nil {
		require.NoError(t, err)
	}

	req.URL.RawQuery = prepareQuery(req, "1", "moscow")

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)
	respCode := responseRecorder.Code
	require.Equal(t, expectedCode, respCode)
	require.NotEmpty(t, responseRecorder.Body)
}

func prepareQuery(req *http.Request, count string, city string) string {
	q := req.URL.Query()
	q.Add("count", count)
	q.Add("city", city)
	return q.Encode()
}
