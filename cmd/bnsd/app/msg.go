package app

import (
	"github.com/iov-one/weave"
	"github.com/iov-one/weave/coin"
	"github.com/iov-one/weave/x/cash"
)

// Fee sets the FeeInfo for this tx
func (tx *Tx) Fee(payer weave.Address, fee coin.Coin) {
	tx.Fees = &cash.FeeInfo{
		Payer: payer,
		Fees:  &fee}
}

//Commented out for a minimal feature-set release
//import (
//	"github.com/iov-one/weave"
//	"github.com/iov-one/weave/errors"
//	"github.com/iov-one/weave/x/batch"
//)
//
//var _ batch.Msg = (*BatchMsg)(nil)
//
//func (*BatchMsg) Path() string {
//	return batch.PathExecuteBatchMsg
//}
//
//func (msg *BatchMsg) MsgList() ([]weave.Msg, error) {
//	messages := make([]weave.Msg, len(msg.Messages))
//	// make sure to cover all messages defined in protobuf
//	for i, m := range msg.Messages {
//		res, err := func() (weave.Msg, error) {
//			switch t := m.GetSum().(type) {
//			case *BatchMsg_Union_SendMsg:
//				return t.SendMsg, nil
//			case *BatchMsg_Union_NewTokenMsg:
//				return t.NewTokenMsg, nil
//			case *BatchMsg_Union_SetNameMsg:
//				return t.SetNameMsg, nil
//			case *BatchMsg_Union_CreateEscrowMsg:
//				return t.CreateEscrowMsg, nil
//			case *BatchMsg_Union_ReleaseEscrowMsg:
//				return t.ReleaseEscrowMsg, nil
//			case *BatchMsg_Union_ReturnEscrowMsg:
//				return t.ReturnEscrowMsg, nil
//			case *BatchMsg_Union_UpdateEscrowMsg:
//				return t.UpdateEscrowMsg, nil
//			case *BatchMsg_Union_CreateContractMsg:
//				return t.CreateContractMsg, nil
//			case *BatchMsg_Union_UpdateContractMsg:
//				return t.UpdateContractMsg, nil
//			case *BatchMsg_Union_SetValidatorsMsg:
//				return t.SetValidatorsMsg, nil
//			case *BatchMsg_Union_AddApprovalMsg:
//				return t.AddApprovalMsg, nil
//			case *BatchMsg_Union_RemoveApprovalMsg:
//				return t.RemoveApprovalMsg, nil
//			case *BatchMsg_Union_IssueUsernameNftMsg:
//				return t.IssueUsernameNftMsg, nil
//			case *BatchMsg_Union_AddUsernameAddressNftMsg:
//				return t.AddUsernameAddressNftMsg, nil
//			case *BatchMsg_Union_RemoveUsernameAddressMsg:
//				return t.RemoveUsernameAddressMsg, nil
//			case *BatchMsg_Union_IssueBlockchainNftMsg:
//				return t.IssueBlockchainNftMsg, nil
//			case *BatchMsg_Union_IssueTickerNftMsg:
//				return t.IssueTickerNftMsg, nil
//			case *BatchMsg_Union_IssueBootstrapNodeNftMsg:
//				return t.IssueBootstrapNodeNftMsg, nil
//			default:
//				return nil, errors.ErrUnknownTxType(t)
//			}
//		}()
//		if err != nil {
//			return messages, err
//		}
//		messages[i] = res.(weave.Msg)
//	}
//	return messages, nil
//}
//
//func (msg *BatchMsg) Validate() error {
//	return batch.Validate(msg)
//}
