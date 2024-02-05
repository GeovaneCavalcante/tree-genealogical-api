package database

import (
	"github.com/GeovaneCavalcante/tree-genealogical/internal/entity"
	"github.com/google/uuid"
)

type Database struct {
	Persons       []entity.Person
	Relationships []entity.Relationship
}

var database *Database

func New() *Database {
	if database == nil {
		database = &Database{
			Persons:       []entity.Person{},
			Relationships: []entity.Relationship{},
		}

		loadGeovaneFamily(database)
		loadDefaultFamily(database)
	}

	return database
}

func NewPerson(db *Database, name, gender, fatherID, motherID string) entity.Person {
	person := entity.Person{
		ID:     uuid.New().String(),
		Name:   name,
		Gender: gender,
	}

	if fatherID != "" {
		NewRelationshipAndLoadDb(db, person.ID, fatherID)
	}

	if motherID != "" {
		NewRelationshipAndLoadDb(db, person.ID, motherID)
	}

	db.Persons = append(db.Persons, person)

	return person
}

func NewRelationshipAndLoadDb(db *Database, mainPersonID, secundePersonID string) entity.Relationship {

	relationship := entity.Relationship{
		ID:              uuid.New().String(),
		MainPersonID:    mainPersonID,
		SecundePersonID: secundePersonID,
	}

	db.Relationships = append(db.Relationships, relationship)
	return relationship

}

func loadDefaultFamily(db *Database) []entity.Person {
	martin := NewPerson(db, "Martin", "M", "", "")
	anastasia := NewPerson(db, "Anastasia", "F", "", "")
	phoebe := NewPerson(db, "Phoebe", "F", martin.ID, anastasia.ID)
	advik := NewPerson(db, "Advik", "M", "", "")
	sonny := NewPerson(db, "Sonny", "M", "", "")
	ann := NewPerson(db, "Ann", "F", sonny.ID, "")
	dunny := NewPerson(db, "Dunny", "M", advik.ID, ann.ID)
	NewRelationshipAndLoadDb(db, dunny.ID, phoebe.ID)
	bruce := NewPerson(db, "Bruce", "M", advik.ID, phoebe.ID)
	NewRelationshipAndLoadDb(db, bruce.ID, ann.ID)
	clark := NewPerson(db, "Clark", "M", "", anastasia.ID)
	oprah := NewPerson(db, "Oprah", "F", "", "")
	ellen := NewPerson(db, "Ellen", "F", "", "")
	eric := NewPerson(db, "Eric", "M", ellen.ID, oprah.ID)
	jacqueline := NewPerson(db, "Jacqueline", "F", clark.ID, eric.ID)
	ariel := NewPerson(db, "Ariel", "F", "", "")
	melody := NewPerson(db, "Melody", "F", eric.ID, ariel.ID)

	persons := []entity.Person{
		martin,
		anastasia,
		phoebe,
		advik,
		sonny,
		ann,
		dunny,
		bruce,
		clark,
		eric,
		jacqueline,
		ariel,
		melody,
	}

	return persons

}

func loadGeovaneFamily(db *Database) []entity.Person {

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

	persons := []entity.Person{
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
