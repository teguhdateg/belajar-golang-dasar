package main

import "fmt"

func main() {
	var arr [3]string
	arr[0] = "dateg"
	arr[1] = "teguh"
	arr[2] = "kurniawan"

	fmt.Println(arr[0])
	fmt.Println(arr[1])
	fmt.Println(arr[2])

	var valuearray = [3]int{
		100,
		200,
		300,
	}

	fmt.Println(valuearray)

	// function array:
	// len(array) => menghitung panjang array #bukan jumlah data
	// array[index] => get array di [index]
	// array[index] = value => mengubah nilai array [index]

	fmt.Println(len(arr))
	fmt.Println(len(valuearray))
}
