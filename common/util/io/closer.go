package io

import log "github.com/sirupsen/logrus"

var (
	NopCloser = NewCloser(func() error {
		return nil
	})
)

type Closer interface {
	Close() error
}

type inlineCloser struct {
	close func() error
}

func (c *inlineCloser) Close() error {
	return c.close()
}

func NewCloser(close func() error) Closer {
	return &inlineCloser{close: close}
}

func Close(c Closer) {
	if err := c.Close(); err != nil {
		log.Warnf("failed to close %v: %v", c, err)
	}
}
