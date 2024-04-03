package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func downloadFile(url string, filename string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}

func main() {
	args := os.Args[1:]

	if len(args) < 1 {
		fmt.Println("Usage: task <url>")
		return
	}

	url := args[0]

	tokens := strings.Split(url, "/")
	fileName := tokens[len(tokens)-1]
	if fileName == "" {
		fileName = "index.html"
	}

	fmt.Printf("Downloading %s to %s\n", url, fileName)
	if err := downloadFile(url, fileName); err != nil {
		panic(err)
	}
	fmt.Println("Download complete")
}
