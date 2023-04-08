package ReservCopy

import (
	"encoding/json"
	"fmt"
)

func (r *ReservDB) WriteToSQLite(data []byte) {
	var m Message
	err := json.Unmarshal(data, &m)
	if err != nil {
		fmt.Println(err)
		//return
	}
	r.UpdateAllianceChat(m.AllianceChat)
	r.UpdateMember(m.AllianceMember)
	r.UpdateWs1Chat(m.Ws1Chat)

}
