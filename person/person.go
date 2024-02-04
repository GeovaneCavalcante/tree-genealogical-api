package person

import (
	"context"

	"github.com/GeovaneCavalcante/tree-genealogical/internal/entity"
)

type Repository interface {
	Create(ctx context.Context, person *entity.Person) error
	Get(ctx context.Context, ID string) (*entity.Person, error)
	GetByName(ctx context.Context, name string) (*entity.Person, error)
	List(ctx context.Context, filters map[string]interface{}) ([]*entity.Person, error)
	ListWithRelationships(ctx context.Context, filters map[string]interface{}) ([]*entity.Person, error)
	Update(ctx context.Context, ID string, person *entity.Person) error
	Delete(ctx context.Context, ID string) error
}

type UseCase interface {
	Create(ctx context.Context, person *entity.Person) error
	Get(ctx context.Context, ID string) (*entity.Person, error)
	List(ctx context.Context, filters map[string]interface{}) ([]*entity.Person, error)
	Update(ctx context.Context, ID string, person *entity.Person) error
	Delete(ctx context.Context, ID string) error
}
