package main

import "fmt"

type (
	Collection interface {
		createIterator() Iterator
	}

	Iterator interface {
		hasNext() bool
		getNext() *User
	}

	UserCollection struct {
		users []*User
	}

	UserIterator struct {
		index int
		users []*User
	}

	User struct {
		name string
		age  int
	}
)

// UserCollection's collection of methods
func (s *UserCollection) createIterator() Iterator {
	return &UserIterator{users: s.users}
}

// UserIterator's collection of methods
func (s *UserIterator) hasNext() bool {
	return s.index < len(s.users)
}

func (s *UserIterator) getNext() *User {
	if s.hasNext() {
		user := s.users[s.index]
		s.index++
		return user
	}
	return nil
}

// main function
func main() {
	user1 := &User{name: "John", age: 30}
	user2 := &User{name: "Doe", age: 40}
	user3 := &User{name: "Smith", age: 50}
	user4 := &User{name: "Peter", age: 60}

	users := []*User{user1, user2, user3, user4}
	userCollection := &UserCollection{users: users}

	userIterator := userCollection.createIterator()
	for userIterator.hasNext() {
		user := userIterator.getNext()
		fmt.Println(user.name)
	}
}
