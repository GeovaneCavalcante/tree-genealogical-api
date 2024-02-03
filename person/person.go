package person

import (
	"context"

	"github.com/GeovaneCavalcante/tree-genealogical/relationship"
)

type Person struct {
	ID            string
	Name          string
	Sex           string
	Level         int
	Relationships []relationship.Relationship
}

type Repository interface {
	Create(ctx context.Context, person *Person) error
	Get(ctx context.Context, ID string) (*Person, error)
	List(ctx context.Context, filters map[string]interface{}) ([]*Person, error)
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
