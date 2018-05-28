package main 

import (
 "fmt"
)

var userRepository = newUserRepository()
var userIdCounter int = 0

func entryPoint() {
	isRunning := true
	for isRunning {
		headerOptions()
		inputOption := readNumberInput()

		switch inputOption {
			case 1:
				addUserHandler()
			case 2:
				removeUserHandler()
			case 3:
				findUserHandler()
			case 4:
				getEmailListHandler()
			case 5: 
				fmt.Println("Bye bye")
				isRunning = false
			case 6:
				printAll()
		}
	}
}

func printAll() {
	userRepository.printUsers()
}

func headerOptions() {
	fmt.Println("1. Add new User")
	fmt.Println("2. Remove user")
	fmt.Println("3. Find user")
	fmt.Println("4. Get users emails")
	fmt.Println("5. Finish")
	fmt.Print(">>> ")
}

func addUserHandler() {
	fmt.Print("Type user name: ")
	userName := readLine()
	fmt.Print("Type user email: ")
	userEmail := readLine()
	fmt.Print("Type user password: ")
	userPassword := readLine()

	user := newUser(userIdCounter, userName, userEmail, userPassword)
	userRepository.addUser(user)
	userIdCounter++
}

func removeUserHandler() {
	fmt.Print("Type id to remove: ")
	idToRemove := readNumberInput()
	userRepository.removeUserById(idToRemove)
}

func findUserHandler() {
	fmt.Print("Type id to be found: ")
	idToSearch := readNumberInput()
	foundUser := userRepository.findUserById(idToSearch)
	if foundUser != nil {
		fmt.Println("Name: ", foundUser.getName())
		fmt.Println("Email: ", foundUser.getEmail())	
		fmt.Println("Password: ", foundUser.getPassword())
	} else {
		fmt.Println("User not found")
	}
	fmt.Println("Press any key to get back....")
	readLine()
}

func getEmailListHandler() {
	users := userRepository.getAll()
	for _, user := range users {
		fmt.Println("Email: ", user.getEmail())	
	}
	fmt.Println("Press any key to get back...")
	readLine()
}
