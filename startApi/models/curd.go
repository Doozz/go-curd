package models

import (
	"fmt"
	"startApi/config"
)

type Data struct {
	ID      int
	Name string
    	Age   int
    	From    string
    	Sex int
    	Memo    string
}

func (m *Data) GetOne(id int) error {
	err := config.DBHandle.QueryRow("SELECT id, name FROM tt WHERE id = ? AND status = ? LIMIT 1", id, 1).
		Scan(&m.ID, &m.Name)

	if err != nil {
		return err
	}

	return nil
}

func (m *Data) GetAll() (res []map[string]string, err error) {
	var (
		id string
		name string
	)
	rows, err := config.DBHandle.Query("SELECT id, name WHERE status = ? FROM tt ORDER BY id DESC", 1)

	if err != nil {
		return nil,err
	}

	for rows.Next() {
		err = rows.Scan(&id, &name)
		if err != nil {
			return nil,err
		}
		res = append(res, map[string]string{
			"id": id,
			"name":  name,
		})
	}
	return res,nil
}


func Create(in *Data ) int {
	rs, err := config.DBHandle.Exec(`
	INSERT INTO tt SET
	name = ?,
	memo = ?,
	age = ?,
	status = 1,
	sex = ?,
	address = ?
	`,
		in.Name,
		in.Memo,
		in.Age,
		in.Sex,
		in.From,
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
		memo = ?,
		age = ?,
		sex = ?,
		address = ?
		WHERE id = ?
		`,
		in.Name,
		in.Memo,
		in.Age,
		in.Sex,
		in.From,
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

func Delete(in *Data) int{
	rs, err := config.DBHandle.Exec(`UPDATE tt SET
		status = 3
		WHERE id = ?
		`,
		in.ID,
		)

	if err != nil {
		fmt.Printf("%s", err)
		return -1
	}
	affectId, err := rs.RowsAffected()
	if err != nil {
		fmt.Printf("%s", err)
		return -1
	}
	return int(affectId)
}

func Bussness(){
	Tx, _ := config.DBHandle.Begin()



	err := Tx.Commit()
	if err != nil {
		defer Tx.Rollback()
		return err
	}
	return nil
}
