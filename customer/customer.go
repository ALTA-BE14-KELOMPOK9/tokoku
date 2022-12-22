package customer

import (
	"database/sql"
	"errors"
	"log"
)

// Model Customer
type Customer struct {
	ID        int
	IDPegawai int
	Nama      string
}

// Menu Customer
type MenuCustomer struct {
	DB *sql.DB
}

func (mc *MenuCustomer) ListCustomer() ([]Customer, error) {
	stmt, err := mc.DB.Prepare("SELECT * FROM customer")
	if err != nil {
		log.Println("PREPARE LIST CUSTOMER STATEMENT ERROR: ", err.Error())
		return nil, errors.New("prepare list customer gagal")
	}

	rows, err := stmt.Query()
	if err != nil {
		log.Println("QUERY LIST CUSTOMER STATEMENT ERROR: ", err.Error())
		return nil, errors.New("gagal menampilkan customer")
	}

	var listCustomer []Customer
	for rows.Next() {
		customer := Customer{}
		err = rows.Scan(&customer.ID, &customer.IDPegawai, &customer.Nama)
		if err != nil {
			log.Println("SCAN LIST CUSTOMER ERROR: ", err.Error())
			return nil, errors.New("data customer tidak ditemukan")
		}

		listCustomer = append(listCustomer, customer)
	}

	return listCustomer, nil
}

func (mc *MenuCustomer) CheckDuplicate(nama string) bool {
	row := mc.DB.QueryRow("SELECT id_customer FROM customer WHERE username = ?", nama)

	var id int
	if err := row.Scan(&id); err != nil {
		return false
	}

	return true
}

func (mc *MenuCustomer) TambahCustomer(customer Customer) (int, error) {
	// Check apakah nama customer duplikat
	if mc.CheckDuplicate(customer.Nama) {
		log.Println("CUSTOMER DIPLICATE")
		return 0, errors.New("customer sudah terdaftar")
	}

	stmt, err := mc.DB.Prepare("INSERT INTO customer(id_pegawai, username) VALUES(?, ?)")
	if err != nil {
		log.Println("PREPARE TAMBAH CUSTOMER STATEMENT ERROR: ", err.Error())
		return 0, errors.New("prepare tambah customer gagal")
	}

	result, err := stmt.Exec(customer.IDPegawai, customer.Nama)
	if err != nil {
		log.Println("EXEC TAMBAH CUSTOMER STATEMENT ERROR: ", err.Error())
		return 0, errors.New("tambah customer gagal")
	}

	lastInsertId, err := result.LastInsertId()
	if err != nil {
		return 0, errors.New("tambah customer gagal")
	}

	return int(lastInsertId), nil
}

func (mc *MenuCustomer) CariCustomer(name string) (int, error) {
	stmt, err := mc.DB.Prepare("SELECT id_customer FROM customer WHERE username = ?")
	if err != nil {
		log.Println("PREPARE FIND CUSTOMER STATEMENT ERROR: ", err.Error())
		return 0, errors.New("gagal find customere")
	}

	var id int
	err = stmt.QueryRow(name).Scan(&id)
	if err != nil {
		log.Println("SCAN FIND CUSTOMER STATEMENT ERROR: ", err.Error())
		return 0, errors.New("data customer tidak ditemukan")
	}

	return id, nil
}

// Method hapus customer
func (hc *MenuCustomer) HapusCustomer(id_customer int) (int, error) {

	stmt, err := hc.DB.Prepare("delete from customer where id_customer=?")
	if err != nil {
		log.Println("Hapus Barang gagal: ", err.Error())
		return 0, errors.New("gagal hapus barang")
	}

	result, err := stmt.Exec(id_customer)
	if err != nil {
		log.Println("Gagal hapus data", err.Error())
	}

	rowsAffected, err := result.RowsAffected()
	return int(rowsAffected), nil

}
