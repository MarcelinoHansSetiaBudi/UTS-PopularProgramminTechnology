package main

import (
	"Nomor_3/Pengaturan"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

/* Main program
 * disinilah tempat pemanggilan beberapa function untuk dijalankan pada poin awal / entry poin. */
func main() {
	// Panggil function ClearScreen() dari folder Utils/
	
	// Set port untuk menjalankan program berdasarkan port yang telah diset!
	const port = ":5050"

	// Menampilkan output variasi, supaya terlihat lebih bagus (output header)
	fmt.Printf("\n Link Server: localhost%s", port)

	// Set variable router untuk menampung object dari NewRouter
	router := mux.NewRouter()

	/* Set route program saya sesuai dengan SOAL
	 * pada /nama menggunakan method POST 		=> berfungsi untuk POST data (Nama NIM Alamat)
	 * untuk /semuadata menggunakan method GET	=> berfungsi untuk GET data (tampilkan semua data dari array of struct)
	 *
	 * /nama akan meneruskan request ke function CreateData() yang berada di folder Controllers
	 * hasilnya berupa response yang akan di teruskan ke user
	 *
	 * /semuadata sama seperti /nama, hanya beda fungsi yaitu menggunakan fungsi ReadData()
	 */
	router.HandleFunc("/nama", Pengaturan.Create).Methods("POST")
	router.HandleFunc("/semuadata", Pengaturan.Baca).Methods("GET")

	/* Untuk handle jika route tidak ditemukan, misal :
	 * ketika user request root domain / selain '/nama' & '/semuadata'
	 * maka akan dihandle disini, dengan cara : memanggil function yang ada di folder Controllers/
	 */
	router.NotFoundHandler = http.HandlerFunc(Pengaturan.HandleNotFound)

	/* variable server saya set untuk membuat sebuat object / instance dari http.Server
	 * yang bertujuan untuk membuat server HTTP */
	server := &http.Server{Addr: port, Handler: router}

	/* Logic Handle CTRL + C, References : https://pkg.go.dev/os/signal
	 * disini saya membuat sebuah program untuk handle jika developer menekan CTRL + C
	 * supaya tidak ada output yang menggangu ketika kita menekan CTRL + C saat program dijalankan.
	 *
	 * SIGINT (interrupt signal) akan dihasilkan ketika user menekan CTRL + C
	 *
	 * pada variable isPressedCTRLC akan menunggu adanya sinyal CTRL+C (SIGINT)
	 * pada function signal.Notify() akan memberitahu bahwa ada sinyal yang ditangkap ketika CTRL + C ditekan
	 */

	/* pada code dibawah akan menjalankan server dengan port yang telah di tentukan diatas (line code 58)
	 * serta handle jika terdapat error pada server maka akan exit (function utils.logger params 4 [FATAL])
	 */
	 server.ListenAndServe(); 
}