package errors

import (
	"fmt"
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	echo "github.com/labstack/echo/v4"
)

// AppError standard app error struct
type AppError struct {
	HTTPCode         int                `json:"http_code"`
	Code             int                `json:"error_code"`
	Message          string             `json:"message"`
	Payload          interface{}        `json:"payload,omitempty"`
	ValidationErrors *validation.Errors `json:"validation_errors,omitempty"`
}

func (e AppError) String() string {
	if e.Message == "" {
		e.Message = "undefined internal error"
	}
	return fmt.Sprintf("%d %s %v", e.Code, e.Message, e.Payload)
}

func (e AppError) Error() string {
	return e.String()
}

// NewAppError create error object
func NewAppError(httpCode, code int, err error, payload interface{}) *AppError {
	appError := &AppError{
		HTTPCode: httpCode,
		Code:     code,
		Message:  "internal error: check console logs to get more info",
	}
	
	if err != nil {
		appError.Message = err.Error()
		if vErrs, ok := err.(validation.Errors); ok {
			appError.ValidationErrors = &vErrs
		}
	}
	if payload != nil {
		appError.Payload = payload
	}

	return appError
}

// Prep prepares app error object by error code
func Prep(code int, payload interface{}) *AppError {
	if err, ok := payload.(error); ok {
		payload = err.Error()
	}

	return NewAppError(getHTTPCode(code), code, fmt.Errorf(getTextCode(code)), payload)
}

func getHTTPCode(code int) int {
	httpCode := http.StatusInternalServerError
	if code > 200 && code < http.StatusNetworkAuthenticationRequired { // last 511 http error code
		httpCode = code
	}
	if val, ok := CodeMap[code]; ok {
		httpCode = val
	}
	return httpCode
}

func getTextCode(code int) string {
	err := "unspecified error"
	if text, ok := CodeText[code]; ok {
		err = text
	}
	return err
}

// Response http with app error object
func (e *AppError) Response(ctx echo.Context) error {
	return ctx.JSON(e.HTTPCode, e)
}