package entity

type Person struct {
	ID            string
	Name          string
	Sex           string
	Level         int
	Relationships []Relationship
}
