package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"
	"orderbook/x/dex/types"
)

var _ = strconv.Itoa(0)

func CmdCreateOrder() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-order [source-token] [dest-token]",
		Short: "Broadcast message create-order",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argSourceToken, err := sdk.ParseCoinNormalized(args[0])
			if err != nil {
				return err
			}
			argDestToken, err := sdk.ParseCoinNormalized(args[1])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateOrder(
				clientCtx.GetFromAddress().String(),
				argSourceToken,
				argDestToken,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
