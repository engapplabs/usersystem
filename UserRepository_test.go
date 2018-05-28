package main 

import "testing"

func TestUserInsertion(test *testing.T) {
	const EXPECTED_ID = 505
	const EXPECTED_NAME = "Aurelio"
	const EXPECTED_EMAIL = "abmf"
	const EXPECTED_PASSWORD = "vaca"

	userRepository := newUserRepository()
	
	user := newUser(EXPECTED_ID, EXPECTED_NAME, EXPECTED_EMAIL, EXPECTED_PASSWORD)

	userRepository.addUser(user)

	if userRepository.usersInside() == 0 {
		test.Errorf("Expected number of users inside be greater than zero")
	}
} 

func TestUserInsertionAndSearch(test *testing.T) {
	const EXPECTED_ID = 505
	const EXPECTED_NAME = "Aurelio"
	const EXPECTED_EMAIL = "abmf"
	const EXPECTED_PASSWORD = "vaca"

	userRepository := newUserRepository()
	
	user := newUser(EXPECTED_ID, EXPECTED_NAME, EXPECTED_EMAIL, EXPECTED_PASSWORD)

	userRepository.addUser(user)

	foundUser := userRepository.findUserById(EXPECTED_ID)
	if foundUser == nil {
		test.Errorf("Expected to found user with 505, but got nil")
	}
}

func TestUserDeletion(test *testing.T) {
	const EXPECTED_ID = 505
	const EXPECTED_NAME = "Aurelio"
	const EXPECTED_EMAIL = "abmf"
	const EXPECTED_PASSWORD = "vaca"

	userRepository := newUserRepository()
	
	user := newUser(EXPECTED_ID, EXPECTED_NAME, EXPECTED_EMAIL, EXPECTED_PASSWORD)

	userRepository.addUser(user)

	userRepository.removeUserById(EXPECTED_ID)

	foundUser := userRepository.findUserById(EXPECTED_ID)
	if foundUser != nil {
		test.Errorf("User with id 505 was expected to be deleted")
	}	
}
