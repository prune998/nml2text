package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"io"
	"os"
)

var nmlFile string

func init() {
	flag.StringVar(&nmlFile, "nmlFile", "playlist.nml", "path to the NML file")
}

func main() {

	flag.Parse()

	xmlFile, err := os.Open(nmlFile)
	if err != nil {
		fmt.Println(err)
	}

	defer xmlFile.Close()

	byteValue, err := io.ReadAll(xmlFile)
	if err != nil {
		fmt.Println(err)
	}

	var nml NML
	err = xml.Unmarshal(byteValue, &nml)
	if err != nil {
		fmt.Println(err)
	}

	for c, entry := range nml.COLLECTION.ENTRY {
		fmt.Printf("%d) %s - %s\n", c+1, entry.TITLE, entry.ARTIST)
	}

}
