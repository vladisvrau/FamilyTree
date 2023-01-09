//go:generate go run github.com/golang/mock/mockgen -package=service -source=$GOFILE -destination=../../test/mock/$GOPACKAGE/$GOFILE

package repository

import (
	"context"

	"github.com/vladisvrau/FamilyTree/lib/database"
	"github.com/vladisvrau/FamilyTree/pkg/entity"
)

type Kinship interface {
	Create(ctx context.Context, kinship *entity.Kinship) error
	Delete(ctx context.Context, kinship *entity.Kinship) error
	FetchByChildID(ctx context.Context, id int) ([]*entity.Kinship, error)
	FetchByParentID(ctx context.Context, id int) ([]*entity.Kinship, error)
}

type kinship struct {
	kinshipTable *entity.KinshipCollection
}

func NewKinship() Kinship {
	return &kinship{
		kinshipTable: &entity.KinshipCollection{
			{
				Parent: 2,
				Child:  1,
			},
			{
				Parent: 3,
				Child:  1,
			},
		},
	}
}

func (r *kinship) Create(ctx context.Context, k *entity.Kinship) error {
	err := r.validateInsert(ctx, k)
	if err != nil {
		return err
	}
	// insertion
	index := r.kinshipTable.Push(k)
	if index < 0 {
		return database.ErrTableDoesntExist
	}
	return nil
}

func (r *kinship) validateInsert(ctx context.Context, k *entity.Kinship) error {
	// validation
	index := r.kinshipTable.IndexOf(k)
	if index >= 0 {
		return database.ErrInvalidInsert
	}

	index = r.kinshipTable.IndexOf(&entity.Kinship{Parent: k.Child, Child: k.Parent})
	if index >= 0 {
		return database.ErrInvalidInsert
	}

	parents, err := r.FetchByChildID(ctx, k.Child)
	if err != nil && err != database.ErrEntityNotFound {
		return err
	}

	if len(parents) > 1 {
		return database.ErrInvalidInsert
	}

	// check for incest
	if len(parents) == 1 {
		grandParents, err := r.FetchByChildID(ctx, k.Parent)
		if err != nil && err != database.ErrEntityNotFound {
			return err
		}

		otherGrandParents, err := r.FetchByChildID(ctx, parents[0].Parent)
		if err != nil && err != database.ErrEntityNotFound {
			return err
		}

		for _, g := range grandParents {
			for _, og := range otherGrandParents {
				if g.Parent == og.Parent {
					return database.ErrInvalidInsert
				}
			}
		}
	}

	grandChildren, err := r.FetchByParentID(ctx, k.Child)
	if err != nil && err != database.ErrEntityNotFound {
		return err
	}

	for _, gc := range grandChildren {
		ps, err := r.FetchByChildID(ctx, gc.Child)
		if err != nil && err != database.ErrEntityNotFound {
			return err
		}
		for _, p := range ps {
			if p.Parent != k.Child {
				gs, err := r.FetchByChildID(ctx, p.Parent)
				if err != nil && err != database.ErrEntityNotFound {
					return err
				}
				for _, g := range gs {
					if g.Parent == k.Parent {
						return database.ErrInvalidInsert
					}
				}

			}
		}
	}

	return nil
}

func (r *kinship) Delete(ctx context.Context, k *entity.Kinship) error {
	popped := r.kinshipTable.Pop(k)
	if popped == nil {
		return database.ErrEntityNotFound
	}

	return nil
}

func (r *kinship) FetchByChildID(ctx context.Context, id int) ([]*entity.Kinship, error) {
	if r.kinshipTable == nil {
		return nil, database.ErrTableDoesntExist
	}

	collection := *r.kinshipTable
	result := []*entity.Kinship{}
	for _, k := range collection {
		if k.Child == id {
			result = append(result, &entity.Kinship{Parent: k.Parent, Child: k.Child})
		}
	}
	if len(result) == 0 {
		return nil, database.ErrEntityNotFound
	}
	return result, nil
}

func (r *kinship) FetchByParentID(ctx context.Context, id int) ([]*entity.Kinship, error) {
	if r.kinshipTable == nil {
		return nil, database.ErrTableDoesntExist
	}

	collection := *r.kinshipTable
	result := []*entity.Kinship{}
	for _, k := range collection {
		if k.Parent == id {
			result = append(result, &entity.Kinship{Parent: k.Parent, Child: k.Child})
		}
	}
	if len(result) == 0 {
		return nil, database.ErrEntityNotFound
	}
	return result, nil
}
