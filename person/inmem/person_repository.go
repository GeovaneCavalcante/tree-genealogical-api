package inmem

import (
	"context"
	"fmt"
	"strings"

	"github.com/GeovaneCavalcante/tree-genealogical/database"
	"github.com/GeovaneCavalcante/tree-genealogical/person"
	"github.com/GeovaneCavalcante/tree-genealogical/pkg/logger"
	"github.com/GeovaneCavalcante/tree-genealogical/relationship"
	"github.com/google/uuid"
)

type PersonRepository struct {
	InmenDB *database.Database
}

func NewPersonRepository(inmenDB *database.Database) *PersonRepository {
	return &PersonRepository{
		InmenDB: inmenDB,
	}
}

func (r *PersonRepository) Create(ctx context.Context, person *person.Person) error {
	logger.Info("[Repository] Create person started")
	person.ID = uuid.New().String()
	person.Relationships = []relationship.Relationship{}
	r.InmenDB.Persons = append(r.InmenDB.Persons, person)
	logger.Info("[Repository] Create person finished")
	return nil
}

func (r *PersonRepository) Get(ctx context.Context, personID string) (*person.Person, error) {
	logger.Info(fmt.Sprintf("[Repository] Get person by personID: %s", personID))
	for _, p := range r.InmenDB.Persons {
		if p.ID == personID {
			return p, nil
		}
	}
	logger.Info(fmt.Sprintf("[Repository] Get person by personID: %s not found", personID))
	return nil, nil
}

func (r *PersonRepository) GetByName(ctx context.Context, name string) (*person.Person, error) {
	logger.Info(fmt.Sprintf("[Repository] Get person by name: %s", name))
	for _, p := range r.InmenDB.Persons {
		if strings.EqualFold(p.Name, name) {
			for _, r := range r.InmenDB.Relationships {
				if r.MainPersonID == p.ID {
					p.Relationships = append(p.Relationships, *r)
				}
			}
			return p, nil
		}
	}
	logger.Info(fmt.Sprintf("[Repository] Get person by name: %s not found", name))
	return nil, nil
}

func (r *PersonRepository) List(ctx context.Context, filters map[string]interface{}) ([]*person.Person, error) {
	logger.Info("[Repository] List person started")

	persons := r.InmenDB.Persons

	logger.Info("[Repository] List person finished")
	return persons, nil
}

func (r *PersonRepository) ListWithRelationships(ctx context.Context, filters map[string]interface{}) ([]*person.Person, error) {
	logger.Info("[Repository] List person with relationships started")

	var persons []*person.Person
	for _, p := range r.InmenDB.Persons {
		for _, r := range r.InmenDB.Relationships {
			if r.MainPersonID == p.ID {
				p.Relationships = append(p.Relationships, *r)
			}
		}
		persons = append(persons, p)
	}
	return persons, nil

}

func (r *PersonRepository) Update(ctx context.Context, personID string, person *person.Person) error {
	logger.Info(fmt.Sprintf("[Repository] Update person started by personID: %s", personID))
	for i, p := range r.InmenDB.Persons {
		if p.ID == personID {
			person.ID = p.ID
			r.InmenDB.Persons[i] = person
			return nil
		}
	}
	logger.Info(fmt.Sprintf("[Repository] Update person by personID: %s not found", personID))
	return nil
}

func (r *PersonRepository) Delete(ctx context.Context, personID string) error {
	logger.Info(fmt.Sprintf("[Repository] Delete person started by personID: %s", personID))
	for i, p := range r.InmenDB.Persons {
		if p.ID == personID {
			r.InmenDB.Persons = append(r.InmenDB.Persons[:i], r.InmenDB.Persons[i+1:]...)
			return nil
		}
	}
	logger.Info(fmt.Sprintf("[Repository] Delete person by personID: %s not found", personID))
	return nil
}
