package store

import (
	"github.com/gin-gonic/gin"
)

func UsersGet(c *gin.Context, users interface{}, stream string, streamRole string, project string, projectRole string, permission string) ConnectionError {

	return nil

}

func UsersIdGet(c *gin.Context, user interface{}, id string) ConnectionError {

	return nil

}
