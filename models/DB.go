package models

type ApiUsage struct {
	Token  string `json:"token"`
	Ip     string `json:"ip"`
	Method string `json:"method"`
	Date   string `json:"date"`
}
