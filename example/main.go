package main

import (
	"fmt"

	"github.com/chrisjoyce911/goprowifi"
)

func main() {
	debug := true

	hero := goprowifi.NewClient(debug)

	CameraStatus, _, _ := hero.GetCameraStatus()
	fmt.Println(CameraStatus)
}
