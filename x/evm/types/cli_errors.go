package types

// module errors
const (
	ErrFDepositAddress = "could not get the deposit address"
	ErrFMasterKey      = "could not resolve master key: %s\n"
	ErrFKeyID          = "could not resolve key ID: %s\n"
	ErrFGatewayAddress = "could not resolve gateway address: %s\n"
	ErrFTokenAddress   = "could not resolve token address: %s\n"
	ErrFDeployTx       = "could not send the command transaction with txID %s"
	ErrFSignedTx       = "could not get transaction with txID %s"
	ErrFBytecode       = "could not get the bytecodes for contract %s"
	ErrFSendTx         = "could not send the transaction with txID %s"
	ErrFSendCommandTx  = "could not send %s transaction executing command %s"
)
