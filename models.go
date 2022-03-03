package main

type ExpectedQuotes[] struct {
	AUTHOR string `json:"author"`
	QUOTE  string `json:"text"`
}

type GivenQuotes struct {
	AUTHOR string   `json:"author"`
	QUOTES []string `json:"quotes"`
}
