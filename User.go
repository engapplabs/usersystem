package main

import "fmt"

type User struct {
	id int
	email string
	name string 
	password string
}

func newUser(id int, name, email, password string) *User {
	return &User {
		id: id,
		name: name,
		email: email,
		password: password,
	}
}

func (user *User) setName(name string) {
	user.name = name
}

func (user *User) getName() string {
	return user.name
}

func (user *User) setEmail(email string) {
	user.email = email
}

func (user *User) getEmail() string {
	return user.email
}

func (user *User) setPassword(password string) {
	user.password = password
}

func (user *User) getPassword() string {
	return user.password
}

func (user *User) getId() int {	
	return user.id
}

func (user *User) toString() string {
	return fmt.Sprintf("{id: %v,name:'%v', email:'%v', password:'%v'}", user.id, user.name, user.email, user.password)
}
