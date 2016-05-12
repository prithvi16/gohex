package main

import (
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func usage() {
	fmt.Println("usage: gohex  <filename> or gohex <filename> -o <output_file>")

	os.Exit(1)

}
func datain() []byte {
	filename := os.Args[1]

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	return data

}
func main() {
	if len(os.Args) == 2 {

		fmt.Println(hex.Dump(datain()))

	} else if len(os.Args) == 4 {
		if os.Args[2] != "-o" && os.Args[2] != "-O" {

			usage()
		}
		bytedump := []byte(hex.Dump(datain()))
		err := ioutil.WriteFile(os.Args[3], bytedump, 0777)
		if err != nil {
			log.Fatal(err)
		}

	} else {
		usage()
	}

}
