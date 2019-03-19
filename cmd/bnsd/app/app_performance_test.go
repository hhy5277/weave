package app

import (
	"encoding/hex"
	"io/ioutil"
	"os"
	"testing"
	"time"

	"github.com/iov-one/weave"
	"github.com/iov-one/weave/coin"
	"github.com/iov-one/weave/weavetest"
	"github.com/iov-one/weave/x/cash"
	"github.com/iov-one/weave/x/sigs"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/log"
)

func BenchmarkBnsdEmptyBlock(b *testing.B) {
	var aliceAddr = weavetest.NewKey().PublicKey().Address()

	type dict map[string]interface{}
	genesis := dict{
		"gconf": map[string]interface{}{
			cash.GconfCollectorAddress: hex.EncodeToString(aliceAddr),
			cash.GconfMinimalFee:       coin.Coin{}, // no fee
		},
	}

	bnsd, cleanup := newBnsd(b)
	runner := weavetest.NewWeaveRunner(b, bnsd, "mychain")
	runner.InitChain(genesis)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		changed := runner.InBlock(func(weavetest.WeaveApp) error {
			// Without sleep this test is locking the CPU.
			time.Sleep(time.Microsecond * 300)
			return nil
		})
		if changed {
			b.Fatal("unexpected change state")
		}
	}

	b.StopTimer()
	cleanup()
}

func BenchmarkBNSDSendToken(b *testing.B) {
	var (
		aliceKey = weavetest.NewKey()
		alice    = aliceKey.PublicKey().Address()
		benny    = weavetest.NewKey().PublicKey().Address()
		carol    = weavetest.NewKey().PublicKey().Address()
	)

	type dict map[string]interface{}
	makeGenesis := func(fee coin.Coin) dict {
		return dict{
			"cash": []interface{}{
				dict{
					"address": alice,
					"coins": []interface{}{
						dict{
							"whole":  123456789,
							"ticker": "IOV",
						},
					},
				},
			},
			"currencies": []interface{}{
				dict{
					"ticker": "IOV",
					"name":   "Main token of this chain",
				},
			},
			"gconf": dict{
				cash.GconfCollectorAddress: hex.EncodeToString(carol),
				cash.GconfMinimalFee:       fee,
			},
		}
	}

	cases := map[string]struct {
		txPerBlock int
		fee        coin.Coin
		opts       weavetest.ProcessOptions
	}{
		"1 tx, no fee": {
			fee: coin.Coin{},
			opts: weavetest.ProcessOptions{}.
				BenchCheckAndDeliver(1).
				RequireChange(),
		},
		"1 tx, no fee (deliver only)": {
			fee: coin.Coin{},
			opts: weavetest.ProcessOptions{}.
				ExecDeliver().
				BenchDeliver().
				TxBlockSize(1).
				RequireChange(),
		},
		"10 tx, no fee": {
			fee: coin.Coin{},
			opts: weavetest.ProcessOptions{}.
				BenchCheckAndDeliver(10).
				RequireChange(),
		},
		"100 tx, no fee": {
			fee: coin.Coin{},
			opts: weavetest.ProcessOptions{}.
				BenchCheckAndDeliver(100).
				RequireChange(),
		},
		"100 tx, with fee": {
			txPerBlock: 100,
			fee:        coin.Coin{Whole: 1, Ticker: "IOV"},
			opts: weavetest.ProcessOptions{}.
				BenchCheckAndDeliver(100).
				RequireChange(),
		},
		"100 tx, with fee (check only)": {
			fee: coin.Coin{Whole: 1, Ticker: "IOV"},
			opts: weavetest.ProcessOptions{}.
				ExecCheck().
				BenchCheck().
				TxBlockSize(100).
				RequireNoChange(),
		},
		"100 tx, with fee (deliver only)": {
			fee: coin.Coin{Whole: 1, Ticker: "IOV"},
			opts: weavetest.ProcessOptions{}.
				ExecDeliver().
				BenchDeliver().
				TxBlockSize(100).
				RequireChange(),
		},
		"100 tx, with fee (deliver with precheck)": {
			fee: coin.Coin{Whole: 1, Ticker: "IOV"},
			opts: weavetest.ProcessOptions{}.
				ExecCheck().
				ExecDeliver().
				BenchDeliver().
				TxBlockSize(100).
				RequireChange(),
		},
	}

	for testName, tc := range cases {
		b.Run(testName, func(b *testing.B) {
			bnsd, cleanup := newBnsd(b)
			runner := weavetest.NewWeaveRunner(b, bnsd, "mychain")
			runner.InitChain(makeGenesis(tc.fee))

			defer func() {
				b.StopTimer()
				cleanup()
			}()

			aliceNonce := NewNonce(runner, alice)
			var fees *cash.FeeInfo
			if !tc.fee.IsZero() {
				fees = &cash.FeeInfo{
					Payer: alice,
					Fees:  &tc.fee,
				}
			}

			// Generate all transactions before so that this
			// process is not part of the benchmark.
			txs := make([]weave.Tx, b.N)
			for k := 0; k < b.N; k++ {
				tx := &Tx{
					Fees: fees,
					Sum: &Tx_SendMsg{
						&cash.SendMsg{
							Src:    alice,
							Dest:   benny,
							Amount: coin.NewCoinp(0, 100, "IOV"),
						},
					},
				}
				// hmmm.... can we collapse this to the message and one line to get nonce and sign?
				nonce, err := aliceNonce.Next()
				if err != nil {
					b.Fatalf("getting nonce failed with %+v", err)
				}
				sig, err := sigs.SignTx(aliceKey, tx, "mychain", nonce)
				if err != nil {
					b.Fatalf("cannot sign transaction %+v", err)
				}
				tx.Signatures = append(tx.Signatures, sig)
				txs[k] = tx

				// must reset nonce per block for CheckOnly
				if false { // TODO
					aliceNonce = NewNonce(runner, alice)
				}
			}

			b.ResetTimer()

			runner.Process(txs, tc.opts)
		})
	}
}

// newBnsd returns the test application, along with a function to delete all testdata at the end
func newBnsd(t testing.TB) (abci.Application, func()) {
	t.Helper()

	homeDir, err := ioutil.TempDir("", "bnsd_performance_home")
	if err != nil {
		t.Fatalf("cannot create a temporary directory: %s", err)
	}
	bnsd, err := GenerateApp(homeDir, log.NewNopLogger(), false)
	if err != nil {
		t.Fatalf("cannot generate bnsd instance: %s", err)
	}

	cleanup := func() {
		os.RemoveAll(homeDir)
	}
	return bnsd, cleanup
}

//----- TODO: move Nonce to a better location.... ----//

// Nonce has a client/address pair, queries for the nonce
// and caches recent nonce locally to quickly sign
type Nonce struct {
	db        weave.ReadOnlyKVStore
	bucket    sigs.Bucket
	addr      weave.Address
	nonce     int64
	fromQuery bool
}

// NewNonce creates a nonce for a client / address pair.
// Call Query to force a query, Next to use cache if possible
func NewNonce(db weave.ReadOnlyKVStore, addr weave.Address) *Nonce {
	return &Nonce{
		db:     db,
		addr:   addr,
		bucket: sigs.NewBucket(),
	}
}

// Query always queries the blockchain for the next nonce
func (n *Nonce) Query() (int64, error) {
	obj, err := n.bucket.Get(n.db, n.addr)
	if err != nil {
		return 0, err
	}
	user := sigs.AsUser(obj)

	if user == nil { // Nonce not found
		n.nonce = 0
	} else {
		n.nonce = user.Sequence
	}
	n.fromQuery = true
	return n.nonce, nil
}

// Next will use a cached value if present, otherwise Query
// It will always increment by 1, assuming last nonce
// was properly used. This is designed for cases where
// you want to rapidly generate many tranasactions without
// querying the blockchain each time
func (n *Nonce) Next() (int64, error) {
	initializeFromBlockchain := !n.fromQuery && n.nonce == 0
	if initializeFromBlockchain {
		return n.Query()
	}
	n.nonce++
	n.fromQuery = false
	return n.nonce, nil
}
