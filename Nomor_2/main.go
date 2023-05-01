package main

import (
	"fmt"
)

func CekArr (array1 []string, array2 []string){
	isSame := true
    for i := 0; i < len(array1); i++ {
        if array1[i] != array2[i] {
            isSame = false
            fmt.Printf("index ke %d berbeda \n", i)
        }
    }
    if isSame {
        fmt.Println("Kedua array memiliki value tiap indeks yang sama")
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
        fmt.Scan(&arr1[i])
    }

    arr2 := make([]string, n1)
    fmt.Println("Masukkan elemen array kedua:")
    for i := 0; i < n1; i++ {
        fmt.Scan(&arr2[i])
    }

    // cek apakah array sama atau tidak
	CekArr(arr1, arr2)
}
