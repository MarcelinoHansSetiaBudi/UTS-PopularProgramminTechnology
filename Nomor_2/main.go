package main

import (
	"fmt"
)

func CekArr (array1 []string, array2 []string){
	if len(array1) != len(array2) {
        fmt.Println("Panjang kedua array tidak sama")
    } else {
        isSame := true
        for i := 0; i < len(array1); i++ {
            if array1[i] != array2[i] {
                isSame = false
                fmt.Printf("index ke %d berbeda \n", i)
            }
        }
        if isSame {
            fmt.Println("Kedua array memiliki indeks yang sama")
        }
    }
}

func main() {
    // input array pertama
    fmt.Println("Masukkan panjang array pertama:")
    var n1 int
    fmt.Scanln(&n1)

    arr1 := make([]string, n1)
    fmt.Println("Masukkan elemen array pertama:")
    for i := 0; i < n1; i++ {
        fmt.Scanln(&arr1[i])
    }
	fmt.Println("\n")
    for i := 0; i < n1; i++ {
        fmt.Println(arr1[i])
    }

    // input array kedua
    fmt.Println("Masukkan panjang array kedua:")
    var n2 int
    fmt.Scanln(&n2)

    arr2 := make([]string, n2)
    fmt.Println("Masukkan elemen array kedua:")
    for i := 0; i < n2; i++ {
        fmt.Scanln(&arr2[i])
    }
	fmt.Println("\n")
    for i := 0; i < n1; i++ {
        fmt.Println(arr2[i])
    }

    // cek apakah array sama atau tidak
	if len(arr1) != len(arr2) {
        fmt.Println("Panjang kedua array tidak sama")
    } else {
		CekArr(arr1, arr2)
	}
}