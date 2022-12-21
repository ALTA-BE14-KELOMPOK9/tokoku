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
