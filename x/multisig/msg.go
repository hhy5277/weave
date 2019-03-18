package multisig

import (
	"github.com/iov-one/weave/errors"
)

const (
	pathCreateContractMsg = "multisig/create"
	pathUpdateContractMsg = "multisig/update"

	creationCost int64 = 300 // 3x more expensive than SendMsg
	updateCost   int64 = 150 // Half the creation cost
)

// Path fulfills weave.Msg interface to allow routing
func (CreateContractMsg) Path() string {
	return pathCreateContractMsg
}

// Validate enforces sigs and threshold boundaries
func (c *CreateContractMsg) Validate() error {
	return validateWeights(errors.ErrInvalidMsg,
		c.Participants, c.ActivationThreshold, c.AdminThreshold)
}

// Path fulfills weave.Msg interface to allow routing
func (UpdateContractMsg) Path() string {
	return pathUpdateContractMsg
}

// Validate enforces sigs and threshold boundaries
func (c *UpdateContractMsg) Validate() error {
	return validateWeights(errors.ErrInvalidMsg,
		c.Participants, c.ActivationThreshold, c.AdminThreshold)
}

// validateWeights returns an error if given participants and thresholds
// configuration is not valid. This check is done on model and messages so
// instead of copying the code it is extracted into this function.
func validateWeights(
	baseErr error,
	ps []*Participant,
	activationThreshold Weight,
	adminThreshold Weight,
) error {
	if len(ps) == 0 {
		return errors.Wrap(baseErr, "missing participants")
	}

	for _, p := range ps {
		if err := p.Power.Validate(); err != nil {
			return errors.Wrapf(err, "participant %s", p.Signature)
		}
		if err := p.Signature.Validate(); err != nil {
			return errors.Wrapf(err, "participant %s", p.Signature)
		}
	}
	if err := activationThreshold.Validate(); err != nil {
		return errors.Wrap(err, "activation threshold")
	}
	if err := adminThreshold.Validate(); err != nil {
		return errors.Wrap(err, "admin threshold")
	}

	var total Weight
	for _, p := range ps {
		total += p.Power
	}

	if activationThreshold > total {
		return errors.Wrap(baseErr, "activation threshold greater than total power")
	}
	if adminThreshold > total {
		return errors.Wrap(baseErr, "admin threshold greater than total power")
	}
	return nil
}
