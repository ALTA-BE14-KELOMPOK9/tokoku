package transaksi

import (
	"database/sql"
	"errors"
	"log"
	"time"
)

// Model Transaksi
type Transaksi struct {
	ID          int
	IDPegawai   int
	IDCustomer  int
	CreatedDate time.Time
}

// Model Transaksi Barang
type TransaksiBarang struct {
	IDTransaksi int
	IDBarang    int
}

// Model Nota
type Nota struct {
	IDTransaksi  int
	NamaPegawai  string
	NamaCustomer string
	CreatedDate  time.Time
	Barang       []string
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

func (mt *MenuTransaksi) TambahTransaksiBarang(idTransaksi int, idBarang int) error {
	stmt, err := mt.DB.Prepare("INSERT INTO transaksi_barang(id_transaksi, id_barang) VALUES(?, ?)")
	if err != nil {
		log.Println("PREPARE TAMBAH TRANSAKSI BARANG ERROR: ", err.Error())
		return errors.New("prepare tambah transaksi barang gagal")
	}

	_, err = stmt.Exec(idTransaksi, idBarang)
	if err != nil {
		log.Println("EXEC TAMBAH TRANSAKSI BARANG ERROR: ", err.Error())
		return errors.New("gagal menambahkan transaksi barang")
	}

	// rowAffected, err := result.RowsAffected()
	// if err != nil {
	// 	return errors.New("tambah transaksi barang gagal")
	// }

	return nil
}

func (mt *MenuTransaksi) ListNotaTransaksi() ([]Nota, error) {
	stmtNota, err := mt.DB.Prepare("SELECT transaksi.id_transaksi, pegawai.username, customer.username, transaksi.created_date FROM transaksi_barang JOIN transaksi  ON transaksi.id_transaksi = transaksi_barang.id_transaksi JOIN barang  ON barang.id_barang = transaksi_barang.id_barang JOIN pegawai ON pegawai.id_pegawai = transaksi.id_pegawai JOIN customer ON customer.id_customer = transaksi.id_customer WHERE transaksi.id_transaksi = ?")
	if err != nil {
		log.Println("PREPARE LIST NOTA TRANSAKSI ERROR: ", err.Error())
		return nil, errors.New("prepare list nota transaksi gagal")
	}

	stmtIDTransaksi, err := mt.DB.Prepare("SELECT id_transaksi FROM transaksi")
	if err != nil {
		log.Println("PREPARE ID TRANSAKSI ERROR: ", err.Error())
		return nil, errors.New("prepare id transaksi gagal")
	}

	stmtBarang, err := mt.DB.Prepare("SELECT barang.nama FROM transaksi_barang JOIN transaksi ON transaksi.id_transaksi = transaksi_barang.id_transaksi JOIN barang ON barang.id_barang = transaksi_barang.id_barang WHERE transaksi.id_transaksi = ?")
	if err != nil {
		log.Println("PREPARE BARANG TRANSAKSI ERROR: ", err.Error())
		return nil, errors.New("prepare barang transaksi gagal")
	}

	// QUERY ID TRANSAKSI
	rowsIDTransaksi, err := stmtIDTransaksi.Query()
	if err != nil {
		log.Println("QUERY ID TRANSAKSI ERROR: ", err.Error())
		return nil, errors.New("gagal menampilkan id transaksi")
	}

	var listNota []Nota

	for rowsIDTransaksi.Next() {
		var id int
		// SCAN ID TRANSAKSI
		err := rowsIDTransaksi.Scan(&id)
		if err != nil {
			return nil, err
		}

		// QUERY NOTA
		rowsNota, err := stmtNota.Query(id)
		if err != nil {
			log.Println("QUERY LIST NOTA TRANSAKSI ERROR: ", err.Error())
			return nil, errors.New("gagal menampilkan list nota  transaksi")
		}

		nota := Nota{}
		for rowsNota.Next() {
			// SCAN NOTA
			err := rowsNota.Scan(&nota.IDTransaksi, &nota.NamaPegawai, &nota.NamaCustomer, &nota.CreatedDate)
			if err != nil {
				return nil, err
			}
		}

		// QUERY BARANG
		rowsStmtBarang, err := stmtBarang.Query(id)
		if err != nil {
			log.Println("QUERY BARANG TRANSAKSI ERROR: ", err.Error())
			return nil, errors.New("gagal menampilkan BARANG transaksi")
		}

		for rowsStmtBarang.Next() {
			var barang string
			// SCAN BARANG
			rowsStmtBarang.Scan(&barang)
			nota.Barang = append(nota.Barang, barang)
		}

		listNota = append(listNota, nota)

	}

	return listNota, nil
}

// Method hapus transaksi
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
