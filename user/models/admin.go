package models

import (
	"fmt"
	"user/config"
)

type Data struct {
	ID int
	Name string
	Mobile string
    	Email string
    	Status int
    	Password string
    	Createtime int64
    	Logintime int64
}

	func (m *Data) GetOne(name string) error {
		err := config.DBHandle.QueryRow("SELECT id,name,password FROM tt WHERE name = ? AND status = ? LIMIT 1", name, 1).
			Scan(&m.ID, &m.Name, &m.Password)

		if err != nil {
			return err
		}

		return nil
	}

	func (m *Data) GetRowById(id string) error {
		err := config.DBHandle.QueryRow("SELECT id,name,mobile,email FROM tt WHERE id = ? AND status = ? LIMIT 1", id, 1).
			Scan(&m.ID, &m.Name, &m.Mobile, &m.Email)

		if err != nil {
			return err
		}

		return nil
	}

	// func (m *Data) GetAll() (res []map[string]string, err error) {
	// 	var (
	// 		id string
	// 		name string
	// 	)
	// 	rows, err := config.DBHandle.Query("SELECT id, name WHERE status = ? FROM tt ORDER BY id DESC", 1)

	// 	if err != nil {
	// 		return nil,err
	// 	}

	// 	for rows.Next() {
	// 		err = rows.Scan(&id, &name)
	// 		if err != nil {
	// 			return nil,err
	// 		}
	// 		res = append(res, map[string]string{
	// 			"id": id,
	// 			"name":  name,
	// 		})
	// 	}
	// 	return res,nil
	// }


	func Create(in *Data ) int {
		rs, err := config.DBHandle.Exec(`
		INSERT INTO tt SET
		name = ?,
		status = 1,
		mobile = ?,
		email = ?,
		password = ?,
		createtime = ?
		`,
			in.Name,
			in.Mobile,
			in.Email,
			in.Password,
			in.Createtime,
		)

		if err != nil {
			fmt.Printf("%s", err)
			return -1
		}

		insertId, err := rs.LastInsertId()
		if err != nil {
			fmt.Printf("%s", err)
			return -1
		}

		return int(insertId)
	}


func Update(in *Data) error{
	rs, err := config.DBHandle.Exec(`UPDATE tt SET
		name = ?,
		mobile = ?,
		email = ?
		WHERE id = ?
		`,
		in.Name,
		in.Mobile,
		in.Email,
		in.ID,
		)

	if err != nil {
		return err
	}
	_, err = rs.RowsAffected()
	if err != nil {
		return err
	}
	return nil
}

// func Delete(in *Data) int{
// 	rs, err := config.DBHandle.Exec(`UPDATE tt SET
// 		status = 3
// 		WHERE id = ?
// 		`,
// 		in.ID,
// 		)

// 	if err != nil {
// 		fmt.Printf("%s", err)
// 		return -1
// 	}
// 	affectId, err := rs.RowsAffected()
// 	if err != nil {
// 		fmt.Printf("%s", err)
// 		return -1
// 	}
// 	return int(affectId)
// }

// func Bussness(){
// 	Tx, _ := config.DBHandle.Begin()



// 	err := Tx.Commit()
// 	if err != nil {
// 		defer Tx.Rollback()
// 		return err
// 	}
// 	return nil
// }
