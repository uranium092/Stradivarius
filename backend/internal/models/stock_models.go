package models

import "time"

type ItemStock struct {
	Id         string
	Ticker     string  `json:"ticker"`
	TargetFrom string  `json:"target_from"`
	TargetTo   string  `json:"target_to"`
	Company    string  `json:"company"`
	Action     string  `json:"action"`
	Brokerage  string  `json:"brokerage"`
	RatingFrom string  `json:"rating_from"`
	RatingTo   string  `json:"rating_to"`
	Time       time.Time `json:"time"`
}

// model for API Response
type StockResponse struct {
	Items    []ItemStock `json:"items"`
	NextPage string      `json:"next_page"`
}

// Request Queries for API
type RequestQueries struct {
	Page   int
	Search string
	Sort   string
}