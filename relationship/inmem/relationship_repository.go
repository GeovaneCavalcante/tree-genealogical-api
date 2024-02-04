package inmem

import (
	"context"
	"fmt"

	"github.com/GeovaneCavalcante/tree-genealogical/database"
	"github.com/GeovaneCavalcante/tree-genealogical/pkg/logger"
	"github.com/GeovaneCavalcante/tree-genealogical/relationship"
	"github.com/google/uuid"
)

type RelationshipRepository struct {
	InmenDB *database.Database
}

func NewRelationshipRepository(inmenDB *database.Database) *RelationshipRepository {
	return &RelationshipRepository{
		InmenDB: inmenDB,
	}
}

func (r *RelationshipRepository) Create(ctx context.Context, relationship *relationship.Relationship) error {
	logger.Info("[Repository] Create relationship started")
	relationship.ID = uuid.New().String()
	r.InmenDB.Relationships = append(r.InmenDB.Relationships, relationship)
	logger.Info("[Repository] Create relationship finished")
	return nil
}

func (r *RelationshipRepository) Get(ctx context.Context, relationshipID string) (*relationship.Relationship, error) {
	logger.Info(fmt.Sprint("[Repository] Get relationship by relationshipID: ", relationshipID))
	for _, r := range r.InmenDB.Relationships {
		if r.ID == relationshipID {
			return r, nil
		}
	}
	logger.Info(fmt.Sprintf("[Repository] Get relationship by relationshipID: %s not found", relationshipID))
	return nil, nil
}

func (r *RelationshipRepository) List(ctx context.Context, filters map[string]interface{}) ([]*relationship.Relationship, error) {
	logger.Info("[Repository] List relationship started")

	relationships := r.InmenDB.Relationships

	logger.Info("[Repository] List relationship finished")
	return relationships, nil
}

func (r *RelationshipRepository) Update(ctx context.Context, relationshipID string, relationship *relationship.Relationship) error {
	logger.Info(fmt.Sprintf("[Repository] Update relationship started by relationshipID: %s", relationshipID))
	for i, rr := range r.InmenDB.Relationships {
		if rr.ID == relationshipID {
			relationship.ID = rr.ID
			r.InmenDB.Relationships[i] = relationship
			return nil
		}
	}
	logger.Info(fmt.Sprintf("[Repository] Update relationship by relationshipID: %s not found", relationshipID))
	return nil
}

func (r *RelationshipRepository) Delete(ctx context.Context, relationshipID string) error {
	logger.Info(fmt.Sprintf("[Repository] Delete relationship started by relationshipID: %s", relationshipID))
	for i, rr := range r.InmenDB.Relationships {
		if rr.ID == relationshipID {
			r.InmenDB.Relationships = append(r.InmenDB.Relationships[:i], r.InmenDB.Relationships[i+1:]...)
			return nil
		}
	}
	logger.Info(fmt.Sprintf("[Repository] Delete relationship by relationshipID: %s not found", relationshipID))
	return nil
}
