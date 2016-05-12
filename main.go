package main

import (
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("usage: gohex  <filename>")

		os.Exit(1)

	} else {
		filename := os.Args[1]

		data, err := ioutil.ReadFile(filename)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(hex.Dump(data))

	}

}
