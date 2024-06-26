package pin

const (
	//ProtocolID string = "746573746964" //testid(HEX16)
	ProtocolID    string = "6d6574616964" //metaid
	CompliantPath string = "info;file;protocols;nft;ft;mrc20"
)

type PinInscription struct {
	Id                 string `json:"id"`
	Number             int64  `json:"number"`
	MetaId             string `json:"metaid"`
	Address            string `json:"address"`
	CreateAddress      string `json:"createAddress"`
	Output             string `json:"output"`
	OutputValue        int64  `json:"outputValue"`
	Timestamp          int64  `json:"timestamp"`
	GenesisFee         int64  `json:"genesisFee"`
	GenesisHeight      int64  `json:"genesisHeight"`
	GenesisTransaction string `json:"genesisTransaction"`
	TxIndex            int    `json:"txIndex"`
	TxInIndex          uint32 `json:"txInIndex"`
	Offset             uint64 `json:"offset"`
	Location           string `json:"location"`
	Operation          string `json:"operation"`
	Path               string `json:"path"`
	ParentPath         string `json:"parentPath"`
	OriginalPath       string `json:"originalPath"`
	Encryption         string `json:"encryption"`
	Version            string `json:"version"`
	ContentType        string `json:"contentType"`
	ContentTypeDetect  string `json:"contentTypeDetect"`
	ContentBody        []byte `json:"contentBody"`
	ContentLength      uint64 `json:"contentLength"`
	ContentSummary     string `json:"contentSummary"`
	Status             int    `json:"status"`
	OriginalId         string `json:"originalId"`
	IsTransfered       bool   `json:"isTransfered"`
	Preview            string `json:"preview"`
	Content            string `json:"content"`
	Pop                string `json:"pop"`
	PopLv              int    `json:"popLv"`
}
type PinTransferInfo struct {
	Address     string `json:"address"`
	Output      string `json:"output"`
	OutputValue int64  `json:"outputValue"`
	Offset      uint64 `json:"offset"`
	Location    string `json:"location"`
}
type PersonalInformationNode struct {
	Operation     string `json:"operation"`
	Path          string `json:"path"`
	Encryption    string `json:"encryption"`
	Version       string `json:"cersion"`
	ContentType   string `json:"contentType"`
	ContentBody   []byte `json:"contentBody"`
	ContentLength uint64 `json:"contentLength"`
	ParentPath    string `json:"parentPath"`
	Protocols     bool   `json:"protocols"`
}
type FollowData struct {
	MetaId        string `json:"metaId"`
	FollowMetaId  string `json:"followMetaId"`
	FoloowTime    int64  `json:"foloowTime"`
	FollowPinId   string `json:"followPinId"`
	UnFollowPinId string `json:"unFollowPinId"`
	Status        bool   `json:"status"`
}
type MetaIdInfo struct {
	Number        int64  `json:"number"`
	MetaId        string `json:"metaid"`
	Name          string `json:"name"`
	NameId        string `json:"nameId"`
	Address       string `json:"address"`
	Avatar        string `json:"avatar"`
	AvatarId      string `json:"avatarId"`
	Bio           string `json:"bio"`
	BioId         string `json:"bioId"`
	SoulbondToken string `json:"soulbondToken"`
	IsInit        bool   `json:"isInit"`
}

type PinTreeCatalog struct {
	RootTxId string `json:"rootTxId"`
	TreePath string `json:"treePath"`
}

type PinMsg struct {
	Content   string `json:"content"`
	Number    int64  `json:"number"`
	Operation string `json:"operation"`
	Height    int64  `json:"height"`
	Id        string `json:"id"`
	Type      string `json:"type"`
	Path      string `json:"path"`
	MetaId    string `json:"metaid"`
	Pop       string `json:"pop"`
}

type BlockMsg struct {
	BlockHash      string   `json:"blockHash"`
	Target         string   `json:"target"`
	Timestamp      string   `json:"timestamp"`
	Size           int64    `json:"size"`
	Weight         int64    `json:"weight"`
	TransactionNum int      `json:"transactionNum"`
	Transaction    []string `json:"transaction"`
}
type PinCount struct {
	Block  int64 `json:"block"`
	Pin    int64 `json:"Pin"`
	MetaId int64 `json:"metaId"`
	App    int64 `json:"app"`
}
