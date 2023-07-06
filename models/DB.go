package models

type ApiUsage struct {
	Sub       string `json:"sub"`
	Token     string `json:"token"`
	Ip        string `json:"ip"`
	Method    string `json:"method"`
	Date      string `json:"date"`
	Cnt       int    `json:"cnt"`
	ChainName string `json:"chain_name"`
}
