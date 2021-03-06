package escrow

import (
	"context"
	"testing"
	"time"

	"github.com/iov-one/weave"
)

func TestIsExpired(t *testing.T) {
	now := time.Now()
	ctx := weave.WithBlockTime(context.Background(), now)

	future := now.Add(5 * time.Minute)
	if isExpired(ctx, future) {
		t.Error("future is expired")
	}

	past := now.Add(-5 * time.Minute)
	if !isExpired(ctx, past) {
		t.Error("past is not expired")
	}
}

func TestIsExpiredRequiresBlockTime(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Fatal("wanted a panic")
		}
	}()

	// Calling isExpected with a context without a block height
	// attached is expected to panic.
	isExpired(context.Background(), time.Now())
}
