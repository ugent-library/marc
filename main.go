package main

import (
	"log"
	"os"

	"github.com/ugent-library/go-marc/marc"
)

func main() {
	dec := marc.NewDecoder("marcxml", os.Stdin)
	for {
		rec, err := dec.Decode()
		if err != nil {
			log.Panic(err)
		}
		if rec == nil {
			break
		}
		log.Printf("%+v", rec)
	}
}
