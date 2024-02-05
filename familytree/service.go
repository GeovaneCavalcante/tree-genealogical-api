package familytree

import (
	"context"
	"fmt"
	"strings"

	"github.com/GeovaneCavalcante/tree-genealogical/internal/entity"
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

func (s *Service) GetAllFamilyMembers(ctx context.Context, personName string) ([]*entity.Relative, error) {
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

func (s *Service) DetermineRelationship(ctx context.Context, firstPersonName, secondPersonName string) (relationship string, err error) {
	logger.Info(fmt.Sprintf("[Service] DetermineRelationship started for firstPersonName: %s and secondPersonName: %s", firstPersonName, secondPersonName))

	firstPerson, err := s.PersonRepo.GetByName(ctx, firstPersonName)
	if err != nil {
		logger.Error(fmt.Sprintf("[Service] DetermineRelationship error for firstPersonName: %s and secondPersonName: %s", firstPersonName, secondPersonName), err)
		return "", fmt.Errorf("get person error: %w", err)
	}

	persons, err := s.PersonRepo.ListWithRelationships(ctx, nil)
	if err != nil {
		logger.Error(fmt.Sprintf("[Service] DetermineRelationship error for firstPersonName: %s and secondPersonName: %s", firstPersonName, secondPersonName), err)
		return "", fmt.Errorf("get person error: %w", err)
	}
	relatives := s.Genealogy.BuildFamilyTree(ctx, firstPerson, persons, 1)

	if len(relatives) == 0 {
		logger.Info(fmt.Sprintf("[Service] DetermineRelationship finished for firstPersonName: %s and secondPersonName: %s", firstPersonName, secondPersonName))
		return "unrelated", nil
	}

	for _, relative := range relatives {
		if strings.EqualFold(relative.Person.Name, secondPersonName) {
			logger.Info(fmt.Sprintf("[Service] DetermineRelationship finished for firstPersonName: %s and secondPersonName: %s", firstPersonName, secondPersonName))
			return relative.Type, nil
		}
	}

	return "", nil
}

func (s *Service) CalculateKinshipDistance(ctx context.Context, firstPersonName, secondPersonName string) (int, error) {
	logger.Info(fmt.Sprintf("[Service] CalculateKinshipDistance started for firstPersonName: %s and secondPersonName: %s", firstPersonName, secondPersonName))
	firstPerson, err := s.PersonRepo.GetByName(ctx, firstPersonName)
	if err != nil {
		logger.Error(fmt.Sprintf("[Service] CalculateKinshipDistance error for firstPersonName: %s and secondPersonName: %s", firstPersonName, secondPersonName), err)
		return 0, fmt.Errorf("get person error: %w", err)
	}
	persons, err := s.PersonRepo.ListWithRelationships(ctx, nil)
	if err != nil {
		logger.Error(fmt.Sprintf("[Service] CalculateKinshipDistance error for firstPersonName: %s and secondPersonName: %s", firstPersonName, secondPersonName), err)
		return 0, fmt.Errorf("get person error: %w", err)
	}

	relatives := s.Genealogy.BuildFamilyTree(ctx, firstPerson, persons, 1)

	if len(relatives) == 0 {
		logger.Info(fmt.Sprintf("[Service] CalculateKinshipDistance finished for firstPersonName: %s and secondPersonName: %s", firstPersonName, secondPersonName))
		return 0, nil
	}

	for _, relative := range relatives {
		if strings.EqualFold(relative.Person.Name, secondPersonName) {
			logger.Info(fmt.Sprintf("[Service] CalculateKinshipDistance finished for firstPersonName: %s and secondPersonName: %s", firstPersonName, secondPersonName))
			return relative.Level, nil
		}
	}

	return 0, nil
}
