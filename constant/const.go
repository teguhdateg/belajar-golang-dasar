package main

import "fmt"

// konstanta / const sama dengan var tapi isi merupakan nilai mutlak tidak dapat diubah
// konstanta boleh tidak langsung dipakai dan tidak akan menyebabkan error. berbeda dengan variable. ex : value

func main() {
	// bentuk declarasi 1
	const fullname string = "Teguh Kurniawan"
	const nickname = "Dateg"
	const value = 8000

	fmt.Println(fullname)
	fmt.Println(nickname)
	fmt.Println(fullname, nickname)

	// bentuk declarasi 2
	const (
		nama = "Teguh"
		umur = 26
		telp = 12345678
	)

	fmt.Println(nama)
	fmt.Println(umur)
	fmt.Println(telp)

	// jika di aktifkan akan error karena tidak dapat dirubah
	// fullname = "asasas"
	// nickname = "jjkjk"
	// umur = 17

}
