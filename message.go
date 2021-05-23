package main

type Message struct {
	Type                    string   `json:"type"`
	APIKey                  string   `json:"apikey"`
	Heartbeat               bool     `json:"heartbeat"`
	SubscribeDataType       []string `json:"subscribe_data_type"`
	SubscribeFilterSymbolId []string `json:"subscribe_filter_symbol_id"`
}
