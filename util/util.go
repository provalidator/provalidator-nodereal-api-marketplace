package util

import (
	"crypto/rsa"
	"net/url"
	"strings"
	"time"

	"github.com/provalidator-nodereal-api-marketplace/log"

	ratelimit "github.com/JGLTechnologies/gin-rate-limit"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func MakeQueryString(paramPairs url.Values) string {
	queryParams := "?"
	for k, v := range paramPairs {
		if queryParams == "?" {
			queryParams += k + "=" + v[0]
		} else {
			queryParams += "&" + k + "=" + v[0]
		}
	}
	return queryParams
}

func KeyFunc(c *gin.Context) string {
	return c.ClientIP()
}

func ErrorHandler(c *gin.Context, info ratelimit.Info) {
	log.Logger.Error.Println("Too many requests. Try again in " + time.Until(info.ResetTime).String())
	c.String(429, "Too many requests. Try again in "+time.Until(info.ResetTime).String())
}

func Verify(tokenString string, parsedPubKey *rsa.PublicKey) error {
	parts := strings.Split(tokenString, ".")
	method := jwt.GetSigningMethod("RS256")
	err := method.Verify(strings.Join(parts[0:2], "."), parts[2], parsedPubKey)

	if err != nil {
		log.Logger.Error.Println("Verifying key fail")
		log.Logger.Error.Println(err)
		return err
	}
	return nil
}
