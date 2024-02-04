package familytree

import (
	"context"
	"fmt"

	"github.com/GeovaneCavalcante/tree-genealogical/person"
	"github.com/GeovaneCavalcante/tree-genealogical/pkg/logger"
	"github.com/GeovaneCavalcante/tree-genealogical/relationship"
)

type Service struct {
	Genealogy        GenealogyInterface
	PersonRepo       person.Repository
	RelationshipRepo relationship.Repository
}

func NewService(genealogy GenealogyInterface, personRepo person.Repository, relationshipRepo relationship.Repository) *Service {
	return &Service{
		Genealogy:        genealogy,
		PersonRepo:       personRepo,
		RelationshipRepo: relationshipRepo,
	}
}

func (s *Service) GetAllFamilyMembers(ctx context.Context, personName string) ([]*Relative, error) {
	logger.Info(fmt.Sprintf("[Service] GetAllFamilyMembers started for personName: %s", personName))

	person, err := s.PersonRepo.GetByName(ctx, personName)
	if err != nil {
		logger.Error(fmt.Sprintf("[Service] GetAllFamilyMembers error for personName: %s", personName), err)
		return nil, fmt.Errorf("get person error: %w", err)
	}

	persons, err := s.PersonRepo.ListWithRelationships(ctx, nil)

	if err != nil {
		logger.Error(fmt.Sprintf("[Service] GetAllFamilyMembers error for personName: %s", personName), err)
		return nil, fmt.Errorf("get person error: %w", err)
	}

	relatives := s.Genealogy.BuildFamilyTree(ctx, person, persons, 0)

	logger.Info(fmt.Sprintf("[Service] GetAllFamilyMembers finished for personName: %s", personName))
	return relatives, nil
}

func (s *Service) CalculateKinshipDistance(ctx context.Context, firstPersonID, secondPersonID string) (int, error) {
	// logger.Info(fmt.Sprintf("[Service] CalculateKinshipDistance started for firstPersonID: %s and secondPersonID: %s", firstPersonID, secondPersonID))

	// relatives, err := s.GetRelatives(ctx, firstPersonID, secondPersonID)
	// if err != nil {
	// 	logger.Error(fmt.Sprintf("[Service] CalculateKinshipDistance error for firstPersonID: %s and secondPersonID: %s", firstPersonID, secondPersonID), err)
	// 	return 0, fmt.Errorf("get relatives error: %w", err)
	// }

	// distance := s.calculateDistance(relatives)

	// logger.Info(fmt.Sprintf("[Service] CalculateKinshipDistance finished for firstPersonID: %s and secondPersonID: %s", firstPersonID, secondPersonID))
	// return distance, nil

	return 0, nil
}

func (s *Service) DetermineRelationship(ctx context.Context, firstPersonID, secondPersonID string) (relationship string, err error) {
	// logger.Info(fmt.Sprintf("[Service] DetermineRelationship started for firstPersonID: %s and secondPersonID: %s", firstPersonID, secondPersonID))

	// relatives, err := s.GetRelatives(ctx, firstPersonID, secondPersonID)
	// if err != nil {
	// 	logger.Error(fmt.Sprintf("[Service] DetermineRelationship error for firstPersonID: %s and secondPersonID: %s", firstPersonID, secondPersonID), err)
	// 	return "", fmt.Errorf("get relatives error: %w", err)
	// }

	// relationship = s.determineRelationship(relatives)

	// logger.Info(fmt.Sprintf("[Service] DetermineRelationship finished for firstPersonID: %s and secondPersonID: %s", firstPersonID, secondPersonID))
	// return relationship, nil

	return "", nil
}
