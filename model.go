package main

import "time"

type priceRequest struct {
	Success bool      `json:"success"`
	Payload []payload `json:"payload"`
}

type payload struct {
	High      string    `json:"high"`
	Last      string    `json:"last"`
	CreatedAt time.Time `json:"created_at"`
	Book      string    `json:"book"`
	Volume    string    `json:"volume"`
	Vwap      string    `json:"vwap"`
	Low       string    `json:"low"`
	Ask       string    `json:"ask"`
	Bid       string    `json:"bid"`
	Change24  string    `json:"change_24"`
}
