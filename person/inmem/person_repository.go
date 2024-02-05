package inmem

import (
	"context"
	"fmt"
	"strings"

	"github.com/GeovaneCavalcante/tree-genealogical/database"
	"github.com/GeovaneCavalcante/tree-genealogical/internal/entity"
	"github.com/GeovaneCavalcante/tree-genealogical/pkg/logger"
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

func (r *PersonRepository) Create(ctx context.Context, person *entity.Person) error {
	logger.Info("[Repository] Create person started")
	person.ID = uuid.New().String()
	person.Relationships = []*entity.Relationship{}
	r.InmenDB.Persons = append(r.InmenDB.Persons, *person)
	logger.Info("[Repository] Create person finished")
	return nil
}

func (r *PersonRepository) Get(ctx context.Context, personID string) (*entity.Person, error) {
	logger.Info(fmt.Sprintf("[Repository] Get person by personID: %s", personID))

	p := findByID(r.InmenDB.Persons, personID)
	if p == nil {
		logger.Info(fmt.Sprintf("[Repository] Get person by personID: %s not found", personID))
		return nil, fmt.Errorf("person not found")
	}
	logger.Info(fmt.Sprintf("[Repository] Get person by personID: %s not found", personID))
	return p, nil
}

func (r *PersonRepository) GetByName(ctx context.Context, name string) (*entity.Person, error) {
	logger.Info(fmt.Sprintf("[Repository] Get person by name: %s", name))
	for _, p := range r.InmenDB.Persons {
		if strings.EqualFold(p.Name, name) {
			person := p
			for _, rr := range r.InmenDB.Relationships {
				if rr.MainPersonID == person.ID {
					rr.MainPerson = &person
					rr.SecundePerson = findByID(r.InmenDB.Persons, rr.SecundePersonID)
					relationship := rr
					person.Relationships = append(person.Relationships, &relationship)
				}
			}
			return &person, nil
		}
	}
	logger.Info(fmt.Sprintf("[Repository] Get person by name: %s not found", name))
	return nil, nil
}

func (r *PersonRepository) List(ctx context.Context, filters map[string]interface{}) ([]*entity.Person, error) {
	logger.Info("[Repository] List person started")

	var persons []*entity.Person

	for _, p := range r.InmenDB.Persons {
		person := p
		persons = append(persons, &person)
	}

	logger.Info("[Repository] List person finished")
	return persons, nil
}

func (r *PersonRepository) ListWithRelationships(ctx context.Context, filters map[string]interface{}) ([]*entity.Person, error) {
	logger.Info("[Repository] List person with relationships started")

	var persons []*entity.Person
	for _, p := range r.InmenDB.Persons {
		person := p
		for _, rr := range r.InmenDB.Relationships {
			if rr.MainPersonID == person.ID {
				rr.MainPerson = &person
				rr.SecundePerson = findByID(r.InmenDB.Persons, rr.SecundePersonID)
				relationship := rr
				person.Relationships = append(person.Relationships, &relationship)
			}
		}
		persons = append(persons, &person)
	}
	return persons, nil

}

func (r *PersonRepository) Update(ctx context.Context, personID string, person *entity.Person) error {
	logger.Info(fmt.Sprintf("[Repository] Update person started by personID: %s", personID))
	for i, p := range r.InmenDB.Persons {
		if p.ID == personID {
			person.ID = p.ID
			r.InmenDB.Persons[i] = *person
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

func findByID(persons []entity.Person, id string) *entity.Person {
	for _, p := range persons {
		if p.ID == id {
			return &p
		}
	}
	return nil
}
