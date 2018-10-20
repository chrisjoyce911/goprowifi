package main

import (
	"fmt"
	"log"
	"time"

	"github.com/chrisjoyce911/goprowifi"
)

func main() {
	debug := true

	start := time.Now()
	hero := goprowifi.NewClient(debug)
	elapsed := time.Since(start)
	log.Printf("NewClient took %s", elapsed)

	start = time.Now()
	hero.ModePhotoSingle()
	hero.CaptureStart()
	elapsed = time.Since(start)
	log.Printf("Shot took %s", elapsed)

	photos, _ := hero.MediaPhotoList()

	for _, x := range photos {
		fmt.Printf("%+v", x)
	}

}
