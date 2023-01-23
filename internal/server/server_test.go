package server_test

import (
	"testing"

	"github.com/schufeli/minute-word/internal/server"
)

func TestApiKeyCheck(t *testing.T) {
	cases := []struct {
		desc             string
		actual, expected string
	}{
		{"TestApiKeysMatch", "1234", "1234"},
		{"TestApiKeysDontMatch", "1234", "5678"},
	}

	for _, tc := range cases {
		actual := server.ApiKeyCheck(tc.actual, tc.expected)
		if actual && tc.actual != tc.expected {
			t.Errorf("%s: mismatch got: actual: %s and expected: %s", tc.desc, tc.actual, tc.expected)
		}
	}
}
