package auth

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func VerifyToken(c *gin.Context) (string, error) {
	bearerToken := c.GetHeader("Authorization")

	if bearerToken == "" {
		return "", errors.New("No Authentication Token provided.")
	}

	return "", nil
}

func AuthorizeUser(token string, c *gin.Context) error {

	return nil
}

func AuthenticationRequired() gin.HandlerFunc {
	return func(c *gin.Context) {

		token, err := VerifyToken(c)

		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		err = AuthorizeUser(token, c)

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		// add session verification here, like checking if the user and authType
		// combination actually exists if necessary. Try adding caching this (redis)
		// since this middleware might be called a lot
		c.Next()
	}
}
