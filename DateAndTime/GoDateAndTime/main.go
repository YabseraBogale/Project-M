package main

import (
	"strconv"
	"time"
)

func main() {
	tim := strconv.Itoa(time.Now().Hour()) + " " + strconv.Itoa(time.Now().Minute()) + " " + strconv.Itoa(time.Now().Second())
	println(tim)
}
