package ReservCopy

import (
	"encoding/json"
)

type Message struct {
	AllianceChat   []Chat
	AllianceMember []Member
	Ws1Chat        []Chat
}

func (r *ReservDB) NewMessageWriteToPostgres() []byte {
	m := &Message{}
	m.AllianceChat = r.ReadAllianceChat()
	m.AllianceMember = r.ReadMember()
	m.Ws1Chat = r.ReadWs1Chat()
	marshal, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}
	return marshal
}
