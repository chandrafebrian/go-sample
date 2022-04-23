package main

import "fmt"

// cara membuat user input
// konsep input user

func main() {
	var userName string
	var lastName string
	var email string
	var userTiket uint
	var stokTiket uint = 100

	fmt.Println("Masukan Nama Depan")
	fmt.Scan(&userName)

	fmt.Println("Masukan Nama Belakang")
	fmt.Scan(&lastName)

	fmt.Println("Masukan Jumlah Pesanan Tiket")
	fmt.Scan(&userTiket)

	fmt.Println("Masukan Email Anda")
	fmt.Scan(&email)

	stokTiket = stokTiket - userTiket

	fmt.Printf("Terimakasih %v %v Sudah Booking  %v Tiket Untuk Nonton Bioskop Selanjutnya akan kami konfrimasi Berhasil booking di kirim ke email %v \n",
		userName, lastName, userTiket, email)

	fmt.Printf("sisa stok tinggal %v \n tiket", stokTiket)
}
