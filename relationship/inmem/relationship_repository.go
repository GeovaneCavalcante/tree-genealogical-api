package inmem

import (
	"fmt"

	"github.com/GeovaneCavalcante/tree-genealogical/pkg/logger"
	"github.com/GeovaneCavalcante/tree-genealogical/relationship"
)

type RelationshipRepository struct {
	Relationships []*relationship.Relationship
}

func NewInmemRepository(relationships []*relationship.Relationship) *RelationshipRepository {
	return &RelationshipRepository{
		Relationships: relationships,
	}
}

func (r *RelationshipRepository) Create(relationship *relationship.Relationship) error {
	logger.Info("[Repository] Create relationship started")
	r.Relationships = append(r.Relationships, relationship)
	logger.Info("[Repository] Create relationship finished")
	return nil
}

func (r *RelationshipRepository) Get(relationshipID string) (*relationship.Relationship, error) {
	logger.Info(fmt.Sprint("[Repository] Get relationship by relationshipID: ", relationshipID))
	for _, r := range r.Relationships {
		if r.ID == relationshipID {
			return r, nil
		}
	}
	logger.Info(fmt.Sprintf("[Repository] Get relationship by relationshipID: %s not found", relationshipID))
	return nil, nil
}

func (r *RelationshipRepository) List(filters map[string]interface{}) ([]*relationship.Relationship, error) {
	logger.Info("[Repository] List relationship started")

	relationships := r.Relationships

	logger.Info("[Repository] List relationship finished")
	return relationships, nil
}

func (r *RelationshipRepository) Update(relationshipID string, relationship *relationship.Relationship) error {
	logger.Info(fmt.Sprintf("[Repository] Update relationship started by relationshipID: %s", relationshipID))
	for i, rr := range r.Relationships {
		if rr.ID == relationshipID {
			r.Relationships[i] = relationship
			return nil
		}
	}
	logger.Info(fmt.Sprintf("[Repository] Update relationship by relationshipID: %s not found", relationshipID))
	return nil
}

func (r *RelationshipRepository) Delete(relationshipID string) error {
	logger.Info(fmt.Sprintf("[Repository] Delete relationship started by relationshipID: %s", relationshipID))
	for i, rr := range r.Relationships {
		if rr.ID == relationshipID {
			r.Relationships = append(r.Relationships[:i], r.Relationships[i+1:]...)
			return nil
		}
	}
	logger.Info(fmt.Sprintf("[Repository] Delete relationship by relationshipID: %s not found", relationshipID))
	return nil
}
