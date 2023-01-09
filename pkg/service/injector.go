//  go:generate: go run github.com/golang/mock/mockgen -package=service -source=$GOFILE -destination=../../test/mock/$GOPACKAGE/$GOFILE

package service

import "github.com/vladisvrau/FamilyTree/pkg/service/internal/repository"

type Services struct {
	Person  Person
	Kinship Kinship
	Tree    Tree
}

func NewServices() *Services {
	repo := repository.NewRepositories()

	return &Services{
		Person:  NewPerson(repo),
		Kinship: NewKinship(repo),
		Tree:    NewTree(repo),
	}
}
