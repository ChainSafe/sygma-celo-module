package cli

import (
	"github.com/ChainSafe/chainbridge-celo-module/transaction"
	"github.com/ChainSafe/chainbridge-core/chains/evm/cli/bridge"
	"github.com/spf13/cobra"
)

var BridgeCeloCmd = &cobra.Command{
	Use:   "bridge",
	Short: "Bridge-related instructions",
	Long:  "Bridge-related instructions",
}

var registerResourceCeloCmd = &cobra.Command{
	Use:   "register-resource",
	Short: "Register a resource ID",
	Long:  "Register a resource ID with a contract address for a handler",
	RunE: func(cmd *cobra.Command, args []string) error {
		txFabric := transaction.NewCeloTransaction
		return bridge.RegisterResourceCmd(cmd, args, txFabric)
	},
}

var setBurnCeloCmd = &cobra.Command{
	Use:   "set-burn",
	Short: "Set a token contract as mintable/burnable",
	Long:  "Set a token contract as mintable/burnable in a handler",
	RunE: func(cmd *cobra.Command, args []string) error {
		txFabric := transaction.NewCeloTransaction
		return bridge.SetBurnCmd(cmd, args, txFabric)
	},
}

// TODO:
// cancel-proposal
// query-proposal
// query-resource
// register-generic-resource

func init() {
	bridge.BindRegisterResourceCmdFlags(registerResourceCeloCmd)
	bridge.BindSetBurnCmdFlags(setBurnCeloCmd)

	BridgeCeloCmd.AddCommand(registerResourceCeloCmd, setBurnCeloCmd)
}
