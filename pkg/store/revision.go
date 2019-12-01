package store

import (
	"github.com/gin-gonic/gin"
)

func RevisionsGet(c *gin.Context, revision interface{}, stream string) ConnectionError {

	return nil

}

func RevisionsIdDelete(c *gin.Context, id string) ConnectionError {

	return nil

}

func RevisionsIdDiffGet(c *gin.Context, diff interface{}, id string, to string) ConnectionError {

	return nil

}

func RevisionsIdGet(c *gin.Context, revision interface{}, id string) ConnectionError {

	return nil

}

func RevisionsPost(c *gin.Context, created interface{}, new interface{}) ConnectionError {

	return nil

}
