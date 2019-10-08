package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func createEntries(domain string) {
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
	list := flag.Bool("l", false, "list countries")
	flag.Parse()

	if *list {
		doc, err := goquery.NewDocument("https://ipv4.fetus.jp/")
		if err != nil {
			log.Fatal(err)
		}

		doc.Find(".stupidtable tr").Each(func(i int, s *goquery.Selection) {
			td := s.Find("td")
			if td.Size() < 2 {
				return
			}
			fmt.Printf("%s: %s\n", td.Eq(0).Text(), td.Eq(1).Text())
		})
		return
	}
	for _, arg := range flag.Args() {
		createEntries(arg)
	}
}
