package usecase

import (
	"MiraeTechnicalTest-Golang/api/repository"
	"MiraeTechnicalTest-Golang/api/usecase/uc_order"
	"database/sql"
)

type UseCase struct {
	OrderUC uc_order.OrderUCInterface
}

func NewUseCase(db *sql.DB) UseCase {
	repo := repository.NewRepository(db)
	return UseCase{
		OrderUC: uc_order.NewOrderUC(repo),
	}
}
