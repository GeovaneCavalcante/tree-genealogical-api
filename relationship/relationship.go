package relationship

import (
	"context"

	"github.com/GeovaneCavalcante/tree-genealogical/internal/entity"
)

type Repository interface {
	Create(ctx context.Context, relationship *entity.Relationship) error
	Get(ctx context.Context, ID string) (*entity.Relationship, error)
	List(ctx context.Context, filters map[string]interface{}) ([]*entity.Relationship, error)
	Update(ctx context.Context, ID string, relationship *entity.Relationship) error
	Delete(ctx context.Context, ID string) error
}

type UseCase interface {
	Create(ctx context.Context, relationship *entity.Relationship) error
	Get(ctx context.Context, ID string) (*entity.Relationship, error)
	List(ctx context.Context, filters map[string]interface{}) ([]*entity.Relationship, error)
	Update(ctx context.Context, ID string, relationship *entity.Relationship) error
	Delete(ctx context.Context, ID string) error
}
