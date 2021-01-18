package hooks

import (
	"net/http"
)

// Wrap wraps rt.
func Wrap(rt http.RoundTripper) *RoundTripper {
	return &RoundTripper{wrapped: rt}
}

// A RoundTripper wraps an http.RoundTripper.
type RoundTripper struct {
	wrapped http.RoundTripper

	BeforeRoundTrip func(req *http.Request) error
	AfterRoundTrip  func(resp *http.Response, err error) error
}

// RoundTrip implements http.RoundTripper.
func (rt *RoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	if rt.BeforeRoundTrip != nil {
		if err := rt.BeforeRoundTrip(req); err != nil {
			return nil, err
		}
	}
	resp, err := rt.wrapped.RoundTrip(req)
	if rt.AfterRoundTrip != nil {
		if err := rt.AfterRoundTrip(resp, err); err != nil {
			return nil, err
		}
	}
	return resp, err
}
