package beli

import "fmt"

func Pengecekan(d, c int) {
	if d >= c {
		fmt.Println("Terima Kasih Sudah Berbelanja")
	} else if d == 0 {
		fmt.Println("Pulang Aja Deh Kalo Gabawa Uang")
	} else {
		fmt.Println("Maaf Uang Kamu Kurang Manies")
	}
}

func Nike() {
	var a, c, d, e int
	b := 2000000

	fmt.Println("1. Nike")
	fmt.Println("Harga: Rp", b)
	fmt.Print("Jumlah Unit: ")
	fmt.Scan(&a)

	c = a * b

	fmt.Println("Total Harga: Rp", c)
	fmt.Print("Jumlah Yang Dibayarkan: Rp ")
	fmt.Scan(&d)

	e = d - c

	fmt.Println("Kembali: Rp", e)
	Pengecekan(d, c)
}

func Vans() {
	var a, c, d, e int
	b := 950000

	fmt.Println("2. Vans")
	fmt.Println("Harga: Rp", b)
	fmt.Print("Jumlah Unit: ")
	fmt.Scan(&a)

	c = a * b

	fmt.Println("Total Harga: Rp", c)
	fmt.Print("Jumlah Yang Dibayarkan: Rp ")
	fmt.Scan(&d)

	e = d - c

	fmt.Println("Kembali: Rp", e)
	Pengecekan(d, c)
}

func Adidas() {
	var a, c, d, e int
	b := 1500000

	fmt.Println("3. Adidas")
	fmt.Println("Harga: Rp", b)
	fmt.Print("Jumlah Unit: ")
	fmt.Scan(&a)

	c = a * b

	fmt.Println("Total Harga: Rp", c)
	fmt.Print("Jumlah Yang Dibayarkan: Rp ")
	fmt.Scan(&d)

	e = d - c

	fmt.Println("Kembali: Rp", e)
	Pengecekan(d, c)
}

func Reebok() {
	var a, c, d, e int
	b := 800000

	fmt.Println("4. Reebok")
	fmt.Println("Harga: Rp", b)
	fmt.Print("Jumlah Unit: ")
	fmt.Scan(&a)

	c = a * b

	fmt.Println("Total Harga: Rp", c)
	fmt.Print("Jumlah Yang Dibayarkan: Rp ")
	fmt.Scan(&d)

	e = d - c

	fmt.Println("Kembali: Rp", e)
	Pengecekan(d, c)
}

func Sketcher() {
	var a, c, d, e int
	b := 700000

	fmt.Println("5. Sketcher")
	fmt.Println("Harga: Rp", b)
	fmt.Print("Jumlah Unit: ")
	fmt.Scan(&a)

	c = a * b

	fmt.Println("Total Harga: Rp", c)
	fmt.Print("Jumlah Yang Dibayarkan: Rp ")
	fmt.Scan(&d)

	e = d - c

	fmt.Println("Kembali: Rp", e)
	Pengecekan(d, c)
}
