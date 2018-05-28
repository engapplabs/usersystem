package main 

import "fmt"

type UserRepository struct {
	users map[int]*User
}

func newUserRepository() *UserRepository {
	return &UserRepository{
		users: make(map[int]*User),
	}
}

func (respository *UserRepository) getAll() map[int]*User {
	return respository.users
}

func (respository *UserRepository) addUser(user *User) {
	respository.users[user.getId()] = user
}

func (respository *UserRepository) removeUserById(id int) {
	delete(respository.users, id)
}

func (respository *UserRepository) findUserById(id int) *User {
	if user, exist := respository.users[id]; exist {
		return user 
	} else {
		return nil
	}
}

func (respository *UserRepository) printUsers() {
	for userId, user := range respository.users {
		fmt.Printf("Id: %v, User: %v\n", userId, user.toString())
	}
}

func (respository *UserRepository) usersInside() int {
	return len(respository.users)
}
