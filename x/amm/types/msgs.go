package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	_ sdk.Msg = (*MsgAddLiquidity)(nil)
	_ sdk.Msg = (*MsgRemoveLiquidity)(nil)
)

// Message types for the module
const (
	TypeMsgAddLiquidity    = "add_liquidity"
	TypeMsgRemoveLiquidity = "remove_liquidity"
	TypeMsgSwapExactIn     = "swap_exact_in"
	TypeMsgSwapExactOut    = "swap_exact_out"
)

func NewMsgAddLiquidity(sender sdk.AccAddress, coins sdk.Coins) *MsgAddLiquidity {
	return &MsgAddLiquidity{
		Sender: sender.String(),
		Coins:  coins,
	}
}

func (m MsgAddLiquidity) Route() string { return RouterKey }
func (m MsgAddLiquidity) Type() string  { return TypeMsgAddLiquidity }

func (m MsgAddLiquidity) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&m))
}

func (m MsgAddLiquidity) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(m.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{addr}
}

func (m MsgAddLiquidity) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Sender); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address: %v", err)
	}
	if err := m.Coins.Validate(); err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, err.Error())
	}
	if len(m.Coins) != 2 {
		return ErrWrongCoinNumber
	}
	return nil
}

func NewMsgRemoveLiquidity(sender sdk.AccAddress, share sdk.Coin) *MsgRemoveLiquidity {
	return &MsgRemoveLiquidity{
		Sender: sender.String(),
		Share:  share,
	}
}

func (m MsgRemoveLiquidity) Route() string { return RouterKey }
func (m MsgRemoveLiquidity) Type() string  { return TypeMsgRemoveLiquidity }

func (m MsgRemoveLiquidity) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&m))
}

func (m MsgRemoveLiquidity) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(m.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{addr}
}

func (m MsgRemoveLiquidity) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Sender); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address: %v", err)
	}
	if _, err := ParseShareDenom(m.Share.Denom); err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}
	return nil
}

func NewMsgSwapExactIn(sender sdk.AccAddress, coinIn, minCoinOut sdk.Coin) *MsgSwapExactIn {
	return &MsgSwapExactIn{
		Sender:     sender.String(),
		CoinIn:     coinIn,
		MinCoinOut: minCoinOut,
	}
}

func (m MsgSwapExactIn) Route() string { return RouterKey }
func (m MsgSwapExactIn) Type() string  { return TypeMsgSwapExactIn }

func (m MsgSwapExactIn) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&m))
}

func (m MsgSwapExactIn) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(m.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{addr}
}

func (msg MsgSwapExactIn) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Sender); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address: %v", err)
	}
	if err := msg.CoinIn.Validate(); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidCoins, "invalid coin in: %v", err)
	}
	if err := msg.MinCoinOut.Validate(); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidCoins, "invalid min coin out: %v", err)
	}
	return nil
}

func NewMsgSwapExactOut(sender sdk.AccAddress, coinOut, maxCoinIn sdk.Coin) *MsgSwapExactOut {
	return &MsgSwapExactOut{
		Sender:    sender.String(),
		CoinOut:   coinOut,
		MaxCoinIn: maxCoinIn,
	}
}

func (msg MsgSwapExactOut) Route() string { return RouterKey }
func (msg MsgSwapExactOut) Type() string  { return TypeMsgSwapExactOut }

func (msg MsgSwapExactOut) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

func (msg MsgSwapExactOut) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{addr}
}

func (msg MsgSwapExactOut) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Sender); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address: %v", err)
	}
	if err := msg.CoinOut.Validate(); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidCoins, "invalid coin out: %v", err)
	}
	if err := msg.MaxCoinIn.Validate(); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidCoins, "invalid max coin in: %v", err)
	}
	return nil
}
