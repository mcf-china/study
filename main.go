package main

import (
	"fmt"
	"log"

	"github.com/evilsocket/mp4ff/mp4"
)

func main() {
	mp4File, err := mp4.OpenReader("example.mp4")
	if err != nil {
		log.Fatal(err)
	}
	defer mp4File.Close()

	moov := mp4File.Moov
	if moov == nil {
		log.Fatal("moov box not found")
	}

	udta := moov.Udta
	if udta != nil {
		for _, box := range udta.SubBoxes {
			if box.Boxtype() == "udta" {
				fmt.Println(box)
			}
		}
	}
}
