package familytree

import (
	"context"

	"github.com/GeovaneCavalcante/tree-genealogical/internal/entity"
)

type Relative struct {
	Type   string
	Level  int
	Person *entity.Person
}

type GenealogyInterface interface {
	BuildFamilyTree(ctx context.Context, parente *entity.Person, persons []*entity.Person, level int) []*Relative
	GetRelatives(ctx context.Context) []*Relative
}

type UseCase interface {
	GetAllFamilyMembers(ctx context.Context, personName string) ([]*Relative, error)
	CalculateKinshipDistance(ctx context.Context, firstPersonName, secondPersonName string) (int, error)
	DetermineRelationship(ctx context.Context, firstPersonName, secondPersonName string) (relationship string, err error)
}
