package store

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func ProjectsGet(c *gin.Context, project interface{}, stream string, name string, permission string) ConnectionError {

	fmt.Println(project)
	fmt.Println(stream)
	fmt.Println(name)
	fmt.Println(permission)

	return nil

}

func ProjectsIdDelete(c *gin.Context, id string) ConnectionError {

	fmt.Println(id)

	return nil
}

func ProjectsIdGet(c *gin.Context, project interface{}, id string) ConnectionError {

	return nil
}

func ProjectsIdPut(c *gin.Context, id string, update interface{}) ConnectionError {

	return nil

}

func ProjectsPost(c *gin.Context, created interface{}, newProject interface{}) ConnectionError {

	return nil

}
