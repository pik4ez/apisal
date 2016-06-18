package main

import (
	lib "github.com/pik4ez/apisal/apisal"
	"io"
	"log"
	"os"
)

func main() {
	reader := lib.NewPointsReader(os.Stdin)
	writer := lib.NewPointsWriter(os.Stdout)

	for {
		point, err := reader.ReadNext()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		// TODO:

		writer.Write(point)
	}
}
