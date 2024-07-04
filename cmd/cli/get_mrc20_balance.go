package cli

import "github.com/spf13/cobra"

var getMrc20BalanceCmd = &cobra.Command{
	Use:   "balance",
	Short: "Display balance for a given address",
	Long:  `Display balance for a given address`,
	Run: func(cmd *cobra.Command, args []string) {
		getMrc20Balance()
	},
}

func getMrc20Balance() {
	if err := checkWallet(); err != nil {
		return
	}

}
