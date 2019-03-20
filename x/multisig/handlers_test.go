package multisig

import (
	"testing"

	"github.com/iov-one/weave/app"
	"github.com/iov-one/weave/errors"
	"github.com/iov-one/weave/weavetest"
)

func TestHandlers(t *testing.T) {
	ab := weavetest.ActionBuilder{
		ChainID: "my-chain",
		Auth:    &weavetest.CtxAuth{Key: "auth"},
	}

	rt := app.NewRouter()
	RegisterRoutes(rt, ab.Auth)

	alice := weavetest.NewCondition().Address()
	bobby := weavetest.NewCondition().Address()
	cindy := weavetest.NewCondition().Address()

	cases := map[string]struct {
		actions []weavetest.Action
	}{
		"successfully create a contract": {
			actions: ab.Actions(
				weavetest.Action{
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
			),
		},
		"cannot create a contract without participants": {
			actions: ab.Actions(
				weavetest.Action{
					Msg: &CreateContractMsg{
						Participants:        []*Participant{},
						ActivationThreshold: 2,
						AdminThreshold:      3,
					},
					WantCheckErr: errors.ErrInvalidInput,
				},
			),
		},
		"cannot create if activation threshold is too high": {
			actions: ab.Actions(
				weavetest.Action{
					Msg: &CreateContractMsg{
						Participants: []*Participant{
							{Power: 1, Signature: alice},
							{Power: 2, Signature: bobby},
							{Power: 3, Signature: cindy},
						},
						ActivationThreshold: 7, // higher than total
						AdminThreshold:      3,
					},
					WantCheckErr: errors.ErrInvalidInput,
				},
			),
		},
	}

	for testName, _ := range cases {
		t.Run(testName, func(t *testing.T) {

		})
	}
}
