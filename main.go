package main

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
)

const father string = "Father"
const son string = "Son"
const brother string = "Brother"
const grandFather string = "GrandFather"
const greatUncle string = "GreatUncle"
const greatGrandfather string = "GreatGrandFather"
const greatGreatGrandfather string = "GreatGreatGrandFather"
const uncle string = "Uncle"
const cousin string = "Cousin"
const nephew string = "Nephew"
const grandSon string = "GrandSon"
const greatGrandSon string = "GreatGrandSon"
const unknownRelation string = "Unknown Relation"

var kinshipTypes = map[string]map[string]string{
	father: {
		"F": "Mother",
		"M": "Father",
	},
	son: {
		"F": "Daughter",
		"M": "Son",
	},
	brother: {
		"F": "Sister",
		"M": "Brother",
	},
	grandFather: {
		"F": "GrandMother",
		"M": "GrandFather",
	},
	greatUncle: {
		"F": "GreatAunt",
		"M": "GreatUncle",
	},
	greatGrandfather: {
		"F": "GreatGrandMother",
		"M": "GreatGrandFather",
	},
	greatGreatGrandfather: {
		"F": "GreatGreatGrandMother",
		"M": "GreatGreatGrandFather",
	},
	uncle: {
		"F": "Aunt",
		"M": "Uncle",
	},
	cousin: {
		"F": "Cousin",
		"M": "Cousin",
	},
	nephew: {
		"F": "Niece",
		"M": "Nephew",
	},
	grandSon: {
		"F": "Granddaughter",
		"M": "GrandSon",
	},
	greatGrandSon: {
		"F": "GreatGranddaughter",
		"M": "GreatGrandson",
	},
}

type Relationship struct {
	MainPerson    string
	SecundePerson string
}

type Person struct {
	ID            string
	Name          string
	Sex           string
	Level         int
	Relationships []Relationship
}

var persons []*Person

type Relative struct {
	Type   string
	Level  int
	Person *Person
}

type TreeGenealogical struct {
	Root      *Person
	Relatives []*Relative
}

func NewPerson(name, sex string, pai, mae string) *Person {
	person := &Person{
		ID:   uuid.New().String(),
		Name: name,
		Sex:  sex,
	}

	if pai != "" {
		person.Relationships = append(person.Relationships, Relationship{
			MainPerson:    person.ID,
			SecundePerson: pai,
		})
	}

	if mae != "" {
		person.Relationships = append(person.Relationships, Relationship{
			MainPerson:    person.ID,
			SecundePerson: mae,
		})
	}

	return person
}

func main() {

	geruza := NewPerson("Geruza", "F", "", "")
	geova := NewPerson("Geova", "M", "", "")
	pedro := NewPerson("Pedro", "M", "", "")
	iraci := NewPerson("Iraci", "F", "", geruza.ID)
	tereza := NewPerson("Tereza", "F", "", geruza.ID)
	alsimar := NewPerson("Alsimar", "F", pedro.ID, iraci.ID)
	suzamar := NewPerson("Suzamar", "F", pedro.ID, iraci.ID)
	araujo := NewPerson("Araujo", "M", "", "")
	debora := NewPerson("debora", "F", araujo.ID, suzamar.ID)
	bruna := NewPerson("Bruna", "F", "", debora.ID)
	gean := NewPerson("Gean", "M", geova.ID, alsimar.ID)
	bruno := NewPerson("Bruno", "M", gean.ID, "")
	soraia := NewPerson("Soraia", "F", pedro.ID, iraci.ID)

	geovane := NewPerson("Geovane", "M", geova.ID, alsimar.ID)
	victoria := NewPerson("Victoria", "M", "", "")

	cr7 := NewPerson("Cristiano Ronaldo", "M", geovane.ID, victoria.ID)
	oceane := NewPerson("Oceane", "F", geovane.ID, victoria.ID)

	vicovane := NewPerson("Vicovane", "M", cr7.ID, "")

	neymar := NewPerson("Neymar", "M", vicovane.ID, "")

	persons = []*Person{
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

	enviado := geovane
	treeGenealogical := NewFamilyTree(enviado, persons)

	fmt.Printf("My name is %s\n", enviado.Name)
	fmt.Println("My family tree :")

	for _, relative := range treeGenealogical.Relatives {
		formater := "--"
		multipliedCharacter := strings.Repeat(formater, relative.Level)
		if relative.Person == nil {
			fmt.Println(relative)
			continue
		}
		fmt.Printf("%s %s: %s\n", multipliedCharacter, relative.Person.Name, relative.Type)
	}

}

func (tg *TreeGenealogical) findParents(relative *Person) *Relative {
	if relative == nil {
		return nil
	}
	for _, p := range tg.Relatives {
		if p.Person == nil {
			continue
		}
		for _, relationship := range p.Person.Relationships {
			if relative.ID == relationship.SecundePerson {
				return p
			}
		}
	}
	return nil
}

func (tg *TreeGenealogical) findChildren(relative *Person) *Relative {
	if relative == nil {
		return nil
	}
	for _, p := range tg.Relatives {
		for _, relationship := range relative.Relationships {
			if p.Person == nil {
				continue
			}
			if p.Person.ID == relationship.SecundePerson {
				return p
			}
		}
	}
	return nil
}

func (tg *TreeGenealogical) directRelationDescription(relative *Person) string {
	if relative == nil {
		return ""
	}
	for _, relationship := range relative.Relationships {
		if tg.Root.ID == relationship.SecundePerson {
			return descriptionBySex(son, tg.Root.Sex)
		}
	}

	for _, relationship := range tg.Root.Relationships {
		if relationship.SecundePerson == relative.ID {
			return descriptionBySex(father, relative.Sex)
		}
		for _, relationshipParent := range relative.Relationships {
			if relationshipParent.SecundePerson == relationship.SecundePerson {
				secundePerson := findPerson(relationship.SecundePerson, persons)
				if secundePerson != nil {
					return descriptionBySex(brother, secundePerson.Sex)
				}
			}
		}
	}
	return ""
}

func descriptionBySex(relative, sex string) string {
	if relation, ok := kinshipTypes[relative]; ok {
		if desc, ok := relation[sex]; ok {
			return desc
		}
		return unknownRelation
	}
	return unknownRelation
}

func isTypeOfTypeKinship(typeKinship, p string) bool {
	if _, ok := kinshipTypes[p]; !ok {
		return false
	}
	for _, value := range kinshipTypes[p] {
		if typeKinship == value {
			return true
		}
	}
	return false
}

func (tg *TreeGenealogical) relationshipDescription(relative *Person) string {

	description := tg.directRelationDescription(relative)

	if description != "" {
		return description
	}

	parent := tg.findParents(relative)
	rulesParents := map[string]string{
		father:      grandFather,
		grandFather: greatGrandfather,
	}

	if parent != nil {
		for key, value := range rulesParents {
			if isTypeOfTypeKinship(parent.Type, key) {
				return descriptionBySex(value, relative.Sex)
			}
		}
	}

	child := tg.findChildren(relative)

	rulesChild := map[string]string{
		grandFather:      uncle,
		uncle:            cousin,
		cousin:           nephew,
		brother:          nephew,
		greatGrandfather: greatUncle,
		nephew:           nephew,
		son:              grandSon,
		grandSon:         greatGrandSon,
	}

	if child != nil {
		for key, value := range rulesChild {
			if child != nil && isTypeOfTypeKinship(child.Type, key) {
				return descriptionBySex(value, relative.Sex)
			}
		}
	}

	return unknownRelation
}

func findPerson(ID string, persons []*Person) *Person {
	for _, person := range persons {
		if person.ID == ID {
			return person
		}
	}
	return nil
}

func NewFamilyTree(root *Person, persons []*Person) *TreeGenealogical {
	tg := &TreeGenealogical{
		Root: root,
	}
	tg.BuildFamilyTree(root, persons, 1)
	return tg
}

func (tg *TreeGenealogical) BuildFamilyTree(parente *Person, persons []*Person, level int) {
	var relatives []*Relative
	relatives = tg.searchDescendants(tg.Root, persons, level, relatives)
	fmt.Printf("My name is %s\n", tg.Root.Name)
	fmt.Println("My family tree :")
	for _, relative := range relatives {
		formater := "--"
		multipliedCharacter := strings.Repeat(formater, relative.Level)
		if relative.Person == nil {
			fmt.Println(relative)
			continue
		}
		fmt.Printf("%s %s: %s\n", multipliedCharacter, relative.Person.Name, relative.Type)
	}
	//tg.searchAncestors(tg.Root, persons, level)
}

// func (tg *TreeGenealogical) searchAncestors(relative *Person, persons []*Person, level int) {
// 	if relative == nil {
// 		return
// 	}
// 	for _, relationship := range relative.Relationships {
// 		secundePerson := findPerson(relationship.SecundePerson, persons)
// 		tg.addRelative(secundePerson, level)
// 		tg.searchForRelatives(secundePerson, persons, level+1)
// 		tg.searchAncestors(secundePerson, persons, level+1)
// 	}

// }

func (tg *TreeGenealogical) searchDescendants(relative *Person, persons []*Person, level int, relatives []*Relative) []*Relative {
	if relative == nil {
		return relatives
	}
	for _, person := range persons {
		for _, relationship := range person.Relationships {
			if relationship.SecundePerson == relative.ID {
				re := tg.addRelative(person, level, relatives)
				relatives = append(relatives, re...)
				r := tg.searchDescendants(person, persons, level+1, relatives)
				relatives = append(relatives, r...)
			}
		}
	}
	return relatives
}

// func (tg *TreeGenealogical) searchForRelatives(relative *Person, persons []*Person, level int) {
// 	if relative == nil {
// 		return
// 	}
// 	for _, person := range persons {
// 		for _, relationship := range person.Relationships {
// 			if relationship.SecundePerson == relative.ID {
// 				tg.addRelative(person, level)
// 				tg.searchForRelatives(person, persons, level+1)
// 			}
// 		}
// 	}
// }

func (tg *TreeGenealogical) addRelative(person *Person, level int, relatives []*Relative) []*Relative {

	relative := &Relative{
		Type:   tg.relationshipDescription(person),
		Level:  level,
		Person: person,
	}
	tg.Relatives = append(tg.Relatives, relative)
	if !tg.alreadyInFamily(person, relatives) {
		relatives = append(relatives, relative)
	}
	return relatives
	// return nil
}

func (tg *TreeGenealogical) alreadyInFamily(relative *Person, relatives []*Relative) bool {
	for _, p := range relatives {
		if p.Person == nil {
			continue
		}
		if p.Person.ID == relative.ID {
			return true
		}
	}
	return false
}
