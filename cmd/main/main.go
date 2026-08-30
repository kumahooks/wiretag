// Package wiretagexample uses the wiretag taglib binding example.
package main

import (
	"fmt"

	"wiretag"
)

func main() {
	musicFile, err := wiretag.Open("/home/kuma/Music/Sweet Trip/You Will Never Know Why/05 Milk.flac")
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		return
	}

	fmt.Printf("musicFile: %x\n", musicFile)

	musicFile.Close()
}
