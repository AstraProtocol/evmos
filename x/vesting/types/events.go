package types

// vesting events
const (
	EventTypeCreateClawbackVestingAccount = "create_clawback_vesting_account"
	EventTypeClawback                     = "clawback"

	AttributeKeyCoins       = "coins"
	AttributeKeyStartTime   = "start_time"
	AttributeKeyAccount     = "account"
	AttributeKeyFunder      = "funder"
	AttributeKeyDestination = "destination"
)
