package store

import (
	"github.com/gin-gonic/gin"
)

func StreamsGet(c *gin.Context, streams interface{}, name string, project string, schema string) ConnectionError {

	return nil

}

func StreamsIdBranchPost(c *gin.Context, stream interface{}, id string, branchPayload interface{}) ConnectionError {

	return nil

}

func StreamsIdDelete(c *gin.Context, id string) ConnectionError {

	return nil

}

func StreamsIdGet(c *gin.Context, stream interface{}, id string) ConnectionError {

	return nil

}

func StreamsIdMergePost(c *gin.Context, stream interface{}, id string, mergePayload interface{}) ConnectionError {

	return nil

}

func StreamsIdPut(c *gin.Context, id string, update interface{}) ConnectionError {

	return nil

}

func StreamsPost(c *gin.Context, created interface{}, new interface{}) ConnectionError {

	return nil

}
