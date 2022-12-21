package main

import (
	"fmt"
	"tokoku/barang"
	"tokoku/config"
	"tokoku/pegawai"
)

func main() {

	// Setup database
	cfg := config.ReadConfig()
	db := config.OpenConnection(*cfg)

	menuPegawai := pegawai.MenuPegawai{DB: db}
	menuBarang := barang.MenuBarang{DB: db}

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
				fmt.Println(employee)
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
					fmt.Println("7. Tampil List Transaksi")
					fmt.Println("8. Buat Nota Transaksi")
					fmt.Println("9. Logout")
					fmt.Print("Masukkan Input Menu: ")
					fmt.Scanln(&inputMenu)

					switch inputMenu {
					case 1:
						// Tampil List Barang
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

					case 2:
						// Tambah Barang
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

					case 3:
						// Ubah Nama Barang
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

					case 4:
						// Ubah Stok Barang
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

					case 9:
						isLogin = false
						employee = pegawai.Pegawai{}
					}
				}

			}
		}
	}
}
