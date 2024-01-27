package tls

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBestEffortSystemCertPool(t *testing.T) {
	pool := BestEffortSystemCertPool()
	assert.NotNil(t, pool)
}
