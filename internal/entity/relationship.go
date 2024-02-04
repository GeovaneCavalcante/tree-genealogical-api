package entity

type Relationship struct {
	ID              string
	MainPersonID    string
	MainPerson      *Person
	SecundePersonID string
	SecundePerson   *Person
}
