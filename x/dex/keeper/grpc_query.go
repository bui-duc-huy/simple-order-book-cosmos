package keeper

import (
	"orderbook/x/dex/types"
)

var _ types.QueryServer = Keeper{}
