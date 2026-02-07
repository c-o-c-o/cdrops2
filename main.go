package main

import (
	"cdrops/gcmz"
	"cdrops/internal/app"
	"flag"
)

func main() {
	flag.Parse()
	if err := app.Run(flag.Args(), gcmz.DropFiles); err != nil {
		panic(err.Error())
	}
}
