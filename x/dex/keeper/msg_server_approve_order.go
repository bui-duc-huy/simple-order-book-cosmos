package keeper

import (
	"context"
  "fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"orderbook/x/dex/types"
)

func (k msgServer) ApproveOrder(goCtx context.Context, msg *types.MsgApproveOrder) (*types.MsgApproveOrderResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

  order, found := k.GetOrder(ctx, msg.Id)

  if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("Key %d doesn't exist", msg.Id))
  }

  if order.State != "created" {
		return nil, sdkerrors.Wrapf(types.ErrWrongOrderState, "%v", order.State)
  }

  seller, _ := sdk.AccAddressFromBech32(order.Creator)
  buyer, _ := sdk.AccAddressFromBech32(msg.Creator)

  k.bankKeeper.SendCoins(ctx, buyer, seller, sdk.NewCoins(order.SourceToken))
	k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, buyer, sdk.NewCoins(order.DestToken))

  order.Buyer = msg.Creator
	order.State = "approved"
	k.SetOrder(ctx, order)

	return &types.MsgApproveOrderResponse{}, nil
}
