package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, Anda berada di halaman : %s\n", r.URL.Path)
		fmt.Fprintf(w, "Halaman / adalah halaman utama, seperti halaman index pada php.\n")
		fmt.Fprintf(w, "Halaman ini diakses pada port 8080 atau http://localhost:8080")
	})
	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, Anda berada di halaman : %s\n", r.URL.Path)
		fmt.Fprintf(w, "Halaman ini diakses dengan alamat http://localhost:8080/home\n")
		fmt.Fprintf(w, "Ini adalah tugas praktikum TCCLanjut minggu ke -7")

	})

	http.ListenAndServe(":8080", nil)
}