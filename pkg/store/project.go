package store

import (
	model "github.com/antoinedao/aecdelta-go-server/pkg/model"
	"github.com/gin-gonic/gin"
)

func ProjectsGet(c *gin.Context) ([]model.Project, ConnectionError) {

	resource := []model.Project{}

	return resource, nil

}

func ProjectsIdDelete(c *gin.Context) (string, ConnectionError) {

	resource := "Accepted"

	return resource, nil

}

func ProjectsIdGet(c *gin.Context) (model.Project, ConnectionError) {

	resource := model.Project{}

	return resource, nil

}

func ProjectsIdPut(c *gin.Context) (model.Project, ConnectionError) {

	resource := model.Project{}

	return resource, nil

}

func ProjectsPost(c *gin.Context) (model.Project, ConnectionError) {

	resource := model.Project{}

	return resource, nil

}
