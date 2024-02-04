package main

// import (
// 	"fmt"
// 	"strings"

// 	"github.com/GeovaneCavalcante/tree-genealogical/person"
// 	"github.com/GeovaneCavalcante/tree-genealogical/pkg/genealogy"
// 	"github.com/GeovaneCavalcante/tree-genealogical/relationship"
// 	"github.com/google/uuid"
// )

// func NewPerson(name, gender, fatherID, motherID string) *person.Person {
// 	person := &person.Person{
// 		ID:     uuid.New().String(),
// 		Name:   name,
// 		Gender: gender,
// 	}

// 	if fatherID != "" {
// 		person.Relationships = append(person.Relationships, relationship.Relationship{MainPersonID: person.ID, SecundePersonID: fatherID})
// 	}

// 	if motherID != "" {
// 		person.Relationships = append(person.Relationships, relationship.Relationship{MainPersonID: person.ID, SecundePersonID: motherID})
// 	}

// 	return person
// }

// func main() {

// 	geruza := NewPerson("Geruza", "F", "", "")
// 	geova := NewPerson("Geova", "M", "", "")
// 	pedro := NewPerson("Pedro", "M", "", "")
// 	iraci := NewPerson("Iraci", "F", "", geruza.ID)
// 	tereza := NewPerson("Tereza", "F", "", geruza.ID)
// 	alsimar := NewPerson("Alsimar", "F", pedro.ID, iraci.ID)
// 	suzamar := NewPerson("Suzamar", "F", pedro.ID, iraci.ID)
// 	araujo := NewPerson("Araujo", "M", "", "")
// 	debora := NewPerson("debora", "F", araujo.ID, suzamar.ID)
// 	bruna := NewPerson("Bruna", "F", "", debora.ID)
// 	gean := NewPerson("Gean", "M", geova.ID, alsimar.ID)
// 	bruno := NewPerson("Bruno", "M", gean.ID, "")
// 	soraia := NewPerson("Soraia", "F", pedro.ID, iraci.ID)

// 	geovane := NewPerson("Geovane", "M", geova.ID, alsimar.ID)
// 	victoria := NewPerson("Victoria", "M", "", "")

// 	cr7 := NewPerson("Cristiano Ronaldo", "M", geovane.ID, victoria.ID)
// 	oceane := NewPerson("Oceane", "F", geovane.ID, victoria.ID)

// 	vicovane := NewPerson("Vicovane", "M", cr7.ID, "")

// 	neymar := NewPerson("Neymar", "M", vicovane.ID, "")

// 	persons := []*person.Person{
// 		geova,
// 		iraci,
// 		pedro,
// 		alsimar,
// 		suzamar,
// 		soraia,
// 		gean,
// 		geovane,
// 		bruna,
// 		geruza,
// 		bruno,
// 		tereza,
// 		debora,
// 		victoria,
// 		cr7,
// 		vicovane,
// 		neymar,
// 		oceane,
// 		araujo,
// 	}

// 	enviado := geovane

// 	treeGenealogical := genealogy.NewFamilyTree()

// 	fmt.Printf("My name is %s\n", enviado.Name)
// 	fmt.Println("My family tree :")

// 	for _, relative := range treeGenealogical.Relatives {
// 		formater := "--"
// 		multipliedCharacter := strings.Repeat(formater, relative.Level)
// 		if relative.Person == nil {
// 			fmt.Println(relative)
// 			continue
// 		}
// 		fmt.Printf("%s %s: %s\n", multipliedCharacter, relative.Person.Name, relative.Type)
// 	}

// }
