package main

import (
	"fmt"
	"net/http"
	"os"
	"runtime"

	"github.com/go-pg/pg/v10"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/weijinnx/boosty-test/lib/db"
	"github.com/weijinnx/boosty-test/lib/errors"
	"github.com/weijinnx/boosty-test/lib/util"
	"github.com/weijinnx/boosty-test/lib/web"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		panic("fail to load .env file")
	}
}

func appRecover() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// defer the protection function
			defer func() {
				if r := recover(); r != nil {
					err, ok := r.(error)
					if !ok { // create an error if we don't have one
						err = fmt.Errorf("%v", r)
					}
					stackSize := 4 << 10 // 4KB
					stack := make([]byte, stackSize)
					length := runtime.Stack(stack, true)
					fmt.Printf(
						"[PANIC RECOVER] err: %v\n%s\n", err, stack[:length],
					)
					c.Error(err)
				}
			}()
			return next(c)
		}
	}
}

func main() {
	// new echo framework instance
	e := echo.New()
	e.HTTPErrorHandler = errors.ErrorHandler

	conn := pg.Connect(&pg.Options{
		Addr:     "db:"+os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASS"),
		Database: os.Getenv("DB_NAME"),
	})
	defer conn.Close()

	// pretty recover
	e.Use(appRecover())
	// attach custom app context
	e.Use(util.AttachContextMiddleware(conn))

	// routes
	e.GET("/", func(c echo.Context) error {
		cctx := c.Get("cctx").(*util.AppContext)
		wallets, err := db.GetWallets(cctx.DB)
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"success": true,
			"wallets": wallets,
		})
	})
	e.POST("/transaction", web.TransferFunds)

	// start server
	e.Logger.Fatal(e.Start(":8080"))
}