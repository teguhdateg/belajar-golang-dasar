package main

import "fmt"

// sensitif
//4 tipe data Floating Point in go:
// TipeData		| Minimum Value			| Maximum Value

// Float (desimal)
// sering digunakan
// float32		| 1.18 x 10^-38			| 3.4 x 10^38
// float64		| 2.23 x 10^-308		| 1.80 x 10^308

// complex()
// hanya untuk sistem yg sangat complex di hitungan math ex: statistik, dsb.
// untuk ecommerce atau backend jarang digunakan
// complex64	| complex number dengan
// complex128	|
func main() {
	fmt.Println("Tiga Koma Lima = ", 3.5)
	fmt.Println("Enam Koma Dua = ", 6.2)
	fmt.Println("Tiga Koma Lima = ", 3.5)
	fmt.Println("Enam Koma Dua = ", 6.2)
}
