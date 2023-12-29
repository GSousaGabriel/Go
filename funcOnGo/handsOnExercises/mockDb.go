package main

import (
	"fmt"
	"log"
)

// User represents a user with an id and first name
type User struct {
	ID    int
	First string
}

// MockDatastore is a temporary service that stores retrievable data
type MockDatastore struct {
	Users map[int]User
}

func (md MockDatastore) GetUser(id int) (User, error) {

	if user, ok := md.Users[id]; !ok {
		return User{}, fmt.Errorf("User %v not found", id)
	} else {
		return user, nil
	}
}

func (md MockDatastore) SaveUser(u User) error {

	if _, ok := md.Users[u.ID]; ok {
		return fmt.Errorf("User %v already exists", u.ID)
	} else {
		md.Users[u.ID] = u
		return nil
	}
}

type Datastore interface {
	GetUser(id int) (User, error)
	SaveUser(u User) error
}

type Service struct {
	ds Datastore
}

func (s Service) GetUser(id int) (User, error) {
	return s.ds.GetUser(id)
}

func (s Service) SaveUser(u User) error {
	return s.ds.SaveUser(u)
}

func mockDb() {
	db := MockDatastore{
		Users: make(map[int]User),
	}

	srvc := Service{
		ds: db,
	}

	u1 := User{
		ID:    1,
		First: "Roger",
	}

	err := srvc.SaveUser(u1)

	if err != nil {
		log.Fatalf("Error %s", err)
	}

	u1Returned, err := srvc.GetUser(u1.ID)

	if err != nil {
		log.Fatalf("Error %s", err)
	}

	fmt.Println(u1)
	fmt.Println(u1Returned)
}
