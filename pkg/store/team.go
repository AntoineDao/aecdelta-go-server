package store

import (
	"github.com/gin-gonic/gin"
)

func TeamsGet(c *gin.Context, teams interface{}, stream string, project string, role string, permission string) ConnectionError {

	return nil

}

func TeamsIdDelete(c *gin.Context, id string) ConnectionError {

	return nil

}

func TeamsIdGet(c *gin.Context, team interface{}, id string) ConnectionError {

	return nil

}

func TeamsIdPut(c *gin.Context, id string, update interface{}) ConnectionError {

	return nil

}

func TeamsPost(c *gin.Context, created interface{}, new interface{}) ConnectionError {

	return nil

}
