package models

import (
	"net/http"

	"example.com/m/db"
)

type Pegawai struct {
	Id       int    `json:"id"`
	Nama     string `json:"nama"`
	Alamat   string `json:"alamat"`
	Telepon  string `json:"telepon"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func FetchAllPegawai() (Response, error) {

	var obj Pegawai
	var arrobj []Pegawai
	var res Response

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM pegawai"

	rows, err := con.Query(sqlStatement)

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.Id, &obj.Nama, &obj.Alamat, &obj.Telepon, &obj.Username, &obj.Password)
		if err != nil {
			return res, err
		}
		arrobj = append(arrobj, obj)
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = arrobj

	return res, nil

}

func StorePegawai(nama string, alamat string, telepon string, username string, password string) (Response, error) {

	var res Response
	con := db.CreateCon()

	sqlStatement := "INSERT pegawai (nama, alamat, telepon, username, password) VALUES ( ?, ?, ?, ?, ?)"

	stmt, err := con.Prepare(sqlStatement)

	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(nama, alamat, telepon)

	if err != nil {
		return res, err
	}

	lastInsertedId, err := result.LastInsertId()

	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int64{
		"last_inserted_id": lastInsertedId,
	}

	return res, nil
}

func UpdatePegawai(nama string, alamat string, telepon string, username string, password string, id int) (Response, error) {

	var res Response
	con := db.CreateCon()

	sqlStatement := "UPDATE pegawai SET nama=?, alamat =? , telepon =?, username= ?, password = ?WHERE id = ?"

	stmt, err := con.Prepare(sqlStatement)

	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(nama, alamat, telepon, username, password, id)

	if err != nil {
		return res, err
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int64{
		"rows_affected": rowsAffected,
	}

	return res, nil
}

func DeletePegawai(id int) (Response, error) {

	var res Response

	con := db.CreateCon()

	sqlStatement := "DELETE FROM pegawai WHERE id = ?"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(id)
	if err != nil {
		return res, err
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int64{
		"rows_affected": rowsAffected,
	}
	return res, nil

}
