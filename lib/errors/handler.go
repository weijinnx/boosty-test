package errors

import (
	"net/http"

	echo "github.com/labstack/echo/v4"
)

// ErrorHandler is a default http errors handler
func ErrorHandler(err error, ctx echo.Context) {
	// handle default echo error
	if echoErr, ok := err.(*echo.HTTPError); ok {
		Prep(echoErr.Code, err).Response(ctx)
		return
	}

	// handle AppError
	if appError, ok := err.(*AppError); ok {
		appError.Response(ctx)
		return
	}

	// unhandled error
	Prep(http.StatusInternalServerError, err).Response(ctx)
}
