package store

import (
	model "github.com/antoinedao/aecdelta-go-server/pkg/model"
	"github.com/gin-gonic/gin"
)

func TeamsGet(c *gin.Context) (model.Team, ConnectionError) {

	resource := model.Team{}

	return resource, nil

}

func TeamsIdDelete(c *gin.Context) (string, ConnectionError) {

	resource := "resource"

	return resource, nil

}

func TeamsIdGet(c *gin.Context) (model.Team, ConnectionError) {

	resource := model.Team{}

	return resource, nil

}

func TeamsIdPut(c *gin.Context) (model.Team, ConnectionError) {

	resource := model.Team{}

	return resource, nil

}

func TeamsPost(c *gin.Context) (model.Team, ConnectionError) {

	resource := model.Team{}

	return resource, nil

}
