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
	var bookings [50]string

	fmt.Println("Masukan Nama Depan")
	fmt.Scan(&userName)

	fmt.Println("Masukan Nama Belakang")
	fmt.Scan(&lastName)

	fmt.Println("Masukan Jumlah Pesanan Tiket")
	fmt.Scan(&userTiket)

	fmt.Println("Masukan Email Anda")
	fmt.Scan(&email)

	stokTiket = stokTiket - userTiket
	bookings[0] = userName + " " + lastName

	fmt.Printf("Terimakasih %v %v Sudah Booking  %v Tiket Untuk Nonton Bioskop Selanjutnya akan kami konfrimasi Berhasil booking di kirim ke email %v \n",
		userName, lastName, userTiket, email)

	fmt.Printf("sisa stok tinggal %v \n tiket", stokTiket)
	fmt.Printf("hasil seluruh array:  %v \n", bookings)
	fmt.Printf("hasil nilai pertama booking: %v \n", bookings[0])
	fmt.Printf("tipe data array: %T \n", bookings)
	fmt.Printf("panjang data array: %v \n", len(bookings))
}
