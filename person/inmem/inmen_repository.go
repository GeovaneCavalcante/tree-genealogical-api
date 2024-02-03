package inmem

import (
	"context"
	"fmt"

	"github.com/GeovaneCavalcante/tree-genealogical/person"
	"github.com/GeovaneCavalcante/tree-genealogical/pkg/logger"
)

type inmemRepository struct {
	Persons []*person.Person
}

func NewInmemRepository(persons []*person.Person) *inmemRepository {
	return &inmemRepository{
		Persons: persons,
	}
}

func (r *inmemRepository) Creat(ctx context.Context, person *person.Person) error {
	logger.Info("[Repository] Create person started")
	r.Persons = append(r.Persons, person)
	logger.Info("[Repository] Create person finished")
	return nil
}

func (r *inmemRepository) Get(ctx context.Context, personID string) (*person.Person, error) {
	logger.Info(fmt.Sprintf("[Repository] Get person by personID: %s", personID))
	for _, p := range r.Persons {
		if p.ID == personID {
			return p, nil
		}
	}
	logger.Info(fmt.Sprintf("[Repository] Get person by personID: %s not found", personID))
	return nil, nil
}

func (r *inmemRepository) List(ctx context.Context, filters map[string]interface{}) ([]*person.Person, error) {
	logger.Info("[Repository] List person started")

	persons := r.Persons

	logger.Info("[Repository] List person finished")
	return persons, nil
}

func (r *inmemRepository) Update(ctx context.Context, personID string, person *person.Person) error {
	logger.Info(fmt.Sprintf("[Repository] Update person started by personID: %s", personID))
	for i, p := range r.Persons {
		if p.ID == personID {
			r.Persons[i] = person
			return nil
		}
	}
	logger.Info(fmt.Sprintf("[Repository] Update person by personID: %s not found", personID))
	return nil
}

func (r *inmemRepository) Delete(ctx context.Context, personID string) error {
	logger.Info(fmt.Sprintf("[Repository] Delete person started by personID: %s", personID))
	for i, p := range r.Persons {
		if p.ID == personID {
			r.Persons = append(r.Persons[:i], r.Persons[i+1:]...)
			return nil
		}
	}
	logger.Info(fmt.Sprintf("[Repository] Delete person by personID: %s not found", personID))
	return nil
}
