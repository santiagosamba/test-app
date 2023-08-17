package repository

import (
	"bytes"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"net/http"
	"test-app/defines"
	"test-app/utils"
	"testing"
)

func TestCharacterNotFound(t *testing.T) {
	httpClientMock := &utils.HTTPClientMock{
		Response: &http.Response{
			StatusCode: http.StatusNotFound,
			Body:       ioutil.NopCloser(bytes.NewReader([]byte(`{"data":{}}`))),
			Header:     make(http.Header),
		},
	}
	repository := NewCharacterRepository(httpClientMock)

	resp, respErr := repository.GetCharacter(defines.TestID)

	require.Nil(t, resp)
	require.NotNil(t, respErr)
}

func TestCharacterOk(t *testing.T) {
	httpClientMock := &utils.HTTPClientMock{
		Response: &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewReader([]byte(`{"data":{"name":"Mickey Mouse","imageUrl":"https://example.com/mickey.jpg"}}`))),
			Header:     make(http.Header),
		},
	}
	repository := NewCharacterRepository(httpClientMock)

	resp, respErr := repository.GetCharacter(defines.TestID)

	require.Nil(t, respErr)
	require.NotNil(t, resp)
	require.Equal(t, resp.Data.Name, "Mickey Mouse")
}
