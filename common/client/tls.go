package client

import (
	"crypto/tls"
	"fmt"

	tls_util "go.datalift.io/admiral/common/util/tls"
)

func (c *client) tlsConfig() (*tls.Config, error) {
	var tlsConfig tls.Config
	if len(c.config.Settings.CertPEMData) > 0 {
		cp := tls_util.BestEffortSystemCertPool()
		if !cp.AppendCertsFromPEM(c.config.Settings.CertPEMData) {
			return nil, fmt.Errorf("credentials: failed to append certificates")
		}
		tlsConfig.RootCAs = cp
	}
	if c.config.Settings.ClientCert != nil {
		tlsConfig.Certificates = append(tlsConfig.Certificates, *c.config.Settings.ClientCert)
	}
	if c.config.Settings.Insecure {
		tlsConfig.InsecureSkipVerify = true
	}
	return &tlsConfig, nil
}
