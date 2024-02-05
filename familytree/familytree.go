package familytree

import (
	"context"

	"github.com/GeovaneCavalcante/tree-genealogical/internal/entity"
)

type GenealogyInterface interface {
	BuildFamilyTree(ctx context.Context, rootPerson *entity.Person, persons []*entity.Person, level int) []*entity.Relative
	GetRelatives(ctx context.Context) []*entity.Relative
}

type UseCase interface {
	GetAllFamilyMembers(ctx context.Context, personName string) ([]*entity.Relative, error)
	CalculateKinshipDistance(ctx context.Context, firstPersonName, secondPersonName string) (int, error)
	DetermineRelationship(ctx context.Context, firstPersonName, secondPersonName string) (relationship string, err error)
}
