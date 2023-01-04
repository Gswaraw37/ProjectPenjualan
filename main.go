package main

import (
	"errors"
	"fmt"
	"personal-project/penjualan/beli"
	"personal-project/penjualan/jual"
)

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
		break
	case 2:
		beli.Vans()
		break
	case 3:
		beli.Adidas()
		break
	case 4:
		beli.Reebok()
		break
	case 5:
		beli.Sketcher()
		break
	default:
		defer main()
		fmt.Println(errors.New("Maaf, Sepatu Tidak Tersedia\n"))
	}
}

func main() {
	Input()
}
