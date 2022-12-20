package main

import (
	"fmt"
	"log"
	"tokoku/config"
)

func main() {
	cfg := config.ReadConfig()
	db := config.OpenConnection(*cfg)

	var inputMenu = 1
	for inputMenu != 0 {
		var username, password string

		fmt.Println()
		fmt.Println("===== Tokoku ======")
		fmt.Print("Masukkan Username: ")
		fmt.Scanln(&username)
		fmt.Print("Masukkan Password: ")
		fmt.Scanln(&password)

		// Menu Admin
		if username == "admin" && password == "admin" {
			fmt.Println()
			fmt.Println("===== Admin ======")
			fmt.Println("1. Daftarkan  Pegawai")
			fmt.Println("2. Hapus Data Pegawai")
			fmt.Println("3. Hapus Data Barang")
			fmt.Println("4. Hapus Data Transaksi")
			fmt.Println("5. Hapus Data Customer")
			fmt.Println("5. Hapus Data Transaksi Barang")
			fmt.Println("9. Logout")
			fmt.Println("0. Exit")
			fmt.Print("Masukkan Input: ")
			fmt.Scanln(&inputMenu)

			// switch inputMenu {
			// case 1:
			// }

			// Menu Pegawai
		} else {
			row := db.QueryRow("SELECT id_pegawai FROM pegawai WHERE username = ? && password = ?", username, password)

			var idPegawai int
			err := row.Scan(&idPegawai)
			if err != nil {
				log.Println("Gagal melakukan scan id")
			}

			var isLogin = false
			if idPegawai >= 1 {
				isLogin = true
			}

			for isLogin {
				fmt.Println("===== GO-TODO ======")
				fmt.Println("1. Insert Barang")
				fmt.Println("2. Edit Informasi Barang")
				fmt.Println("3. Update Stok Barang")
				fmt.Println("4. Menambahkan Customer")
				fmt.Println("5. Membuat Nota Transaksi")
				fmt.Println("9. Logout")
				fmt.Print("Masukkan Input: ")
				fmt.Scanln(&inputMenu)

				switch inputMenu {
				case 1: // Insert barang
					var (
						nama     string
						quantity int
					)

					fmt.Print("Masukkan Nama Barang: ")
					fmt.Scanln(&nama)
					fmt.Print("Masukkan Jumlah Barang: ")
					fmt.Scanln(&quantity)

					result, err := db.Exec("INSERT INTO barang(id_pegawai, nama, quantity) VALUES(?, ?, ?)", idPegawai, nama, quantity)
					if err != nil {
						log.Println("Tambah barang gagal", err.Error())
					}

					fmt.Println(result.LastInsertId())
				}
			}
		}
	}
}
