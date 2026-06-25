package main

import "strings"

type Student struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

func (s Student) Validate() []string {

	var errors []string

	if strings.TrimSpace(s.Name) == "" {
		errors = append(errors, "Name is required")
	}

	if strings.TrimSpace(s.Email) == "" {
		errors = append(errors, "Email is required")
	}

	if s.Age <= 0 {
		errors = append(errors, "Age must be greater than zero")
	}

	return errors
}