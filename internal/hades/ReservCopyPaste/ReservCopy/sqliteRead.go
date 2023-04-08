package ReservCopy

import "fmt"

func (r *ReservDB) ReadAllianceChat() []Chat {
	rows, err := r.db.Query("SELECT * FROM AllianceChat")
	if err != nil {
		return nil
	}
	defer rows.Close()
	var alliace []Chat
	for rows.Next() {
		var a Chat
		if err := rows.Scan(&a.id, &a.MesId, &a.CorpName); err != nil {
			fmt.Println(err)
		}
		alliace = append(alliace, a)
	}
	return alliace
}
func (r *ReservDB) ReadWs1Chat() []Chat {
	rows, err := r.db.Query("SELECT * FROM Ws1Chat")
	if err != nil {
		return nil
	}
	defer rows.Close()
	var Ws1 []Chat
	for rows.Next() {
		var a Chat
		if err := rows.Scan(&a.id, &a.MesId, &a.CorpName); err != nil {
			fmt.Println(err)
		}
		Ws1 = append(Ws1, a)
	}
	return Ws1
}
func (r *ReservDB) ReadMember() []Member {
	rows, err := r.db.Query("SELECT * FROM AllianceMember")
	if err != nil {
		return nil
	}
	defer rows.Close()
	var Members []Member
	for rows.Next() {
		var a Member
		if err := rows.Scan(&a.id, &a.CorpName, &a.UserName, &a.Rang); err != nil {
			fmt.Println(err)
		}
		Members = append(Members, a)
	}
	return Members
}
