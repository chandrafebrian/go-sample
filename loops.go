package main

// import (
// 	"fmt"
// 	"strings"
// )

// //konsep loop dengan for

// func main() {

// 	var stokTiket uint = 100
// 	bookings := []string{}

// 	for {

// 		var userName string
// 		var lastName string
// 		var email string
// 		var userTiket uint

// 		fmt.Println("Masukan Nama Depan")
// 		fmt.Scan(&userName)

// 		fmt.Println("Masukan Nama Belakang")
// 		fmt.Scan(&lastName)

// 		fmt.Println("Masukan Jumlah Pesanan Tiket")
// 		fmt.Scan(&userTiket)

// 		fmt.Println("Masukan Email Anda")
// 		fmt.Scan(&email)

// 		stokTiket = stokTiket - userTiket
// 		bookings = append(bookings, userName+" "+lastName)

// 		fmt.Printf("Terimakasih %v %v Sudah Booking  %v Tiket Untuk Nonton Bioskop Selanjutnya akan kami konfrimasi Berhasil booking di kirim ke email %v  \n",
// 			userName, lastName, userTiket, email)

// 		firstName := []string{}
// 		for _, booking := range bookings {
// 			var names = strings.Fields(booking)

// 			firstName = append(firstName, names[0])

// 		}
// 		fmt.Printf("the first name of all collect member : %v \n", firstName)

// 		fmt.Printf("the first name profile: %v \n", userName)

// 		fmt.Printf("sisa stok available tinggal %v \n", stokTiket)

// 		fmt.Printf("ini adalah hasil pengambilan data nama dari seluruh user yg sudah booking : %v \n", bookings)
// 	}

// }
