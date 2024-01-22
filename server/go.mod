module go.datalift.io/datalift/server

go 1.21.6

replace go.datalift.io/datalift/common => ../common

require (
	github.com/envoyproxy/protoc-gen-validate v1.0.2
	github.com/gobwas/glob v0.2.3
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.19.0
	github.com/joho/godotenv v1.5.1
	github.com/stretchr/testify v1.8.4
	github.com/uber-go/tally/v4 v4.1.10
	go.datalift.io/datalift/common v0.0.0-00010101000000-000000000000
	go.uber.org/zap v1.26.0
	google.golang.org/grpc v1.60.1
	google.golang.org/protobuf v1.32.0
	gopkg.in/yaml.v3 v3.0.1
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/rogpeppe/go-internal v1.12.0 // indirect
	github.com/twmb/murmur3 v1.1.5 // indirect
	go.uber.org/atomic v1.7.0 // indirect
	go.uber.org/multierr v1.10.0 // indirect
	golang.org/x/net v0.19.0 // indirect
	golang.org/x/sys v0.15.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	google.golang.org/genproto v0.0.0-20240116215550-a9fa1716bcac // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20240102182953-50ed04b92917 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240102182953-50ed04b92917 // indirect
)
