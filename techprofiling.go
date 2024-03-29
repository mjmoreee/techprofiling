package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	wappalyzer "github.com/projectdiscovery/wappalyzergo"
)

func main() {
	// Baca input dari stdin
	reader := bufio.NewReader(os.Stdin)
	url, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	// Hapus karakter newline dari URL
	url = url[:len(url)-1]

	// Lakukan permintaan HTTP ke URL
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// Baca respons body
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Inisialisasi Wappalyzer client
	wappalyzerClient, err := wappalyzer.New()
	if err != nil {
		log.Fatal(err)
	}

	// Lakukan fingerprinting terhadap header dan body respons
	fingerprints := wappalyzerClient.Fingerprint(resp.Header, data)
	fmt.Printf("%v\n", fingerprints)
}
