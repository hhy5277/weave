package cash

import (
	"fmt"
	"testing"

	"github.com/iov-one/weave"
	coin "github.com/iov-one/weave/coin"
	"github.com/iov-one/weave/errors"
	"github.com/iov-one/weave/orm"
	"github.com/iov-one/weave/store"
	"github.com/iov-one/weave/weavetest"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type checkErr func(error) bool

func noErr(err error) bool { return err == nil }

func TestSend(t *testing.T) {
	foo := coin.NewCoin(100, 0, "FOO")
	some := coin.NewCoin(300, 0, "SOME")

	perm := weave.NewCondition("sig", "ed25519", []byte{1, 2, 3})
	perm2 := weave.NewCondition("sig", "ed25519", []byte{4, 5, 6})

	cases := []struct {
		signers       []weave.Condition
		initState     []orm.Object
		msg           weave.Msg
		expectCheck   checkErr
		expectDeliver checkErr
	}{
		0: {nil, nil, nil, errors.ErrInvalidMsg.Is, errors.ErrInvalidMsg.Is},
		1: {nil, nil, new(SendMsg), errors.ErrInvalidAmount.Is, errors.ErrInvalidAmount.Is},
		2: {nil, nil, &SendMsg{Amount: &foo}, errors.ErrInvalidInput.Is, errors.ErrInvalidInput.Is},
		3: {
			nil,
			nil,
			&SendMsg{Amount: &foo, Src: perm.Address(), Dest: perm2.Address()},
			errors.ErrUnauthorized.Is,
			errors.ErrUnauthorized.Is,
		},
		// sender has no account
		4: {
			[]weave.Condition{perm},
			nil,
			&SendMsg{Amount: &foo, Src: perm.Address(), Dest: perm2.Address()},
			noErr, // we don't check funds
			errors.ErrEmpty.Is,
		},
		// sender too poor
		5: {
			[]weave.Condition{perm},
			[]orm.Object{must(WalletWith(perm.Address(), &some))},
			&SendMsg{Amount: &foo, Src: perm.Address(), Dest: perm2.Address()},
			noErr, // we don't check funds
			errors.ErrInsufficientAmount.Is,
		},
		// sender got cash
		6: {
			[]weave.Condition{perm},
			[]orm.Object{must(WalletWith(perm.Address(), &foo))},
			&SendMsg{Amount: &foo, Src: perm.Address(), Dest: perm2.Address()},
			noErr,
			noErr,
		},
	}

	for i, tc := range cases {
		t.Run(fmt.Sprintf("case-%d", i), func(t *testing.T) {
			auth := &weavetest.Auth{Signers: tc.signers}
			controller := NewController(NewBucket())
			h := NewSendHandler(auth, controller)

			kv := store.MemStore()
			bucket := NewBucket()
			for _, wallet := range tc.initState {
				err := bucket.Save(kv, wallet)
				require.NoError(t, err)
			}

			tx := &weavetest.Tx{Msg: tc.msg}

			_, err := h.Check(nil, kv, tx)
			assert.True(t, tc.expectCheck(err), "%+v", err)
			_, err = h.Deliver(nil, kv, tx)
			assert.True(t, tc.expectDeliver(err), "%+v", err)
		})
	}
}
