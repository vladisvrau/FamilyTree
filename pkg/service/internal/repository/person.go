//go:generate go run github.com/golang/mock/mockgen -package=service -source=$GOFILE -destination=../../test/mock/$GOPACKAGE/$GOFILE

package repository

import (
	"context"
	"errors"

	"github.com/vladisvrau/FamilyTree/lib/database"
	"github.com/vladisvrau/FamilyTree/lib/util"
	"github.com/vladisvrau/FamilyTree/pkg/entity"
)

type Person interface {
	FetchByID(ctx context.Context, id int) (*entity.Person, error)
	FetchByName(ctx context.Context, name string) ([]*entity.Person, error)
	Create(ctx context.Context, obj *entity.Person) (*entity.Person, error)
	Delete(ctx context.Context, id int) error
	Edit(ctx context.Context, new *entity.Person) (*entity.Person, error)
}

type person struct {
	people         map[int]*entity.Person // Mocked for testing
	nameToID       map[string][]int       // In case more than one person has the same name (would be an extra query in mysql)
	lastInsertedID int                    // AUTOINCREMENT Simulation
}

func NewPerson() Person {
	return &person{
		people: map[int]*entity.Person{
			1: {
				ID:   1,
				Name: "Arthur",
			},
			2: {
				ID:   2,
				Name: "Geraldo",
			},
			3: {
				ID:   3,
				Name: "Valeria",
			},
		},
		nameToID: map[string][]int{
			"Arthur":  {1},
			"Geraldo": {2},
			"Valeria": {3},
		},
		lastInsertedID: 3,
	}
}

func (r *person) FetchByID(ctx context.Context, id int) (*entity.Person, error) {
	p, ok := r.people[id]
	if !ok {
		return nil, database.ErrEntityNotFound
	}

	result := entity.Person{ID: p.ID, Name: p.Name}
	return &result, nil
}

func (r *person) FetchByName(ctx context.Context, name string) ([]*entity.Person, error) {
	ids, ok := r.nameToID[name]
	if !ok {
		return nil, database.ErrEntityNotFound
	}

	var people []*entity.Person
	for _, id := range ids {
		p, err := r.FetchByID(ctx, id)
		if err != nil {
			if err != database.ErrEntityNotFound {
				return nil, errors.New("wtf")
			}
			continue
		}

		people = append(people, p)
	}

	return people, nil
}

func (r *person) Create(ctx context.Context, obj *entity.Person) (*entity.Person, error) {
	if obj.ID != 0 {
		return nil, database.ErrInvalidInsert
	}

	r.lastInsertedID++
	newPerson := entity.Person{
		ID:   r.lastInsertedID,
		Name: obj.Name,
	}
	r.people[r.lastInsertedID] = &newPerson
	obj.ID = newPerson.ID

	ids, ok := r.nameToID[newPerson.Name]
	if !ok {
		r.nameToID[newPerson.Name] = []int{newPerson.ID}
	} else {
		r.nameToID[newPerson.Name] = append(ids, newPerson.ID)
	}

	return obj, nil
}

func (r *person) Delete(ctx context.Context, id int) error {
	p, ok := r.people[id]
	if !ok {
		return database.ErrEntityNotFound
	}

	delete(r.people, id)
	ids, ok := r.nameToID[p.Name]
	if ok {
		if len(ids) == 1 {
			delete(r.nameToID, p.Name)
		} else {
			ids = util.SlicePopComparable(ids, id)
			r.nameToID[p.Name] = ids
		}
	}

	return nil
}

func (r *person) Edit(ctx context.Context, new *entity.Person) (*entity.Person, error) {
	old, ok := r.people[new.ID]
	if !ok {
		return nil, database.ErrEntityNotFound
	}

	ids, ok := r.nameToID[old.Name]
	if ok {
		if len(ids) == 1 {
			delete(r.nameToID, old.Name)
		} else {
			ids = util.SlicePopComparable(ids, new.ID)
			r.nameToID[old.Name] = ids
		}
	}

	r.people[new.ID].Name = new.Name

	ids, ok = r.nameToID[new.Name]
	if !ok {
		r.nameToID[new.Name] = []int{new.ID}
	} else {
		r.nameToID[new.Name] = append(ids, new.ID)
	}

	newObj := entity.Person{
		ID:   new.ID,
		Name: new.Name,
	}
	return &newObj, nil
}
