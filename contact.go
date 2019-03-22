package main

import (
	"fmt"
)

// Contact represents a single record in the database.
type Contact struct {
	ID        int    `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Phone     string `json:"phone"`
	Address   string `json:"address"`
	Email     string `json:"email"`
}

// Valid checks if the contact information is valid or not.
func (c *Contact) Valid() error {
	if c.Firstname == "" {
		return fmt.Errorf("firstname is empty")
	}

	if c.Lastname == "" {
		return fmt.Errorf("lastname is empty")
	}

	if c.Phone == "" {
		return fmt.Errorf("phone is empty")
	}

	if c.Address == "" {
		return fmt.Errorf("address is empty")
	}

	if c.Email == "" {
		return fmt.Errorf("email is empty")
	}

	return nil
}
