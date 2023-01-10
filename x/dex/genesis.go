package dex

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"orderbook/x/dex/keeper"
	"orderbook/x/dex/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the order
	for _, elem := range genState.OrderList {
		k.SetOrder(ctx, elem)
	}

	// Set order count
	k.SetOrderCount(ctx, genState.OrderCount)
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.OrderList = k.GetAllOrder(ctx)
	genesis.OrderCount = k.GetOrderCount(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
