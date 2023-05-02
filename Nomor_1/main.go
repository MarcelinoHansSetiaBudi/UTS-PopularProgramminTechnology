package main

import "fmt"

func total(items ...int) int {
	jumlah := 0
	for _, item := range items {
		jumlah += item
	}
	return jumlah
}

func main() {
	nasgor := 10000
	miegor := 30000

	jml_item_1 := 2
	jml_item_2 := 3

	Bayar := total(nasgor*jml_item_1, miegor*jml_item_2)
	fmt.Printf("Total: %d \n", Bayar)
}
