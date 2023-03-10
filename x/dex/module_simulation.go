package dex

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"orderbook/testutil/sample"
	dexsimulation "orderbook/x/dex/simulation"
	"orderbook/x/dex/types"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = dexsimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgCreateOrder = "op_weight_msg_create_order"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateOrder int = 100

	opWeightMsgCancelOrder = "op_weight_msg_cancel_order"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCancelOrder int = 100

	opWeightMsgApproveOrder = "op_weight_msg_approve_order"
	// TODO: Determine the simulation weight value
	defaultWeightMsgApproveOrder int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	dexGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&dexGenesis)
}

// ProposalContents doesn't return any content functions for governance proposals
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// RandomizedParams creates randomized  param changes for the simulator
func (am AppModule) RandomizedParams(_ *rand.Rand) []simtypes.ParamChange {

	return []simtypes.ParamChange{}
}

// RegisterStoreDecoder registers a decoder
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreateOrder int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateOrder, &weightMsgCreateOrder, nil,
		func(_ *rand.Rand) {
			weightMsgCreateOrder = defaultWeightMsgCreateOrder
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateOrder,
		dexsimulation.SimulateMsgCreateOrder(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCancelOrder int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCancelOrder, &weightMsgCancelOrder, nil,
		func(_ *rand.Rand) {
			weightMsgCancelOrder = defaultWeightMsgCancelOrder
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCancelOrder,
		dexsimulation.SimulateMsgCancelOrder(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgApproveOrder int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgApproveOrder, &weightMsgApproveOrder, nil,
		func(_ *rand.Rand) {
			weightMsgApproveOrder = defaultWeightMsgApproveOrder
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgApproveOrder,
		dexsimulation.SimulateMsgApproveOrder(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
