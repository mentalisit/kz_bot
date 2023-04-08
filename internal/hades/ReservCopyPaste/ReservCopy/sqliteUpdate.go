package ReservCopy

import "fmt"

func (r *ReservDB) UpdateAllianceChat(m []Chat) {
	stmt, err := r.db.Prepare("UPDATE AllianceChat SET MesId = ? WHERE CorpName = ?")
	if err != nil {
		fmt.Println(err)
		// Обработка ошибки
	}
	defer stmt.Close()

	for _, c := range m {
		resault, err := stmt.Exec(c.MesId, c.CorpName)
		if err != nil {
			fmt.Println(err)
			// Обработка ошибки
		}
		res, _ := resault.RowsAffected()
		if res == 0 {
			stmtInsert, err := r.db.Prepare("INSERT INTO AllianceChat(MesId, CorpName) VALUES(?, ?)")
			if err != nil {
				// Обработка ошибки
			}
			defer stmtInsert.Close()

			_, err = stmtInsert.Exec(c.MesId, c.CorpName)
			if err != nil {
				// Обработка ошибки
			}
		}
	}
}
func (r *ReservDB) UpdateWs1Chat(m []Chat) {
	stmt, err := r.db.Prepare("UPDATE Ws1Chat SET MesId = ? WHERE CorpName = ?")
	if err != nil {
		fmt.Println(err)
		// Обработка ошибки
	}
	defer stmt.Close()

	for _, c := range m {
		resault, err := stmt.Exec(c.MesId, c.CorpName)
		if err != nil {
			fmt.Println(err)
			// Обработка ошибки
		}
		res, _ := resault.RowsAffected()
		if res == 0 {
			stmtInsert, err := r.db.Prepare("INSERT INTO Ws1Chat(MesId, CorpName) VALUES(?, ?)")
			if err != nil {
				// Обработка ошибки
			}
			defer stmtInsert.Close()

			_, err = stmtInsert.Exec(c.MesId, c.CorpName)
			if err != nil {
				// Обработка ошибки
			}
		}
	}
}
func (r *ReservDB) UpdateMember(m []Member) {
	stmt, err := r.db.Prepare("UPDATE AllianceMember SET rang = ? WHERE CorpName = ? AND UserName = ?")
	if err != nil {
		fmt.Println(err)
		// Обработка ошибки
	}
	defer stmt.Close()

	for _, c := range m {
		resault, err := stmt.Exec(c.Rang, c.CorpName, c.UserName)
		if err != nil {
			fmt.Println(err)
			// Обработка ошибки
		}
		res, _ := resault.RowsAffected()
		if res == 0 {
			stmtInsert, err := r.db.Prepare("INSERT INTO AllianceMember (CorpName,UserName,rang) VALUES(?, ?, ?)")
			if err != nil {
				// Обработка ошибки
			}
			defer stmtInsert.Close()

			_, err = stmtInsert.Exec(c.CorpName, c.UserName, c.Rang)
			if err != nil {
				// Обработка ошибки
			}
		}
	}
}
