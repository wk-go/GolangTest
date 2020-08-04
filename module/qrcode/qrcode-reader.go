package main

import (
	"fmt"
	"github.com/tuotoo/qrcode"
	"os"
)

func main() {

	fi, err := os.Open("qrcode_google.png")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer fi.Close()
	qrmatrix, err := qrcode.Decode(fi)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(qrmatrix.Content)
}
