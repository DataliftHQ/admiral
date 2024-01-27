package http

import (
	"log"
	"net/http"
	"net/http/httputil"
)

type DebugTransport struct {
	RoundTripper http.RoundTripper
}

func (d DebugTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	reqDump, err := httputil.DumpRequest(req, true)
	if err != nil {
		return nil, err
	}
	log.Printf("%s", reqDump)

	resp, err := d.RoundTripper.RoundTrip(req)
	if err != nil {
		return nil, err
	}

	respDump, err := httputil.DumpResponse(resp, true)
	if err != nil {
		_ = resp.Body.Close()
		return nil, err
	}
	log.Printf("%s", respDump)
	return resp, nil
}

type TransportWithHeader struct {
	RoundTripper http.RoundTripper
	Header       http.Header
}

func (rt *TransportWithHeader) RoundTrip(r *http.Request) (*http.Response, error) {
	if rt.Header != nil {
		headers := rt.Header.Clone()
		for k, vs := range r.Header {
			for _, v := range vs {
				headers.Add(k, v)
			}
		}
		r.Header = headers
	}
	return rt.RoundTripper.RoundTrip(r)
}
