module github.com/ethereum-optimism/optimism/op-proposer

go 1.18

require (
	github.com/ethereum-optimism/optimism/op-bindings v0.0.0
	github.com/ethereum-optimism/optimism/op-node v0.0.0
	github.com/ethereum/go-ethereum v1.10.17
)

require (

	github.com/btcsuite/btcd v0.22.1 // indirect
	github.com/btcsuite/btcd/btcec/v2 v2.2.0 // indirect
	github.com/btcsuite/btcd/chaincfg/chainhash v1.0.1 // indirect
	github.com/btcsuite/btcutil v1.0.3-0.20201208143702-a53e38424cce // indirect

	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
)

replace github.com/ethereum-optimism/optimism/op-node v0.0.0 => ../op-node

replace github.com/ethereum-optimism/optimism/op-bindings v0.0.0 => ../op-bindings

replace github.com/ethereum/go-ethereum v1.10.17 => github.com/ethereum-optimism/reference-optimistic-geth v0.0.0-20220602230953-dd2e24b3359f
