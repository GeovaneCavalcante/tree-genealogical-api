package database

import (
	"github.com/GeovaneCavalcante/tree-genealogical/person"
	"github.com/GeovaneCavalcante/tree-genealogical/relationship"
	"github.com/google/uuid"
)

type Database struct {
	Persons       []*person.Person
	Relationships []*relationship.Relationship
}

var database *Database

func New() *Database {
	if database == nil {
		database = &Database{
			Persons:       []*person.Person{},
			Relationships: []*relationship.Relationship{},
		}

		loadGeovaneFamily(database)
	}

	return database
}

func NewPerson(db *Database, name, gender, fatherID, motherID string) *person.Person {
	person := &person.Person{
		ID:     uuid.New().String(),
		Name:   name,
		Gender: gender,
	}

	if fatherID != "" {
		relationship := relationship.Relationship{ID: uuid.New().String(), MainPersonID: person.ID, SecundePersonID: fatherID}
		db.Relationships = append(db.Relationships, &relationship)

	}

	if motherID != "" {
		relationship := relationship.Relationship{ID: uuid.New().String(), MainPersonID: person.ID, SecundePersonID: motherID}
		db.Relationships = append(db.Relationships, &relationship)
	}

	db.Persons = append(db.Persons, person)

	return person
}

func loadGeovaneFamily(db *Database) []*person.Person {

	geruza := NewPerson(db, "Geruza", "F", "", "")
	geova := NewPerson(db, "Geova", "M", "", "")
	pedro := NewPerson(db, "Pedro", "M", "", "")
	iraci := NewPerson(db, "Iraci", "F", "", geruza.ID)
	tereza := NewPerson(db, "Tereza", "F", "", geruza.ID)
	alsimar := NewPerson(db, "Alsimar", "F", pedro.ID, iraci.ID)
	suzamar := NewPerson(db, "Suzamar", "F", pedro.ID, iraci.ID)
	araujo := NewPerson(db, "Araujo", "M", "", "")
	debora := NewPerson(db, "debora", "F", araujo.ID, suzamar.ID)
	bruna := NewPerson(db, "Bruna", "F", "", debora.ID)
	gean := NewPerson(db, "Gean", "M", geova.ID, alsimar.ID)
	bruno := NewPerson(db, "Bruno", "M", gean.ID, "")
	soraia := NewPerson(db, "Soraia", "F", pedro.ID, iraci.ID)

	geovane := NewPerson(db, "Geovane", "M", geova.ID, alsimar.ID)
	victoria := NewPerson(db, "Victoria", "M", "", "")

	cr7 := NewPerson(db, "Cristiano Ronaldo", "M", geovane.ID, victoria.ID)
	oceane := NewPerson(db, "Oceane", "F", geovane.ID, victoria.ID)

	vicovane := NewPerson(db, "Vicovane", "M", cr7.ID, "")

	neymar := NewPerson(db, "Neymar", "M", vicovane.ID, "")

	persons := []*person.Person{
		geova,
		iraci,
		pedro,
		alsimar,
		suzamar,
		soraia,
		gean,
		geovane,
		bruna,
		geruza,
		bruno,
		tereza,
		debora,
		victoria,
		cr7,
		vicovane,
		neymar,
		oceane,
		araujo,
	}

	return persons
}
