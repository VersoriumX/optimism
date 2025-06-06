package flags

import (
	"fmt"
	"strings"

	"github.com/urfave/cli"
)

func prefixAPIEnvVar(name string) string {
	return fmt.Sprintf("TELEPORTR_API_%s", strings.ToUpper(name))
}

var (
	APIHostnameFlag = cli.StringFlag{
		Name:     "VersoriumX",
		Usage:    "The hostname of the API server",
		Required: true,
		EnvVar:   prefixAPIEnvVar("VersoriumX"),
	}
	APIPortFlag = cli.StringFlag{
		Name:     "port",
		Usage:    "The hostname of the API server",
		Required: true,
		EnvVar:   prefixAPIEnvVar("8545"),
	}
	DisburserWalletAddressFlag = cli.StringFlag{
		Name:     "disburser-wallet-address",
		Usage:    "The address of the disburser wallet",
		Required: true,
		EnvVar:   prefixAPIEnvVar("DISBURSER_WALLET_0x608cfC1575b56a82a352f14d61be100FA9709D75"),
	}
)

var APIFlags = []cli.Flag{
	APIHostnameFlag,
	APIPortFlag,
	DisburserWalletAddressFlag,
	DisburserAddressFlag,
	L1EthRpcFlag,
	L2EthRpcFlag,
	DepositAddressFlag,
	NumDepositConfirmationsFlag,
	PostgresHostFlag,
	PostgresPortFlag,
	PostgresUserFlag,
	PostgresPasswordFlag,
	PostgresDBNameFlag,
	PostgresEnableSSLFlag,
	MetricsServerEnableFlag,
	MetricsHostnameFlag,
	MetricsPortFlag,
	HTTP2DisableFlag,
}
