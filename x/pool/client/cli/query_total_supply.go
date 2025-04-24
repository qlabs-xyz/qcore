package cli

import (
	"context"

	"github.com/qlabs-xyz/qcore/x/pool/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

func CmdQueryTotalSupply() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get-total_supply",
		Short: "shows token total supply amount",
		//Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {

			clientCtx := client.GetClientContextFromCmd(cmd)
			queryClient := types.NewQueryClient(clientCtx)

			req := &types.QueryTotalSupplyRequest{}

			res, err := queryClient.GetTotalSupply(context.Background(), req)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}
