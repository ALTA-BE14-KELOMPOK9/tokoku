package transaksi

import (
	"database/sql"
	"errors"
	"log"
)

// Model Transaksi
type Transaksi struct {
	ID           int
	IDPegawai    int
	IDCustomer   int
	NamaPegawai  string
	NamaCustomer string
	CreatedDate  string
}

// Model Transaksi Barang
type TransaksiBarang struct {
	IDTransaksi int
	IDBarang    int
	NamaBarang  string
	Quantity    int
}

// Menu Transaksi
type MenuTransaksi struct {
	DB *sql.DB
}

// Method list transaksi
func (mt *MenuTransaksi) ListTransaksi() ([]Transaksi, error) {
	stmt, err := mt.DB.Prepare("select * from transaksi")
	if err != nil {
		log.Println("Prepare list transaksi gagal: ", err.Error())
		return nil, errors.New("prepare list gagal")
	}

	rows, err := stmt.Query()
	if err != nil {
		log.Println("Query list transaksi gagal: ", err.Error())
		return nil, errors.New("gagal menampilkan list")
	}

	var listTransaksi []Transaksi
	for rows.Next() {
		transaksi := Transaksi{}
		err = rows.Scan(&transaksi.ID, &transaksi.IDPegawai, &transaksi.IDCustomer, &transaksi.CreatedDate)
		if err != nil {
			log.Println("Scan List transaksi gagal: ", err.Error())
			return nil, errors.New("data tidak ditemukan")
		}

		listTransaksi = append(listTransaksi, transaksi)
	}

	return listTransaksi, nil
}

func (mt *MenuTransaksi) ListTransaction() ([]Transaksi, error) {
	stmt, err := mt.DB.Prepare("SELECT transaksi.id_transaksi, pegawai.id_pegawai, pegawai.username, customer.id_customer, customer.username, transaksi.created_date FROM transaksi JOIN pegawai ON pegawai.id_pegawai = transaksi.id_pegawai JOIN customer ON customer.id_customer = transaksi.id_customer ORDER BY transaksi.created_date ASC")
	if err != nil {
		log.Println("PREPARE LIST TRANSACTION ERROR: ", err.Error())
		return nil, errors.New("prepare list transaksi gagal")
	}

	rows, err := stmt.Query()
	if err != nil {
		log.Println("QUERY LIST TRANSACTION ERROR: ", err.Error())
		return nil, errors.New("gagal menampilkan list nota  transaksi")
	}

	var listTransaksi []Transaksi
	for rows.Next() {
		transaksi := Transaksi{}
		err := rows.Scan(&transaksi.ID, &transaksi.IDPegawai, &transaksi.NamaPegawai, &transaksi.IDCustomer, &transaksi.NamaCustomer, &transaksi.CreatedDate)
		if err != nil {
			return nil, err
		}

		listTransaksi = append(listTransaksi, transaksi)
	}

	return listTransaksi, nil
}

func (mt *MenuTransaksi) ListItemTransaction(id int) ([]TransaksiBarang, error) {
	stmt, err := mt.DB.Prepare("SELECT tb.id_transaksi, b.id_barang, b.nama, tb.quantity FROM transaksi_barang tb JOIN barang b ON b.id_barang = tb.id_barang WHERE tb.id_transaksi = ?")
	if err != nil {
		log.Println("PREPARE LIST ITEM TRANSACTION ERROR: ", err.Error())
		return nil, errors.New("prepare list transaksi barang gagal")
	}

	rows, err := stmt.Query(id)
	if err != nil {
		log.Println("QUERY LIST ITEM TRANSACTION ERROR: ", err.Error())
		return nil, errors.New("gagal menampilkan list nota  transaksi")
	}

	var listTransaksiBarang []TransaksiBarang
	for rows.Next() {
		transaksiBarang := TransaksiBarang{}
		err := rows.Scan(&transaksiBarang.IDTransaksi, &transaksiBarang.IDBarang, &transaksiBarang.NamaBarang, &transaksiBarang.Quantity)
		if err != nil {
			return nil, err
		}

		listTransaksiBarang = append(listTransaksiBarang, transaksiBarang)
	}

	return listTransaksiBarang, nil
}

func (mt *MenuTransaksi) ListTransaksiBarang() ([]TransaksiBarang, error) {
	stmt, err := mt.DB.Prepare("select * from transaksi_barang")
	if err != nil {
		log.Println("Prepare list gagal: ", err.Error())
		return nil, errors.New("prepare list gagal")
	}

	rows, err := stmt.Query()
	if err != nil {
		log.Println("Query list gagal: ", err.Error())
		return nil, errors.New("gagal menampilkan list")
	}

	var listTransaksiBarang []TransaksiBarang
	for rows.Next() {
		transaksiBarang := TransaksiBarang{}
		err = rows.Scan(&transaksiBarang.IDTransaksi, &transaksiBarang.IDBarang, &transaksiBarang.NamaBarang)
		if err != nil {
			log.Println("Scan List transaksi gagal: ", err.Error())
			return nil, errors.New("data tidak ditemukan")
		}

		listTransaksiBarang = append(listTransaksiBarang, transaksiBarang)
	}

	return listTransaksiBarang, nil
}

func (mt *MenuTransaksi) TambahTransaksi(idPegawai int, idCustomer int) (int, error) {
	stmt, err := mt.DB.Prepare("INSERT INTO transaksi(id_pegawai, id_customer) VALUES(?, ?)")
	if err != nil {
		log.Println("PREPARE TAMBAH TRANSAKSI ERROR: ", err.Error())
		return 0, errors.New("prepare tambah transaksi gagal")
	}

	result, err := stmt.Exec(idPegawai, idCustomer)
	if err != nil {
		log.Println("EXEC TAMBAH TRANSAKSI ERROR: ", err.Error())
		return 0, errors.New("gagal menambahkan transaksi")
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, errors.New("tambah transaksi gagal")
	}

	return int(id), nil
}

func (mt *MenuTransaksi) TambahTransaksiBarang(idTransaksi int, idBarang int, quantity int) error {
	stmt, err := mt.DB.Prepare("INSERT INTO transaksi_barang(id_transaksi, id_barang, quantity) VALUES(?, ?, ?)")
	if err != nil {
		log.Println("PREPARE TAMBAH TRANSAKSI BARANG ERROR: ", err.Error())
		return errors.New("prepare tambah transaksi barang gagal")
	}

	_, err = stmt.Exec(idTransaksi, idBarang, quantity)
	if err != nil {
		log.Println("EXEC TAMBAH TRANSAKSI BARANG ERROR: ", err.Error())
		return errors.New("gagal menambahkan transaksi barang")
	}
	return nil
}

func (mt *MenuTransaksi) HapusTransaksi(id_transaksi int) (int, error) {

	stmt, err := mt.DB.Prepare("delete from transaksi where id_transaksi=?")
	if err != nil {
		log.Println("Hapus Transaksi gagal: ", err.Error())
		return 0, errors.New("gagal hapus transaksi")
	}

	result, err := stmt.Exec(id_transaksi)
	if err != nil {
		log.Println("Gagal hapus data", err.Error())
	}

	rowsAffected, err := result.RowsAffected()
	return int(rowsAffected), nil
}

func (mt *MenuTransaksi) HapusTransaksiBarang(id_transaksi int) (int, error) {

	stmt, err := mt.DB.Prepare("delete from transaksi_barang where id_transaksi=?")
	if err != nil {
		log.Println("Hapus gagal: ", err.Error())
		return 0, errors.New("gagal hapus")
	}

	result, err := stmt.Exec(id_transaksi)
	if err != nil {
		log.Println("Gagal hapus data", err.Error())
	}

	rowsAffected, err := result.RowsAffected()
	return int(rowsAffected), nil
}
