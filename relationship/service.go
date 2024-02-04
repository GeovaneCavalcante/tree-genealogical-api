package relationship

import (
	"context"
	"fmt"

	"github.com/GeovaneCavalcante/tree-genealogical/internal/entity"
	"github.com/GeovaneCavalcante/tree-genealogical/pkg/logger"
)

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) Create(ctx context.Context, relationship *entity.Relationship) error {
	logger.Info("[Service] Create relationship started")

	err := s.repo.Create(ctx, relationship)
	if err != nil {
		logger.Error("[Service] Create relationship error: ", err)
		return fmt.Errorf("create relationship error: %w", err)
	}

	logger.Info("[Service] Create relationship finished")
	return nil
}

func (s *Service) Get(ctx context.Context, relationshipID string) (*entity.Relationship, error) {
	logger.Info(fmt.Sprintf("[Service] Get relationship by relationshipID: %s", relationshipID))

	relationship, err := s.repo.Get(ctx, relationshipID)
	if err != nil {
		logger.Error(fmt.Sprintf("[Service] Get relationship by relationshipID: %s error ", relationshipID), err)
		return nil, fmt.Errorf("get relationship error: %w", err)
	}

	logger.Info(fmt.Sprintf("[Service] Get relationship service finished for relationshipID: %s", relationshipID))
	return relationship, nil
}

func (s *Service) List(ctx context.Context, filters map[string]interface{}) ([]*entity.Relationship, error) {
	logger.Info("[Service] List relationship started")

	relationships, err := s.repo.List(ctx, filters)
	if err != nil {
		logger.Error("[Service] List relationship error: ", err)
		return nil, fmt.Errorf("list relationship error: %w", err)
	}

	logger.Info("[Service] List relationship finished")
	return relationships, nil
}

func (s *Service) Update(ctx context.Context, relationshipID string, relationship *entity.Relationship) error {
	logger.Info(fmt.Sprintf("[Service] Update relationship started by relationshipID: %s", relationshipID))

	r, err := s.Get(ctx, relationshipID)
	if err != nil {
		logger.Error(fmt.Sprintf("[Service] Update relationship by relationshipID: %s error ", relationshipID), err)
		return fmt.Errorf("update relationship error: %w", err)
	}

	if r == nil {
		logger.Error(fmt.Sprintf("[Service] Update relationship by relationshipID %s error not found ", relationshipID), nil)
		return fmt.Errorf("relationship not found")
	}

	err = s.repo.Update(ctx, relationshipID, relationship)
	if err != nil {
		logger.Error(fmt.Sprintf("[Service] Update relationship by relationshipID: %s error ", relationshipID), err)
		return fmt.Errorf("update relationship error: %w", err)
	}

	logger.Info(fmt.Sprintf("[Service] Update relationship service finished for relationshipID: %s", relationshipID))
	return nil
}

func (s *Service) Delete(ctx context.Context, relationshipID string) error {
	logger.Info(fmt.Sprintf("[Service] Delete relationship started by relationshipID: %s", relationshipID))

	r, err := s.Get(ctx, relationshipID)
	if err != nil {
		logger.Error(fmt.Sprintf("[Service] Delete relationship by relationshipID: %s error ", relationshipID), err)
		return fmt.Errorf("delete relationship error: %w", err)
	}

	if r == nil {
		logger.Error(fmt.Sprintf("[Service] Delete relationship by relationshipID %s error not found ", relationshipID), nil)
		return fmt.Errorf("relationship not found")
	}

	err = s.repo.Delete(ctx, relationshipID)
	if err != nil {
		logger.Error(fmt.Sprintf("[Service] Delete relationship by relationshipID: %s error ", relationshipID), err)
		return fmt.Errorf("delete relationship error: %w", err)
	}

	logger.Info(fmt.Sprintf("[Service] Delete relationship service finished for relationshipID: %s", relationshipID))
	return nil
}
