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
	req, err := http.NewRequest("GET", "/cafe", nil)
	if err != nil {
		require.NoError(t, err)
	}

	q := req.URL.Query()
	q.Add("count", "5")
	q.Add("city", "moscow")
	req.URL.RawQuery = q.Encode()

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	// здесь нужно добавить необходимые проверки
	resp := responseRecorder.Body.String()
	body := strings.Split(resp, ",")
	assert.Equal(t, totalCount, len(body))
}

func TestMainHandlerWhenUnsupportedCity(t *testing.T) {
	expectedError := "wrong city value"
	expectedCode := 400
	req, err := http.NewRequest("GET", "/cafe", nil)
	if err != nil {
		require.NoError(t, err)
	}

	q := req.URL.Query()
	q.Add("count", "4")
	q.Add("city", "ufa")
	req.URL.RawQuery = q.Encode()

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)
	resCode := responseRecorder.Code
	require.Equal(t, expectedCode, resCode)
	resp := responseRecorder.Body.String()
	require.Equal(t, resp, expectedError)
}

func TestMainHandlerWhenQueryIsCorrect(t *testing.T) {
	expectedCode := 200
	req, err := http.NewRequest("GET", "/cafe", nil)
	if err != nil {
		require.NoError(t, err)
	}

	q := req.URL.Query()
	q.Add("count", "1")
	q.Add("city", "moscow")
	req.URL.RawQuery = q.Encode()

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)
	respCode := responseRecorder.Code
	require.Equal(t, expectedCode, respCode)
	require.NotEmpty(t, responseRecorder.Body)
}
