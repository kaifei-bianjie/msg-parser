package service

import (
	. "github.com/kaifei-bianjie/msg-parser/modules"
)

type (
	DocMsgWithdrawEarnedFees struct {
		Owner    string
		Provider string
	}
)

func (m *DocMsgWithdrawEarnedFees) GetType() string {
	return MsgTypeWithdrawEarnedFees
}

func (m *DocMsgWithdrawEarnedFees) BuildMsg(v interface{}) {
	msg := v.(*MsgWithdrawEarnedFees)

	m.Owner = msg.Owner
	m.Provider = msg.Provider
}

func (m *DocMsgWithdrawEarnedFees) HandleTxMsg(v SdkMsg) MsgDocInfo {
	var (
		addrs []string
		msg   MsgWithdrawEarnedFees
	)

	ConvertMsg(v, &msg)
	addrs = append(addrs, msg.Owner, msg.Provider)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(v, handler)
}
