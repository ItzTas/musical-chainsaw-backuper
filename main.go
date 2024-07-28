package main

import (
	"flag"
	"fmt"
	"log"
)

func main() {
	src := flag.String("src", "", "source of the directory to copy")
	dst := flag.String("dst", "", "destination of the directory to copy")
	flag.Parse()

	if *src == "" || *dst == "" {
		log.Fatal("src and dst params required")
	}

	fmt.Println(*src, *dst)
}
