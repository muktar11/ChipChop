package postgres

import (
    "context"
    "database/sql"
    "github.com/your_project/internal/domain"
    "github.com/your_project/internal/repository"
    "github.com/your_project/internal/repository/sqlc"
)

type EstimationRepository struct {
    db *sql.DB
    q  *sqlc.Queries
}

func NewEstimationRepository(db *sql.DB) *EstimationRepository {
    return &EstimationRepository{
        db: db,
        q:  sqlc.New(db),
    }
}

func (r *EstimationRepository) CreateEstimation(ctx context.Context, estimation domain.Estimation) error {
    return r.q.CreateEstimation(ctx, sqlc.CreateEstimationParams{
        ProductName: estimation.ProductName,
        Price:       estimation.Price,
        Location:    estimation.Location, // Assuming Location is a JSON type
        UserID:      estimation.UserID,
    })
}

func (r *EstimationRepository) GetEstimationByID(ctx context.Context, id int64) (domain.Estimation, error) {
    row, err := r.q.GetEstimationByID(ctx, id)
    if err != nil {
        return domain.Estimation{}, err
    }
    return domain.Estimation{
        ID:          row.ID,
        ProductName: row.ProductName,
        Price:       row.Price,
        Location:    row.Location,
        UserID:      row.UserID,
    }, nil
}

func (r *EstimationRepository) ListEstimations(ctx context.Context, limit, offset int32) ([]domain.Estimation, error) {
    rows, err := r.q.ListEstimations(ctx, sqlc.ListEstimationsParams{Limit: limit, Offset: offset})
    if err != nil {
        return nil, err
    }
    estimations := make([]domain.Estimation, len(rows))
    for i, row := range rows {
        estimations[i] = domain.Estimation{
            ID:          row.ID,
            ProductName: row.ProductName,
            Price:       row.Price,
            Location:    row.Location,
            UserID:      row.UserID,
        }
    }
    return estimations, nil
}

func (r *EstimationRepository) UpdateEstimation(ctx context.Context, id int64, price *float64) error {
    return r.q.UpdateEstimation(ctx, sqlc.UpdateEstimationParams{
        ID:    id,
        Price: price,
    })
}

func (r *EstimationRepository) DeleteEstimation(ctx context.Context, id int64) error {
    return r.q.DeleteEstimation(ctx, id)
}
