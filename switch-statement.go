package main

// import (
// 	"fmt"
// 	"strings"
// )

// //konsep loop dengan for

// func main() {

// 	var stokTiket uint = 50
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

// 		if userTiket > stokTiket {

// 			fmt.Printf("maaf !!!  anda tidak bisa booking %v  tiket, saat ini hanya tersedia %v tiket\n", userTiket, stokTiket)
// 			break

// 		}

// 		isValidName := len(userName) >= 2 && len(lastName) >= 2
// 		isValidEmail := strings.Contains(email, "@")
// 		isValidTiket := userTiket > 0 && userTiket < stokTiket

// 		if isValidName && isValidEmail && isValidTiket {

// 			stokTiket = stokTiket - userTiket
// 			bookings = append(bookings, userName+" "+lastName)

// 			fmt.Printf("Terimakasih %v %v Sudah Booking  %v Tiket Untuk Nonton Bioskop Selanjutnya akan kami konfrimasi Berhasil booking di kirim ke email %v  \n",
// 				userName, lastName, userTiket, email)

// 			// (firstName) adalah nama identifier untuk memanggil nilai value dr variable lain
// 			firstName := []string{}
// 			for _, booking := range bookings {
// 				var names = strings.Fields(booking)

// 				firstName = append(firstName, names[0])

// 			}
// 			fmt.Printf("the first name of all collect member : %v \n", firstName)

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
// 	// switch statement //
// 	city := "London"
// 	switch city {
// 	case "New york ":
// 		// execute code for booking new york conference ticket
// 	case "Singapore":
// 	// execute code for booking singapore conference ticket
// 	case "London":
// 	// some code
// 	case "Berlin":
// 	// some code here
// 	case "Mexico":
// 	// some code here
// 	case "Indonesia":
// 		// some code here

// 	default:
// 		fmt.Print("No city selected")
// 	}

// }
