package response

import (
	"github.com/gin-gonic/gin"
	internalError "go-ewallet/pkg/error"
	"net/http"
)

type HttpResponse struct {
	Data      interface{} `json:"data,omitempty"`
	Message   string      `json:"message,omitempty"`
	IsSuccess bool        `json:"is_success"`
}

func Error(c *gin.Context, err error) {
	httpStatusCode := http.StatusInternalServerError
	hr := HttpResponse{
		IsSuccess: false,
		Message:   err.Error(),
	}

	if err, ok := err.(*internalError.Err); ok {
		hr.Message = err.Error()
		httpStatusCode = err.GetHTTPStatusCode()
	}
	c.AbortWithStatusJSON(httpStatusCode, hr)
}

func Success(c *gin.Context, data any, httpStatusCode int) {
	httpStatusCodeResponse := http.StatusOK
	if httpStatusCode != 0 {
		httpStatusCodeResponse = httpStatusCode
	}
	hr := HttpResponse{
		IsSuccess: true,
		Data:      data,
		Message:   "success",
	}

	c.JSON(httpStatusCodeResponse, hr)
}
