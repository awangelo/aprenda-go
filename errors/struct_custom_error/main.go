package main

import "fmt"

type User struct {
	name  string
	age   int
	alive bool
}

type ValidationError struct {
	msg   string
	field string
}

func (v *ValidationError) Error() string {
	return fmt.Sprintf("%s: (wrong field: %s)", v.msg, v.field)
}

func main() {
	test_user := User{
		name:  "Abraham Lincoln",
		age:   -121,
		alive: false,
	}

	validationErrors := validateUser(test_user)

	if len(validationErrors) > 0 {
		for _, err := range validationErrors {
			fmt.Println(err)
		}
	} else {
		fmt.Println("User is valid")
	}
}

func validateUser(u User) []error {
	var possibleErrors []error
	if !u.alive {
		possibleErrors = append(possibleErrors, &ValidationError{"User can't be dead", "alive"})
	}
	if u.age < 0 {
		possibleErrors = append(possibleErrors, &ValidationError{"User can't have negative age", "age"})
	}
	return possibleErrors
}
