module github.com/ethereum-optimism/optimism/batch-submitter

go 1.18

replace github.com/ethereum-optimism/optimism/bss-core v0.0.0 => ../bss-core

replace github.com/ethereum-optimism/optimism/l2geth v0.0.0 => ../l2geth

require (
	github.com/ethereum-optimism/optimism/bss-core v0.0.0
	github.com/ethereum-optimism/optimism/l2geth v0.0.0
	github.com/ethereum/go-ethereum v1.10.17
	github.com/getsentry/sentry-go v0.12.0
	github.com/prometheus/client_golang v1.11.0
	github.com/stretchr/testify v1.7.0
	github.com/urfave/cli v1.22.5
)

require (
	github.com/btcsuite/btcd v0.22.1 // indirect
	github.com/btcsuite/btcd/btcec/v2 v2.1.2 // indirect
	github.com/go-ole/go-ole v1.2.6 // indirect
	github.com/go-stack/stack v1.8.1 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/golang/snappy v0.0.4 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/prometheus/client_model v0.2.0 // indirect
	github.com/prometheus/common v0.30.0 // indirect
	github.com/prometheus/procfs v0.7.3 // indirect
	github.com/prometheus/tsdb v0.10.0 // indirect
	golang.org/x/crypto v0.0.0-20220307211146-efcb8507fb70 // indirect
	golang.org/x/sys v0.0.0-20220310020820-b874c991c1a5 // indirect
	google.golang.org/protobuf v1.27.1 // indirect
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
)
