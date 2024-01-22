package gateway

import (
	"net/http"
	"path/filepath"

	"github.com/joho/godotenv"
	"go.uber.org/zap"

	gatewayv1 "go.datalift.io/datalift/server/config/gateway/v1"
	"go.datalift.io/datalift/server/endpoint"
	"go.datalift.io/datalift/server/middleware"
	"go.datalift.io/datalift/server/service"
)

type ComponentFactory struct {
	Services   service.Factory
	Middleware middleware.Factory
	Endpoints  endpoint.Factory
}

func loadEnv(f *Flags) {
	// Order is important as godotenv will NOT overwrite existing environment variables.
	envFiles := f.EnvFiles

	for _, filename := range envFiles {
		// Use a temporary logger to parse the environment files
		tmpLogger := newTmpLogger().With(zap.String("file", filename))

		p, err := filepath.Abs(filename)
		if err != nil {
			tmpLogger.Fatal("parsing .env file failed", zap.Error(err))
		}
		// Ignore lint below as it is ok to to ignore dotenv loads as not all env files are guaranteed
		// to be present.
		// nolint
		err = godotenv.Load(p)
		if err == nil {
			tmpLogger.Info("successfully loaded environment variables")
		}
	}
}

func Run(f *Flags, cf *ComponentFactory, assets http.FileSystem) {
	loadEnv(f)
	cfg := MustReadOrValidateConfig(f)
	RunWithConfig(f, cfg, cf, assets)
}

func RunWithConfig(f *Flags, cfg *gatewayv1.Config, cf *ComponentFactory, assets http.FileSystem) {
}
