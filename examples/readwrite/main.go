package main

import (
	"log"
	"os"

	"github.com/riadevatix/ini"
)

func main() {
	f, err := os.ReadFile("./input.ini")
	if err != nil {
		log.Fatalln(err)
	}

	conf, err := ini.Load(f)
	if err != nil {
		log.Fatalln(err)
	}

	err = conf.SaveTo("output.ini")
	if err != nil {
		log.Fatalln(err)
	}
}
