package factory

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/go-playground/validator"
	"github.com/labstack/echo"
)

var port int64 = 4000
var validate = validator.New()

func init() {
	if envPort := os.Getenv("PORT"); envPort != "" {
		parsedPort, err := strconv.ParseInt(envPort, 10, 64)
		if err != nil {
			log.Fatalf("Invalid PORT: %d\n", parsedPort)
		}

		port = parsedPort
	}
}

type (
	Context struct {
		echo.Context
	}
)

func (c *Context) BindValidate(ptr interface{}) error {
	err := c.Bind(c)
	if err != nil {
		return err
	}

	return c.Validate(ptr)
}

func Start() {

	e := echo.New()

	e.HTTPErrorHandler = func(err error, c echo.Context) {
		switch err := err.(type) {
		case validator.ValidationErrors:
			e.DefaultHTTPErrorHandler(&echo.HTTPError{
				Code:    http.StatusBadRequest,
				Message: err.Error(),
			}, c)
		default:
			e.DefaultHTTPErrorHandler(err, c)
		}
	}

	e.GET("/ping", func(c echo.Context) (err error) {
		c.String(http.StatusOK, fmt.Sprintf("hello:", time.Now()))
		return
	})

	{
		g := e.Group("/api/v1")
		g.GET("/product_inspections", getProductInspections)
		g.POST("/product_inspections", postProductInspection)
	}

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", port)))
}
