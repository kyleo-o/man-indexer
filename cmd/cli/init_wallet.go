package cli

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"manindexer/man"
	"os"
)

var (
	configFileName = "config.json"
)

var initWalletCmd = &cobra.Command{
	Use:   "init-wallet",
	Short: "Init Wallet for CLI",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		initCliWallet()
	},
}

func initCliWallet() {
	if _, err := os.Stat(configFileName); os.IsNotExist(err) {
		fmt.Println("Config file not found, creating a new one...")

		name := "mywallet5"
		//_, err := CreateWallet(name)
		////err := CreateAccount(name)
		//if err != nil {
		//	fmt.Printf("Failed to create wallet: %v\n", err)
		//	return
		//}
		address, err := GetNewAddress(name)
		if err != nil {
			fmt.Printf("Failed to get new address: %v\n", err)
			return
		}
		privateKey, err := DumpPrivKeyHex(address)
		if err != nil {
			fmt.Printf("Failed to dump private key: %v\n", err)
			return
		}

		wallet = &CliWallet{
			account:    name,
			privateKey: privateKey,
			address:    address,
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

func checkManDbAdapter() error {
	if man.DbAdapter == nil {
		return fmt.Errorf("MAN DB adapter is not initialized")
	}
	return nil
}
