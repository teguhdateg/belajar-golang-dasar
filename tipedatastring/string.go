package main

import "fmt"

// Function untuk String:
// Function | Keterangan
// len("string") | menghitung jumlah karakter
// "string"[number(index)] | menagambil karakter yang ditentukan yang ditampilkan adalah byte dari abjad tsb.

func main() {
	fmt.Println("ini adalah string biasa yg didalam kutip dua")
	fmt.Println(len("ini adalah string yg didalam kutip dua"))
	fmt.Println("ini adalah string yg didalam kutip dua"[1])
	fmt.Println("ini adalah string yg didalam kutip dua"[8])
}
