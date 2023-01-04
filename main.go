package main

import (
	"errors"
	"fmt"
	"personal-project/penjualan/beli"
	"personal-project/penjualan/jual"
	"strings"
)

func Ulang() {
	var Pilihan string
	fmt.Println("\nApakah Anda Mau Membeli Lagi? (Y/T)")
	fmt.Scan(&Pilihan)
	fmt.Println()

	if strings.ToUpper(Pilihan) == "Y" {
		defer main()
	} else if strings.ToUpper(Pilihan) == "T" {
		fmt.Println("Hati - Hati Di Jalan")
		return
	} else {
		fmt.Println("Pilih Yang Bener Dong")
	}
}

func Input() {
	jual.DaftarJual()

	var a int
	fmt.Print("\nSepatu Yang Ingin Dibeli: ")
	fmt.Scan(&a)
	fmt.Println()
	Pilihan(a)
}

func Pilihan(a int) {
	switch a {
	case 1:
		beli.Nike()
		Ulang()
	case 2:
		beli.Vans()
		Ulang()
	case 3:
		beli.Adidas()
		Ulang()
	case 4:
		beli.Reebok()
		Ulang()
	case 5:
		beli.Sketcher()
		Ulang()
	default:
		defer main()
		fmt.Println(errors.New("Maaf, Sepatu Tidak Tersedia\n"))
	}
}

func main() {
	Input()
}
