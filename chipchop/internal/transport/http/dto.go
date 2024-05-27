package http

import "github.com/your_project/internal/domain"

type CreateEstimationRequest struct {
    ProductName string         `json:"product_name" binding:"required"`
    Price       float64        `json:"price" binding:"required"`
    Location    domain.Location `json:"location" binding:"required"`
    UserID      int64          `json:"user_id" binding:"required"`
}

type UpdateEstimationRequest struct {
    Price *float64 `json:"price"`
}
