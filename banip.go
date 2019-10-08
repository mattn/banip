package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net/http"
	"strings"
)

func list(domain string) {
	resp, err := http.Get("https://ipv4.fetus.jp/" + domain + ".txt")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		text := strings.TrimSpace(scanner.Text())
		if text == "" || text[0] == '#' {
			continue
		}
		fmt.Printf("### tuple ### deny any any 0.0.0.0/0 any %s in\n", text)
		fmt.Printf("-A banip-list -s %s -j DROP\n", text)
		fmt.Println()
	}
}

func main() {
	flag.Parse()
	for _, arg := range flag.Args() {
		list(arg)
	}
}
