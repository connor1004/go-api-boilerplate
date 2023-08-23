package app

import (
	"encoding/json"
	"fmt"
	"os"
)

// Users - a struct to import json file content
type Users struct {
	Objects []map[string]interface{} `json:"objects"`
}

// ImportUsers - a function to import users from json file and keep them in database
func (a *App) ImportUsers() {
	jsonFile, err := os.Open("ExportJson.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Successfully Opened ExportJson.json")

	defer jsonFile.Close()

	users := &Users{}

	err = json.NewDecoder(jsonFile).Decode(&users)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("total %v users", len(users.Objects))

	for _, user := range users.Objects {
		_, err = a.DB.Exec(
			"call add_user(?, ?, ?, ?, ?, ?, ?, ?, ?)",
			user["job_title"],
			user["first_name"],
			user["last_name"],
			user["gender"],
			user["birth_date"],
			user["department_name"],
			user["badge_id"],
			user["phone"],
			user["email"],
		)
		if err != nil {
			fmt.Println(err)
		}
	}
}
