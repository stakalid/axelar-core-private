package cli

import (
	"fmt"
	"io/ioutil"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authUtils "github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/spf13/cobra"

	"github.com/axelarnetwork/axelar-core/utils"
	"github.com/axelarnetwork/axelar-core/x/ethereum/types"
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd(cdc *codec.Codec) *cobra.Command {
	ethTxCmd := &cobra.Command{
		Use:                        "ethereum",
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		TraverseChildren:           true,
		RunE:                       client.ValidateCmd,
	}

	ethTxCmd.AddCommand(
		flags.PostCommands(
			GetCmdSignTx(cdc),
			GetCmdVerifyTx(cdc),
			GetCmdSignPendingTransfersTx(cdc),
			GetCmdSignDeployToken(cdc),
		)...,
	)

	return ethTxCmd
}

// GetCmdSignTx returns the cli command to sign the given transaction
func GetCmdSignTx(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "sign [tx json file path]",
		Short: "sign a raw Ethereum transaction",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx, txBldr := utils.PrepareCli(cmd.InOrStdin(), cdc)

			json, err := ioutil.ReadFile(args[0])
			if err != nil {
				return err
			}
			var tx *ethTypes.Transaction
			cdc.MustUnmarshalJSON(json, &tx)

			msg := types.NewMsgSignTx(cliCtx.GetFromAddress(), json)

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return authUtils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}

	return cmd
}

// GetCmdVerifyTx returns the cli command to verify the given transaction
func GetCmdVerifyTx(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "verify [tx json file path]",
		Short: "Verify an Ethereum transaction",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {

			cliCtx, txBldr := utils.PrepareCli(cmd.InOrStdin(), cdc)

			json, err := ioutil.ReadFile(args[0])
			if err != nil {
				return err
			}
			var tx *ethTypes.Transaction
			cdc.MustUnmarshalJSON(json, &tx)

			msg := types.NewMsgVerifyTx(cliCtx.GetFromAddress(), json)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return authUtils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}

// GetCmdSignPendingTransfersTx returns the cli command to sign all pending token transfers to Ethereum
func GetCmdSignPendingTransfersTx(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "sign-pending-transfers",
		Short: "Sign all pending transfers to Ethereum",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx, txBldr := utils.PrepareCli(cmd.InOrStdin(), cdc)

			msg := types.NewMsgSignPendingTransfersTx(cliCtx.GetFromAddress())
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return authUtils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}

// GetCmdSignDeployToken returns the cli command to sign deploy-token command data for Ethereum
func GetCmdSignDeployToken(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "sign-deploy-token [name] [symbol] [decimals] [capacity]",
		Short: "Signs the call data to deploy a token with the AxelarGateway contract",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx, txBldr := utils.PrepareCli(cmd.InOrStdin(), cdc)

			decs, err := strconv.ParseUint(args[2], 10, 8)
			if err != nil {
				return fmt.Errorf("could not parse decimals")
			}

			capacity, ok := sdk.NewIntFromString(args[3])
			if !ok {
				return fmt.Errorf("could not parse capacity")
			}
			msg := types.NewMsgSignDeployToken(cliCtx.GetFromAddress(), args[0], args[1], uint8(decs), capacity)
			if err = msg.ValidateBasic(); err != nil {
				return err
			}

			return authUtils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}