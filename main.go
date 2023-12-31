package main

import (
	"gofeeds/scan"
	"fmt"
	"os"
	"log"
)

func main() {
	log.Println(fmt.Sprintf("%v", scan.Walk(os.DirFS("."))))
}


