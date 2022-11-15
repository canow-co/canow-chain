package keeper_test

import (
	"context"
	"testing"

	keepertest "canow-chain/testutil/keeper"
	"canow-chain/x/canowchain/keeper"
	"canow-chain/x/canowchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.CanowchainKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
