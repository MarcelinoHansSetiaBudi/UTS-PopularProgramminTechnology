package Pesan

import (
	"os"

	"fmt"
	"time"

	"github.com/fatih/color"
)

/* Function Pesan
 * Fungsi ini untuk menampilkan pesan aktivitas. Seperti Success, error, dkk */
 
func Psn(opt int, msg string) {
        now := time.Now()
		red := color.New(color.FgHiRed)

        fmt.Printf(" [%s] ", now.Format("15:04:05 2006/01/02"))

	 switch (opt) {
	 case 1:
		fmt.Printf("INFO")
	 case 2:
		red.Printf("ERROR")
	 case 3:
		fmt.Printf("SUCCESS")
	 case 4:
		fmt.Printf("FATAL")
		red.Println(": "+msg)
		os.Exit(0)
	 }
        fmt.Println(": "+msg)
}
