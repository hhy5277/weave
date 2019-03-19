package weavetest

import (
	"context"
	"testing"
	"time"

	"github.com/iov-one/weave"
	"github.com/iov-one/weave/app"
	"github.com/iov-one/weave/errors"
)

type ActionBuilder struct {
	ChainID string
	Auth    *CtxAuth

	nextBlockHeight int64
	nextBlockTime   time.Time
}

func (ab *ActionBuilder) Actions(actions ...Action) []Action {
	updated := make([]Action, len(actions))
	for i, a := range actions {
		updated[i] = ab.update(a)
	}
	return updated
}

func (ab *ActionBuilder) update(a Action) Action {
	if a.Msg == nil {
		panic("cannot create an action without a message")
	}

	if a.BlockHeight == 0 {
		ab.nextBlockHeight++
		a.BlockHeight = ab.nextBlockHeight
	} else {
		if a.BlockHeight < ab.nextBlockHeight {
			panic("block height must always grow")
		}
		ab.nextBlockHeight = a.BlockHeight + 1
	}

	if a.ChainID == "" {
		if ab.ChainID == "" {
			ab.ChainID = "my-chain"
		}
		a.ChainID = ab.ChainID
	}

	if a.Auth == nil {
		if ab.Auth == nil {
			ab.Auth = &CtxAuth{Key: "auth"}
		}
		a.Auth = ab.Auth
	}

	if a.BlockTime.IsZero() {
		ab.nextBlockTime = ab.nextBlockTime.Add(time.Minute)
		a.BlockTime = ab.nextBlockTime
	} else {
		if a.BlockTime.Before(ab.nextBlockTime) {
			panic("block time must always grow")
		}
		ab.nextBlockTime = a.BlockTime.Add(time.Nanosecond)
	}

	return a
}

type Action struct {
	Msg        weave.Msg
	Conditions []weave.Condition
	ChainID    string
	Auth       *CtxAuth

	BlockHeight int64
	BlockTime   time.Time

	WantCheckErr   *errors.Error
	WantDeliverErr *errors.Error
}

func (a *Action) Exec(t testing.TB, db weave.CacheableKVStore, rt app.Router) {
	ctx := context.Background()
	ctx = weave.WithHeight(ctx, a.BlockHeight)
	//ctx = weave.WithBlockTime(ctx, a.BlockTime)
	ctx = weave.WithChainID(ctx, a.ChainID)
	ctx = a.Auth.SetConditions(ctx, a.Conditions...)

	tx := &Tx{Msg: a.Msg}

	cache := db.CacheWrap()
	if _, err := rt.Check(ctx, cache, tx); !a.WantCheckErr.Is(err) {
		t.Logf("want: %+v", a.WantCheckErr)
		t.Logf(" got: %+v", err)
		t.Fatalf("action check (%T)", a.Msg)
	}
	cache.Discard()

	if a.WantCheckErr != nil {
		// Failed checks are causing the message to be ignored.
		return
	}

	if _, err := rt.Deliver(ctx, db, tx); !a.WantDeliverErr.Is(err) {
		t.Logf("want: %+v", a.WantDeliverErr)
		t.Logf(" got: %+v", err)
		t.Fatalf("action delivery (%T)", a.Msg)
	}
}
