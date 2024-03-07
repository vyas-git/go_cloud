package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	args := os.Args

	if len(args) < 2 {
		log.Fatal("Not enough arguements passed , need to pass url argument")
	}

	url := args[1]

	resp, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	bytes, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("HTTP status code : %d \n Body : %s\n", resp.StatusCode, (bytes))
}
