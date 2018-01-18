package main

import (
	"flag"
	"fmt"
	"time"
)

type 

func main() {
	var period = flag.Duration("period", 1*time.Second, "sleep period")
	flag.Parse()
	flag.IVar()
	fmt.Printf("sleep for %v...\n", *period)
	time.Sleep(*period)
	fmt.Println()
}
