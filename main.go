package main

import (
	"gofeeds/scan"
	"gofeeds/validate"
	"os"
	"log"
	"reflect"
)

func main() {
	files := scan.Walk(os.DirFS("samples/"))
	
	conditions := []validate.Condition{
		validate.StartsWith{ Prefix: "sample-" },
	}

	valid, rejected := validate.ValidateFiles(files, conditions)

	log.Printf("valid: %v", valid)
	log.Printf("invalid: %v", reflect.ValueOf(rejected).MapKeys())
}


