package presenter

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

func TestSuite(t *testing.T) {
	suite.Run(t, new(FamilyTreePresenerTestSuite))
	suite.Run(t, new(PersonPresenerTestSuite))
	suite.Run(t, new(RelationshipPresenerTestSuite))
}
