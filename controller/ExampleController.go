package controller

import (
	example "ginexample/model/Example"
	"ginexample/utils/log"

	"github.com/gin-gonic/gin"
)

func HandleGetExample(c *gin.Context) (any, error) {
	params, err := example.HandleGetExampleParams(c)
	if err != nil {
		log.Errorf("get example params failed, err: %v", err)
		return nil, err
	}

	examples, err := example.GetExamples(params)
	if err != nil {
		log.Errorf("get examples failed, err: %v", err)
		return nil, err
	}
	return examples, nil
}

func HandleGetExampleDetail(c *gin.Context) (any, error) {
	id := c.Param("id")
	example, err := example.GetExample(id)
	if err != nil {
		log.Errorf("get example failed, err: %v", err)
		return nil, err
	}
	return example, nil
}
