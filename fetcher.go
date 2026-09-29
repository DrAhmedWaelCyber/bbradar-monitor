package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/PuerkitoBio/goquery"
)

const TargetURL = "https://bbradar.io/?tags=wildcard%2Cdomain%2Capi&platform=bugcrowd%2Cyeswehack%2Chackerone"

func FetchPrograms() ([]Program, error) {
	client := &http.Client{
		Timeout: 15 * time.Second,
	}

	req, err := http.NewRequest("GET", TargetURL, nil)
	if err != nil {
		return nil, err
	}
	
	req.Header.Set("User-Agent", "curl/8.4.0")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
	req.Header.Set("Accept-Language", "en-US,en;q=0.9")

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch target URL: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to parse HTML: %w", err)
	}

	var programs []Program
	scriptText := doc.Find("#dynamic-structured-data-itemlist").Text()
	if scriptText == "" {
		return nil, fmt.Errorf("could not find ld+json script tag with id 'dynamic-structured-data-itemlist'")
	}

	var itemList ItemList
	if err := json.Unmarshal([]byte(scriptText), &itemList); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON: %w", err)
	}

	for _, el := range itemList.ItemListElement {
		programs = append(programs, el.Item)
	}

	return programs, nil
}
