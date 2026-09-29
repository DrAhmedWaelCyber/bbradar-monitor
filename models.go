package main

type Program struct {
	Name       string `json:"name"`
	URL        string `json:"url"`
	Identifier string `json:"identifier"`
}

type ListItem struct {
	Type     string  `json:"@type"`
	Position int     `json:"position"`
	Item     Program `json:"item"`
}

type ItemList struct {
	Type            string     `json:"@type"`
	Name            string     `json:"name"`
	URL             string     `json:"url"`
	NumberOfItems   int        `json:"numberOfItems"`
	ItemListElement []ListItem `json:"itemListElement"`
}
