package main

// import (
// 	"fmt"
// 	"strings"
// )

// //konsep loop dengan for

// func main() {

// 	var stokTiket uint = 50
// 	bookings := []string{}

// 	namaFungsi1(888)

// 	for {

// 		userName, lastName, email, userTiket := getUserfunction()

// 		isValidName, isValidEmail, isValidTiket := validasiInput(userName, lastName, userTiket, email, stokTiket)

// 		if isValidName && isValidEmail && isValidTiket {

// 			bokingTiket(stokTiket, bookings, userTiket, userName, lastName, email)

// 			firstName := namaFungsi2(bookings)
// 			fmt.Printf("the first names of booking are : %v \n", firstName)
// 			// ..................................... //

// 			fmt.Printf("the first name profile: %v \n", userName)

// 			fmt.Printf("sisa stok available tinggal %v \n", stokTiket)

// 			fmt.Printf("ini adalah hasil pengambilan data nama dari seluruh user yg sudah booking : %v \n", bookings)

// 			var tiketKosong bool = stokTiket == 0

// 			if tiketKosong {

// 				fmt.Println("tiket sudah habis terjual")
// 				break

// 			} else if stokTiket < userTiket {
// 				println("pesanan Berhasil")
// 				break
// 			}

// 		} else {
// 			if !isValidName {
// 				fmt.Println("error nama depan atau nama belakang terlalu singkat minimal 2 karakter")

// 			}
// 			if !isValidEmail {
// 				fmt.Println("error email anda tidak mengandung karakter @")
// 			}
// 			if !isValidTiket {
// 				fmt.Println("tiket tidak boleh kosong atau 0 ")
// 			}

// 		}

// 	}

// }

// func namaFungsi1(namaParamaterAngka int) {
// 	fmt.Printf("ini hasil nama fungsi 1 %v \n", namaParamaterAngka)
// }

// func namaFungsi2(bookings []string) []string {

// 	firstName := []string{}
// 	for _, booking := range bookings {
// 		var names = strings.Fields(booking)

// 		firstName = append(firstName, names[0])

// 	}
// 	return firstName
// }

// func validasiInput(userName string, lastName string, userTiket uint, email string, stokTiket uint) (bool, bool, bool) {
// 	isValidName := len(userName) >= 2 && len(lastName) >= 2
// 	isValidEmail := strings.Contains(email, "@")
// 	isValidTiket := userTiket > 0 && userTiket < stokTiket

// 	return isValidName, isValidEmail, isValidTiket

// }

// func getUserfunction() (string, string, string, uint) {
// 	var userName string
// 	var lastName string
// 	var email string
// 	var userTiket uint

// 	fmt.Println("Masukan Nama Depan")
// 	fmt.Scan(&userName)

// 	fmt.Println("Masukan Nama Belakang")
// 	fmt.Scan(&lastName)

// 	fmt.Println("Masukan Jumlah Pesanan Tiket")
// 	fmt.Scan(&userTiket)

// 	fmt.Println("Masukan Email Anda")
// 	fmt.Scan(&email)

// 	return userName, lastName, email, userTiket

// }

// func bokingTiket(stokTiket uint, bookings []string, userTiket uint, userName string, lastName string, email string) {

// 	stokTiket = stokTiket - userTiket
// 	bookings = append(bookings, userName+" "+lastName)

// 	fmt.Printf("Terimakasih %v %v Sudah Booking  %v Tiket Untuk Nonton Bioskop Selanjutnya akan kami konfrimasi Berhasil booking di kirim ke email %v  \n",
// 		userName, lastName, userTiket, email)
// 	fmt.Printf("sisa stok %v dan tersedia stok %v \n", stokTiket, bookings)
// }
