package genealogy

import "github.com/GeovaneCavalcante/tree-genealogical/internal/entity"

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

type Relative struct {
	Type   string
	Level  int
	Person *entity.Person
}

type TreeGenealogical struct {
	Root      *entity.Person
	Persons   []*entity.Person
	Relatives []*Relative
}

// Cria uma nova árvore genealógica com base no parente e na lista de pessoas.
func NewFamilyTree(root *entity.Person, persons []*entity.Person) *TreeGenealogical {
	tg := &TreeGenealogical{
		Root: root,
	}
	tg.BuildFamilyTree(root, persons, 1)
	return tg
}

// Constrói a árvore genealógica com base no parente e na lista de pessoas.
func (tg *TreeGenealogical) BuildFamilyTree(parente *entity.Person, persons []*entity.Person, level int) []*Relative {
	var relatives []*Relative
	// Busca por descendentes.
	relatives = tg.searchDescendants(tg.Root, persons, level, relatives)
	// Busca por ancestrais e seus parentes.
	relatives = tg.searchAncestors(tg.Root, persons, level, relatives)
	tg.Relatives = relatives
	return relatives
}

// Descrição da relação com base no parente e no sexo.
func (tg *TreeGenealogical) relationshipDescription(relative *entity.Person, relatives []*Relative, persons []*entity.Person) string {

	// Verifica se a relação é direta.
	description := tg.directRelationDescription(relative, persons)
	if description != "" {
		return description
	}

	// Busca os pais da pessoa com base nos parentes ja catalogados.
	parent := tg.findParents(relative, relatives)

	// Regras para determinar o novo parente com base no parente encontrado.
	rulesParents := map[string]string{
		father:      grandFather,
		grandFather: greatGrandfather,
	}

	// Aplica as regras para determinar o novo parente.
	if parent != nil {
		for key, value := range rulesParents {
			if isTypeOfTypeKinship(parent.Type, key) {
				return descriptionBySex(value, relative.Sex)
			}
		}
	}

	// Busca os filhos da pessoa com base nos parentes ja catalogados.
	child := tg.findChildren(relative, relatives)

	// Regras para determinar o novo parente com base no parente encontrado.
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

	// Aplica as regras para determinar o novo parente.
	if child != nil {
		for key, value := range rulesChild {
			if child != nil && isTypeOfTypeKinship(child.Type, key) {
				return descriptionBySex(value, relative.Sex)
			}
		}
	}

	return unknownRelation
}

// Busca por ancestrais de maneira recursiva.
func (tg *TreeGenealogical) searchAncestors(relative *entity.Person, persons []*entity.Person, level int, relatives []*Relative) []*Relative {
	if relative == nil {
		return relatives
	}
	for _, relationship := range relative.Relationships {
		secundePerson := findPerson(relationship.SecundePerson, persons)
		if secundePerson == nil || tg.alreadyInFamily(secundePerson, relatives) {
			continue // Pula para o próximo relacionamento se o parente já estiver na lista ou não for encontrado
		}
		// Adiciona o parente encontrado (ancestral) apenas se não estiver já na lista
		re := tg.newRelative(secundePerson, level, relatives, persons)
		relatives = append(relatives, re)

		// Recursivamente busca por mais ancestrais deste parente encontrado
		relatives = tg.searchForRelatives(secundePerson, persons, level+1, relatives)
		// Recursivamente busca por mais ancestrais
		relatives = tg.searchAncestors(secundePerson, persons, level+1, relatives)

	}

	return relatives

}

// Busca por descendentes de maneira recursiva.
func (tg *TreeGenealogical) searchDescendants(relative *entity.Person, persons []*entity.Person, level int, relatives []*Relative) []*Relative {
	if relative == nil {
		return relatives
	}
	for _, person := range persons {
		for _, relationship := range person.Relationships {
			if relationship.SecundePerson == relative.ID {
				// Verifica se a pessoa já foi adicionada e não é o Root
				if !tg.alreadyInFamily(person, relatives) {
					re := tg.newRelative(person, level, relatives, persons) // Assumindo que newRelative agora aceita relatives
					relatives = append(relatives, re)                       // Adiciona o parente apenas uma vez
				}
				// Continua a busca por descendentes de maneira recursiva
				relatives = tg.searchDescendants(person, persons, level+1, relatives)
			}
		}
	}
	return relatives
}

// Busca por parentes de maneira recursiva.
func (tg *TreeGenealogical) searchForRelatives(relative *entity.Person, persons []*entity.Person, level int, relatives []*Relative) []*Relative {
	if relative == nil {
		return relatives
	}
	for _, person := range persons {
		for _, relationship := range person.Relationships {
			if relationship.SecundePerson == relative.ID {
				// Verifica se a pessoa já foi adicionada e não é o Root
				if !tg.alreadyInFamily(person, relatives) && tg.Root.ID != person.ID {
					// Se não estiver na lista, adicione e continue a busca recursiva
					re := tg.newRelative(person, level, relatives, persons)
					relatives = append(relatives, re) // Adiciona uma única vez

					// Continua a busca por mais parentes sem passar o mesmo slice modificado
					relatives = tg.searchForRelatives(person, persons, level+1, relatives)
				}
			}
		}
	}
	return relatives
}

// Encontra o pais do relative (Parente interado no momento).
func (tg *TreeGenealogical) findParents(relative *entity.Person, relatives []*Relative) *Relative {
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

func (tg *TreeGenealogical) directRelationDescription(relative *entity.Person, persons []*entity.Person) string {
	if relative == nil {
		return ""
	}
	// Verifica se o relative é filho do Root.
	if tg.isChildOfRoot(relative) {
		return descriptionBySex(son, relative.Sex)
	}

	// Verifica se o relative é pai ou mae do Root.
	if tg.isParentOfRoot(relative) {
		return descriptionBySex(father, relative.Sex)
	}

	// Verifica se o relative é irmão do Root.
	return tg.checkSiblingRelation(relative, persons)
}

// Verifica se a pessoa é filho(a) do Root.
func (tg *TreeGenealogical) isChildOfRoot(relative *entity.Person) bool {
	for _, relationship := range relative.Relationships {
		if relationship.SecundePerson == tg.Root.ID {
			return true
		}
	}
	return false
}

// Verifica se a pessoa é pai/mãe do Root.
func (tg *TreeGenealogical) isParentOfRoot(relative *entity.Person) bool {
	for _, relationship := range tg.Root.Relationships {
		if relationship.SecundePerson == relative.ID {
			return true
		}
	}
	return false
}

// Verifica se a pessoa é irmão do Root.
func (tg *TreeGenealogical) checkSiblingRelation(relative *entity.Person, persons []*entity.Person) string {
	for _, relationship := range tg.Root.Relationships {
		for _, relationshipParent := range relative.Relationships {
			if relationshipParent.SecundePerson == relationship.SecundePerson {
				// Se ambos compartilham o mesmo pai/mãe, são irmãos.
				secundePerson := findPerson(relationship.SecundePerson, persons)
				if secundePerson != nil {
					return descriptionBySex(brother, relative.Sex) // Usa o sexo do relative para determinar a relação.
				}
			}
		}
	}
	return ""
}

// Encontra o parente que é filho do relative (Parente interado no momento).
func (tg *TreeGenealogical) findChildren(relative *entity.Person, relatives []*Relative) *Relative {
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

// Cria um novo parente com base no relative e o adiciona à lista de parentes.
func (tg *TreeGenealogical) newRelative(person *entity.Person, level int, relatives []*Relative, persons []*entity.Person) *Relative {
	relative := &Relative{
		Type:   tg.relationshipDescription(person, relatives, persons),
		Level:  level,
		Person: person,
	}
	return relative
}

// Verifica se o parente já está na lista de parentes.
func (tg *TreeGenealogical) alreadyInFamily(relative *entity.Person, relatives []*Relative) bool {
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

// Retorna a descrição da relação com base no parente e no sexo.
func descriptionBySex(relative, sex string) string {
	if relation, ok := kinshipTypes[relative]; ok {
		if desc, ok := relation[sex]; ok {
			return desc
		}
		return unknownRelation
	}
	return unknownRelation
}

// Verifica se o tipo de parentesco é válido.
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

// Encontra a pessoa com base no ID.
func findPerson(ID string, persons []*entity.Person) *entity.Person {
	for _, person := range persons {
		if person.ID == ID {
			return person
		}
	}
	return nil
}
