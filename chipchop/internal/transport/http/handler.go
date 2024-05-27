package http

import (
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
    "github.com/your_project/internal/service"
)

type Handler struct {
    service *service.EstimationService
}

func NewHandler(service *service.EstimationService) *Handler {
    return &Handler{service: service}
}

func (h *Handler) CreateEstimation(c *gin.Context) {
    var req CreateEstimationRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    estimation := domain.Estimation{
        ProductName: req.ProductName,
        Price:       req.Price,
        Location:    req.Location,
        UserID:      req.UserID,
    }
    if err := h.service.CreateEstimation(c.Request.Context(), estimation); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.Status(http.StatusCreated)
}

func (h *Handler) GetEstimationByID(c *gin.Context) {
    id, err := strconv.ParseInt(c.Param("id"), 10, 64)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
        return
    }
    estimation, err := h.service.GetEstimationByID(c.Request.Context(), id)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, estimation)
}

func (h *Handler) ListEstimations(c *gin.Context) {
    limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
    offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))
    estimations, err := h.service.ListEstimations(c.Request.Context(), int32(limit), int32(offset))
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, estimations)
}

func (h *Handler) UpdateEstimation(c *gin.Context) {
    id, err := strconv.ParseInt(c.Param("id"), 10, 64)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
        return
    }
    var req UpdateEstimationRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    if err := h.service.UpdateEstimation(c.Request.Context(), id, req.Price); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.Status(http.StatusNoContent)
}

func (h *Handler) DeleteEstimation(c *gin.Context) {
    id, err := strconv.ParseInt(c.Param("id"), 10, 64)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
        return
    }
    if err := h.service.DeleteEstimation(c.Request.Context(), id); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.Status(http.StatusNoContent)
}
