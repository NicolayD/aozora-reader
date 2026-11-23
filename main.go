package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

func main() {
	// terminal utility in go to read aozora - index authors and works, add command to search for word and call jisho API or fetch jmdict,
	// # use similar to vim keybindings. add I feel lucky button to open random work. save searched words.

	client := http.Client{}
	req, err := http.NewRequest("GET", "https://www.aozora.gr.jp/cards/000305/files/50379_39084.html", nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header = http.Header{
		"Accept-Charset": {"utf-8"},
		"Content-Type":   {"application/html"},
	}

	resp, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	if err != nil {
		log.Fatal(err)
	}
	// defer resp.Body.Close()

	reader := transform.NewReader(resp.Body, japanese.ShiftJIS.NewDecoder())
	body, err := io.ReadAll(reader)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(body))
}
