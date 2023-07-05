package models

type ApiUsage struct {
	Token  string `json:"token"`
	Method string `json:"method"`
	Date   string `json:"date"`
}
