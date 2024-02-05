package genealogy

import (
	"context"
	"testing"

	"github.com/GeovaneCavalcante/tree-genealogical/internal/entity"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type GenealogyTestSuite struct {
	suite.Suite
	persons []*entity.Person
	root    *entity.Person
}

func NewPerson(name, gender, fatherID, motherID string) *entity.Person {
	person := &entity.Person{
		ID:     uuid.New().String(),
		Name:   name,
		Gender: gender,
	}

	if fatherID != "" {
		person.Relationships = append(person.Relationships, &entity.Relationship{MainPersonID: person.ID, SecundePersonID: fatherID})
	}

	if motherID != "" {
		person.Relationships = append(person.Relationships, &entity.Relationship{MainPersonID: person.ID, SecundePersonID: motherID})
	}

	return person
}

func loadPersons() ([]*entity.Person, *entity.Person) {
	martin := NewPerson("Martin", "M", "", "")
	anastasia := NewPerson("Anastasia", "F", "", "")
	phoebe := NewPerson("Phoebe", "F", martin.ID, anastasia.ID)
	bruce := NewPerson("Bruce", "M", "", phoebe.ID)

	return []*entity.Person{martin, anastasia, phoebe, bruce}, phoebe
}

func (suite *GenealogyTestSuite) SetupTest() {

	persons, root := loadPersons()
	suite.persons = persons
	suite.root = root

}

func (suite *GenealogyTestSuite) TestBuildFamilyTree() {
	ctx := context.Background()

	suite.Run("should return the populated Root and Relative properties", func() {
		familytree := NewFamilyTree()
		family := familytree.BuildFamilyTree(ctx, suite.root, suite.persons, 0)
		assert.Equal(suite.T(), "Phoebe", familytree.Root.Name)
		assert.NotEmpty(suite.T(), family)
		assert.NotEmpty(suite.T(), familytree.Relatives)
	})

	suite.Run("should return the Relative and family property only with root", func() {
		familytree := NewFamilyTree()
		family := familytree.BuildFamilyTree(ctx, suite.root, []*entity.Person{}, 0)
		assert.Len(suite.T(), family, 1)
		assert.Len(suite.T(), familytree.Relatives, 1)
		assert.Equal(suite.T(), "Phoebe", family[0].Person.Name)
		assert.Equal(suite.T(), "Root", family[0].Type)
	})
}

func (suite *GenealogyTestSuite) TestGetRelatives() {
	ctx := context.Background()

	suite.Run("should return the relatives of the root", func() {
		familytree := NewFamilyTree()
		_ = familytree.BuildFamilyTree(ctx, suite.root, suite.persons, 0)
		relatives := familytree.GetRelatives(ctx)
		assert.Len(suite.T(), relatives, 4)
	})
}

func (suite *GenealogyTestSuite) TestSearchDescendants() {
	ctx := context.Background()

	suite.Run("should return the descendants of the root", func() {
		familytree := NewFamilyTree()
		familytree.Root = suite.root
		relatives := []*entity.Relative{}
		descendants := familytree.searchDescendants(ctx, familytree.Root, suite.persons, 0, relatives)
		assert.Len(suite.T(), descendants, 1)
		assert.Equal(suite.T(), "Bruce", descendants[0].Person.Name)
	})

	suite.Run("should must return the original relatives when the alanised relative was nil", func() {
		familytree := NewFamilyTree()
		familytree.Root = suite.root
		relatives := []*entity.Relative{}
		descendants := familytree.searchDescendants(ctx, nil, suite.persons, 0, relatives)
		assert.Len(suite.T(), descendants, 0)
	})
}

func (suite *GenealogyTestSuite) TestSearchAncestors() {
	ctx := context.Background()
	suite.Run("should return the ancestors of the root", func() {
		familytree := NewFamilyTree()
		familytree.Root = suite.root
		relatives := []*entity.Relative{}
		ancestors := familytree.searchAncestors(ctx, suite.root, suite.persons, 0, relatives)
		assert.Len(suite.T(), ancestors, 2)
		assert.Equal(suite.T(), "Martin", ancestors[0].Person.Name)
		assert.Equal(suite.T(), "Anastasia", ancestors[1].Person.Name)
	})

	suite.Run("should must return the original relatives when the alanised relative was nil", func() {
		familytree := NewFamilyTree()
		familytree.Root = suite.root
		relatives := []*entity.Relative{}
		ancestors := familytree.searchAncestors(ctx, nil, suite.persons, 0, relatives)
		assert.Len(suite.T(), ancestors, 0)
	})
}

func (suite *GenealogyTestSuite) TestSearchForRelatives() {
	ctx := context.Background()
	suite.Run("should return the relatives of the root", func() {
		familytree := NewFamilyTree()
		familytree.Root = suite.root
		relatives := []*entity.Relative{}
		relatives = familytree.searchForRelatives(ctx, suite.root, suite.persons, 0, relatives)
		assert.Len(suite.T(), relatives, 1)
		assert.Equal(suite.T(), "Bruce", relatives[0].Person.Name)
	})

	suite.Run("should must return the original relatives when the alanised relative was nil", func() {
		familytree := NewFamilyTree()
		familytree.Root = suite.root
		relatives := []*entity.Relative{}
		relatives = familytree.searchForRelatives(ctx, nil, suite.persons, 0, relatives)
		assert.Len(suite.T(), relatives, 0)
	})
}

func (suite *GenealogyTestSuite) TestRelationshipDescription() {
	roberta := &entity.Person{
		ID:     "4",
		Name:   "Roberta",
		Gender: "F",
	}

	anastasia := &entity.Person{
		ID:     "2",
		Name:   "Anastasia",
		Gender: "F",
		Relationships: []*entity.Relationship{
			{
				MainPersonID:    "2",
				SecundePersonID: roberta.ID,
			},
		},
	}
	frida := &entity.Person{
		ID:     "5",
		Name:   "Frida",
		Gender: "F",
		Relationships: []*entity.Relationship{
			{
				MainPersonID:    "5",
				SecundePersonID: roberta.ID,
			},
		},
	}
	suite.Run("should return the value of the direct relationship with relative", func() {

		suite.persons = append(suite.persons, anastasia)

		suite.root.Relationships = []*entity.Relationship{
			{
				MainPersonID:    "3",
				SecundePersonID: "2",
			},
		}

		relatives := []*entity.Relative{}

		familytree := NewFamilyTree()
		familytree.Root = suite.root
		description := familytree.relationshipDescription(anastasia, relatives, suite.persons)
		assert.Equal(suite.T(), "Mother", description)
	})

	suite.Run("should returns root's relationship to a new relative based on parent rules", func() {

		anastasia.Relationships = []*entity.Relationship{
			{
				MainPersonID:    "2",
				SecundePersonID: roberta.ID,
			},
		}

		suite.persons = append(suite.persons, anastasia, roberta)

		familytree := NewFamilyTree()
		familytree.Root = suite.root
		familytree.Root.Relationships = []*entity.Relationship{
			{
				MainPersonID:    suite.root.ID,
				SecundePersonID: anastasia.ID,
			},
		}
		relatives := []*entity.Relative{
			{
				Type:   "Mother",
				Person: anastasia,
				Level:  1,
			},
		}
		description := familytree.relationshipDescription(roberta, relatives, suite.persons)
		assert.Equal(suite.T(), "GrandMother", description)
	})

	suite.Run("should returns root's relationship to a new relative based on parent rules", func() {

		suite.persons = append(suite.persons, anastasia, roberta, frida)

		familytree := NewFamilyTree()
		familytree.Root = suite.root
		familytree.Root.Relationships = []*entity.Relationship{
			{
				MainPersonID:    suite.root.ID,
				SecundePersonID: anastasia.ID,
			},
		}
		relatives := []*entity.Relative{
			{
				Type:   "Mother",
				Person: anastasia,
				Level:  1,
			},
			{
				Type:   "GrandMother",
				Person: roberta,
				Level:  2,
			},
		}
		description := familytree.relationshipDescription(frida, relatives, suite.persons)
		assert.Equal(suite.T(), "Aunt", description)
	})

	suite.Run("should returns unknown relation when the relative is not in the family", func() {

		suite.persons = append(suite.persons, anastasia, roberta)

		suite.root.Relationships = []*entity.Relationship{
			{
				MainPersonID:    suite.root.ID,
				SecundePersonID: anastasia.ID,
			},
		}

		familytree := NewFamilyTree()
		familytree.Root = suite.root
		relatives := []*entity.Relative{}
		description := familytree.relationshipDescription(roberta, relatives, suite.persons)
		assert.Equal(suite.T(), "Unknown Relation", description)
	})
}

func (suite *GenealogyTestSuite) TestFindParents() {
	ruff := &entity.Person{
		ID:     "6",
		Name:   "Ruff",
		Gender: "M",
	}
	suite.Run("should return the parents of the root", func() {
		familytree := NewFamilyTree()
		relatives := familytree.BuildFamilyTree(context.Background(), suite.root, suite.persons, 0)
		suite.root.Relationships[0].SecundePersonID = ruff.ID
		parents := familytree.findParents(ruff, relatives)
		assert.Equal(suite.T(), "Root", parents.Type)
	})

	suite.Run("should return empty when the root has no parents", func() {
		familytree := NewFamilyTree()
		relatives := familytree.BuildFamilyTree(context.Background(), suite.root, suite.persons, 0)
		relatives[0].Person = nil
		suite.root.Relationships[0].SecundePersonID = ruff.ID
		parents := familytree.findParents(ruff, relatives)
		assert.Empty(suite.T(), parents)
	})

	suite.Run("should return empty aa when the root has no parents", func() {
		familytree := NewFamilyTree()
		relatives := []*entity.Relative{}
		parents := familytree.findParents(nil, relatives)
		assert.Empty(suite.T(), parents)
	})
}

func (suite *GenealogyTestSuite) TestDirectRelationDescription() {
	suite.Run("should must return an empty string when parent is nil", func() {
		familytree := NewFamilyTree()
		description := familytree.directRelationDescription(nil, suite.persons)
		assert.Equal(suite.T(), "", description)
	})
}

func (suite *GenealogyTestSuite) TestCheckSiblingRelation() {
	ctx := context.Background()
	suite.Run("should return root's sibling", func() {
		familytree := NewFamilyTree()
		_ = familytree.BuildFamilyTree(ctx, suite.root, suite.persons, 0)
		ruff := &entity.Person{
			ID:            "6",
			Name:          "Ruff",
			Gender:        "M",
			Relationships: suite.root.Relationships,
		}
		sibling := familytree.checkSiblingRelation(ruff, suite.persons)
		assert.Equal(suite.T(), "Brother", sibling)
	})

}

func (suite *GenealogyTestSuite) TestFindChildren() {
	suite.Run("should return the parents of the root", func() {
		familytree := NewFamilyTree()
		relatives := familytree.BuildFamilyTree(context.Background(), suite.root, suite.persons, 0)
		relatives[0].Person = nil
		parents := familytree.findChildren(suite.root, relatives)
		assert.NotNil(suite.T(), parents)
	})

	suite.Run("should return empty aa when the root has no parents", func() {
		familytree := NewFamilyTree()
		relatives := []*entity.Relative{}
		parents := familytree.findChildren(nil, relatives)
		assert.Empty(suite.T(), parents)
	})
}

func (suite *GenealogyTestSuite) TestAlreadyInFamily() {
	suite.Run("should return the parents of the root", func() {
		familytree := NewFamilyTree()
		relatives := []*entity.Relative{
			{
				Type:   "Mother",
				Person: nil,
				Level:  1,
			},
		}
		result := familytree.alreadyInFamily(suite.root, relatives)
		assert.False(suite.T(), result)
	})
}

func (suite *GenealogyTestSuite) TestDescriptionBySex() {
	suite.Run("should return the description unknown", func() {
		description := descriptionBySex("greatgretgrandfather", "F")
		assert.Equal(suite.T(), "Unknown Relation", description)
	})

	suite.Run("should return the description to gender unknown", func() {
		description := descriptionBySex("Father", "a")
		assert.Equal(suite.T(), "Unknown Relation", description)
	})
}

func (suite *GenealogyTestSuite) TestIsTypeOfTypeKinship() {
	suite.Run("should return the type invalid", func() {
		check := isTypeOfTypeKinship("greatgretgrandfather", "F")
		assert.Equal(suite.T(), false, check)
	})

	suite.Run("should return the type unknown", func() {
		check := isTypeOfTypeKinship("Father", "a")
		assert.Equal(suite.T(), false, check)
	})

}

func TestSuite(t *testing.T) {
	suite.Run(t, new(GenealogyTestSuite))
}
