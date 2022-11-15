package keeper

import (
	"canow-chain/x/canowchain/types"
)

var _ types.QueryServer = Keeper{}
