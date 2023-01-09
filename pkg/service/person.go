//go:generate go run github.com/golang/mock/mockgen -package=service -source=$GOFILE -destination=../../test/mock/$GOPACKAGE/$GOFILE
package service

import (
	"context"

	"github.com/vladisvrau/FamilyTree/pkg/entity"
	"github.com/vladisvrau/FamilyTree/pkg/service/internal/repository"
)

type Person interface {
	GetByID(ctx context.Context, id int) (*entity.Person, error)
	GetByName(ctx context.Context, name string) ([]*entity.Person, error)
	Save(ctx context.Context, p *entity.Person) (*entity.Person, error)
	Edit(ctx context.Context, p *entity.Person) (*entity.Person, error)
	Delete(ctx context.Context, id int) error
}

type person struct {
	repository *repository.Repositories
}

func NewPerson(repo *repository.Repositories) Person {
	return &person{
		repository: repo,
	}
}

func (s *person) GetByID(ctx context.Context, id int) (*entity.Person, error) {
	return s.repository.Person.FetchByID(ctx, id)
}
func (s *person) GetByName(ctx context.Context, name string) ([]*entity.Person, error) {
	return s.repository.Person.FetchByName(ctx, name)
}

func (s *person) Save(ctx context.Context, p *entity.Person) (*entity.Person, error) {
	return s.repository.Person.Create(ctx, p)
}

func (s *person) Edit(ctx context.Context, p *entity.Person) (*entity.Person, error) {
	return s.repository.Person.Edit(ctx, p)
}

func (s *person) Delete(ctx context.Context, id int) error {
	return s.repository.Person.Delete(ctx, id)
}
