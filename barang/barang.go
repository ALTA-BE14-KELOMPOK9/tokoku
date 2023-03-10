package barang

import (
	"database/sql"
	"errors"
	"log"
)

// Model Barang
type Barang struct {
	ID          int
	IDPegawai   int
	Nama        string
	Stok        int
	NamaPegawai string
	CreatedDate string
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
		log.Println("BARANG DUPLICATE")
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

func (mb *MenuBarang) TambahStokBarang(id int, quantity int) (int, error) {
	stmt, err := mb.DB.Prepare("UPDATE barang SET quantity = quantity + ? WHERE id_barang = ?")
	if err != nil {
		log.Println("PREPARE TAMBAH STOK BARANG STATEMENT ERROR: ", err.Error())
		return 0, errors.New("prepare tambah stok barang gagal")
	}

	result, err := stmt.Exec(quantity, id)
	if err != nil {
		log.Println("EXEC TAMBAH STOK BARANG STATEMENT ERROR: ", err.Error())
		return 0, errors.New("tambah stok barang gagal")
	}

	rowAffected, err := result.RowsAffected()
	if err != nil {
		return 0, errors.New("tambah stok barang gagal")
	}

	return int(rowAffected), nil
}

func (mb *MenuBarang) KurangiStokBarang(id int, quantity int) (int, error) {
	stmt, err := mb.DB.Prepare("UPDATE barang SET quantity = quantity - ? WHERE id_barang = ?")
	if err != nil {
		log.Println("PREPARE KURANGI STOK BARANG STATEMENT ERROR: ", err.Error())
		return 0, errors.New("prepare kurangi stok barang gagal")
	}

	result, err := stmt.Exec(quantity, id)
	if err != nil {
		log.Println("EXEC KURANGI STOK BARANG STATEMENT ERROR: ", err.Error())
		return 0, errors.New("kurangi stok barang gagal")
	}

	rowAffected, err := result.RowsAffected()
	if err != nil {
		return 0, errors.New("kurangi stok barang gagal")
	}

	return int(rowAffected), nil
}

// Method cari barang
func (mb *MenuBarang) CariBarang(nama string) (int, error) {
	stmt, err := mb.DB.Prepare("SELECT id_barang FROM barang WHERE nama = ?")
	if err != nil {
		log.Println("PREPARE CARI BARANG STATEMENT ERROR: ", err.Error())
		return 0, errors.New("gagal cari barang")
	}

	var id int
	err = stmt.QueryRow(nama).Scan(&id)
	if err != nil {
		log.Println("SCAN CARI BARANG STATEMENT ERROR: ", err.Error())
		return 0, errors.New("data barang tidak ditemukan")
	}

	return id, nil
}

// Method hapus barang
func (mb *MenuBarang) HapusBarang(id_barang int) (int, error) {

	stmt, err := mb.DB.Prepare("delete from barang where id_barang=?")
	if err != nil {
		log.Println("Hapus Barang gagal: ", err.Error())
		return 0, errors.New("gagal hapus barang")
	}

	result, err := stmt.Exec(id_barang)
	if err != nil {
		log.Println("Gagal hapus data", err.Error())
	}

	rowsAffected, err := result.RowsAffected()
	return int(rowsAffected), nil

}
