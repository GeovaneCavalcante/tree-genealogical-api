package main

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
)

const (
	father                string = "Father"
	son                   string = "Son"
	brother               string = "Brother"
	grandFather           string = "GrandFather"
	greatUncle            string = "GreatUncle"
	greatGrandfather      string = "GreatGrandFather"
	greatGreatGrandfather string = "GreatGreatGrandFather"
	uncle                 string = "Uncle"
	cousin                string = "Cousin"
	nephew                string = "Nephew"
	grandSon              string = "GrandSon"
	greatGrandSon         string = "GreatGrandSon"
	unknownRelation       string = "Unknown Relation"
)

var kinshipTypes = map[string]map[string]string{
	father:                {"F": "Mother", "M": "Father"},
	son:                   {"F": "Daughter", "M": "Son"},
	brother:               {"F": "Sister", "M": "Brother"},
	grandFather:           {"F": "GrandMother", "M": "GrandFather"},
	greatUncle:            {"F": "GreatAunt", "M": "GreatUncle"},
	greatGrandfather:      {"F": "GreatGrandMother", "M": "GreatGrandFather"},
	greatGreatGrandfather: {"F": "GreatGreatGrandMother", "M": "GreatGreatGrandFather"},
	uncle:                 {"F": "Aunt", "M": "Uncle"},
	cousin:                {"F": "Cousin", "M": "Cousin"},
	nephew:                {"F": "Niece", "M": "Nephew"},
	grandSon:              {"F": "Granddaughter", "M": "GrandSon"},
	greatGrandSon:         {"F": "GreatGranddaughter", "M": "GreatGrandson"},
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

func NewPerson(name, sex, fatherID, motherID string) *Person {
	person := &Person{
		ID:   uuid.New().String(),
		Name: name,
		Sex:  sex,
	}

	if fatherID != "" {
		person.Relationships = append(person.Relationships, Relationship{MainPerson: person.ID, SecundePerson: fatherID})
	}

	if motherID != "" {
		person.Relationships = append(person.Relationships, Relationship{MainPerson: person.ID, SecundePerson: motherID})
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

func (tg *TreeGenealogical) findParents(relative *Person, relatives []*Relative) *Relative {
	if relative == nil {
		return nil
	}
	for _, p := range relatives {
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

func (tg *TreeGenealogical) findChildren(relative *Person, relatives []*Relative) *Relative {
	if relative == nil {
		return nil
	}
	for _, p := range relatives {
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

func (tg *TreeGenealogical) relationshipDescription(relative *Person, relatives []*Relative) string {

	description := tg.directRelationDescription(relative)

	if description != "" {
		return description
	}

	parent := tg.findParents(relative, relatives)
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

	child := tg.findChildren(relative, relatives)

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

func (tg *TreeGenealogical) BuildFamilyTree(parente *Person, persons []*Person, level int) []*Relative {
	var relatives []*Relative
	relatives = tg.searchDescendants(tg.Root, persons, level, relatives)
	relatives = tg.searchAncestors(tg.Root, persons, level, relatives)
	tg.Relatives = relatives
	return relatives
}

func (tg *TreeGenealogical) searchAncestors(relative *Person, persons []*Person, level int, relatives []*Relative) []*Relative {
	if relative == nil {
		return relatives
	}
	for _, relationship := range relative.Relationships {
		secundePerson := findPerson(relationship.SecundePerson, persons)
		if secundePerson == nil || tg.alreadyInFamily(secundePerson, relatives) {
			continue // Pula para o próximo relacionamento se o parente já estiver na lista ou não for encontrado
		}

		// Adiciona o parente encontrado (ancestral) apenas se não estiver já na lista
		re := tg.addRelative(secundePerson, level, relatives)
		relatives = append(relatives, re)

		// Recursivamente busca por mais ancestrais deste parente encontrado
		relatives = tg.searchForRelatives(secundePerson, persons, level+1, relatives)
		relatives = tg.searchAncestors(secundePerson, persons, level+1, relatives)

	}

	return relatives

}

func (tg *TreeGenealogical) searchDescendants(relative *Person, persons []*Person, level int, relatives []*Relative) []*Relative {
	if relative == nil {
		return relatives
	}
	for _, person := range persons {
		for _, relationship := range person.Relationships {
			if relationship.SecundePerson == relative.ID {
				// Verifica se a pessoa já foi adicionada e não é o Root
				if !tg.alreadyInFamily(person, relatives) {
					re := tg.addRelative(person, level, relatives) // Assumindo que addRelative agora aceita relatives
					relatives = append(relatives, re)              // Adiciona o parente apenas uma vez
				}
				// Continua a busca por descendentes de maneira recursiva
				relatives = tg.searchDescendants(person, persons, level+1, relatives)
			}
		}
	}
	return relatives
}

func (tg *TreeGenealogical) searchForRelatives(relative *Person, persons []*Person, level int, relatives []*Relative) []*Relative {
	if relative == nil {
		return relatives
	}
	for _, person := range persons {
		for _, relationship := range person.Relationships {
			if relationship.SecundePerson == relative.ID {
				// Verifica se a pessoa já foi adicionada e não é o Root
				if !tg.alreadyInFamily(person, relatives) && tg.Root.ID != person.ID {
					// Se não estiver na lista, adicione e continue a busca recursiva
					re := tg.addRelative(person, level, relatives)
					relatives = append(relatives, re) // Adiciona uma única vez

					// Continua a busca por mais parentes sem passar o mesmo slice modificado
					relatives = tg.searchForRelatives(person, persons, level+1, relatives)
				}
			}
		}
	}
	return relatives
}

func (tg *TreeGenealogical) addRelative(person *Person, level int, relatives []*Relative) *Relative {
	relative := &Relative{
		Type:   tg.relationshipDescription(person, relatives),
		Level:  level,
		Person: person,
	}

	return relative
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
