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

// model for population API Response
type StockResponsePopulation struct {
	Items    []ItemStock `json:"items"`
	NextPage string      `json:"next_page"`
}

// model for own Endpoint-API Response
type StockResponse struct {
	TotalPages float64 `json:"totalPages"`
	DataStock []ItemStock   `json:"dataStock"`
}

// model to know if the stock population is finished success
type StockStatus struct {
	Done bool
	NextPage string
}

// Request Queries for own Endpoint-API
type RequestQueries struct {
	Page   int
	Search string
	Sort   string
}