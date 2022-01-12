package main

import "fmt"

func main() {
	var (
		nilai32 int32  = 100000
		nilai64 int64  = int64(nilai32)
		nilai8  int8   = int8(nilai32)
		str            = "ini adalah string yg didalam kutip dua"
		e       byte   = str[1]
		eStr    string = string(e)
	)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai8)
	fmt.Println(str)
	fmt.Println(e)
	fmt.Println(eStr)
}
