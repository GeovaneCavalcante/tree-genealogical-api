package person

import (
	"context"

	"github.com/GeovaneCavalcante/tree-genealogical/relationship"
)

type Person struct {
	ID            string                      `json:"id"`
	Name          string                      `json:"name"`
	Gender        string                      `json:"gender"`
	Level         int                         `json:"level"`
	Relationships []relationship.Relationship `json:"relationships"`
}

type Repository interface {
	Create(ctx context.Context, person *Person) error
	Get(ctx context.Context, ID string) (*Person, error)
	GetByName(ctx context.Context, name string) (*Person, error)
	List(ctx context.Context, filters map[string]interface{}) ([]*Person, error)
	ListWithRelationships(ctx context.Context, filters map[string]interface{}) ([]*Person, error)
	Update(ctx context.Context, ID string, person *Person) error
	Delete(ctx context.Context, ID string) error
}

type UseCase interface {
	Create(ctx context.Context, person *Person) error
	Get(ctx context.Context, ID string) (*Person, error)
	List(ctx context.Context, filters map[string]interface{}) ([]*Person, error)
	Update(ctx context.Context, ID string, person *Person) error
	Delete(ctx context.Context, ID string) error
}
