package relationship

import "context"

type Relationship struct {
	ID            string
	MainPerson    string
	SecundePerson string
}

type Repository interface {
	Create(ctx context.Context, relationship *Relationship) error
	Get(ctx context.Context, ID string) (*Relationship, error)
	List(ctx context.Context, filters map[string]interface{}) ([]*Relationship, error)
	Update(ctx context.Context, ID string, relationship *Relationship) error
	Delete(ctx context.Context, ID string) error
}

type UseCase interface {
	Create(ctx context.Context, relationship *Relationship) error
	Get(ctx context.Context, ID string) (*Relationship, error)
	List(ctx context.Context, filters map[string]interface{}) ([]*Relationship, error)
	Update(ctx context.Context, ID string, relationship *Relationship) error
	Delete(ctx context.Context, ID string) error
}
