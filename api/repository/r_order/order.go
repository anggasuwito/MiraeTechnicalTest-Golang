package r_order

import (
	"MiraeTechnicalTest-Golang/api/models"
	"database/sql"
)

type OrderRepo struct {
	DB *sql.DB
}

type OrderRepoInterface interface {
	GetAllOrderRepo(offset int, limit int) (res []models.OrderResponse, err error)
	CountTotalOrderRepo() (total int, err error)
}

func NewOrderRepo(db *sql.DB) OrderRepoInterface {
	return &OrderRepo{DB: db}
}
