package req

import (
	"product_svc/pkg/utils"
	"time"
)

type DeleteId struct {
	ProductID uint `json:"productid" binding:"required,numeric"`
}

type ReqSalesReport struct {
	StartDate  time.Time        `json:"start_date"`
	EndDate    time.Time        `json:"end_date"`
	Pagination utils.Pagination `json:"pagination"`
}
