package familytree

import (
	"context"

	"github.com/GeovaneCavalcante/tree-genealogical/person"
)

type Relative struct {
	Type   string
	Level  int
	Person *person.Person
}

type GenealogyInterface interface {
	BuildFamilyTree(ctx context.Context, parente *person.Person, persons []*person.Person, level int) []*Relative
	GetRelatives(ctx context.Context) []*Relative
}

type UseCase interface {
	GetAllFamilyMembers(ctx context.Context, personName string) ([]*Relative, error)
	CalculateKinshipDistance(ctx context.Context, firstPersonName, secondPersonName string) (int, error)
	DetermineRelationship(ctx context.Context, firstPersonName, secondPersonName string) (relationship string, err error)
}
