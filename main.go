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

	// Setup Service
	menuPegawai := pegawai.MenuPegawai{DB: db}
	menuBarang := barang.MenuBarang{DB: db}
	menuCustomer := customer.MenuCustomer{DB: db}
	menuTransaksi := transaksi.MenuTransaksi{DB: db}

	inputMenu := 1
	isLogin := false

	for inputMenu != 0 {
		fmt.Println()
		fmt.Println("===== Tokoku ======")
		fmt.Println("1. Login")
		fmt.Println("0. Exit")
		fmt.Print("Masukkan Input Menu: ")
		fmt.Scanln(&inputMenu)

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

				// PEGAWAI ADMIN

			} else {
				// PEGAWAI BIASA
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
					fmt.Println("0. Logout")
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
							fmt.Println("-----------------------------------------------------")
							fmt.Println("ID Barang: ", barang.ID)
							fmt.Println("Nama Barang: ", barang.Nama)
							fmt.Println("Stok Barang: ", barang.Stok)
							fmt.Println("Nama Pegawai: ", barang.NamaPegawai)
							fmt.Println("Tanggal dibuat: ", barang.CreatedDate)
						}
						fmt.Println()
						fmt.Println("Total Data: ", len(listBarang))
						fmt.Print("Press enter to return to the menu...")
						fmt.Scanln()

					case 2: // Tambah Barang
						product := barang.Barang{}

						fmt.Print("Masukkan Nama Barang: ")
						fmt.Scanln(&product.Nama)
						fmt.Print("Masukkan Nama Barang: ")
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

						fmt.Print("Input [tambah / kurang]: ")
						fmt.Scanln(&condition)
						fmt.Print("Masukkan ID Barang: ")
						fmt.Scanln(&id)
						fmt.Print("Masukkan Stok Barang: ")
						fmt.Scanln(&quantity)

						res, err := menuBarang.UbahStokBarang(id, quantity, condition)
						if err != nil {
							fmt.Println(err)
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
							fmt.Println("-----------------------------------------------------")
							fmt.Println("ID Customer: ", customer.ID)
							fmt.Println("ID Pegawai: ", customer.IDPegawai)
							fmt.Println("Nama Customer: ", customer.Nama)
						}
						fmt.Println()
						fmt.Println("Total Data: ", len(listCustomer))
						fmt.Print("Press enter to return to the menu...")
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

					case 7: // Tampil List Nota
						listNota, err := menuTransaksi.ListNotaTransaksi()
						if err != nil {
							fmt.Println(err)
						}

						fmt.Println()
						fmt.Println("List Nota Transaksi")
						for _, nota := range listNota {
							fmt.Println("-----------------------------------------------------")
							fmt.Println("ID Transaksi: ", nota.IDTransaksi)
							fmt.Println("Nama Pegawai: ", nota.NamaPegawai)
							fmt.Println("Nama Customer: ", nota.NamaCustomer)
							fmt.Println("Tanggal Dibuat: ", nota.CreatedDate)
							fmt.Println("Barang transaksi: ")
							for i, barang := range nota.Barang {
								fmt.Printf("%d. %s\n", i+1, barang)
							}
						}
						fmt.Println()
						fmt.Println("Total Data: ", len(listNota))
						fmt.Print("Press enter to return to the menu...")
						fmt.Scanln()

					case 8: // Membuat Nota Transaksi
						var (
							namaCustomer string
							idCustomer   int
							listIdBarang []int
							listQuantity []int
						)

						// Ambil nama customer
						fmt.Print("Masukkan Nama Customer: ")
						fmt.Scanln(&namaCustomer)
						idCustomer, err = menuCustomer.CariCustomer(namaCustomer)
						if err != nil {
							fmt.Println(err)
							continue
						}

						// Ambil data barang
						inputTransaksi := 1
						for inputTransaksi != 0 {
							fmt.Println("Tambahkan Barang: ")
							fmt.Println("1. Tambahkan barang")
							fmt.Println("0. Selesai")
							fmt.Print("Masukkan Input: ")
							fmt.Scanln(&inputTransaksi)

							if inputTransaksi == 1 {
								var namaBarang string
								var quantity int
								fmt.Print("Masukkan nama barang: ")
								fmt.Scanln(&namaBarang)
								fmt.Print("Masukkan jumlah barang: ")
								fmt.Scanln(&quantity)

								id, err := menuBarang.CariBarang(namaBarang)
								if err != nil {
									fmt.Println(err)
									continue
								}

								listIdBarang = append(listIdBarang, id)
								listQuantity = append(listQuantity, quantity)
							}
						}

						// Tambah Transaksi
						idTrarnsaksi, err := menuTransaksi.TambahTransaksi(employee.ID, idCustomer)
						if err != nil {
							fmt.Println(err)
							continue
						}

						res := 0 // Jumlah barang yang berhasil ditambah
						// Tambah Transaksi Barang
						for i := range listIdBarang {
							err := menuTransaksi.TambahTransaksiBarang(idTrarnsaksi, listIdBarang[i])
							if err != nil {
								fmt.Println(err)
								break
							}

							// Kurangi barang ketika berhasil menambah barang ke transaksi
							_, err = menuBarang.UbahStokBarang(listIdBarang[i], listQuantity[i], "kurang")
							if err != nil {
								fmt.Println(err)
								break
							}

							res++
						}

						if res > 0 {
							fmt.Println("Berhasil menambahkan transaksi barang!")
						} else {
							fmt.Println("Gagal menambahkan transaksi barang!")
						}

					case 0:
						isLogin = false
						employee = pegawai.Pegawai{}
					}
				}

			}
		}
	}
}
