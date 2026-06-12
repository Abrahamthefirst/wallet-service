package custom_http

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var (
	RecordNotFound        = errors.New("Record Not Found")
	ErrConflict           = errors.New("")
	BadRequestException   = errors.New("Bad request")
	ErrInvalidCredentials = errors.New("Invalid Credentials")
	ErrInternal           = errors.New("Internal Server Error")
)

func ValidateRequest(c *gin.Context, dto any) error {
	validate := validator.New()

	if err := c.ShouldBindJSON(dto); err != nil {

		slog.WarnContext(c.Request.Context(), "request validation failed",
			"error", err.Error(),
			"path", c.Request.URL.Path,
		)

		c.JSON(http.StatusBadRequest, ErrorResponse{
			Message:    "all fields required",
			StatusCode: http.StatusBadRequest,
		})
		return err
	}

	if err := validate.Struct(dto); err != nil {

		switch errs := err.(type) {
		case validator.ValidationErrors:
			for _, err := range errs {
				slog.WarnContext(c.Request.Context(), "request validation failed",
					"error", err.Error(),
					"path", c.Request.URL.Path,
				)

				c.JSON(http.StatusBadRequest, ErrorResponse{
					Message:    fmt.Sprintf("%s validation failed", err.Field()),
					StatusCode: http.StatusBadRequest,
				})
				return err
			}

		case *validator.InvalidValidationError:
			slog.Error("Validator received invalid input", "error", err.Error())
			c.JSON(http.StatusInternalServerError, ErrorResponse{
				Message:    "Internal validation error",
				StatusCode: http.StatusInternalServerError,
			})
			return err

		default:
			slog.Error("Validator received invalid input", "error", err.Error())

			c.JSON(http.StatusBadRequest, ErrorResponse{
				Message:    "Validation failed",
				StatusCode: http.StatusBadRequest,
			})
			return err

		}

	}

	return nil
}



func HandleError(c *gin.Context, err error) {
	switch {

	case errors.Is(err, ErrInvalidCredentials):
		c.JSON(http.StatusUnauthorized, ErrorResponse{
			Message:    err.Error(),
			StatusCode: http.StatusUnauthorized,
		})
	case errors.Is(err, ErrConflict):
		c.JSON(http.StatusConflict, ErrorResponse{
			Message:    err.Error(),
			StatusCode: http.StatusConflict,
		})

	case errors.Is(err, context.DeadlineExceeded):
		c.JSON(http.StatusRequestTimeout, ErrorResponse{
			Message:    "The server took too long to respond. Please try again.",
			StatusCode: 408,
		})

	default:
		c.JSON(http.StatusInternalServerError,
			ErrorResponse{
				Message:    "Internal Server Error",
				StatusCode: 500,
			})
	}
}
