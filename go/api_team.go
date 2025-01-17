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

// TeamsGet - Query teams and get a list. User will only see lists they have &#x60;read&#x60; permissions to.
func TeamsGet(c *gin.Context) {
	// Get stream query parameter
	stream := c.DefaultQuery("stream", "")

	// Get project query parameter
	project := c.DefaultQuery("project", "")

	// Get role query parameter
	role := c.DefaultQuery("role", "")

	// Get permission query parameter
	permission := c.DefaultQuery("permission", "")

	// Initialise response object
	response := []Team{}

	// Execute operation from the store package
	err := store.TeamsGet(c, &response, stream, project, role, permission)

	if err != nil {
		c.AbortWithStatusJSON(err.StatusCode(), err)
	}

	c.JSON(200, response)
}

// TeamsIdDelete - Delete a team. User must have admin access to the team.
func TeamsIdDelete(c *gin.Context) {
	// Get id path parameter
	id := c.Param("id")

	// Set default response
	response := "Accepted"

	// Execute operation from the store package
	err := store.TeamsIdDelete(c, id)

	if err != nil {
		c.AbortWithStatusJSON(err.StatusCode(), err)
	}

	c.JSON(202, response)
}

// TeamsIdGet - Get a single team by ID. User must have &#x60;read&#x60; access to the team.
func TeamsIdGet(c *gin.Context) {
	// Get id path parameter
	id := c.Param("id")

	// Initialise response object
	response := Team{}

	// Execute operation from the store package
	err := store.TeamsIdGet(c, &response, id)

	if err != nil {
		c.AbortWithStatusJSON(err.StatusCode(), err)
	}

	c.JSON(200, response)
}

// TeamsIdPut - Update a team&#39;s information. User must have &#x60;write&#x60; access to the team.
func TeamsIdPut(c *gin.Context) {
	// Get id path parameter
	id := c.Param("id")

	// Get newTeam payload
	var newTeam NewTeam
	c.BindJSON(&newTeam)

	// Set default response
	response := "Accepted"

	// Execute operation from the store package
	err := store.TeamsIdPut(c, id, newTeam)

	if err != nil {
		c.AbortWithStatusJSON(err.StatusCode(), err)
	}

	c.JSON(202, response)
}

// TeamsPost - Create a new team.
func TeamsPost(c *gin.Context) {
	// Get newTeam payload
	var newTeam NewTeam
	c.BindJSON(&newTeam)

	// Initialise response object
	response := Created{}

	// Execute operation from the store package
	err := store.TeamsPost(c, &response, newTeam)

	if err != nil {
		c.AbortWithStatusJSON(err.StatusCode(), err)
	}

	c.JSON(201, response)
}
