package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/connor1004/go-api-boilerplate/config"
	"github.com/connor1004/go-api-boilerplate/models"
	"github.com/connor1004/go-api-boilerplate/utils"
)

// UserController - controller for processing routes related to users
type UserController struct {
	db *sql.DB
}

// NewUserController - create a new UserController
func NewUserController(db *sql.DB) *UserController {
	return &UserController{db: db}
}

// AddUser - Add a new user
func (x UserController) AddUser(ctx *utils.Context) {
	var user models.User
	response := map[string]interface{}{}

	err := json.NewDecoder(ctx.Request.Body).Decode(&user)
	if err != nil {
		response["message"] = "Bad request payload: " + fmt.Sprintf("%+v", err)
		ctx.Respond(http.StatusBadRequest, response)
		return
	}

	errors := user.ValidateUser()
	if errors != nil {
		response["message"] = "Bad request"
		response["errors"] = errors
		ctx.Respond(http.StatusBadRequest, response)
		return
	}
	var userInfo string
	var status int
	err = x.db.QueryRow(
		"call add_user(?, ?, ?, ?, ?, ?, ?, ?, ?)",
		user.JobTitle,
		user.FirstName,
		user.LastName,
		user.Gender,
		user.BirthDate,
		user.DepartmentName,
		user.BadgeID,
		user.Phone,
		user.Email,
	).Scan(&userInfo, &status)

	if err != nil {
		response["message"] = "Internal serve error: sql"
		response["error"] = fmt.Sprintf("%+v", err)
		ctx.Respond(http.StatusInternalServerError, response)
		return
	}

	if status == 0 {
		message := fmt.Sprintf("%v", userInfo)
		response["message"] = "Bad request"
		response["errors"] = map[string]string{"badge_id": message}
		ctx.Respond(http.StatusBadRequest, response)
		return
	}

	response["message"] = "A user was successfully added"
	response["data"] = map[string]interface{}{"id": userInfo}

	ctx.Respond(http.StatusCreated, response)
}

// GetUserByID - find a user with an Id and add Chuck Norries Fact to the response
func (x UserController) GetUserByID(ctx *utils.Context) {
	response := map[string]interface{}{}

	pID, err := strconv.Atoi(ctx.Params[0])
	if err != nil {
		response["message"] = "Bad request. user id should be an integer"
		ctx.Respond(http.StatusBadRequest, response)
		return
	}

	user := &models.User{}

	err = x.db.QueryRow("call get_user_by_id(?)", pID).Scan(
		&user.ID,
		&user.JobTitle,
		&user.FirstName,
		&user.LastName,
		&user.Gender,
		&user.BirthDate,
		&user.DepartmentName,
		&user.BadgeID,
		&user.Phone,
		&user.Email,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			response["message"] = "No user with given id"
			response["user"] = map[string]interface{}{}
			ctx.Respond(http.StatusNotFound, response)
		} else {
			response["message"] = "Internal serve error: sql"
			response["error"] = fmt.Sprintf("%+v", err)
			ctx.Respond(http.StatusInternalServerError, response)
		}
		return
	}

	response["message"] = "A user was successfully retrieved"
	response["user"] = user

	url := "https://matchilling-chuck-norris-jokes-v1.p.rapidapi.com/jokes/random"
	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("x-rapidapi-host", "matchilling-chuck-norris-jokes-v1.p.rapidapi.com")
	req.Header.Add("x-rapidapi-key", config.XRapidAPIKey)
	req.Header.Add("accept", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		response["chuck_norris"] = map[string]interface{}{}
	} else {
		defer res.Body.Close()

		var fact map[string]interface{}
		err = json.NewDecoder(res.Body).Decode(&fact)
		if err != nil {
			response["chuck_norris"] = map[string]interface{}{}
		} else {
			response["chuck_norris"] = fact
		}
	}

	ctx.Respond(http.StatusOK, response)
}

// SearchUsers - search users based on their name, job title, department name, phone or email
func (x UserController) SearchUsers(ctx *utils.Context) {
	var searchParams = ctx.Request.URL.Query()

	response := map[string]interface{}{}

	rows, err := x.db.Query(
		"call search_users(?, ?, ?, ?, ?)",
		searchParams.Get("name"),
		searchParams.Get("job_title"),
		searchParams.Get("department_name"),
		searchParams.Get("phone"),
		searchParams.Get("email"),
	)

	if err != nil {
		response["message"] = "Internal serve error: sql"
		response["error"] = fmt.Sprintf("%+v", err)
		ctx.Respond(http.StatusInternalServerError, response)
		return
	}

	users := []*models.User{}
	for rows.Next() {
		user := &models.User{}
		err = rows.Scan(
			&user.ID,
			&user.JobTitle,
			&user.FirstName,
			&user.LastName,
			&user.Gender,
			&user.BirthDate,
			&user.DepartmentName,
			&user.BadgeID,
			&user.Phone,
			&user.Email,
		)
		if err != nil {
			response["message"] = "Internal serve error: sql"
			response["error"] = fmt.Sprintf("%+v", err)
			ctx.Respond(http.StatusInternalServerError, response)
			return
		}
		users = append(users, user)
	}

	response["message"] = "search was done successfully"
	response["users"] = users

	ctx.Respond(http.StatusOK, response)
}
