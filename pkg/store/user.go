package store

import (
	model "github.com/antoinedao/aecdelta-go-server/pkg/model"
	"github.com/gin-gonic/gin"
)

func UsersGet(c *gin.Context) ([]model.User, ConnectionError) {

	resource := []model.User{}

	return resource, nil

}

func UsersIdGet(c *gin.Context) (model.User, ConnectionError) {

	resource := model.User{}

	return resource, nil

}
