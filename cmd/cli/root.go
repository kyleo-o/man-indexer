package cli

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"manindexer/common"
	"manindexer/man"
	"os"
)

var CfgFile = ""

var rootCmd = &cobra.Command{
	Use:   "man-cli",
	Short: "MAN-CLI is a tool to interact with metaid-v2",
	Long:  "This is a MAN-CLI, which is a tool to interact with metaid-v2 in bitcoin chain",
	Run: func(cmd *cobra.Command, args []string) {
		Error(cmd, args, errors.New("unrecognized command"))
	},
}

func initConfig() {
	man.InitAdapter(common.Chain, common.Db, common.TestNet, common.Server)
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.AddCommand(initWalletCmd)
	rootCmd.AddCommand(getVersionCmd)
	rootCmd.AddCommand(getBalanceCmd)
	rootCmd.AddCommand(getMrc20BalanceCmd)
	rootCmd.AddCommand(mrc20OperationCmd)
}

func Error(cmd *cobra.Command, args []string, err error) {
	fmt.Fprintf(os.Stderr, "execute %s args:%v error:%v\n", cmd.Name(), args, err)
	os.Exit(1)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
