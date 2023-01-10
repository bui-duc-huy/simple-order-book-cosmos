package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreateOrder{}, "dex/CreateOrder", nil)
	cdc.RegisterConcrete(&MsgCancelOrder{}, "dex/CancelOrder", nil)
	cdc.RegisterConcrete(&MsgApproveOrder{}, "dex/ApproveOrder", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateOrder{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCancelOrder{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgApproveOrder{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
