package controller

import (
	"ginexample/server_error"

	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

var isReleaseMode = gin.Mode() == gin.ReleaseMode

type GinAppError struct {
	Msg  string `json:"msg"`
	Code int    `json:"code"`
	Data any    `json:"data"`
}

type GinAppData[T any] struct {
	Msg  string `json:"msg"`
	Code int    `json:"code"`
	Data T      `json:"data"`
}

func GinHandleError(handle func(c *gin.Context) error) gin.HandlerFunc {
	return GinHandleResultError(func(c *gin.Context) (any, error) {
		return nil, handle(c)
	})
}

func GinHandleResultError(handle func(c *gin.Context) (any, error)) gin.HandlerFunc {
	return func(c *gin.Context) {
		result, err := handle(c)
		if err != nil {
			GinAbortWithError(c, err)
			return
		}
		if result != nil {
			GinExitWithResult(c, result)
			return
		}
		GinExitOK(c)
	}
}

func GinHandleDownloadImage(handle func(c *gin.Context) (string, string, error)) gin.HandlerFunc {
	return func(c *gin.Context) {
		imgFilePath, imgFileName, err := handle(c)
		if err != nil {
			GinAbortWithError(c, err)
			return
		}
		c.FileAttachment(filepath.Join(imgFilePath, imgFileName), imgFileName)
	}
}

func GinExitWithResult[T any](c *gin.Context, result T) {
	c.JSON(http.StatusOK, GinAppData[T]{
		Msg:  "OK",
		Code: 1,
		Data: result,
	})
}

func GinExitOK(c *gin.Context) {
	c.JSON(http.StatusOK, &GinAppError{Msg: "OK", Code: 1})
}

func GinAbortWithError(c *gin.Context, err error) {
	if ginError := c.Errors; len(ginError) > 0 {
		for _, e := range ginError {
			panic(e)
		}
		if !isReleaseMode {
			c.Writer.Write([]byte(err.Error()))
			c.Writer.Flush()
		}
		return
	}
	if appErr, ok := err.(server_error.ServerError); ok {
		c.AbortWithStatusJSON(http.StatusOK, GinAppError{
			Code: appErr.Code(),
			Msg:  appErr.Msg(),
			Data: nil,
		})
		panic(err)
	}
	if appErr, ok := err.(server_error.BadRequestError); ok {
		c.AbortWithStatusJSON(http.StatusOK, GinAppError{
			Code: appErr.Code(),
			Msg:  appErr.Msg(),
			Data: nil,
		})
		panic(err)
	}
	c.AbortWithStatusJSON(http.StatusOK, GinAppError{
		Code: -1,
		Msg:  err.Error(),
		Data: nil,
	})
	panic(err)
}
