package canowchain_test

import (
	"testing"

	keepertest "canow-chain/testutil/keeper"
	"canow-chain/testutil/nullify"
	"canow-chain/x/canowchain"
	"canow-chain/x/canowchain/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.CanowchainKeeper(t)
	canowchain.InitGenesis(ctx, *k, genesisState)
	got := canowchain.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
