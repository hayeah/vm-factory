package factory

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/labstack/echo"
	_ "github.com/lib/pq"
)

var port int64 = 4000

func init() {
	if envPort := os.Getenv("PORT"); envPort != "" {
		parsedPort, err := strconv.ParseInt(envPort, 10, 64)
		if err != nil {
			log.Fatalf("Invalid PORT: %d\n", parsedPort)
		}

		port = parsedPort
	}
}

func Start() {

	// engine, err := xorm.NewEngine("postgres", "dbname=vmfactory sslmode=disable")
	// if err != nil {
	// 	log.Fatalln("db connection failure:", err)
	// }

	e := echo.New()

	// e := gin.New()

	// e.Use(gin.Logger())
	// e.Use(gin.Recovery())
	// e.Use(func(c *gin.Context) {
	// 	c.Next()
	// 	if c.IsAborted() {
	// 		c.JSON(5)
	// 	}
	// 	if c.Errors != nil {

	// 	}

	// })

	e.GET("/ping", func(c echo.Context) (err error) {
		c.String(http.StatusOK, fmt.Sprintf("hello. time is: %v", time.Now()))
		return
	})

	e.GET("/error", func(c echo.Context) (err error) {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: c.QueryParam("msg"),
		}
	})

	e.POST("/bind", func(c echo.Context) (err error) {
		var req struct {
			Number int    `query:"number"`
			String string `query:"string"`
		}

		err = c.Bind(&req)
		if err != nil {
			return &echo.HTTPError{
				Code:    http.StatusBadRequest,
				Message: "Invalid request body",
			}
		}

		type H map[string]interface{}

		c.JSON(http.StatusOK, H{
			"request": req,
		})

		return
	})

	// e.GET("/ping", func(c *gin.Context) {
	// 	c.String(200, "pong: %s\n", time.Now())
	// })

	// e.GET("/divide", func(c *gin.Context) {
	// 	n, err := strconv.ParseInt(c.DefaultQuery("n", "0"), 10, 64)
	// 	if err != nil {
	// 		c.AbortWithError(500, err)
	// 		return
	// 	}

	// 	i := 10 / n

	// 	c.String(200, "ok: %d", i)
	// })

	// e.GET("/error", func(c *gin.Context) {
	// 	// c.Error()

	// 	// c.AbortWithStatus()
	// 	c.JSON(http.StatusBadRequest, c.Error(fmt.Errorf("test error")))
	// })

	// e.GET("/product_inspections", func(c *gin.Context) {
	// 	var inspections []model.ProductInspections

	// 	err := engine.Find(&inspections)
	// 	if err != nil {
	// 		c.AbortWithError(500, err)
	// 		return
	// 	}

	// 	c.IndentedJSON(200, inspections)
	// })

	// e.POST("/product_inspections", func(c *gin.Context) {
	// 	// TODO validate

	// 	var req struct {
	// 		Serial string `json:serial`
	// 	}

	// 	if err := c.BindJSON(&req); err != nil {
	// 		c.Error(fmt.Errorf("Invalid request"))
	// 		c.Abort()
	// 		return
	// 	}

	// })

	// e.Run()
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", port)))
}