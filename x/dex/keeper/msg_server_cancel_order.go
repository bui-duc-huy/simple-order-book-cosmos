package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"orderbook/x/dex/types"
)

func (k msgServer) CancelOrder(goCtx context.Context, msg *types.MsgCancelOrder) (*types.MsgCancelOrderResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	order, found := k.GetOrder(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("Key %d doesn't exist", msg.Id))
	}

	if order.Creator != msg.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Cannot cancel: not the creator")
	}

	if order.State != "created" {
		return nil, sdkerrors.Wrapf(types.ErrWrongOrderState, "%v", order.State)
	}

	creator, _ := sdk.AccAddressFromBech32(order.Creator)

	k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, creator, sdk.NewCoins(order.SourceToken))

	order.State = "cancelled"
	k.SetOrder(ctx, order)

	return &types.MsgCancelOrderResponse{}, nil
}
