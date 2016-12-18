package factory

import (
	"net/http"

	"github.com/labstack/echo"
)

func getProductInspections(c echo.Context) (err error) {
	var inspections []ProductInspection

	err = DB.Select("*").From("product_inspections").QueryStructs(&inspections)

	if err != nil {
		return
	}

	c.JSONPretty(http.StatusOK, inspections, "  ")
	return
}

func postProductInspection(c echo.Context) (err error) {
	var req struct {
		Serial string `json:"serial" validate:"required"`
	}

	err = c.Bind(&req)
	if err != nil {
		return
	}

	productInspection, err := UpsertProductInspection(req.Serial)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, productInspection)
	return
}
