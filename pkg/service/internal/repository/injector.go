package repository

type Repositories struct {
	Person  Person
	Kinship Kinship
}

func NewRepositories() *Repositories {
	return &Repositories{
		Person:  NewPerson(),
		Kinship: NewKinship(),
	}
}
