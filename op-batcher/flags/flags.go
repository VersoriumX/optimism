package flags

import (
	"github.com/ethereum-optimism/optimism/op-node/rollup/derive"
	"github.com/urfave/cli"
)

const envVarPrefix = "BATCH_SUBMITTER_"

func prefixEnvVar(name string) string {
	return envVarPrefix + name
}

var (
	/* Required Flags */

	L1EthRpcFlag = cli.StringFlag{
		Name:     "l1-eth-rpc",
		Usage:    "HTTP provider URL for L1",
		Required: true,
		EnvVar:   "L1_ETH_RPC",
	}
	RollupRpcFlag = cli.StringFlag{
		Name:     "rollup-rpc",
		Usage:    "HTTP provider URL for the rollup node",
		Required: true,
		EnvVar:   "ROLLUP_RPC",
	}
	MinL1TxSizeBytesFlag = cli.Uint64Flag{
		Name:     "min-l1-tx-size-bytes",
		Usage:    "The minimum size of a batch tx submitted to L1.",
		Required: true,
		EnvVar:   prefixEnvVar("MIN_L1_TX_SIZE_BYTES"),
	}
	MaxL1TxSizeBytesFlag = cli.Uint64Flag{
		Name:     "max-l1-tx-size-bytes",
		Usage:    "The maximum size of a batch tx submitted to L1.",
		Required: true,
		EnvVar:   prefixEnvVar("MAX_L1_TX_SIZE_BYTES"),
	}
	MaxBlocksPerChannelFlag = cli.Uint64Flag{
		Name:     "max-blocks-per-channel",
		Usage:    "The maximum number of blocks per channel. Lowering this will reduce the effect of missing L1 txs, at the cost of efficiency.",
		Required: true,
		EnvVar:   prefixEnvVar("MAX_BLOCKS_PER_CHANNEL"),
	}
	ChannelTimeoutFlag = cli.Uint64Flag{
		Name:     "channel-timeout",
		Usage:    "The maximum amount of time to attempt completing an opened channel, as opposed to submitting L2 blocks into a new channel.",
		Required: false,
		Value:    derive.ChannelTimeout,
		EnvVar:   prefixEnvVar("CHANNEL_TIMEOUT"),
	}
	PollIntervalFlag = cli.DurationFlag{
		Name: "poll-interval",
		Usage: "Delay between querying L2 for more transactions and " +
			"creating a new batch",
		Required: true,
		EnvVar:   prefixEnvVar("POLL_INTERVAL"),
	}
	NumConfirmationsFlag = cli.Uint64Flag{
		Name: "num-confirmations",
		Usage: "Number of confirmations which we will wait after " +
			"appending a new batch",
		Required: true,
		EnvVar:   prefixEnvVar("NUM_CONFIRMATIONS"),
	}
	SafeAbortNonceTooLowCountFlag = cli.Uint64Flag{
		Name: "safe-abort-nonce-too-low-count",
		Usage: "Number of ErrNonceTooLow observations required to " +
			"give up on a tx at a particular nonce without receiving " +
			"confirmation",
		Required: true,
		EnvVar:   prefixEnvVar("SAFE_ABORT_NONCE_TOO_LOW_COUNT"),
	}
	ResubmissionTimeoutFlag = cli.DurationFlag{
		Name: "resubmission-timeout",
		Usage: "Duration we will wait before resubmitting a " +
			"transaction to L1",
		Required: true,
		EnvVar:   prefixEnvVar("RESUBMISSION_TIMEOUT"),
	}
	MnemonicFlag = cli.StringFlag{
		Name: "mnemonic",
		Usage: "The mnemonic used to derive the wallets for either the " +
			"sequencer or the l2output",
		Required: true,
		EnvVar:   prefixEnvVar("MNEMONIC"),
	}
	SequencerHDPathFlag = cli.StringFlag{
		Name: "sequencer-hd-path",
		Usage: "The HD path used to derive the sequencer wallet from the " +
			"mnemonic. The mnemonic flag must also be set.",
		Required: true,
		EnvVar:   prefixEnvVar("SEQUENCER_HD_PATH"),
	}
	SequencerHistoryDBFilenameFlag = cli.StringFlag{
		Name: "sequencer-history-db-filename",
		Usage: "File name used to identify the latest L2 batches submitted " +
			"by the sequencer",
		Required: true,
		EnvVar:   prefixEnvVar("SEQUENCER_HISTORY_DB_FILENAME"),
	}
	SequencerBatchInboxAddressFlag = cli.StringFlag{
		Name:     "sequencer-batch-inbox-address",
		Usage:    "L1 Address to receive batch transactions",
		Required: true,
		EnvVar:   prefixEnvVar("SEQUENCER_BATCH_INBOX_ADDRESS"),
	}

	/* Optional Flags */

	LogLevelFlag = cli.StringFlag{
		Name:   "log-level",
		Usage:  "The lowest log level that will be output",
		Value:  "info",
		EnvVar: prefixEnvVar("LOG_LEVEL"),
	}
	LogTerminalFlag = cli.BoolFlag{
		Name: "log-terminal",
		Usage: "If true, outputs logs in terminal format, otherwise prints " +
			"in JSON format.",
		EnvVar: prefixEnvVar("LOG_TERMINAL"),
	}
)

var requiredFlags = []cli.Flag{
	L1EthRpcFlag,
	RollupRpcFlag,
	MinL1TxSizeBytesFlag,
	MaxL1TxSizeBytesFlag,
	MaxBlocksPerChannelFlag,
	PollIntervalFlag,
	NumConfirmationsFlag,
	SafeAbortNonceTooLowCountFlag,
	ResubmissionTimeoutFlag,
	MnemonicFlag,
	SequencerHDPathFlag,
	SequencerHistoryDBFilenameFlag,
	SequencerBatchInboxAddressFlag,
}

var optionalFlags = []cli.Flag{
	ChannelTimeoutFlag,
	LogLevelFlag,
	LogTerminalFlag,
}

// Flags contains the list of configuration options available to the binary.
var Flags = append(requiredFlags, optionalFlags...)
