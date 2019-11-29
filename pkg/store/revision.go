package store

import (
	model "github.com/antoinedao/aecdelta-go-server/pkg/model"
	"github.com/gin-gonic/gin"
)

func RevisionsGet(c *gin.Context) ([]model.Revision, ConnectionError) {

	resource := []model.Revision{}

	return resource, nil

}

func RevisionsIdDelete(c *gin.Context) (string, ConnectionError) {

	resource := "resource"

	return resource, nil

}

func RevisionsIdDiffGet(c *gin.Context) (model.Diff, ConnectionError) {

	resource := model.Diff{}

	return resource, nil

}

func RevisionsIdGet(c *gin.Context) (model.Revision, ConnectionError) {

	resource := model.Revision{}

	return resource, nil

}

func RevisionsPost(c *gin.Context) (model.Revision, ConnectionError) {

	resource := model.Revision{}

	return resource, nil

}
