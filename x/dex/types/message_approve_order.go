package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgApproveOrder = "approve_order"

var _ sdk.Msg = &MsgApproveOrder{}

func NewMsgApproveOrder(creator string, id uint64) *MsgApproveOrder {
	return &MsgApproveOrder{
		Creator: creator,
		Id:      id,
	}
}

func (msg *MsgApproveOrder) Route() string {
	return RouterKey
}

func (msg *MsgApproveOrder) Type() string {
	return TypeMsgApproveOrder
}

func (msg *MsgApproveOrder) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgApproveOrder) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgApproveOrder) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
