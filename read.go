package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func reafText() {

    content, err := ioutil.ReadFile("sample.txt")

     if err != nil {
          log.Fatal(err)
     }

    fmt.Println(string(content))
}

