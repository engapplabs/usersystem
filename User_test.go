package main 

import "testing"

func TestUserCreation(test *testing.T) {

	const EXPECTED_ID = 505
	const EXPECTED_NAME = "Aurelio"
	const EXPECTED_EMAIL = "abmf"
	const EXPECTED_PASSWORD = "vaca"
	
	user := newUser(EXPECTED_ID, EXPECTED_NAME, EXPECTED_EMAIL, EXPECTED_PASSWORD)

	if user.getId() != EXPECTED_ID {
		test.Errorf("Expected id to be 505, but got %v", user.getId())
	}

	if user.getName() != EXPECTED_NAME {
		test.Errorf("Expected name to be Aurelio, but got %v", user.getName())
	}

	if user.getEmail() != EXPECTED_EMAIL {
		test.Errorf("Expected email to be abmf, but got %v", user.getEmail())
	}

	if user.getPassword() != EXPECTED_PASSWORD {
		test.Errorf("Expected password to be Aurelio, but got %v", user.getPassword())
	}
}
