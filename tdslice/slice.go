package main

import "fmt"

// slice +> potongan data array
// array = mutlak sedangkan slice dapat dirubah
// array hanya punya length tapi slice ada 3 data (pointer,length,capacity)
//   pointer =>penunjuk data pertama di array para slice
// 	 length => panjang dari slice
// 	 capacity => kapasitas dari slice dimana length tidak boleh lebih dari capcity

// code slice array:
// array[low:high] => mebuat slice array dari index low sampai index sebelum high
// array[low:] =>  mebuat slice array dari index low sampai index akhir di array
// array[:high] =>  mebuat slice array dari index 0 sampai index sebelum high
// array[:] => mebuat slice array dari index 0 sampai index akhir di array
func main() {
	var (
		months = [...]string{
			"Januari",
			"Februari",
			"Maret",
			"April",
			"Mei",
			"Juni",
			"Juli",
			"Agustus",
			"September",
			"Oktober",
			"November",
			"Desember",
		}

		slice1 = months[4:7]
	)

	// function slice:
	// len(slice) => mendapatkan panjang
	// cap(slice) => mendapatkan kapasitas
	// appen(slice, data) => membuat slice baru dengan menambah data ke posisi terakhir slice, jika kapasitas sudah penuhm maka akan membuat array baru
	// make([]TypeData, length, capacity) => membuat array baru
	// copy(destination, source) => menyalin slice dari source ke destination

	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))
}
