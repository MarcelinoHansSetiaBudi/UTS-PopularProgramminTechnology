package Pengaturan

import (
	"Nomor_3/Pesan"
	"encoding/json"
	"fmt"
	"net/http"
)

type Mahasiswa struct {
	Nama   string `json:"Nama"`
	NIM    string `json:"Nim"`
	Alamat string `json:"Alamat"`
}

/* Deklarasi array of struct data dengan format dari struct Mahasiswa
 * variable ini bertujuan menampung beberapa 'data' dari var simpan mahasiswa.
 *
 * Ini menggunakan penyimpanan sederhana Jadi, ketika program di hentikan, maka data akan hilang. */
var data []Mahasiswa

func Create(response http.ResponseWriter, request *http.Request) {
	// Set respon MIME type json sesuai dengan yang diminta soal
	response.Header().Set("Content-Type", "application/json")

	var simpan Mahasiswa

	simpan.Nama   = request.FormValue("Nama")
	simpan.NIM    = request.FormValue("NIM")
	simpan.Alamat = request.FormValue("Alamat")

	if simpan.Nama == "" || simpan.NIM == "" || simpan.Alamat == "" {
		response.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(response, `{"status":"error", "message":"Nama, NIM, dan Alamat tidak boleh kosong!"}`)
		Pesan.Psn(2, "func 'Create()' -> Pengisian data masih ada yang kosong.")
		return
	} else {
		data = append(data, simpan)
		
		if extractData, err := json.Marshal(simpan); err != nil {
			response.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(response, `{"status":"error", "message":"terjadi kesalahan pada server."}`)
			Pesan.Psn(2, "func 'Create()' -> Terjadi error saat marshal data (data->array json).")
			return
		} else {
			fmt.Fprint(response, string(extractData))
			Pesan.Psn(3, "func 'Create()' -> Berhasil menambahkan data.")
		}
	}
}

func Baca(response http.ResponseWriter, request *http.Request) {
	// Set respon MIME type json, bertujuan untuk memberi tau user bahwa output ini merupakan json.
	response.Header().Set("Content-Type", "application/json")

	if len(data) == 0 {
		response.WriteHeader(http.StatusNotFound)
		fmt.Fprint(response, `{"status":"error", "message":"Data Not Found"}`)
		Pesan.Psn(1, "func 'Baca()' -> Data dalam array of struct masih kosong.")
		return
	}
	if extractData, err := json.Marshal(data); err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(response, `{"status":"error", "message":"terjadi kesalahan pada server."}`)
			Pesan.Psn(2, "func 'Baca()' -> Terjadi error saat marshal data (data->array json).")
		return
	} else {
		fmt.Fprint(response, string(extractData))
		Pesan.Psn(3, "func 'Baca()' -> Berhasil menampilkan semua data.")
	}
}

func Redirect(response http.ResponseWriter, request *http.Request){
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusNotFound)

	array := map[string]string{
		"get": "http://localhost:5050/showAll",
		"post": "http://localhost:5050/create",
	}

	jsonData, _ := json.Marshal(map[string]interface{}{
		"status":  "error",
		"message": "route yang anda masukkan salah",
		"route": array,
	})

	fmt.Fprint(response, string(jsonData))
	Pesan.Psn(2, "func 'HandleNotFound()' -> Route yang anda masukkan salah.")
}