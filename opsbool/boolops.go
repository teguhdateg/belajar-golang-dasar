package main

import "fmt"

// end && => 2bool
// OR || => 2bool
// negasi atau kebalikan ! => 1 bool

// end &&
// true && true => true
// true && false => false
// false && true => false
// false && false => false

// OR ||
// true && true => true
// true && false => true
// false && true => true
// false && false => false

// Negasi atau kebalikan !
// !=true => false
// !=false => true

func main() {
	var (
		nilaiakhir = 98
		absensi    = 40

		lulusnilaiakhir bool = nilaiakhir >= 70
		lulusabsensi    bool = absensi >= 80

		lulus bool = lulusnilaiakhir && lulusabsensi
	)

	fmt.Println(lulus)
}
