package errors

import (
	"fmt"
	"net/http"
)
import (
	"github.com/gin-gonic/gin"
)

type Error struct {
	Code       string
	Message    string
	HTTPStatus int
}

func (e *Error) Error() string {
	return fmt.Sprintf("code: %s, message: %s", e.Code, e.Message)
}

func InvalidRequestError(message string) *Error {
	return &Error{
		Code:       "InvalidRequestError",
		Message:    message,
		HTTPStatus: http.StatusBadRequest,
	}
}

func InternalError(message string) *Error {
	return &Error{
		Code:       "InternalError",
		Message:    message,
		HTTPStatus: http.StatusInternalServerError,
	}
}

func NoAuthorityError() *Error {
	return &Error{
		Code:       "NoAuthorityError",
		Message:    "No authority",
		HTTPStatus: http.StatusForbidden,
	}
}

func AbortWithWriteErrorResponse(c *gin.Context, err error) {
	if err == nil {
		return
	}

	statusCode := c.Writer.Status()
	var body map[string]string

	switch e := err.(type) {
	case *Error:
		if e.HTTPStatus != 0 {
			statusCode = e.HTTPStatus
		}
		body = map[string]string{
			"code":    e.Code,
			"massage": e.Message,
		}
	default:
		if statusCode == http.StatusOK || statusCode == http.StatusCreated || statusCode == http.StatusAccepted {
			statusCode = http.StatusInternalServerError
		}
		body = map[string]string{
			"code":    "InternalError",
			"message": err.Error(),
		}
	}

	c.AbortWithStatusJSON(statusCode, body)
}
