/*
 * aecdeltas
 *
 * The AEC Deltas OpenAPI Spec
 *
 * API version: 0.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"net/http"

	"github.com/gin-gonic/gin"

	// WARNING!
	// Change this to a fully-qualified import path
	// once you place this file into your project.
	// For example,
	//
	//    store "github.com/myname/myrepo/pkg/store"
	//
	store "github.com/antoinedao/aecdelta-go-server/pkg/store"

)

// StreamsGet - 
func StreamsGet(c *gin.Context) {
	// Run Auth checks here

	// Run data store operation
	// `resource` must be of type map[string]interface{}
	resource, err := store.StreamsGet(c)

	if err != nil {
		c.AbortWithStatusJSON(err.StatusCode(), err)
	}

	c.JSON(http.StatusOK, resource)
}

// StreamsIdBranchPost - 
func StreamsIdBranchPost(c *gin.Context) {
	// Run Auth checks here

	// Run data store operation
	// `resource` must be of type map[string]interface{}
	resource, err := store.StreamsIdBranchPost(c)

	if err != nil {
		c.AbortWithStatusJSON(err.StatusCode(), err)
	}

	c.JSON(http.StatusOK, resource)
}

// StreamsIdDelete - 
func StreamsIdDelete(c *gin.Context) {
	// Run Auth checks here

	// Run data store operation
	// `resource` must be of type map[string]interface{}
	resource, err := store.StreamsIdDelete(c)

	if err != nil {
		c.AbortWithStatusJSON(err.StatusCode(), err)
	}

	c.JSON(http.StatusOK, resource)
}

// StreamsIdGet - 
func StreamsIdGet(c *gin.Context) {
	// Run Auth checks here

	// Run data store operation
	// `resource` must be of type map[string]interface{}
	resource, err := store.StreamsIdGet(c)

	if err != nil {
		c.AbortWithStatusJSON(err.StatusCode(), err)
	}

	c.JSON(http.StatusOK, resource)
}

// StreamsIdMergePost - 
func StreamsIdMergePost(c *gin.Context) {
	// Run Auth checks here

	// Run data store operation
	// `resource` must be of type map[string]interface{}
	resource, err := store.StreamsIdMergePost(c)

	if err != nil {
		c.AbortWithStatusJSON(err.StatusCode(), err)
	}

	c.JSON(http.StatusOK, resource)
}

// StreamsIdPut - 
func StreamsIdPut(c *gin.Context) {
	// Run Auth checks here

	// Run data store operation
	// `resource` must be of type map[string]interface{}
	resource, err := store.StreamsIdPut(c)

	if err != nil {
		c.AbortWithStatusJSON(err.StatusCode(), err)
	}

	c.JSON(http.StatusOK, resource)
}

// StreamsPost - 
func StreamsPost(c *gin.Context) {
	// Run Auth checks here

	// Run data store operation
	// `resource` must be of type map[string]interface{}
	resource, err := store.StreamsPost(c)

	if err != nil {
		c.AbortWithStatusJSON(err.StatusCode(), err)
	}

	c.JSON(http.StatusOK, resource)
}