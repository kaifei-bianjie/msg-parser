package token

import (
	"fmt"
	. "github.com/kaifei-bianjie/msg-parser/modules"
)

type DocMsgMintToken struct {
	Symbol string `bson:"symbol"`
	Amount string `bson:"amount"`
	To     string `bson:"to"`
	Owner  string `bson:"owner"`
}

func (m *DocMsgMintToken) GetType() string {
	return MsgTypeMintToken
}

func (m *DocMsgMintToken) BuildMsg(v interface{}) {
	msg := v.(*MsgMintToken)

	m.Symbol = msg.Symbol
	m.Amount = fmt.Sprint(msg.Amount)
	m.To = msg.To
	m.Owner = msg.Owner
}

func (m *DocMsgMintToken) HandleTxMsg(v SdkMsg) MsgDocInfo {
	var (
		addrs []string
		msg   MsgMintToken
	)

	ConvertMsg(v, &msg)
	addrs = append(addrs, msg.Owner, msg.To)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(v, handler)
}
