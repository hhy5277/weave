package multisig

import (
	"context"
	"testing"

	"github.com/iov-one/weave"
	"github.com/iov-one/weave/app"
	"github.com/iov-one/weave/errors"
	"github.com/iov-one/weave/store"
	"github.com/iov-one/weave/weavetest"
)

func TestCreateContractHandler(t *testing.T) {
	alice := weavetest.NewCondition().Address()
	bobby := weavetest.NewCondition().Address()
	cindy := weavetest.NewCondition().Address()

	cases := map[string]struct {
		Msg            weave.Msg
		WantCheckErr   *errors.Error
		WantDeliverErr *errors.Error
	}{
		"successfully create a contract": {
			Msg: &CreateContractMsg{
				Participants: []*Participant{
					{Power: 1, Signature: alice},
					{Power: 2, Signature: bobby},
					{Power: 3, Signature: cindy},
				},
				ActivationThreshold: 2,
				AdminThreshold:      3,
			},
		},
		"cannot create a contract without participants": {
			Msg: &CreateContractMsg{
				Participants:        []*Participant{},
				ActivationThreshold: 2,
				AdminThreshold:      3,
			},
			WantCheckErr: errors.ErrInvalidMsg,
		},
		"cannot create if activation threshold is too high": {
			Msg: &CreateContractMsg{
				Participants: []*Participant{
					{Power: 1, Signature: alice},
					{Power: 2, Signature: bobby},
					{Power: 3, Signature: cindy},
				},
				ActivationThreshold: 7, // higher than total
				AdminThreshold:      3,
			},
			WantCheckErr: errors.ErrInvalidMsg,
		},
		"cannot create if activation threshold is higher than admin threshold": {
			Msg: &CreateContractMsg{
				Participants: []*Participant{
					{Power: 2, Signature: alice},
					{Power: 2, Signature: bobby},
				},
				ActivationThreshold: 2,
				AdminThreshold:      1,
			},
			WantCheckErr: errors.ErrInvalidMsg,
		},
	}

	auth := &weavetest.Auth{
		Signer: weavetest.NewCondition(), // Any signer will do.
	}
	rt := app.NewRouter()
	RegisterRoutes(rt, auth)

	for testName, tc := range cases {
		t.Run(testName, func(t *testing.T) {
			db := store.MemStore()
			ctx := context.Background()
			tx := &weavetest.Tx{Msg: tc.Msg}

			cache := db.CacheWrap()
			if _, err := rt.Check(ctx, cache, tx); !tc.WantCheckErr.Is(err) {
				t.Logf("want: %+v", tc.WantCheckErr)
				t.Logf(" got: %+v", err)
				t.Fatalf("check (%T)", tc.Msg)
			}
			cache.Discard()
			if tc.WantCheckErr != nil {
				// Failed checks are causing the message to be ignored.
				return
			}

			if _, err := rt.Deliver(ctx, db, tx); !tc.WantDeliverErr.Is(err) {
				t.Logf("want: %+v", tc.WantDeliverErr)
				t.Logf(" got: %+v", err)
				t.Fatalf("delivery (%T)", tc.Msg)
			}
		})
	}
}

func TestUpdateContractHandler(t *testing.T) {
	aliceCond := weavetest.NewCondition()
	alice := aliceCond.Address()
	bobby := weavetest.NewCondition().Address()
	cindyCond := weavetest.NewCondition()
	cindy := cindyCond.Address()

	cases := map[string]struct {
		Msg            weave.Msg
		Conditions     []weave.Condition
		WantCheckErr   *errors.Error
		WantDeliverErr *errors.Error
	}{
		"successfully update a contract": {
			Conditions: []weave.Condition{
				cindyCond,
			},
			Msg: &UpdateContractMsg{
				ContractID: weavetest.SequenceID(1),
				Participants: []*Participant{
					{Power: 1, Signature: alice},
					{Power: 2, Signature: bobby},
					{Power: 3, Signature: cindy},
				},
				ActivationThreshold: 2,
				AdminThreshold:      3,
			},
		},
		"cannot create a contract without participants": {
			Msg: &UpdateContractMsg{
				ContractID:          weavetest.SequenceID(1),
				Participants:        []*Participant{},
				ActivationThreshold: 2,
				AdminThreshold:      3,
			},
			WantCheckErr: errors.ErrInvalidMsg,
		},
		"cannot create if activation threshold is too high": {
			Msg: &UpdateContractMsg{
				ContractID: weavetest.SequenceID(1),
				Participants: []*Participant{
					{Power: 1, Signature: alice},
					{Power: 2, Signature: bobby},
					{Power: 3, Signature: cindy},
				},
				ActivationThreshold: 7, // higher than total
				AdminThreshold:      3,
			},
			WantCheckErr: errors.ErrInvalidMsg,
		},
		"cannot create if activation threshold is higher than admin threshold": {
			Msg: &UpdateContractMsg{
				ContractID: weavetest.SequenceID(1),
				Participants: []*Participant{
					{Power: 2, Signature: alice},
					{Power: 2, Signature: bobby},
				},
				ActivationThreshold: 2,
				AdminThreshold:      1,
			},
			WantCheckErr: errors.ErrInvalidMsg,
		},
	}

	auth := &weavetest.CtxAuth{Key: "auth"}
	rt := app.NewRouter()
	RegisterRoutes(rt, auth)

	for testName, tc := range cases {
		t.Run(testName, func(t *testing.T) {
			db := store.MemStore()
			ctx := context.Background()
			ctx = auth.SetConditions(ctx, tc.Conditions...)
			tx := &weavetest.Tx{Msg: tc.Msg}

			b := NewContractBucket()
			err := b.Save(db, b.Build(db, &Contract{
				Participants: []*Participant{
					{Power: 1, Signature: alice},
					{Power: 2, Signature: bobby},
					{Power: 3, Signature: cindy},
				},
				ActivationThreshold: 2,
				AdminThreshold:      3,
			}))
			if err != nil {
				t.Fatalf("cannot create a contract")
			}

			cache := db.CacheWrap()
			if _, err := rt.Check(ctx, cache, tx); !tc.WantCheckErr.Is(err) {
				t.Logf("want: %+v", tc.WantCheckErr)
				t.Logf(" got: %+v", err)
				t.Fatalf("check (%T)", tc.Msg)
			}
			cache.Discard()
			if tc.WantCheckErr != nil {
				// Failed checks are causing the message to be ignored.
				return
			}

			if _, err := rt.Deliver(ctx, db, tx); !tc.WantDeliverErr.Is(err) {
				t.Logf("want: %+v", tc.WantDeliverErr)
				t.Logf(" got: %+v", err)
				t.Fatalf("delivery (%T)", tc.Msg)
			}
		})
	}
}

func TestHandlers(t *testing.T) {
	alice := weavetest.NewCondition().Address()
	bobby := weavetest.NewCondition().Address()
	cindy := weavetest.NewCondition().Address()

	type action struct {
		Msg            weave.Msg
		WantCheckErr   *errors.Error
		WantDeliverErr *errors.Error
	}

	cases := map[string]struct {
		actions []action
	}{
		"successfully create a contract": {
			actions: []action{
				{
					Msg: &CreateContractMsg{
						Participants: []*Participant{
							{Power: 1, Signature: alice},
							{Power: 2, Signature: bobby},
							{Power: 3, Signature: cindy},
						},
						ActivationThreshold: 2,
						AdminThreshold:      3,
					},
				},
			},
		},
		"cannot create a contract without participants": {
			actions: []action{
				{
					Msg: &CreateContractMsg{
						Participants:        []*Participant{},
						ActivationThreshold: 2,
						AdminThreshold:      3,
					},
					WantCheckErr: errors.ErrInvalidMsg,
				},
			},
		},
		"cannot create if activation threshold is too high": {
			actions: []action{
				{
					Msg: &CreateContractMsg{
						Participants: []*Participant{
							{Power: 1, Signature: alice},
							{Power: 2, Signature: bobby},
							{Power: 3, Signature: cindy},
						},
						ActivationThreshold: 7, // higher than total
						AdminThreshold:      3,
					},
					WantCheckErr: errors.ErrInvalidMsg,
				},
			},
		},
		"cannot create if activation threshold is higher than admin threshold": {
			actions: []action{
				{
					Msg: &CreateContractMsg{
						Participants: []*Participant{
							{Power: 2, Signature: alice},
							{Power: 2, Signature: bobby},
						},
						ActivationThreshold: 2,
						AdminThreshold:      1,
					},
					WantCheckErr: errors.ErrInvalidMsg,
				},
			},
		},
		"cannot update a contract without participants": {
			actions: []action{
				{
					Msg: &CreateContractMsg{
						Participants: []*Participant{
							{Power: 1, Signature: alice},
							{Power: 2, Signature: bobby},
							{Power: 3, Signature: cindy},
						},
						ActivationThreshold: 2,
						AdminThreshold:      3,
					},
				},
				{
					Msg: &CreateContractMsg{
						Participants:        []*Participant{},
						ActivationThreshold: 2,
						AdminThreshold:      3,
					},
					WantCheckErr: errors.ErrInvalidMsg,
				},
			},
		},
	}

	auth := &weavetest.Auth{
		Signer: weavetest.NewCondition(), // Any signer will do.
	}
	rt := app.NewRouter()
	RegisterRoutes(rt, auth)

	for testName, tc := range cases {
		t.Run(testName, func(t *testing.T) {
			db := store.MemStore()

			for i, a := range tc.actions {
				ctx := context.Background()
				tx := &weavetest.Tx{Msg: a.Msg}

				cache := db.CacheWrap()
				if _, err := rt.Check(ctx, cache, tx); !a.WantCheckErr.Is(err) {
					t.Logf("want: %+v", a.WantCheckErr)
					t.Logf(" got: %+v", err)
					t.Fatalf("action %d check (%T)", i, a.Msg)
				}
				cache.Discard()
				if a.WantCheckErr != nil {
					// Failed checks are causing the message to be ignored.
					continue
				}

				if _, err := rt.Deliver(ctx, db, tx); !a.WantDeliverErr.Is(err) {
					t.Logf("want: %+v", a.WantDeliverErr)
					t.Logf(" got: %+v", err)
					t.Fatalf("action %d delivery (%T)", i, a.Msg)
				}
			}
		})
	}
}
