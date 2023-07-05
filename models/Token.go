package models

import "github.com/golang-jwt/jwt"

type NoderealToken struct {
	Exp   int64  `json:"exp"`
	Iss   string `json:"iss"`
	Scope string `json:"scope"`
	Sub   string `json:"sub"`
	jwt.StandardClaims
}
