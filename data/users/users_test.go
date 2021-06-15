package data

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Init Config
var (
	jsonData string
)

func TestMain(m *testing.M) {
	c := m.Run()
	os.Exit(c)
}

func TestValidID(t *testing.T) {
	jsonData = `[
				{
				"Id": 1,
				"Name": "Adam Admin",
				"Role": 1
				},
				{
				"Id": 2,
				"Name": "Emily Employee",
				"Role": 4
				},
				{
				"Id": 3,
				"Name": "Sam Supervisor",
				"Role": 3
				},
				{
				"Id": 4,
				"Name": "Mary Manager",
				"Role": 2
				},
				{
				"Id": 5,
				"Name": "Steve Trainer",
				"Role": 5
				}
				]`

	var users []Users
	assert.Nil(t, json.Unmarshal([]byte(jsonData), &users))
	for _, role := range users {
		assert.Nil(t, role.Validate())
	}
}

func TestInvalidID(t *testing.T) {
	jsonData = `[{
				 "Id": "1",
				 "Name": "System Administrator",
				 "Role": 2
				 },
				 {
				 "Id": 2,
				 "Name": "Location Manager",
				 "Role": 1
				 }]`

	var users []Users
	assert.Error(t, json.Unmarshal([]byte(jsonData), &users))
}

func TestInvalidID2(t *testing.T) {
	jsonData = `[{
				 "Id": 0,
				 "Name": "Emily System Administrator",
				 "Role": 1
				 },
				 {
				 "Id": 2,
				 "Name": "Andrew Location Manager",
				 "Role": 1
				 }]`

	var users []Users
	assert.Nil(t, json.Unmarshal([]byte(jsonData), &users))
	totalErrors := 0
	for _, user := range users {
		totalErrors += len(user.Validate())
	}
	assert.Equal(t, 1, totalErrors)
}

func TestEmptyID(t *testing.T) {
	jsonData = `[{
				 "Name": "Andrew System Administrator",
				 "Role": 3
				 },
				 {
				 "Name": "Emily Location Manager",
				 "Role": 1
				 }]`

	var users []Users
	assert.Nil(t, json.Unmarshal([]byte(jsonData), &users))
	totalErrors := 0
	for _, user := range users {
		totalErrors += len(user.Validate())
	}
	// First Case - Empty ID, first input
	assert.Equal(t, 2, totalErrors)
}

// Invalid Name, should return 1 Error
func TestInvalidName(t *testing.T) {
	jsonData = `[{
				 "Id": 1,
				 "Name": 3,
				 "Role": 0
				 },
				 {
				 "Id": 2,
				 "Name": "",
				 "Role": 1
				 }]`

	var users []Users
	assert.Error(t, json.Unmarshal([]byte(jsonData), &users))
}

// Invalid Name, should return 1 Error
func TestInvalidNameEmpty(t *testing.T) {
	jsonData = `[{
				 "Id": 1,
				 "Role": 2
				 },
				 {
				 "Id": 2,
				 "Name": "Emily Location Manager",
				 "Role": 1
				 }]`

	var users []Users
	assert.Nil(t, json.Unmarshal([]byte(jsonData), &users))
	totalErrors := 0
	for _, user := range users {
		totalErrors += len(user.Validate())
	}
	// First Case - Empty Name
	assert.Equal(t, 1, totalErrors)
}

func TestRoleID(t *testing.T) {
	jsonData = `[{
				 "Id": 1,
				 "Name": "System Administrator",
				 "Role": "2"
				 },
				 {
				 "Id": 2,
				 "Name": "Location Manager",
				 "Role": 1
				 }]`

	var users []Users
	assert.Error(t, json.Unmarshal([]byte(jsonData), &users))
}

func TestInvalidRoleID2(t *testing.T) {
	jsonData = `[{
				 "Id": 1,
				 "Name": "System Administrator",
				 "Role": 2a
				 },
				 {
				 "Id": 2,
				 "Name": "Location Manager",
				 "Role": 1
				 }]`

	var users []Users
	assert.Error(t, json.Unmarshal([]byte(jsonData), &users))
}

func TestEmptyRoleID(t *testing.T) {
	jsonData = `[{
				 "Name": "System Administrator",
				 "Id": 3
				 },
				 {
				 "Name": "Location Manager",
				 "Id": 1
				 }]`

	var users []Users
	assert.Nil(t, json.Unmarshal([]byte(jsonData), &users))
	totalErrors := 0
	for _, user := range users {
		totalErrors += len(user.Validate())
	}
	// First Case - Empty ID, first input

	assert.Equal(t, 2, totalErrors)
}
