package driver

import (
	"context"
	"math/big"

	"github.com/ethereum-optimism/optimism/op-node/eth"
	"github.com/ethereum-optimism/optimism/op-node/l1"
	"github.com/ethereum-optimism/optimism/op-node/l2"
	"github.com/ethereum-optimism/optimism/op-node/rollup"
	"github.com/ethereum-optimism/optimism/op-node/rollup/derive"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
)

type Driver struct {
	s *state
}

type Downloader interface {
	InfoByHash(ctx context.Context, hash common.Hash) (eth.L1Info, error)
	Fetch(ctx context.Context, blockHash common.Hash) (eth.L1Info, types.Transactions, types.Receipts, error)
}

type L1Chain interface {
	derive.L1Fetcher
	L1HeadBlockRef(context.Context) (eth.L1BlockRef, error)
}

type L2Chain interface {
	derive.Engine
	L2BlockRefByNumber(ctx context.Context, l2Num *big.Int) (eth.L2BlockRef, error)
	L2BlockRefByHash(ctx context.Context, l2Hash common.Hash) (eth.L2BlockRef, error)
}

type DerivationPipeline interface {
	Reset(ctx context.Context, l2SafeHead eth.L2BlockRef, unsafeL2Head eth.L2BlockRef) error
	Step(ctx context.Context) error
	AddUnsafePayload(payload *eth.ExecutionPayload)
	Finalized() eth.L2BlockRef
	SafeL2Head() eth.L2BlockRef
	UnsafeL2Head() eth.L2BlockRef
	CurrentL1() eth.L1BlockRef
}

type outputInterface interface {
	// createNewBlock builds a new block based on the L2 Head, L1 Origin, and the current mempool.
	createNewBlock(ctx context.Context, l2Head eth.L2BlockRef, l2SafeHead eth.BlockID, l2Finalized eth.BlockID, l1Origin eth.L1BlockRef) (eth.L2BlockRef, *eth.ExecutionPayload, error)
}

type Network interface {
	// PublishL2Payload is called by the driver whenever there is a new payload to publish, synchronously with the driver main loop.
	PublishL2Payload(ctx context.Context, payload *eth.ExecutionPayload) error
}

func NewDriver(cfg *rollup.Config, l2 *l2.Source, l1 *l1.Source, network Network, log log.Logger, snapshotLog log.Logger, sequencer bool) *Driver {
	output := &outputImpl{
		Config: cfg,
		dl:     l1,
		l2:     l2,
		log:    log,
	}
	derivationPipeline := derive.NewDerivationPipeline(log, cfg, l1, l2)
	return &Driver{
		s: NewState(log, snapshotLog, cfg, l1, l2, output, derivationPipeline, network, sequencer),
	}
}

func (d *Driver) Emitter() *derive.ChannelEmitter {
	return d.s.Emitter()
}

func (d *Driver) OnL1Head(ctx context.Context, head eth.L1BlockRef) error {
	return d.s.OnL1Head(ctx, head)
}

func (d *Driver) OnUnsafeL2Payload(ctx context.Context, payload *eth.ExecutionPayload) error {
	return d.s.OnUnsafeL2Payload(ctx, payload)
}

func (d *Driver) Start(ctx context.Context) error {
	return d.s.Start(ctx)
}
func (d *Driver) Close() error {
	return d.s.Close()
}
