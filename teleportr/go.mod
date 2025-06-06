module github.com/ethereum-optimism/optimism/teleportr

go 1.18

replace github.com/ethereum-optimism/optimism/bss-core v0.0.0 => ../bss-core

require (
	github.com/ethereum-optimism/optimism/bss-core v0.0.0
	github.com/ethereum/go-ethereum v1.10.17
	github.com/google/uuid v1.3.0
	github.com/gorilla/mux v1.8.0
	github.com/lib/pq v1.10.4
	github.com/prometheus/client_golang v1.11.0
	github.com/rs/cors v1.8.2
	github.com/stretchr/testify v1.7.0
	github.com/urfave/cli v1.22.5
)

require (
	github.com/VictoriaMetrics/fastcache v1.9.0 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/btcsuite/btcd/btcec/v2 v2.1.2 // indirect
	golang.org/x/crypto v0.0.0-20220307211146-efcb8507fb70 // indirect
	golang.org/x/sys v0.0.0-20220310020820-b874c991c1a5 // indirect
	google.golang.org/protobuf v1.27.1 // indirect
	gopkg.in/natefinch/npipe.v2 v2.0.0-20160621034901-c1b8fa8bdcce // indirect
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
)
