package main

import (
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	filename := os.Args[1]
	data, _ := ioutil.ReadFile(filename)
	fmt.Println(hex.Dump(data))
}
