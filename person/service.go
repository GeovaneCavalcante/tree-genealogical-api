package person

import (
	"context"
	"fmt"

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

func (s *Service) Create(ctx context.Context, person *Person) error {
	logger.Info("[Service] Create person started")

	err := s.repo.Create(ctx, person)
	if err != nil {
		logger.Error("[Service] Create person error: ", err)
		return fmt.Errorf("create person error: %w", err)
	}

	logger.Info("[Service] Create person finished")
	return nil
}

func (s *Service) Get(ctx context.Context, personID string) (*Person, error) {
	logger.Info(fmt.Sprintf("[Service] Get person by personID: %s", personID))

	person, err := s.repo.Get(ctx, personID)
	if err != nil {
		logger.Error(fmt.Sprintf("[Service] Get person by personID: %s error ", personID), err)
		return nil, fmt.Errorf("get person error: %w", err)
	}

	logger.Info(fmt.Sprintf("[Service] Get person service finished for personID: %s", personID))
	return person, nil
}

func (s *Service) List(ctx context.Context, filters map[string]interface{}) ([]*Person, error) {
	logger.Info("[Service] List person started")

	persons, err := s.repo.List(ctx, filters)
	if err != nil {
		logger.Error("[Service] List person error: ", err)
		return nil, fmt.Errorf("list person error: %w", err)
	}

	if len(persons) == 0 {
		logger.Info("[Service] List person not found")
		return nil, nil
	}

	logger.Info("[Service] List person finished")
	return persons, nil
}

func (s *Service) Update(ctx context.Context, personID string, person *Person) error {
	logger.Info(fmt.Sprintf("[Service] Update person started by personID: %s", personID))

	p, err := s.Get(ctx, personID)
	if err != nil {
		logger.Error(fmt.Sprintf("[Service] Update person by personID: %s error", personID), err)
		return fmt.Errorf("update person error: %w", err)
	}

	if p == nil {
		logger.Error(fmt.Sprintf("[Service] Update person by personID: %s not found", personID), nil)
		return fmt.Errorf("update person error: not found")
	}

	err = s.repo.Update(ctx, personID, person)
	if err != nil {
		logger.Error("[Service] Update person error: ", err)
		return fmt.Errorf("update person error: %w", err)
	}

	logger.Info(fmt.Sprintf("[Service] Update person finished by personID: %s", personID))
	return nil
}

func (s *Service) Delete(ctx context.Context, personID string) error {
	logger.Info(fmt.Sprintf("[Service] Delete person started by personID: %s", personID))

	p, err := s.Get(ctx, personID)
	if err != nil {
		logger.Error(fmt.Sprintf("[Service] Delete person by personID: %s error", personID), err)
		return fmt.Errorf("delete person error: %w", err)
	}

	if p == nil {
		logger.Error(fmt.Sprintf("[Service] Delete person by personID: %s not found", personID), nil)
		return fmt.Errorf("delete person error: not found")
	}

	err = s.repo.Delete(ctx, personID)
	if err != nil {
		logger.Error(fmt.Sprintf("[Service] Delete person error by personID: %s", personID), err)
		return fmt.Errorf("delete person error: %w", err)
	}

	logger.Info(fmt.Sprintf("[Service] Delete person finished by personID: %s", personID))
	return nil
}
