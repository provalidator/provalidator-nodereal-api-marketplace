package routes

import (
	"crypto/rsa"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/provalidator-nodereal-api-marketplace/log"
	"github.com/provalidator-nodereal-api-marketplace/models"
	"github.com/provalidator-nodereal-api-marketplace/util"
)

var parsedPubKey *rsa.PublicKey

func AuthorizationMiddleware(c *gin.Context) {
	// Get public key
	pubKey, _ := ioutil.ReadFile(os.Getenv("ROOT_PATH") + os.Getenv("PUB_KEY_DIR"))
	parsedPubKey, _ = jwt.ParseRSAPublicKeyFromPEM(pubKey)

	// Without Token
	whiteListIps := strings.Split(os.Getenv("WHITE_LIST_IPS"), ",")

	// Ip free pass
	for _, ip := range whiteListIps {
		if c.ClientIP() == ip {
			log.Logger.Trace.Println("whiteListIps ", c.ClientIP())
			c.Next()
			return
		}
	}

	// Swagger pass
	if strings.HasPrefix(c.Request.URL.Path, "/swagger") {
		c.Next()
		return
	}

	// Get Authorization
	tokenString := c.Request.Header.Get("Authorization")

	// No token
	if tokenString == "" {
		log.Logger.Error.Println("Token is None")
		c.JSON(http.StatusUnauthorized, gin.H{"status": http.StatusUnauthorized, "error": "Token is None"})
		c.Abort()
		return
	}

	// Remove Bearer
	tokenString = strings.Split(tokenString, "Bearer ")[1]

	// Write Log
	models.WriteLog(tokenString, c.ClientIP(), c.Request.URL.Path)

	// Verify
	tokenVerify := util.Verify(tokenString, parsedPubKey)

	if (tokenVerify) != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"status": http.StatusUnauthorized, "error": "Verifying key fail"})
		c.Abort()
		return
	}
	claims := &models.NoderealToken{}
	_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return parsedPubKey, nil
	})

	// Token valid check
	// 1. Expired check
	now := time.Now()
	unixTime := now.Unix()
	// claims.Exp = 1682038514 // Test

	if claims.Exp < unixTime {
		log.Logger.Error.Println("Token is expired")
		log.Logger.Error.Println("UnixTime : ", time.Unix(unixTime, 0))
		log.Logger.Error.Println("Token Exp : ", time.Unix(claims.Exp, 0))
		log.Logger.Error.Println("unixTime - claims.Exp =", unixTime-claims.Exp, "secs")
		c.JSON(http.StatusUnauthorized, gin.H{"status": http.StatusUnauthorized, "error": "Token is expired"})
		c.Abort()
		return
	}

	if err != nil {
		// 2. Sign check
		if err == jwt.ErrSignatureInvalid {
			log.Logger.Error.Println("Token signature is invalid")
			log.Logger.Error.Println(err)
			c.JSON(http.StatusUnauthorized, gin.H{"status": http.StatusUnauthorized, "error": "Token signature is invalid"})
			c.Abort()
			return
		}
		// 3. Etc error
		log.Logger.Error.Println("Authentication failed")
		log.Logger.Error.Println(err)
		c.JSON(http.StatusForbidden, gin.H{"status": http.StatusForbidden, "error": "Authentication failed"})
		c.Abort()
		return
	}
	// Usage check from DB
	// 1. Read usage data ( * find by tokenString)

	// 2. Usage process
	// 3. Write usage data
	c.Next()
}
