package nft

import (
	"fmt"

	"github.com/iov-one/weave"
	"github.com/iov-one/weave/errors"
)

const UnlimitedCount = -1

type ApprovalMeta []Approval
type Approvals map[Action]ApprovalMeta

func (m ActionApprovals) Clone() ActionApprovals {
	return m
}

func (m Approval) Clone() Approval {
	return m
}

func (m ApprovalMeta) Clone() ApprovalMeta {
	return m
}

func (m ApprovalMeta) Validate() error {
	for _, v := range m {
		if err := v.Validate(); err != nil {
			return err
		}
	}
	return nil
}

func (m Approval) Validate() error {
	if err := m.Options.Validate(); err != nil {
		return err
	}
	if err := m.AsAddress().Validate(); err != nil {
		return err
	}

	return m.Options.Validate()
}

func (a Approval) AsAddress() weave.Address {
	return weave.Address(a.Address)
}

func (a Approval) Equals(o Approval) bool {
	return a.AsAddress().Equals(o.AsAddress()) &&
		a.Options.Equals(o.Options)
}

func (a ApprovalOptions) Equals(o ApprovalOptions) bool {
	return a.Immutable == o.Immutable && a.Count == o.Count && a.UntilBlockHeight == o.UntilBlockHeight
}

func (a ApprovalOptions) EqualsAfterUse(used ApprovalOptions) bool {
	if a.Count == UnlimitedCount || a.Immutable {
		return a.Equals(used)
	}

	return a.Count == used.Count+1 &&
		a.Immutable == used.Immutable &&
		a.UntilBlockHeight == used.UntilBlockHeight
}

func (a ApprovalOptions) Validate() error {
	if a.Count == 0 || a.Count < UnlimitedCount {
		return errors.Wrap(errors.ErrInvalidInput, "Approval count should either be unlimited or above zero")
	}
	return nil
}

//This requires all the model-specific actions to be passed here
//TODO: Not sure I'm a fan of array of maps, but it makes sense
//given we validate using protobuf enum value maps
func (m Approvals) Validate(actionMaps ...map[Action]int32) error {
	for action, meta := range m {
		if err := meta.Validate(); err != nil {
			return err
		}

		if !isValidAction(action) {
			return errors.Wrap(errors.ErrInvalidInput, fmt.Sprintf("illegal action: %s", action))
		}
		for _, actionMap := range actionMaps {
			if _, ok := actionMap[action]; ok {
				return errors.Wrap(errors.ErrInvalidInput, fmt.Sprintf("illegal action: %s", action))
			}
		}
	}

	return nil
}

func (m Approvals) FilterExpired(blockHeight int64) Approvals {
	res := make(map[Action]ApprovalMeta, 0)
	for action, approvals := range m {
		for _, approval := range approvals {
			if approval.Options.UntilBlockHeight > 0 && approval.Options.UntilBlockHeight < blockHeight {
				continue
			}

			if approval.Options.Count == 0 {
				continue
			}

			if _, ok := res[action]; !ok {
				res[action] = make([]Approval, 0)
			}

			res[action] = append(res[action], approval)
		}
	}
	return res
}

func (m Approvals) AsPersistable() []ActionApprovals {
	r := make([]ActionApprovals, 0)
	for k, v := range m {
		r = append(r, ActionApprovals{Action: k, Approvals: v})
	}
	return r
}

func (m Approvals) IsEmpty() bool {
	return len(m) == 0
}

func (m Approvals) MetaByAction(action Action) ApprovalMeta {
	return m[action]
}

func (m Approvals) ForAction(action Action) Approvals {
	res := make(map[Action]ApprovalMeta, 0)
	res[action] = m.MetaByAction(action)
	return res
}

func (m Approvals) ForAddress(addr weave.Address) Approvals {
	res := make(map[Action]ApprovalMeta, 0)
	for k, v := range m {
		r := make([]Approval, 0)
		for _, vv := range v {
			if vv.AsAddress().Equals(addr) {
				r = append(r, vv)
			}
		}
		if len(r) > 0 {
			res[k] = r
		}
	}
	return res
}

func (m Approvals) Filter(obsolete Approvals) Approvals {
	res := make(map[Action]ApprovalMeta, 0)

ApprovalsLoop:
	for action, approvals := range m {
		obsoleteApprovals := obsolete[action]
		for _, approval := range approvals {
			for _, obsoleteApproval := range obsoleteApprovals {
				if approval.Equals(obsoleteApproval) {
					continue ApprovalsLoop
				}
			}
			res[action] = append(res[action], approval)
		}
	}
	return res
}

func (m Approvals) Add(action Action, approval Approval) Approvals {
	m[action] = append(m[action], approval)
	return m
}

func (m Approvals) UseCount() Approvals {
	res := make(map[Action]ApprovalMeta, 0)
	for action, approvals := range m {
		for _, approval := range approvals {
			if approval.Options.Count == 0 {
				continue
			}

			if _, ok := res[action]; !ok {
				res[action] = make([]Approval, 0)
			}

			if !approval.Options.Immutable {
				approval.Options.Count--
			}

			res[action] = append(res[action], approval)
		}
	}
	return res
}

func (m Approvals) MergeUsed(used Approvals) Approvals {
	for action, aUsed := range used {
		found := false
		aDest := m[action]
		for _, u := range aUsed {
			for idx, dest := range aDest {
				if u.AsAddress().Equals(dest.AsAddress()) &&
					dest.Options.EqualsAfterUse(u.Options) {
					aDest[idx] = u
					found = true
					break
				}
			}

			if !found {
				m[action] = append(m[action])
			}
		}
	}
	return m
}

func (m Approvals) Intersect(others Approvals) Approvals {
	res := make(map[Action]ApprovalMeta, 0)
	for action, approvals := range others {
		mApprovals := m[action]
		for _, src := range approvals {
			for _, dest := range mApprovals {
				if dest.Equals(src) {
					if _, ok := res[action]; !ok {
						res[action] = make([]Approval, 0)
					}
					res[action] = append(res[action], dest)
				}
			}
		}
	}
	return res
}
