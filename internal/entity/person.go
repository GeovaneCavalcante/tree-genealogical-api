package entity

type Person struct {
	ID            string          `json:"id"`
	Name          string          `json:"name"`
	Gender        string          `json:"gender"`
	Level         int             `json:"level"`
	Relationships []*Relationship `json:"relationships"`
}
