package factory

import "time"

type ProductInspection struct {
	Id        string    `json:"id" db:"id"`
	Serial    string    `json:"serial" db:"serial"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

const (
	ProductInspectionTable = "product_inspections"
)

func UpsertProductInspection(serial string) (productInspection ProductInspection, err error) {
	err = DB.Upsert(ProductInspectionTable).
		Columns("serial", "updated_at").
		Values(serial, time.Now().UTC()).
		Where("serial = $1", serial).
		Returning("*").
		QueryStruct(&productInspection)

	return
}
