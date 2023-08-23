package models

import (
	"regexp"
)

// User - a user model
type User struct {
	ID             int    `json:"id,omitempty"`
	JobTitle       string `json:"job_title"`
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	Gender         string `json:"gender"`
	BirthDate      string `json:"birth_date"`
	DepartmentName string `json:"department_name"`
	BadgeID        int    `json:"badge_id"`
	Phone          string `json:"phone"`
	Email          string `json:"email"`
}

// ValidateUser - check the user info is valid
func (user *User) ValidateUser() map[string]string {
	isValid := true
	errors := map[string]string{}

	if user.JobTitle == "" {
		errors["job_title"] = "Job title is required"
		isValid = false
	}

	if user.FirstName == "" {
		errors["first_name"] = "First name is required"
		isValid = false
	}

	if user.LastName == "" {
		errors["last_name"] = "Last name is required"
		isValid = false
	}

	if user.Gender == "" {
		errors["gender"] = "Gender is required"
		isValid = false
	}

	if user.DepartmentName == "" {
		errors["department_name"] = "Department name is required"
		isValid = false
	}

	if user.Phone != "" {
		re := regexp.MustCompile(`^(?:(?:\(?(?:00|\+)([1-4]\d\d|[1-9]\d?)\)?)?[\-\.\ \\\/]?)?((?:\(?\d{1,}\)?[\-\.\ \\\/]?){0,})(?:[\-\.\ \\\/]?(?:#|ext\.?|extension|x)[\-\.\ \\\/]?(\d+))?$`)
		if !re.MatchString(user.Phone) {
			errors["phone"] = "Phone number is invalid"
			isValid = false
		}
	}

	if user.Email == "" {
		errors["email"] = "Email is required"
		isValid = false
	} else {
		re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
		if !re.MatchString(user.Email) {
			errors["email"] = "Email is invalid"
			isValid = false
		}
	}

	if isValid {
		return nil
	}
	return errors
}
