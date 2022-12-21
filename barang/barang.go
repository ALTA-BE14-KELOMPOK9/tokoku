package barang

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"
)

// Model Barang
type Barang struct {
	ID          int
	IDPegawai   int
	Nama        string
	Stok        int
	NamaPegawai string
	CreatedDate time.Time
}

// Menu Barang
type MenuBarang struct {
	DB *sql.DB
}

func (mb *MenuBarang) ListBarang() ([]Barang, error) {
	stmt, err := mb.DB.Prepare("SELECT barang.id_barang, barang.id_pegawai, barang.nama, barang.quantity, pegawai.username, barang.created_date FROM barang JOIN pegawai ON pegawai.id_pegawai = barang.id_pegawai")
	if err != nil {
		log.Println("PREPARE LIST BARANG STATEMENT ERROR: ", err.Error())
		return nil, errors.New("prepare list barang gagal")
	}

	rows, err := stmt.Query()
	if err != nil {
		log.Println("QUERY LIST BARANG STATEMENT ERROR: ", err.Error())
		return nil, errors.New("gagal menampilkan barang")
	}

	var listBarang []Barang
	for rows.Next() {
		barang := Barang{}
		err = rows.Scan(&barang.ID, &barang.IDPegawai, &barang.Nama, &barang.Stok, &barang.NamaPegawai, &barang.CreatedDate)
		if err != nil {
			log.Println("SCAN LIST BARANG ERROR: ", err.Error())
			return nil, errors.New("data barang tidak ditemukan")
		}

		listBarang = append(listBarang, barang)
	}

	return listBarang, nil
}

func (mb *MenuBarang) CheckDuplicate(namaBarang string) bool {

	row := mb.DB.QueryRow("SELECT id_barang FROM barang WHERE nama = ?", namaBarang)

	var id int
	if err := row.Scan(&id); err != nil {
		return false
	}

	return true
}

func (mb *MenuBarang) TambahBarang(barang Barang) (int, error) {
	// Check apakah barang duplikat
	if mb.CheckDuplicate(barang.Nama) {
		log.Println("BARANG DIPLICATE")
		return 0, errors.New("barang sudah pernah terdaftar")
	}

	stmt, err := mb.DB.Prepare("INSERT INTO barang(id_pegawai, nama, quantity) VALUES(?, ?, ?)")
	if err != nil {
		log.Println("PREPARE TAMBAH BARANG STATEMENT ERROR: ", err.Error())
		return 0, errors.New("prepare tambah barang gagal")
	}

	result, err := stmt.Exec(barang.IDPegawai, barang.Nama, barang.Stok)
	if err != nil {
		log.Println("EXEC TAMBAH BARANG STATEMENT ERROR: ", err.Error())
		return 0, errors.New("tambah barang gagal")
	}

	rowAffected, err := result.RowsAffected()
	if err != nil {
		return 0, errors.New("tambah barang gagal")
	}

	return int(rowAffected), nil
}

func (mb *MenuBarang) UbahNamaBarang(id int, nama string) (int, error) {

	stmt, err := mb.DB.Prepare("UPDATE barang SET nama = ? WHERE id_barang = ?")
	if err != nil {
		log.Println("PREPARE UBAH NAMA BARANG STATEMENT ERROR: ", err.Error())
		return 0, errors.New("prepare ubah nama barang gagal")
	}

	result, err := stmt.Exec(nama, id)
	if err != nil {
		log.Println("EXEC UBAH NAMA BARANG STATEMENT ERROR: ", err.Error())
		return 0, errors.New("ubah nama barang gagal")
	}

	rowAffected, err := result.RowsAffected()
	if err != nil {
		return 0, errors.New("ubah nama barang gagal")
	}

	return int(rowAffected), nil
}

func (mb *MenuBarang) UbahStokBarang(id int, quantity int, condition string) (int, error) {

	if condition == "tambah" {
		condition = "+"
	} else if condition == "kurang" {
		condition = "-"
	} else {
		return 0, errors.New("salah input condition")
	}

	query := fmt.Sprintf("UPDATE barang SET quantity = quantity %s ? WHERE id_barang = ?", condition)
	stmt, err := mb.DB.Prepare(query)
	if err != nil {
		log.Println("PREPARE UBAH UPDATE BARANG STATEMENT ERROR: ", err.Error())
		return 0, errors.New("prepare ubah update barang gagal")
	}

	result, err := stmt.Exec(quantity, id)
	if err != nil {
		log.Println("EXEC UBAH UPDATE BARANG STATEMENT ERROR: ", err.Error())
		return 0, errors.New("ubah update barang gagal")
	}

	rowAffected, err := result.RowsAffected()
	if err != nil {
		return 0, errors.New("ubah update barang gagal")
	}

	return int(rowAffected), nil
}
