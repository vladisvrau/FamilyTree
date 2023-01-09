//go:generate go run github.com/golang/mock/mockgen -package=service -source=$GOFILE -destination=../../test/mock/$GOPACKAGE/$GOFILE
package service

import (
	"context"

	"github.com/vladisvrau/FamilyTree/pkg/entity"
	"github.com/vladisvrau/FamilyTree/pkg/service/internal/repository"
)

type Kinship interface {
	Create(ctx context.Context, k *entity.Kinship) error
	Delete(ctx context.Context, k *entity.Kinship) error
	GetParents(ctx context.Context, child int) ([]*entity.Kinship, error)
	GetChildren(ctx context.Context, parent int) ([]*entity.Kinship, error)
}

func NewKinship(repo *repository.Repositories) Kinship {
	return &kinship{repo}
}

type kinship struct {
	repository *repository.Repositories
}

func (s *kinship) Create(ctx context.Context, k *entity.Kinship) error {

	if _, err := s.repository.Person.FetchByID(ctx, k.Parent); err != nil {
		return err
	}

	if _, err := s.repository.Person.FetchByID(ctx, k.Child); err != nil {
		return err
	}

	return s.repository.Kinship.Create(ctx, k)
}

func (s *kinship) Delete(ctx context.Context, k *entity.Kinship) error {
	return s.repository.Kinship.Delete(ctx, k)
}

func (s *kinship) GetParents(ctx context.Context, child int) ([]*entity.Kinship, error) {
	return s.repository.Kinship.FetchByChildID(ctx, child)
}

func (s *kinship) GetChildren(ctx context.Context, parent int) ([]*entity.Kinship, error) {
	return s.repository.Kinship.FetchByParentID(ctx, parent)
}
