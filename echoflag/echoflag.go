package main

import (
	"flag"
	"fmt"
	"time"
)

var bval = flag.Bool("bool", false, "bool value for test")
var ival = flag.Int("int", 100, "integer value for test")
var sval = flag.String("string", "null", "string value for test")
var tval = flag.Duration("time", 10*time.Second, "time duration for test")

func main() {

	flag.Parse()

	fmt.Println("bool:\t", *bval)
	fmt.Println("int:\t", *ival)
	fmt.Println("string:\t", *sval)
	fmt.Println("time:\t", *tval)
}
