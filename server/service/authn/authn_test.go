package authn

import (
	"testing"

	"github.com/stretchr/testify/assert"

	apimock "go.datalift.io/admiral/server/mock/api"
)

func TestNotImpl(t *testing.T) {
	a := apimock.AnyFromYAML(`
"@type": types.google.com/admiral.config.service.authn.v1.Config
session_secret: my_session_secret
`)
	svc, err := New(a, nil, nil)
	assert.Error(t, err)
	assert.Nil(t, svc)
}
