package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"orderbook/x/dex/types"
)

func (k msgServer) CreateOrder(goCtx context.Context, msg *types.MsgCreateOrder) (*types.MsgCreateOrderResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	order := types.Order{
		SourceToken: msg.SourceToken,
		DestToken:   msg.DestToken,
		Creator:     msg.Creator,
		State:       "created",
	}

	creator, _ := sdk.AccAddressFromBech32(msg.Creator)

	sdkError := k.bankKeeper.SendCoinsFromAccountToModule(ctx, creator, types.ModuleName, sdk.NewCoins(msg.SourceToken))
	if sdkError != nil {
		return nil, sdkError
	}

	k.AppendOrder(ctx, order)

	return &types.MsgCreateOrderResponse{}, nil
}
