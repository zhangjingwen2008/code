package main

import (
	"fmt"
	"image/png"
	"log"
	"net/http"
)

func showImageSize(uri string) {
	resp, err := http.Get(uri)
	if err != nil {
		log.Fatalf("http.Get('%s') failed with %s\n", uri, err)
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 400 {
		log.Fatalf("http.Get() failed with '%s'\n", resp.Status)
	}
	img, err := png.Decode(resp.Body)
	if err != nil {
		log.Fatalf("png.Decode() failed with '%s'\n", err)
	}
	size := img.Bounds().Size()
	fmt.Printf("Size of the '%s' image: %dx%d\n", uri, size.X, size.Y)
}

func main() {
	showImageSize("https://www.programming-books.io/covers/Go.png")
}
