package controllers_test

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"testing"

	"github.com/connor1004/go-api-boilerplate/app"
	"github.com/connor1004/go-api-boilerplate/models"
	"github.com/connor1004/go-api-boilerplate/utils"
)

var testApp *app.App

func TestMain(m *testing.M) {
	testApp = app.NewApp()
	testApp.InitializeDB()
	defer testApp.DB.Close()

	testApp.InitializeRoutes()

	code := m.Run()

	os.Exit(code)
}

func getMockUser() *models.User {
	return &models.User{
		JobTitle:       "Dentist",
		FirstName:      "Johnathan",
		LastName:       "Furnell",
		Gender:         "M",
		BirthDate:      "1990-12-21 20:16:05Z",
		DepartmentName: "Human Resources",
		BadgeID:        199999999,
		Phone:          "8-648-350-5156",
		Email:          "Johnathan_Furnell9568@eirey.tech",
	}
}

func addUser() (string, error) {
	user := getMockUser()
	var userInfo string
	var status int
	err := testApp.DB.QueryRow(
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
		log.Fatal("Internal Server Error")
		return "", err
	}

	return userInfo, err
}

func clearTestData() {
	testApp.DB.Exec("DELETE FROM users WHERE badge_id = ?", 199999999)
}

func deleteUser(pID int) {
	testApp.DB.Exec("DELETE FROM users WHERE id = ?", pID)
}

func TestAddUser(t *testing.T) {
	clearTestData()

	user := getMockUser()
	buf := new(bytes.Buffer)
	err := json.NewEncoder(buf).Encode(user)
	if err != nil {
		t.Error("Mock User data is invalid")
		return
	}

	req, err := http.NewRequest("POST", "/api/users", buf)
	if err != nil {
		t.Error("Can't create a request")
		return
	}
	req.Header.Set("Content-Type", "application/json")
	response := httptest.NewRecorder()
	testApp.ServeHTTP(response, req)

	utils.CheckResponseCode(t, http.StatusCreated, response.Code)
	if response.Code == http.StatusCreated {
		m := map[string]interface{}{}
		json.NewDecoder(response.Body).Decode(&m)

		userInfo, ok := m["data"].(map[string]interface{})
		if !ok {
			t.Error("Expected map[string]string but didn't get")
		} else {
			pID, err := strconv.Atoi(userInfo["id"].(string))
			if err != nil {
				t.Errorf("Expected id of type int but failed: %v", err)
			}
			deleteUser(pID)
		}
	}
}

func TestGetUserByID(t *testing.T) {
	clearTestData()
	userID, err := addUser()
	if err != nil {
		t.Error(err)
	}
	if _, err := strconv.Atoi(userID); err != nil {
		t.Errorf("Expected an integer but get: %v", userID)
		return
	}

	req, err := http.NewRequest("GET", "/api/users/"+userID, nil)
	if err != nil {
		t.Error("Can't create a request")
		return
	}
	req.Header.Set("Content-Type", "application/json")
	response := httptest.NewRecorder()
	testApp.ServeHTTP(response, req)

	utils.CheckResponseCode(t, http.StatusOK, response.Code)
	if response.Code == http.StatusOK {
		m := map[string]interface{}{}
		json.NewDecoder(response.Body).Decode(&m)

		user, ok := m["user"].(map[string]interface{})

		if !ok {
			t.Error("Expected user but didn't get")
		} else {
			badgeID, ok := user["badge_id"].(float64)

			if !ok || badgeID != 199999999.0 {
				t.Errorf("Expected 199999999 but get: %v", badgeID)
			}
		}
	}
	clearTestData()
}

func TestSearchUsers(t *testing.T) {
	clearTestData()
	userID, err := addUser()
	if err != nil {
		t.Error(err)
		return
	}
	if _, err := strconv.Atoi(userID); err != nil {
		t.Errorf("Expected an integer but get: %v", userID)
		return
	}

	req, err := http.NewRequest("GET", "/api/search-users", nil)
	if err != nil {
		t.Error("Can't create a request")
		return
	}
	req.Header.Set("Content-Type", "application/json")

	q := req.URL.Query()
	q.Add("name", "Johnathan Furnell")
	q.Add("job_title", "Dentist")
	q.Add("department_name", "Human Resources")
	q.Add("phone", "8-648-350-5156")
	q.Add("email", "Johnathan_Furnell9568@eirey.tech")
	req.URL.RawQuery = q.Encode()

	response := httptest.NewRecorder()
	testApp.ServeHTTP(response, req)

	utils.CheckResponseCode(t, http.StatusOK, response.Code)
	if response.Code == http.StatusOK {
		m := map[string]interface{}{}
		json.NewDecoder(response.Body).Decode(&m)

		users, ok := m["users"].([]interface{})

		if !ok {
			t.Error("Expected users but didn't get")
		} else {
			if len(users) < 1 {
				t.Error("Expected at least 1 user but didn't get")
			}
		}
	}
	clearTestData()
}
