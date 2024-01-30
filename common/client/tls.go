package client

import (
	"crypto/tls"
	"fmt"

	tls_util "go.datalift.io/admiral/common/util/tls"
)

func (c *client) tlsConfig() (*tls.Config, error) {
	var tlsConfig tls.Config
	if len(c.CertPEMData) > 0 {
		cp := tls_util.BestEffortSystemCertPool()
		if !cp.AppendCertsFromPEM(c.CertPEMData) {
			return nil, fmt.Errorf("credentials: failed to append certificates")
		}
		tlsConfig.RootCAs = cp
	}
	if c.ClientCert != nil {
		tlsConfig.Certificates = append(tlsConfig.Certificates, *c.ClientCert)
	}
	if c.Insecure {
		tlsConfig.InsecureSkipVerify = true
	}
	return &tlsConfig, nil
}
