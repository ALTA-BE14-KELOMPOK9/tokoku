package pegawai

import (
	"database/sql"
	"errors"
	"log"
)

// Model Pegawai
type Pegawai struct {
	ID       int
	Username string
	Password string
}

type MenuPegawai struct {
	DB *sql.DB
}

// Method Pegawai
func (mp *MenuPegawai) Login(username string, password string) (Pegawai, error) {
	stmt, err := mp.DB.Prepare("SELECT id_pegawai, username FROM pegawai WHERE username = ? AND password = ?")
	if err != nil {
		log.Println("PREPARE LOGIN STATEMENT ERROR: ", err.Error())
		return Pegawai{}, errors.New("prepare login gagal")
	}

	row := stmt.QueryRow(username, password)
	if row.Err() != nil {
		log.Println("QUERY ROW LOGIN ERROR: ", err.Error())
		return Pegawai{}, errors.New("query login gagal")
	}

	pegawai := Pegawai{}
	err = row.Scan(&pegawai.ID, &pegawai.Username)
	if err != nil {
		log.Println("SCAN LOGIN ERROR: ", err.Error())
		return Pegawai{}, errors.New("data pegawai tidak ditemukan")
	}

	return pegawai, nil
}

// Method List Pegawai
func (mp *MenuPegawai) ListPegawai() ([]Pegawai, error) {
	stmt, err := mp.DB.Prepare("select * from pegawai")
	if err != nil {
		log.Println("Prepare list pegawai gagal: ", err.Error())
		return []Pegawai{}, errors.New("prepare list gagal")
	}

	rows, err := stmt.Query()
	if err != nil {
		log.Println("Query list pegawai gagal: ", err.Error())
		return nil, errors.New("gagal menampilkan list")
	}

	var listPegawai []Pegawai
	for rows.Next() {
		pegawai := Pegawai{}
		err = rows.Scan(&pegawai.ID, &pegawai.Username, &pegawai.Password)
		if err != nil {
			log.Println("Scan List pegawai gagal: ", err.Error())
			return nil, errors.New("data tidak ditemukan")
		}

		listPegawai = append(listPegawai, pegawai)
	}

	return listPegawai, nil
}

// Method Hapus Pegawai
func (mp *MenuPegawai) HapusPegawai(id_pegawai int) (int, error) {

	stmt, err := mp.DB.Prepare("delete from pegawai where id_pegawai=?")
	if err != nil {
		log.Println("Hapus pegawai gagal: ", err.Error())
		return 0, errors.New("gagal hapus pegawai")
	}

	result, err := stmt.Exec(id_pegawai)
	if err != nil {
		log.Println("Gagal hapus data", err.Error())
	}

	rowsAffected, err := result.RowsAffected()
	return int(rowsAffected), nil

}
