package keeper_test

import (
	"testing"

	testkeeper "canow-chain/testutil/keeper"
	"canow-chain/x/canowchain/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.CanowchainKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
