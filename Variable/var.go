package main

import "fmt"

//1 variable hanya menyimpan tipe data yang sama.
// nama variable unik.
// variable yang tidak digunakan akan merah
// key var tidak wajib dapat diganti :=

func main() {
	var name string

	name = "Teguh Kurniawan"
	var nik = 01212.0121215
	city := "padang"
	fmt.Println(name, nik, city)

	name = "Dateg"
	fmt.Println(name)

	var (
		fullname = "Teguh Kurniawan"
		nickname = "Dateg"
	)

	fmt.Println(fullname)
	fmt.Println(nickname)
	fmt.Println(fullname, nickname)
}
