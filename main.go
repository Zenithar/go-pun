package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"golang.org/x/net/idna"
)

var (
	decode bool
	file   string
)

func init() {
	flag.BoolVar(&decode, "d", false, "Decode domain as unicode to ascii")
	flag.StringVar(&file, "f", "", "Defines filename from where to read domain")
	flag.Parse()
}

func main() {
	var domain string
	var err error

	// Defines output writer
	log.SetOutput(os.Stderr)

	// Select input from STDIN or File
	reader := os.Stdin
	if len(strings.TrimSpace(file)) > 0 {
		log.Printf("Reading from file '%s'.\n", file)

		reader, err = os.Open(file)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		log.Println("Reading from STDIN.")
	}
	defer reader.Close()

	// Read line by line
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if decode {
			// Decode PUNYCODE to ASCII
			domain, err = idna.ToASCII(scanner.Text())
		} else {
			// Encode ASCII to PUNYCODE
			domain, err = idna.ToUnicode(scanner.Text())
		}

		if err != nil {
			log.Fatalf("pun: failed to convert a domain: %s", err)
		}
		fmt.Printf("%s\n", domain)
	}

	// Check errors
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
