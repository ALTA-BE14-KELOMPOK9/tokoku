package main

import (
	"fmt"
	"tokoku/barang"
	"tokoku/config"
	"tokoku/customer"
	"tokoku/pegawai"
	"tokoku/transaksi"
)

func main() {

	// Setup database
	cfg := config.ReadConfig()
	db := config.OpenConnection(*cfg)

	// Setup Objek Menu
	menuPegawai := pegawai.MenuPegawai{DB: db}
	menuBarang := barang.MenuBarang{DB: db}
	menuCustomer := customer.MenuCustomer{DB: db}
	menuTransaksi := transaksi.MenuTransaksi{DB: db}

	inputMenu := 1
	isLogin := false

	// Data admin di-insert ke dalam database
	for inputMenu != 0 {
		fmt.Println()
		fmt.Println("===== Tokoku ======")
		fmt.Println("1. Login")
		fmt.Println("0. Exit")
		fmt.Print("Masukkan Input Menu: ")
		_, err := fmt.Scanf("%d\n", &inputMenu)
		if err != nil {
			fmt.Println(err)
		}

		if inputMenu == 1 {
			var username, password string
			fmt.Print("Masukkan Username: ")
			fmt.Scanln(&username)
			fmt.Print("Masukkan Password: ")
			fmt.Scanln(&password)

			employee, err := menuPegawai.Login(username, password)
			if err != nil {
				fmt.Println(err)
				continue
			}

			if employee.Username == "admin" {
				// Menu Admin
				isLogin := true
				for isLogin {
					fmt.Println()
					fmt.Println("===== Menu Admin ======")
					fmt.Println("1. Daftarkan Pegawai")
					fmt.Println("2. Hapus Data Pegawai")
					fmt.Println("3. Hapus Data Barang")
					fmt.Println("4. Hapus Data Transaksi")
					fmt.Println("5. Hapus Data Customer")
					fmt.Println("6. Hapus Data Transaksi Barang")
					fmt.Println("9. Logout")

					fmt.Print("Masukkan Input: ")
					fmt.Scanln(&inputMenu)

					switch inputMenu {

					case 1: // Mendaftarkan Pegawai
						pegawai := pegawai.Pegawai{}

						fmt.Print("Masukkan Username Pegawai: ")
						fmt.Scanln(&pegawai.Username)
						fmt.Print("Masukkan Password Pegawai: ")
						fmt.Scanln(&pegawai.Password)

						rowAffected, err := menuPegawai.DaftarPegawai(pegawai)
						if err != nil {
							fmt.Println(err)
							continue
						}

						if rowAffected > 0 {
							fmt.Println("Berhasil mendaftarkan pegawai baru!")
						}

					case 2: // List dan hapus data pegawai
						// List Pegawai
						listPegawai, err := menuPegawai.ListPegawai()
						if err != nil {
							fmt.Println(err)
						}

						fmt.Println()
						fmt.Println("List Pegawai")
						for _, pegawai := range listPegawai {
							fmt.Println("---------------------------------------")
							fmt.Println("ID Pegawai: ", pegawai.ID)
							fmt.Println("Nama Pegawai: ", pegawai.Username)
							fmt.Println("Password: ", pegawai.Password)
						}
						fmt.Println()
						fmt.Println("Total Data: ", len(listPegawai))

						// Fitur menghapus pegawai
						var id_pegawai int
						fmt.Print("Masukkan ID Pegawai yang akan dihapus: ")
						fmt.Scanln(&id_pegawai)

						hapusPegawai, err := menuPegawai.HapusPegawai(id_pegawai)
						if err != nil {
							fmt.Println(err)
						}

						if hapusPegawai > 0 {
							fmt.Println("Hapus pegawai berhasil")
						} else {
							fmt.Println("Hapus pegawai gagal")
						}

					case 3: // List dan Hapus Data Barang
						// List Barang
						listBarang, err := menuBarang.ListBarang()
						if err != nil {
							fmt.Println(err)
						}

						fmt.Println()
						fmt.Println("List Barang")
						for _, barang := range listBarang {
							fmt.Println("---------------------------------------")
							fmt.Println("ID Barang: ", barang.ID)
							fmt.Println("Nama Barang: ", barang.Nama)
							fmt.Println("Stok Barang: ", barang.Stok)
							fmt.Println("Nama Pegawai: ", barang.NamaPegawai)
							fmt.Println("Tanggal dibuat: ", barang.CreatedDate)
						}
						fmt.Println()
						fmt.Println("Total Data: ", len(listBarang))

						// Fitur menghapus barang
						var id_barang int
						fmt.Print("Masukkan ID Barang yang akan dihapus: ")
						fmt.Scanln(&id_barang)

						hapusBarang, err := menuBarang.HapusBarang(id_barang)
						if err != nil {
							fmt.Println(err)
						}

						if hapusBarang > 0 {
							fmt.Println("Hapus barang berhasil")
						} else {
							fmt.Println("Hapus barang gagal")
						}

					case 4: // List dan hapus data transaksi
						// List transaksi
						listTransaksi, err := menuTransaksi.ListTransaction()
						if err != nil {
							fmt.Println(err)
						}

						fmt.Println()
						fmt.Println("List Transaksi")
						for _, transaksi := range listTransaksi {
							fmt.Println("---------------------------------------")
							fmt.Println("ID Transaksi\t: ", transaksi.ID)
							fmt.Println("Nama Pegawai\t: ", transaksi.NamaPegawai)
							fmt.Println("Nama Customer\t: ", transaksi.NamaCustomer)
							fmt.Println("Tanggal dibuat\t: ", transaksi.CreatedDate)

						}
						fmt.Println("---------------------------------------")
						fmt.Println("Total Data: ", len(listTransaksi))

						// Fitur menghapus transaksi
						var id_transaksi int
						fmt.Print("Masukkan ID Transaksi yang akan dihapus: ")
						fmt.Scanln(&id_transaksi)

						hapusTransaksi, err := menuTransaksi.HapusTransaksi(id_transaksi)
						if err != nil {
							fmt.Println(err)
						}

						if hapusTransaksi > 0 {
							fmt.Println("Hapus transaksi berhasil")
						} else {
							fmt.Println("Hapus transaksi gagal")
						}

					case 5: // List dan Hapus data customer
						// List customer
						listCustomer, err := menuCustomer.ListCustomer()
						if err != nil {
							fmt.Println(err)
						}

						fmt.Println()
						fmt.Println("List Customer")
						for _, customer := range listCustomer {
							fmt.Println("-----------------------------------------------------")
							fmt.Println("ID Customer: ", customer.ID)
							fmt.Println("ID Pegawai: ", customer.IDPegawai)
							fmt.Println("Nama Customer: ", customer.Nama)
						}
						fmt.Println()
						fmt.Println("Total Data: ", len(listCustomer))

						// Hapus Customer
						var id_customer int
						fmt.Print("Masukkan ID customer yang akan dihapus: ")
						fmt.Scanln(&id_customer)

						hapusCustomer, err := menuCustomer.HapusCustomer(id_customer)
						if err != nil {
							fmt.Println(err)
						}

						if hapusCustomer > 0 {
							fmt.Println("Hapus customer berhasil")
						} else {
							fmt.Println("Hapus customer gagal")
						}

					case 6: // List dan Hapus data transaksi barang
						// List transaksi barang
						listTransaksiBarang, err := menuTransaksi.ListTransaksiBarang()
						if err != nil {
							fmt.Println(err)
						}

						fmt.Println()
						fmt.Println("List ")
						for _, transaksiBarang := range listTransaksiBarang {
							fmt.Println("-----------------------------------------------------")
							fmt.Println("ID Transaksi: ", transaksiBarang.IDTransaksi)
							fmt.Println("ID Barang: ", transaksiBarang.IDBarang)
						}
						fmt.Println()
						fmt.Println("Total Data: ", len(listTransaksiBarang))

						// Fitur menghapus transaksi barang
						var id_transaksi int
						fmt.Print("Masukkan ID Transaksi yang akan dihapus: ")
						fmt.Scanln(&id_transaksi)

						hapusTransaksi, err := menuTransaksi.HapusTransaksi(id_transaksi)
						if err != nil {
							fmt.Println(err)
						}

						if hapusTransaksi > 0 {
							fmt.Println("Hapus transaksi berhasil")
						} else {
							fmt.Println("Hapus transaksi gagal")
						}

					case 9: // Logout
						isLogin = false
						employee = pegawai.Pegawai{}
					}
				}

			} else {
				// MENU PEGAWAI
				isLogin = true
				for isLogin {
					fmt.Println()
					fmt.Println("===== MENU PEGAWAI ======")
					fmt.Println("1. Tampil List Barang")
					fmt.Println("2. Tambah Barang")
					fmt.Println("3. Ubah Nama Barang")
					fmt.Println("4. Ubah Stok Barang")
					fmt.Println("5. Tampil List Customer")
					fmt.Println("6. Tambah Customer")
					fmt.Println("7. Tampil List Nota")
					fmt.Println("8. Buat Nota Transaksi")
					fmt.Println("9. Logout")
					fmt.Print("Masukkan Input Menu: ")
					fmt.Scanln(&inputMenu)

					switch inputMenu {
					case 1: // Tampil List Barang
						listBarang, err := menuBarang.ListBarang()
						if err != nil {
							fmt.Println(err)
						}

						fmt.Println()
						fmt.Println("List Barang")
						for _, barang := range listBarang {
							fmt.Println("---------------------------------------")
							fmt.Println("ID Barang\t: ", barang.ID)
							fmt.Println("Nama Barang\t: ", barang.Nama)
							fmt.Println("Stok Barang\t: ", barang.Stok)
							fmt.Println("Nama Pegawai\t: ", barang.NamaPegawai)
							fmt.Println("Tanggal dibuat\t: ", barang.CreatedDate)
						}
						fmt.Println("---------------------------------------")
						fmt.Println("Total Data: ", len(listBarang))
						fmt.Print("Press enter to continue...")
						fmt.Scanln()

					case 2: // Tambah Barang
						product := barang.Barang{}

						fmt.Print("Masukkan Nama Barang: ")
						fmt.Scanln(&product.Nama)
						fmt.Print("Masukkan Jumlah Barang: ")
						fmt.Scanln(&product.Stok)
						product.IDPegawai = employee.ID

						res, err := menuBarang.TambahBarang(product)
						if err != nil {
							fmt.Println(err)
						}

						if res > 0 {
							fmt.Println("Berhasil menambahkan barang!")
						} else {
							fmt.Println("Gagal menambahkan barang!")
						}

					case 3: // Ubah Nama Barang
						var (
							nama string
							id   int
						)

						fmt.Print("Masukkan ID Barang: ")
						fmt.Scanln(&id)
						fmt.Print("Masukkan Nama Barang: ")
						fmt.Scanln(&nama)

						res, err := menuBarang.UbahNamaBarang(id, nama)
						if err != nil {
							fmt.Println(err)
						}

						if res > 0 {
							fmt.Println("Berhasil mengubah nama barang!")
						} else {
							fmt.Println("Gagal mengubah nama barang!")
						}

					case 4: // Ubah Stok Barang
						var (
							quantity  int
							id        int
							condition string
						)

						fmt.Print("Input (tambah / kurang) Stok: ")
						fmt.Scanln(&condition)
						fmt.Print("Masukkan ID Barang: ")
						fmt.Scanln(&id)
						fmt.Print("Masukkan Stok Barang: ")
						fmt.Scanln(&quantity)

						var res int
						if condition == "tambah" {
							res, err = menuBarang.TambahStokBarang(id, quantity)
							if err != nil {
								fmt.Println(err)
								continue
							}
						} else if condition == "kurang" {
							res, err = menuBarang.KurangiStokBarang(id, quantity)
							if err != nil {
								fmt.Println(err)
								continue
							}
						} else {
							fmt.Println("Only accept (kurang / tambah) input")
							continue
						}

						if res > 0 {
							fmt.Println("Berhasil mengubah stok barang!")
						} else {
							fmt.Println("Gagal mengubah stok barang!")
						}

					case 5: // Tampil List Customer
						listCustomer, err := menuCustomer.ListCustomer()
						if err != nil {
							fmt.Println(err)
						}

						fmt.Println()
						fmt.Println("List Customer")
						for _, customer := range listCustomer {
							fmt.Println("---------------------------------------")
							fmt.Println("ID Customer\t: ", customer.ID)
							fmt.Println("Nama Customer\t: ", customer.Nama)
						}
						fmt.Println("---------------------------------------")
						fmt.Println("Total Data: ", len(listCustomer))
						fmt.Print("Press enter to continue...")
						fmt.Scanln()

					case 6: // Tambah Customer
						inputCustomer := customer.Customer{}

						fmt.Print("Masukkan Nama Customer: ")
						fmt.Scanln(&inputCustomer.Nama)
						inputCustomer.IDPegawai = employee.ID

						res, err := menuCustomer.TambahCustomer(inputCustomer)
						if err != nil {
							fmt.Println(err)
						}

						if res > 0 {
							fmt.Println("Berhasil menambahkan customer!")
						} else {
							fmt.Println("Gagal menambahkan customer!")
						}

					case 7: // Tampil List Transaksi
						// Get all transaction
						listTransaction, err := menuTransaksi.ListTransaction()
						if err != nil {
							fmt.Println(err)
							continue
						}

						fmt.Println()
						fmt.Println("List Transaksi")
						for _, transaction := range listTransaction {
							fmt.Println("---------------------------------------")
							fmt.Println("ID Transaksi\t: ", transaction.ID)
							fmt.Println("Nama Pegawai\t: ", transaction.NamaPegawai)
							fmt.Println("Nama Customer\t: ", transaction.NamaCustomer)
							fmt.Println("Tanggal Dibuat\t: ", transaction.CreatedDate)

							// Get all item transaction
							listItemTransaction, err := menuTransaksi.ListItemTransaction(transaction.ID)
							if err != nil {
								break
							}

							fmt.Printf("No. Nama barang\tQty\n")
							for i, itemTransaction := range listItemTransaction {
								fmt.Printf("%d.  %s\t%d\n", i+1, itemTransaction.NamaBarang, itemTransaction.Quantity)
							}
						}
						fmt.Println("---------------------------------------")
						fmt.Println("Total Data: ", len(listTransaction))
						fmt.Print("Press enter to continue...")
						fmt.Scanln()

					case 8: // Membuat Transaksi
						transaction := transaksi.Transaksi{}

						fmt.Print("Masukkan Nama Customer: ")
						fmt.Scanln(&transaction.NamaCustomer)

						transaction.IDCustomer, err = menuCustomer.CariCustomer(transaction.NamaCustomer)
						if err != nil {
							fmt.Println(err)
							continue
						}
						transaction.IDPegawai = employee.ID

						// Membuat transaksi baru
						transaction.ID, err = menuTransaksi.TambahTransaksi(transaction)
						if err != nil {
							fmt.Println(err)
							continue
						}

						if transaction.ID <= 0 {
							fmt.Println("Gagal menambahkan transaksi barang!")
							continue
						}

						inputMenu = 1
						for inputMenu != 0 {
							fmt.Println("Tambahkan Barang: ")
							fmt.Println("1. Tambahkan barang")
							fmt.Println("0. Selesai")
							fmt.Print("Masukkan Input: ")
							fmt.Scanln(&inputMenu)

							if inputMenu == 1 {
								var (
									nama     string
									idBarang int
									quantity int
								)
								fmt.Print("Masukkan nama barang: ")
								fmt.Scanln(&nama)
								fmt.Print("Masukkan jumlah barang: ")
								fmt.Scanln(&quantity)

								// Cari id barang
								idBarang, err = menuBarang.CariBarang(nama)
								if err != nil {
									fmt.Println(err)
									continue
								}

								// Tambah transaksi barang dengan id transaksi dan id barang
								err = menuTransaksi.TambahTransaksiBarang(transaction.ID, idBarang, quantity)
								if err != nil {
									fmt.Println(err)
									continue
								}

								// Kurangi stok barang ketika transaksi barang berhasil
								_, err = menuBarang.KurangiStokBarang(idBarang, quantity)
								if err != nil {
									fmt.Println(err)
									continue
								}
							}
						}

						fmt.Println("Berhasil menambahkan transaksi barang!")

					case 9: // Logout
						isLogin = false
						employee = pegawai.Pegawai{}
					}
				}

			}
		}
	}
}
