package http

import (
    "github.com/gin-gonic/gin"
    "github.com/your_project/internal/service"
)

func NewRouter(service *service.EstimationService) *gin.Engine {
    handler := NewHandler(service)
    router := gin.Default()
    router.GET("/estimations", handler.ListEstimations)
    router.GET("/estimations/:id", handler.GetEstimationByID)
    router.POST("/estimations", handler.CreateEstimation)
    router.PATCH("/estimations/:id", handler.UpdateEstimation)
    router.DELETE("/estimations/:id", handler.DeleteEstimation)
    return router
}
