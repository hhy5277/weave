package app

import (
	abci "github.com/tendermint/abci/types"

	"github.com/confio/weave"
	"github.com/confio/weave/errors"
)

// TODO: what about the init state stuff.... where does that go????

// BaseApp adds DeliverTx, CheckTx, and BeginBlock
// handlers to the storage and query functionality of StoreApp
type BaseApp struct {
	*StoreApp
	decoder weave.TxDecoder
	handler weave.Handler
	ticker  weave.Ticker
}

var _ abci.Application = BaseApp{}

// NewBaseApp constructs a basic abci application
func NewBaseApp(store *StoreApp, decoder weave.TxDecoder,
	handler weave.Handler, ticker weave.Ticker) BaseApp {

	return BaseApp{
		StoreApp: store,
		decoder:  decoder,
		handler:  handler,
		ticker:   ticker,
	}
}

// DeliverTx - ABCI - dispatches to the handler
func (b BaseApp) DeliverTx(txBytes []byte) abci.ResponseDeliverTx {
	tx, err := b.loadTx(txBytes)
	if err != nil {
		return weave.DeliverTxError(err)
	}

	// ignore error here, allow it to be logged
	ctx := weave.WithLogInfo(b.BlockContext(),
		"call", "deliver_tx",
		"path", weave.GetPath(tx))

	res, err := b.handler.Deliver(ctx, b.DeliverStore(), tx)
	return weave.DeliverOrError(res, err)
}

// CheckTx - ABCI - dispatches to the handler
func (b BaseApp) CheckTx(txBytes []byte) abci.ResponseCheckTx {
	tx, err := b.loadTx(txBytes)
	if err != nil {
		return weave.CheckTxError(err)
	}

	ctx := weave.WithLogInfo(b.BlockContext(),
		"call", "check_tx",
		"path", weave.GetPath(tx))

	res, err := b.handler.Check(ctx, b.CheckStore(), tx)
	return weave.CheckOrError(res, err)
}

// BeginBlock - ABCI
func (b BaseApp) BeginBlock(req abci.RequestBeginBlock) (
	res abci.ResponseBeginBlock) {

	// default: set the context properly
	b.StoreApp.BeginBlock(req)

	// call the ticker, if set
	if b.ticker != nil {
		// start := time.Now()
		// Add info to the logger
		ctx := weave.WithLogInfo(b.BlockContext(), "call", "begin_block")
		res, err := b.ticker.Tick(ctx, b.DeliverStore())
		// logDuration(ctx, start, "Ticker", err, false)
		if err != nil {
			panic(err)
		}
		b.StoreApp.AddValChange(res.Diff)
	}
	return
}

// loadTx calls the decoder, and capture any panics
func (b BaseApp) loadTx(txBytes []byte) (tx weave.Tx, err error) {
	defer errors.Recover(&err)
	tx, err = b.decoder(txBytes)
	return
}
