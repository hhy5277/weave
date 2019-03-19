package weavetest

import (
	"testing"

	"github.com/iov-one/weave"
	"github.com/iov-one/weave/errors"
)

func (w *WeaveRunner) Process(txs []weave.Tx, opts ProcessOptions) {
	// For benchmark we want to control the measurement time.
	b, isBench := w.t.(*testing.B)

	blocks := splitTxs(txs, opts.txBlockSize)
	for _, txs := range blocks {
		changed := w.InBlock(func(wapp WeaveApp) error {
			if opts.execCheck {
				if isBench && !opts.benchCheck {
					b.StopTimer()
				}
				for i, tx := range txs {
					if err := wapp.CheckTx(tx); err != nil {
						return errors.Wrapf(err, "check transaction %d", i)
					}
				}
				if isBench && !opts.benchCheck {
					b.StartTimer()
				}
			}

			if opts.execDeliver {
				if isBench && !opts.benchDeliver {
					b.StopTimer()
				}
				for i, tx := range txs {
					if err := wapp.DeliverTx(tx); err != nil {
						return errors.Wrapf(err, "deliver transaction %d", i)
					}
				}
				if isBench && !opts.benchDeliver {
					b.StartTimer()
				}
			}

			return nil
		})

		if opts.mustChange != nil && *opts.mustChange != changed {
			if *opts.mustChange {
				w.t.Fatal("require state change")
			} else {
				w.t.Fatal("unexpected state change")
			}
		}
	}
}

// splitTxs will break one slice of transactions into many slices, one per
// block. It will fill up to txPerBlockx txs in each block The last block may
// have less, if there is not enough for a full block.
func splitTxs(txs []weave.Tx, txPerBlock uint) [][]weave.Tx {
	if txPerBlock == 0 {
		return [][]weave.Tx{txs}
	}

	blocks := numBlocks(len(txs), int(txPerBlock))
	res := make([][]weave.Tx, blocks)

	// Full chunks for all but the last block.
	for i := 0; i < blocks-1; i++ {
		res[i], txs = txs[:txPerBlock], txs[txPerBlock:]
	}

	// Remainder in the last block.
	res[blocks-1] = txs

	return res
}

// numBlocks returns total number of blocks for benchmarks that split b.N
// into many smaller blocks
func numBlocks(totalTx, txPerBlock int) int {
	runs := totalTx / txPerBlock
	if totalTx%txPerBlock > 0 {
		return runs + 1
	}
	return runs
}

type ProcessOptions struct {
	execCheck       bool
	benchCheck      bool
	execDeliver     bool
	benchDeliver    bool
	mustChange      *bool
	requireNoChange bool
	txBlockSize     uint
}

func (o ProcessOptions) ExecCheck() ProcessOptions {
	o.execCheck = true
	return o
}

func (o ProcessOptions) BenchCheck() ProcessOptions {
	o.benchCheck = true
	return o
}

func (o ProcessOptions) ExecDeliver() ProcessOptions {
	o.execDeliver = true
	return o
}

func (o ProcessOptions) BenchDeliver() ProcessOptions {
	o.benchDeliver = true
	return o
}

func (o ProcessOptions) RequireChange() ProcessOptions {
	ok := true
	o.mustChange = &ok
	return o
}

func (o ProcessOptions) RequireNoChange() ProcessOptions {
	ok := false
	o.mustChange = &ok
	return o
}

func (o ProcessOptions) TxBlockSize(blockSize uint) ProcessOptions {
	o.txBlockSize = blockSize
	return o
}

func (o ProcessOptions) BenchCheckAndDeliver(blockSize uint) ProcessOptions {
	return o.ExecCheck().
		BenchCheck().
		ExecDeliver().
		BenchDeliver().
		TxBlockSize(blockSize)
}
