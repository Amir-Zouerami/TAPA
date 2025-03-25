package services

import (
	"github.com/jmoiron/sqlx"
)

type Services struct {
	Dashboard *DashboardService
}

func NewServiceContainer(db *sqlx.DB) *Services {
	return &Services{
		Dashboard: NewDashboardService(db),
	}
}
