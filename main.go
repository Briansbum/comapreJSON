package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/nsf/jsondiff"
)

func main() {

	verbose := flag.Bool("v", false, "`true` provides a verbose output with the result of the comparison")
	help := flag.Bool("help", false, "provides this help message")

	flag.Parse()

	if *help == true {
		flag.PrintDefaults()
		os.Exit(0)
	}

	if len(flag.Args()) != 2 {
		log.Fatal("two json documents must be provided")
	}

	args := flag.Args()

	diff, success := CompareJSON([]byte(args[0]), []byte(args[1]), verbose)
	if success == false {
		fmt.Println(diff.String())
		os.Exit(1)
	} else {
		fmt.Println(diff.String())
		os.Exit(0)
	}
}

func CompareJSON(json1 []byte, json2 []byte, verbose *bool) (jsondiff.Difference, bool) {
	opts := jsondiff.DefaultConsoleOptions()

	result, diff := jsondiff.Compare(json1, json2, &opts)

	if *verbose == true {
		log.Println(result)
		log.Println("----------DIFF----------")
		log.Println(diff)
		log.Println("----------DIFF----------")
	}

	if result.String() == "FullMatch" {
		return result, true
	} else if result.String() == "SupersetMatch" {
		return result, true
	} else {
		return result, false
	}
}
