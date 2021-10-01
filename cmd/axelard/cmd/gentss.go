package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/server"
	genutiltypes "github.com/cosmos/cosmos-sdk/x/genutil/types"

	tssTypes "github.com/axelarnetwork/axelar-core/x/tss/types"

	"github.com/cosmos/cosmos-sdk/x/genutil"
	"github.com/spf13/cobra"
)

const (
	flagKeygen       = "keygen"
	flagCorruption   = "corruption"
	flagAck          = "ack-window"
	flagBondFraction = "bond-fraction"
)

// SetGenesisTSSCmd returns set-genesis-chain-params cobra Command.
func SetGenesisTSSCmd(defaultNodeHome string,
) *cobra.Command {
	var (
		ackWindow int64
	)

	cmd := &cobra.Command{
		Use:   "set-genesis-tss",
		Short: "Set the genesis parameters for the tss module",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			depCdc := clientCtx.JSONMarshaler
			cdc := depCdc.(codec.Marshaler)

			serverCtx := server.GetServerContextFromCmd(cmd)
			config := serverCtx.Config

			config.SetRoot(clientCtx.HomeDir)

			genFile := config.GenesisFile()
			appState, genDoc, err := genutiltypes.GenesisStateFromGenFile(genFile)
			if err != nil {
				return fmt.Errorf("failed to unmarshal genesis state: %w", err)
			}
			genesisTSS := tssTypes.GetGenesisStateFromAppState(cdc, appState)

			if ackWindow > 0 {
				genesisTSS.Params.AckWindowInBlocks = ackWindow
			}

			genesisTSSBz, err := cdc.MarshalJSON(&genesisTSS)
			if err != nil {
				return fmt.Errorf("failed to marshal tss genesis state: %w", err)
			}

			appState[tssTypes.ModuleName] = genesisTSSBz

			appStateJSON, err := json.Marshal(appState)
			if err != nil {
				return fmt.Errorf("failed to marshal application genesis state: %w", err)
			}
			genDoc.AppState = appStateJSON

			return genutil.ExportGenesisFile(genDoc, genFile)
		},
	}

	cmd.Flags().String(flags.FlagHome, defaultNodeHome, "node's home directory")
	cmd.Flags().Int64Var(&ackWindow, flagAck, 0, "A positive integer representing the time period for validators to submit acknowledgments for a keygen/sign in terms of number of blocks")

	return cmd
}
