package main

import "fmt"

//konsep loop dengan for

func main() {

	var stokTiket uint = 100
	bookings := []string{}

	for {

		var userName string
		var lastName string
		var email string
		var userTiket uint

		fmt.Println("Masukan Nama Depan")
		fmt.Scan(&userName)

		fmt.Println("Masukan Nama Belakang")
		fmt.Scan(&lastName)

		fmt.Println("Masukan Jumlah Pesanan Tiket")
		fmt.Scan(&userTiket)

		fmt.Println("Masukan Email Anda")
		fmt.Scan(&email)

		stokTiket = stokTiket - userTiket
		bookings = append(bookings, userName+" "+lastName)

		fmt.Printf("the whole slice: %v \n", bookings)
		fmt.Printf("the first value: %v \n", bookings[0])
		fmt.Printf("slice length: %v \n", len(bookings))

		fmt.Printf("Terimakasih %v %v Sudah Booking  %v Tiket Untuk Nonton Bioskop Selanjutnya akan kami konfrimasi Berhasil booking di kirim ke email %v  \n",
			userName, lastName, userTiket, email)

		fmt.Printf("sisa stok tinggal %v \n tiket", stokTiket)

		fmt.Printf("ini adalah hasil pengambilan data nama dari seluruh user yg sudah booking : %v \n", bookings)
	}

}
