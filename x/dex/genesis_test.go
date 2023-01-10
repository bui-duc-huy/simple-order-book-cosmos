package dex_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "orderbook/testutil/keeper"
	"orderbook/testutil/nullify"
	"orderbook/x/dex"
	"orderbook/x/dex/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		OrderList: []types.Order{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		OrderCount: 2,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.DexKeeper(t)
	dex.InitGenesis(ctx, *k, genesisState)
	got := dex.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.OrderList, got.OrderList)
	require.Equal(t, genesisState.OrderCount, got.OrderCount)
	// this line is used by starport scaffolding # genesis/test/assert
}
