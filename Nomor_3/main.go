package main

import (
	"Nomor_3/Pengaturan"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	
	const port = ":5050"
	fmt.Printf("\n Link Server: localhost%s", port)

	router := mux.NewRouter()
	router.HandleFunc("/create", Pengaturan.Create).Methods("POST")
	router.HandleFunc("/showAll", Pengaturan.Baca).Methods("GET")

	// Jika route yang dimasukkan user salah akan di direct dan menampilkan pesan isi route yang dapat digunakan
	router.NotFoundHandler = http.HandlerFunc(Pengaturan.Redirect)

	server := &http.Server{Addr: port, Handler: router}

	 server.ListenAndServe(); 
}
