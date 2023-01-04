package jual

import "fmt"

func DaftarJual() {
	fmt.Println("Silahkan Pilih Sepatu Yang Ingin Anda Beli:")

	Sepatu := [5]string{
		"1. Nike",
		"2. Vans",
		"3. Adidas",
		"4. Reebok",
		"5. Sketcher",
	}
	fmt.Println(Sepatu[0])
	fmt.Println(Sepatu[1])
	fmt.Println(Sepatu[2])
	fmt.Println(Sepatu[3])
	fmt.Println(Sepatu[4])
}
