//go:generate go run github.com/golang/mock/mockgen -package=service -source=$GOFILE -destination=../../test/mock/$GOPACKAGE/$GOFILE

package service

import (
	"context"

	"github.com/vladisvrau/FamilyTree/lib/database"
	"github.com/vladisvrau/FamilyTree/pkg/entity"
	"github.com/vladisvrau/FamilyTree/pkg/service/internal/repository"
)

type Tree interface {
	BuildFamilyTree(ctx context.Context, personID int) ([]entity.PersonNode, error)
}

func NewTree(repo *repository.Repositories) Tree {
	return &tree{
		repository: repo,
	}
}

type tree struct {
	repository *repository.Repositories
}

func (s *tree) BuildFamilyTree(ctx context.Context, personID int) ([]entity.PersonNode, error) {
	result := []entity.PersonNode{}

	p, err := s.repository.Person.FetchByID(ctx, personID)
	if err != nil {
		return nil, err
	}
	upwardMembers, upwardRelationships, err := s.upwards(ctx, p.ID)
	if err != nil {
		return nil, err
	}

	downwardMembers, downwardRelationships, err := s.downwards(ctx, p.ID)
	if err != nil {
		return nil, err
	}

	currentNode := entity.PersonNode{
		ID:            p.ID,
		Name:          p.Name,
		Relationships: append(upwardRelationships, downwardRelationships...),
	}

	result = append(result, currentNode)
	result = append(result, upwardMembers...)
	result = append(result, downwardMembers...)

	return result, nil
}

func (s *tree) upwards(ctx context.Context, personID int) ([]entity.PersonNode, []entity.Relationship, error) {
	members := []entity.PersonNode{}
	relationships := []entity.Relationship{}
	parents, err := s.repository.Kinship.FetchByChildID(ctx, personID)
	if err != nil && err != database.ErrEntityNotFound {
		return nil, nil, err
	}
	for _, p := range parents {
		person, err := s.repository.Person.FetchByID(ctx, p.Parent)
		if err != nil {
			return nil, nil, err
		}
		relationships = append(relationships, entity.Relationship{
			PersonID: person.ID,
			Name:     person.Name,
			Kind:     "parent",
		})

		upwardMembers, upwardKinships, err := s.upwards(ctx, person.ID)
		if err != nil {
			return nil, nil, err
		}

		node := entity.PersonNode{
			Name:          person.Name,
			ID:            person.ID,
			Relationships: upwardKinships,
		}

		members = append(members, node)
		members = append(members, upwardMembers...)
	}

	return members, relationships, nil
}

func (s *tree) downwards(ctx context.Context, parent int) ([]entity.PersonNode, []entity.Relationship, error) {
	members := []entity.PersonNode{}
	relationships := []entity.Relationship{}
	children, err := s.repository.Kinship.FetchByParentID(ctx, parent)
	if err != nil && err != database.ErrEntityNotFound {
		return nil, nil, err
	}
	for _, c := range children {
		person, err := s.repository.Person.FetchByID(ctx, c.Child)
		if err != nil {
			return nil, nil, err
		}
		relationships = append(relationships, entity.Relationship{
			PersonID: person.ID,
			Name:     person.Name,
			Kind:     "child",
		})

		upwardMembers, personalKinships, err := s.downwards(ctx, person.ID)

		node := entity.PersonNode{
			Name:          person.Name,
			ID:            person.ID,
			Relationships: personalKinships,
		}

		members = append(members, node)
		members = append(members, upwardMembers...)
	}

	return members, relationships, nil
}
