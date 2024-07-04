package cli

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
)

var (
	configFileName = "config.json"
)

var initWalletCmd = &cobra.Command{
	Use:   "init-wallet",
	Short: "Init Wallet",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		initCliWallet()
	},
}

func initCliWallet() {
	if _, err := os.Stat(configFileName); os.IsNotExist(err) {
		fmt.Println("Config file not found, creating a new one...")

		wallet = &CliWallet{
			privateKey: "default_private_key",
			address:    "default_address",
		}

		file, err := os.Create(configFileName)
		if err != nil {
			fmt.Printf("Failed to create config file: %v\n", err)
			return
		}
		defer file.Close()

		jsonData, err := json.MarshalIndent(wallet, "", "  ")
		if err != nil {
			fmt.Printf("Failed to marshal JSON: %v\n", err)
			return
		}

		_, err = file.Write(jsonData)
		if err != nil {
			fmt.Printf("Failed to write to config file: %v\n", err)
			return
		}

		fmt.Println("Config file created successfully.")
	} else {
		// 文件存在，读取并解析
		fmt.Println("Config file found, reading...")

		data, err := ioutil.ReadFile(configFileName)
		if err != nil {
			fmt.Printf("Failed to read config file: %v\n", err)
			return
		}

		err = json.Unmarshal(data, &wallet)
		if err != nil {
			fmt.Printf("Failed to unmarshal JSON: %v\n", err)
			return
		}

		fmt.Println("Config file read successfully.")
	}
}

func checkWallet() error {
	if wallet == nil {
		return fmt.Errorf("wallet is not initialized")
	}
	if wallet.GetAddress() == "" {
		return fmt.Errorf("wallet address is not initialized")
	}
	if wallet.GetPrivateKey() == "" {
		return fmt.Errorf("wallet private key is not initialized")
	}
	return nil
}
