package goprowifi

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

type testDoer struct {
	response     string
	responseCode int
	http.Header
}

func (nd testDoer) Do(*http.Request) (*http.Response, error) {
	return &http.Response{
		Body:       ioutil.NopCloser(bytes.NewReader([]byte(nd.response))),
		StatusCode: nd.responseCode,
		Header:     nd.Header,
	}, nil
}

func CreateTestRestClient(d Doer) *Client {
	client := NewClient(true)
	client.myDoer = d
	return client
}
