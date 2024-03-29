package stats

import (
	tallyprom "github.com/uber-go/tally/v4/prometheus"

	gatewayv1 "go.datalift.io/admiral/server/config/gateway/v1"
)

func NewPrometheusReporter(cfg *gatewayv1.Stats_PrometheusReporter) (tallyprom.Reporter, error) {
	promCfg := tallyprom.Configuration{
		HandlerPath: cfg.HandlerPath,
	}
	return promCfg.NewReporter(tallyprom.ConfigurationOptions{})
}
