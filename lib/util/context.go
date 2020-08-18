package util

import (
	"github.com/go-pg/pg/v10"
	"github.com/labstack/echo/v4"
)

// AppContext is a custom context for echo
type AppContext struct {
	Ctx    echo.Context
	DB     *pg.DB
}

// NewAppContext create new context instance
func NewAppContext(ectx echo.Context, db *pg.DB) *AppContext {
	appCtx := &AppContext{
		Ctx:    ectx,
		DB:     db,
	}
	return appCtx
}

// AttachContextMiddleware attaches context to each request
// that will have a pointer to DB connection
func AttachContextMiddleware(db *pg.DB) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("cctx", NewAppContext(c, db))
			return next(c)
		}
	}
}