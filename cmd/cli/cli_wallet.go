package cli

import "github.com/btcsuite/btcd/chaincfg"

var (
	wallet *CliWallet
)

type CliWallet struct {
	account    string
	privateKey string
	address    string
	net        *chaincfg.Params
	protocolId string
}

type WalletUtxo struct {
	TxId         string `json:"txId"`
	Vout         uint32 `json:"vout"`
	Shatoshi     int64  `json:"shatoshi"`
	ScriptPubKey string `json:"scriptPubKey"`
	Address      string `json:"address"`
}

func (c *CliWallet) toString() string {
	return "privateKey: " + c.privateKey + ", address: " + c.address
}

func (c *CliWallet) GetPrivateKey() string {
	return c.privateKey
}

func (c *CliWallet) GetAddress() string {
	return c.address
}

func (c *CliWallet) GetUtxos() []*WalletUtxo {
	return nil
}

func (c *CliWallet) GetNet() *chaincfg.Params {
	return c.net
}

func (c *CliWallet) GetProtocolId() string {
	return c.protocolId
}

func (c *CliWallet) GetMrc20Utxos() {

}

func (c *CliWallet) GetPins() {

}

func (c *CliWallet) GetBalance() {

}

func (c *CliWallet) GetMrc20Balance() {

}
