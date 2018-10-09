package main

import (
	"errors"
	"log"
	"os"

	"github.com/nsf/jsondiff"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("two json documents must be provided")
	}

	json1 := []byte(os.Args[1])
	json2 := []byte(os.Args[2])

	_, err := CompareJSON(json1, json2)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func CompareJSON(json1 []byte, json2 []byte) (jsondiff.Difference, error) {
	opts := jsondiff.DefaultConsoleOptions()

	result, diff := jsondiff.Compare(json1, json2, &opts)

	log.Println(result)
	log.Println("----------BREAK----------")
	log.Println(diff)
	log.Println("----------BREAK----------")

	if result != jsondiff.Difference(1) && result != jsondiff.Difference(0) {
		return result, errors.New("json documents are different, " + result.String())
	} else {
		return result, nil
	}
}
