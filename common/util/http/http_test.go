package http

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestRoundTripper struct{}

func (rt TestRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	resp := http.Response{}
	resp.Header = http.Header{}
	for k, vs := range req.Header {
		for _, v := range vs {
			resp.Header.Add(k, v)
		}
	}
	return &resp, nil
}

func TestTransportWithHeader(t *testing.T) {
	client := &http.Client{}
	req, _ := http.NewRequest(http.MethodGet, "/foo", nil)
	req.Header.Set("Bar", "req_1")
	req.Header.Set("Foo", "req_1")

	// No default headers.
	client.Transport = &TransportWithHeader{
		RoundTripper: &TestRoundTripper{},
	}
	resp, err := client.Do(req)
	assert.NoError(t, err)
	assert.Equal(t, resp.Header, http.Header{
		"Bar": []string{"req_1"},
		"Foo": []string{"req_1"},
	})

	// with default headers.
	client.Transport = &TransportWithHeader{
		RoundTripper: &TestRoundTripper{},
		Header: http.Header{
			"Foo": []string{"default_1", "default_2"},
		},
	}
	resp, err = client.Do(req)
	assert.NoError(t, err)
	assert.Equal(t, resp.Header, http.Header{
		"Bar": []string{"req_1"},
		"Foo": []string{"default_1", "default_2", "req_1"},
	})
}
