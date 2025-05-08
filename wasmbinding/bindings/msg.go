package bindings

type QCoreMsg struct {
	MsgMintTribute *MsgMintTribute `json:"msg_mint_tribute,omitempty"`
}

type MsgMintTribute struct {
	Creator        string `json:"creator"`
	MintAmount     string `json:"mint_amount"`
	ReceiptAddress string `json:"receipt_address"`
}
