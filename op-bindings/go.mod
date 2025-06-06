module github.com/ethereum-optimism/optimism/op-bindings

go 1.18

require github.com/ethereum/go-ethereum v1.10.17

require (
	github.com/btcsuite/btcd/btcec/v2 v2.2.0 // indirect
	github.com/go-stack/stack v1.8.1 // indirect
	github.com/golang-jwt/jwt/v4 v4.3.0 // indirect
	github.com/google/uuid v1.3.0 // indirect
	golang.org/x/crypto v0.0.0-20220307211146-efcb8507fb70 // indirect
	golang.org/x/sys v0.0.0-20220310020820-b874c991c1a5 // indirect
)

replace github.com/ethereum/go-ethereum v1.10.17 => github.com/ethereum-optimism/reference-optimistic-geth v0.0.0-20220602230953-dd2e24b3359f
