package cli

import (
	"fmt"
	"github.com/spf13/cobra"
	"manindexer/inscribe/mrc20_service"
	"strconv"
)

var mrc20OperationCmd = &cobra.Command{
	Use:   "mrc20op",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if err := checkWallet(); err != nil {
			return
		}
		if err := checkManDbAdapter(); err != nil {
			return
		}

		if len(args) < 1 {
			fmt.Println("mrc20op command required")
			return
		}
		switch args[0] {
		case "deploy":
			//tick := ""

			//mrc20opDeploy("", "", "", "", "", "", "", "", "", "", 0)
			break
		case "mint":
			if len(args) < 3 {
				fmt.Println("mrc20op mint {tickId} {feeRate}")
				return
			}
			tickId := args[1]
			feeRate, _ := strconv.ParseInt(args[2], 10, 64)
			mrc20opMint(tickId, feeRate)
			break
		case "transfer":
			if len(args) < 5 {
				fmt.Println("mrc20op transfer {tickId} {to} {amount} {feeRate}")
				return
			}
			tickId := args[1]
			to := args[2]
			amount := args[3]
			feeRate, _ := strconv.ParseInt(args[4], 10, 64)
			mrc20opTransfer(tickId, to, amount, feeRate)
			return
		}
	},
}

// ./man-cli mrc20op deploy
// ./man-cli mrc20op mint {tickId} {feeRate}
// ./man-cli mrc20op transfer {tickId} {to} {amount} {feeRate}

func mrc20opDeploy(tick, tokenName, decimals, amtPerMint, mintCount, premineCount, blockHeight, qualCreator, qualPath, qualCount, qualLvl string, feeRate int64) {
	var (
		commitTxId, revealTxId string = "", ""
		fee                    int64  = 0
		err                    error
		opRep                  *mrc20_service.Mrc20OpRequest
		payload                string = ""
		fetchCommitUtxoFunc    mrc20_service.FetchCommitUtxoFunc
	)
	opRep = &mrc20_service.Mrc20OpRequest{
		Net:                     wallet.GetNet(),
		MetaIdFlag:              wallet.GetProtocolId(),
		Op:                      "deploy",
		OpPayload:               payload,
		DeployPinOutAddress:     "",
		DeployPremineOutAddress: "",
		Mrc20OutValue:           546,
		ChangeAddress:           "",
	}

	fetchCommitUtxoFunc = func(needAmount int64) ([]*mrc20_service.CommitUtxo, error) {
		return GetBtcUtxoList(wallet.GetAddress(), needAmount)
	}

	commitTxId, revealTxId, fee, err = mrc20_service.Mrc20Deploy(opRep, feeRate, fetchCommitUtxoFunc)
	if err != nil {
		fmt.Printf("Mrc20 deploy err:%s\n", err.Error())
		return
	}
	fmt.Printf("Mrc20 deploy success\n")
	fmt.Printf("Fee:%d\n", fee)
	fmt.Printf("CommitTx:%s\n", commitTxId)
	fmt.Printf("RevealTxId:%s\n", revealTxId)
}

func mrc20opMint(tickId string, feeRate int64) {
	var (
		commitTxId, revealTxId string = "", ""
		fee                    int64  = 0
		err                    error
		opRep                  *mrc20_service.Mrc20OpRequest
		payload                string                      = fmt.Sprintf(`{"id":"%s"}`, tickId)
		mintPins               []*mrc20_service.MintPin    = make([]*mrc20_service.MintPin, 0)
		commitUtxos            []*mrc20_service.CommitUtxo = make([]*mrc20_service.CommitUtxo, 0)
		changeAddress          string                      = wallet.GetAddress()
		fetchCommitUtxoFunc    mrc20_service.FetchCommitUtxoFunc
	)

	for _, v := range wallet.GetUtxos() {
		commitUtxos = append(commitUtxos, &mrc20_service.CommitUtxo{
			PrivateKeyHex: wallet.GetPrivateKey(),
			PkScript:      v.ScriptPubKey,
			Address:       v.Address,
			UtxoTxId:      v.TxId,
			UtxoIndex:     v.Vout,
			UtxoOutValue:  v.Shatoshi,
		})
	}

	mintPins, err = getShovelList(wallet.GetAddress(), tickId)
	if err != nil {
		fmt.Printf("Mrc20 mint err:%s\n", err.Error())
		return
	}

	opRep = &mrc20_service.Mrc20OpRequest{
		Net:           wallet.GetNet(),
		MetaIdFlag:    wallet.GetProtocolId(),
		Op:            "mint",
		OpPayload:     payload,
		CommitUtxos:   commitUtxos,
		MintPins:      mintPins,
		Mrc20OutValue: 546,
		Mrc20OutAddressList: []string{
			wallet.GetAddress(),
		},
		ChangeAddress: changeAddress,
	}

	fetchCommitUtxoFunc = func(needAmount int64) ([]*mrc20_service.CommitUtxo, error) {
		return GetBtcUtxoList(wallet.GetAddress(), needAmount)
	}

	commitTxId, revealTxId, fee, err = mrc20_service.Mrc20Mint(opRep, feeRate, fetchCommitUtxoFunc)
	if err != nil {
		fmt.Printf("Mrc20 mint err:%s\n", err.Error())
		return
	}
	fmt.Printf("Mrc20 mint success\n")
	fmt.Printf("Fee:%d\n", fee)
	fmt.Printf("CommitTx:%s\n", commitTxId)
	fmt.Printf("RevealTxId:%s\n", revealTxId)
}

func mrc20opTransfer(tickId, to, amount string, feeRate int64) {
	var (
		commitTxId, revealTxId string = "", ""
		fee                    int64  = 0
		err                    error
		toPkScript, _                 = mrc20_service.AddressToPkScript(wallet.GetNet(), to)
		changeAddress          string = wallet.GetAddress()
		opRep                  *mrc20_service.Mrc20OpRequest
		commitUtxos            []*mrc20_service.CommitUtxo    = make([]*mrc20_service.CommitUtxo, 0)
		transferMrc20s         []*mrc20_service.TransferMrc20 = make([]*mrc20_service.TransferMrc20, 0)
		mrc20Outs              []*mrc20_service.Mrc20OutInfo  = []*mrc20_service.Mrc20OutInfo{
			{
				Amount:   amount,
				Address:  to,
				PkScript: toPkScript,
				OutValue: 546,
			},
		}
		payload             string = ""
		fetchCommitUtxoFunc mrc20_service.FetchCommitUtxoFunc
	)
	commitUtxos, err = GetBtcUtxoList(wallet.GetAddress(), 0)

	transferMrc20s, err = getMrc20Utxos(wallet.GetAddress(), tickId, amount)
	if err != nil {
		fmt.Printf("Mrc20 transfer err:%s\n", err.Error())
		return
	}

	for _, v := range wallet.GetUtxos() {
		commitUtxos = append(commitUtxos, &mrc20_service.CommitUtxo{
			PrivateKeyHex: wallet.GetPrivateKey(),
			PkScript:      v.ScriptPubKey,
			Address:       v.Address,
			UtxoTxId:      v.TxId,
			UtxoIndex:     v.Vout,
			UtxoOutValue:  v.Shatoshi,
		})
	}

	payload, err = mrc20_service.MakeTransferPayload(tickId, transferMrc20s, mrc20Outs)
	if err != nil {
		fmt.Printf("Mrc20 transfer err:%s\n", err.Error())
		return
	}
	opRep = &mrc20_service.Mrc20OpRequest{
		Net:            wallet.GetNet(),
		MetaIdFlag:     wallet.GetProtocolId(),
		Op:             "transfer",
		OpPayload:      payload,
		CommitUtxos:    commitUtxos,
		TransferMrc20s: transferMrc20s,
		Mrc20Outs:      mrc20Outs,
		ChangeAddress:  changeAddress,
	}

	fetchCommitUtxoFunc = func(needAmount int64) ([]*mrc20_service.CommitUtxo, error) {
		return GetBtcUtxoList(wallet.GetAddress(), needAmount)
	}

	commitTxId, revealTxId, fee, err = mrc20_service.Mrc20Transfer(opRep, feeRate, fetchCommitUtxoFunc)
	if err != nil {
		fmt.Printf("Mrc20 transfer err:%s\n", err.Error())
		return
	}
	fmt.Printf("Mrc20 transfer success\n")
	fmt.Printf("Fee:%d\n", fee)
	fmt.Printf("CommitTx:%s\n", commitTxId)
	fmt.Printf("RevealTxId:%s\n", revealTxId)
}
