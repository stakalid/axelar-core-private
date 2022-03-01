package types

// Event types
const (
	EventTypeChain           = "chain"
	EventTypeChainMaintainer = "chainMaintainer"
)

// Event attribute keys
const (
	AttributeKeyChain                  = "chain"
	AttributeKeyChainMaintainerAddress = "chainMaintainerAddress"
	AttributeKeyAsset                  = "asset"
)

// Event attribute values
const (
	AttributeValueRegister    = "register"
	AttributeValueDeregister  = "deregister"
	AttributeValueActivated   = "activated"
	AttributeValueDeactivated = "deactivated"
)
