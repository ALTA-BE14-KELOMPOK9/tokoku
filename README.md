# TokoKu

...

## Table of content

1. [Tokoku](#tokoku)
2. [Table of content](#table-of-content)
3. [Setup project](#setup-project)
4. [Entitiy Relationship Diagram](#entity-relationship-diagram)
5. [Project structure](#project-structure)
6. [Flow](#flow)

## Setup project

1. Clone repository
   `git clone https://github.com/ALTA-BE14-KELOMPOK9/tokoku.git`
2. Delete .git
   `rm -rf .git`
3. Install package
   `go get github.com/joho/godotenv`
   `github.com/go-sql-driver/mysql`
4. Execute tokoku.sql
5. Run program
   `go run .`

## Entity Relationship Diagram

![erd tokoku](https://github.com/ALTA-BE14-KELOMPOK9/tokoku/blob/main/erd.png?raw=true)

## Project structure

## Flow

Admin

1. Input 1 untuk login
2. Input username = admin dan password = admin
3. Input 1 untuk mendaftarkan pegawai

-   Tulis username dan password pegawai

4. Input 2 untuk menampilkan list dan menghapus pegawai

-   Input ID pegawai yang akan dihapus

5. Input 3 untuk menampilkan list dan menghapus barang

-   Input ID barang yang akan dihapus

6. Input 4 untuk menampilkan list dan menghapus transaksi

-   Input ID transaksi yang akan dihapus

7. Input 5 untuk menampilkan list dan menghapus customer

-   Input ID customer yang akan dihapus

8. Input 6 untuk menampilkan list dan menghapus transaksi

-   Input ID yang akan dihapus

9. Input 9 untuk logout
10. Input 0 untuk exit program

User

1. Input 1 untuk login
2. Input username dan password sesuai yang ada pada database
3. Input 1 untuk menampilkan list barang
4. Input 2 untuk menambahkan barang

-   Masukkan nama dan jumlah barang

5. Input 3 untuk mengubah nama barang

-   Masukkan ID barang yang akan diubah
-   Input perubahan nama barang

6. Input 4 untuk mengubah stok barang

-   Pilih kondisi perubahan, bertambah atau berkurang dengan input (tambah/kurang)
-   Masukkan ID barang yang akan diubah stoknya
-   Masukkan perubahan jumlah barang

7. Input 5 untuk menampilkan list customer
8. Input 6 untuk menambahkan customer

-   Masukkan nama customer

9. Input 7 untuk menampilkan list nota
10. Input 8 untuk membuat nota transaksi

-   Masukkan nama customer yang akan bertransaksi
-   Input 1 untuk tambahkan barang
-   Masukkan nama barang yang akan dibeli
-   Masukkan jumlah barang yang akan dibeli
-   Input 1 untuk melakukan transaksi ulang
-   Input 2 untuk selesai transaksi
