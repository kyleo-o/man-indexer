package cli

import (
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/btcsuite/btcd/btcjson"
	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/rpcclient"
	"manindexer/common"
)

var (
	client *rpcclient.Client
)

func init() {
	btc := common.Config.Btc
	rpcConfig := &rpcclient.ConnConfig{
		Host:                 btc.RpcHost,
		User:                 btc.RpcUser,
		Pass:                 btc.RpcPass,
		HTTPPostMode:         btc.RpcHTTPPostMode, // Bitcoin core only supports HTTP POST mode
		DisableTLS:           btc.RpcDisableTLS,   // Bitcoin core does not provide TLS by default
		DisableAutoReconnect: true,
		DisableConnectOnNew:  true,
	}
	var err error
	client, err = rpcclient.New(rpcConfig, nil)
	if err != nil {
		panic(err)
	}
}

func GetBtcUtxo() ([]btcjson.ListUnspentResult, error) {
	return client.ListUnspent()
}

func CreateWallet(name string) error {
	return client.CreateNewAccount(name)
}

func GetNewAddress(accountName string) (string, error) {
	newAddress, err := client.GetNewAddress(accountName)
	if err != nil {
		return "", errors.New(fmt.Sprintf("Error creating new address: %v", err))
	}
	return newAddress.EncodeAddress(), nil
}
func DumpPrivKeyHex(newAddress string) (string, error) {
	addr, err := btcutil.DecodeAddress(newAddress, GetNetParams("livenet"))
	if err != nil {
		return "", err
	}
	dumpPrivKey, err := client.DumpPrivKey(addr)
	if err != nil {
		return "", errors.New(fmt.Sprintf("Error dumping private key: %v", err))

	}
	return hex.EncodeToString(dumpPrivKey.PrivKey.Serialize()), nil
}
