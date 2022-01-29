// package main

// import "fmt"

// // konsep array

// func main() {
// 	var userName = "chandra"
// 	var lastName = "febrian"
// 	var email = "chandra@gmail.com"
// 	var userTiket uint
// 	var stokTiket uint = 100
// 	var bookings [50]string

// 	stokTiket = stokTiket - userTiket
// 	bookings[0] = userName + " " + lastName

// 	fmt.Printf("Terimakasih %v %v Sudah Booking  %v Tiket Untuk Nonton Bioskop Selanjutnya akan kami konfrimasi Berhasil booking di kirim ke email %v \n",
// 		userName, lastName, userTiket, email)

// 	fmt.Printf("sisa stok tinggal %v \n tiket", stokTiket)
// 	fmt.Printf("hasil seluruh array:  %v \n", bookings)
// 	fmt.Printf("hasil nilai pertama booking: %v \n", bookings[0])
// 	fmt.Printf("tipe data array: %T \n", bookings)
// 	fmt.Printf("panjang data array: %v \n", len(bookings))
// }

// -------------------------------------------------------------------------

// konsep slice
// slice adalah abstrasi dari array
package main

// import "fmt"

// func main() {
// 	var userName = "chandra"
// 	var lastName = "febrian"
// 	var email = "chandra@gmail.com"
// 	var userTiket uint
// 	// var stokTiket uint = 100
// 	bookings := []string{}

// 	// stokTiket = stokTiket - userTiket

// 	bookings = append(bookings, userName+" "+lastName)

// 	fmt.Printf("Terimakasih %v %v Sudah Booking  %v Tiket Untuk Nonton Bioskop Selanjutnya akan kami konfrimasi Berhasil booking di kirim ke email %v \n",
// 		userName, lastName, userTiket, email)

// 	fmt.Printf("ini adalah seluruh hasil booking an tiket : %v \n", bookings)
// }
