package postgresql

import (
	"manindexer/mrc20"
	"manindexer/pin"
)

func (pg *Postgresql) GetMrc20TickInfo(tick string) (info mrc20.Mrc20DeployInfo, err error) {
	return
}

func (pg *Postgresql) SaveMrc20Pin(data []mrc20.Mrc20Utxo) (err error) {
	return
}
func (pg *Postgresql) SaveMrc20Tick(data []mrc20.Mrc20DeployInfo) (err error) {
	return
}
func (pg *Postgresql) GetMrc20TickPageList(cursor int64, size int64, order string) (total int64, list []mrc20.Mrc20DeployInfo, err error) {
	return
}
func (pg *Postgresql) AddMrc20Shovel(shovelList []string, pinId string) (err error) {
	return
}
func (pg *Postgresql) GetMrc20Shovel(shovels []string) (data map[string]mrc20.Mrc20Shovel, err error) {
	return
}
func (pg *Postgresql) UpdateMrc20TickInfo(tickId string, minted int64) (err error) {
	return
}
func (pg *Postgresql) GetMrc20ByAddressAndTick(address string, tickId string) (list []mrc20.Mrc20Utxo, err error) {
	return
}
func (pg *Postgresql) GetMrc20HistoryPageList(tickId string, isPage bool, page int64, size int64) (list []mrc20.Mrc20Utxo, total int64, err error) {
	return
}
func (pg *Postgresql) GetMrc20UtxoByOutPutList(outputList []string) (list []*mrc20.Mrc20Utxo, err error) {
	return
}
func (pg *Postgresql) UpdateMrc20Utxo(list []*mrc20.Mrc20Utxo) (err error) {
	return
}
func (pg *Postgresql) GetHistoryByAddress(tickId string, address string, page int64, size int64) (list []mrc20.Mrc20Utxo, total int64, err error) {
	return
}
func (pg *Postgresql) GetMrc20BalanceByAddress(address string) (list []mrc20.Mrc20Balance, total int64, err error) {
	return
}
func (pg *Postgresql) GetHistoryByTx(txId string, index int64, cursor int64, size int64) (list []mrc20.Mrc20Utxo, total int64, err error) {
	return
}
func (pg *Postgresql) GetShovelListByAddress(address string, lv int, path, query, key, operator, value string, cursor int64, size int64) (list []*pin.PinInscription, total int64, err error) {
	return
}
