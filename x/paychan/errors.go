package paychan

// ABCI Response Codes
//
// paychan takes 1021-1029
const (
	CodeMissingCondition uint32 = 1021
	CodeInvalidCondition        = 1022
	CodeNotFound                = 1023

	CodeTODO = 666 // Used while refactoring as a temporary replacement to internal codes
)
