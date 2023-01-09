package entity

type Person struct {
	ID   int    `json: "id"`
	Name string `json: "name"`
}

type Relationship struct {
	PersonID int    `json:"person_id"`
	Name     string `json:"name"`
	Kind     string `json:"kind"`
}

type PersonNode struct {
	ID            int            `json:id`
	Name          string         `json:"name"`
	Relationships []Relationship `json:"relationships"`
}
