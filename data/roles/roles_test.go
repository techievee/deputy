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
	jsonData = `[{
				 "Id": 1,
				 "Name": "System Administrator",
				 "Parent": 0
				 },
				 {
				 "Id": 2,
				 "Name": "Location Manager",
				 "Parent": 1
				 },
				 {
				 "Id": 3,
				 "Name": "Supervisor",
				 "Parent": 2
				 },
				 {
				 "Id": 4,
				 "Name": "Employee",
				 "Parent": 3
				 },
				 {
				 "Id": 5,
				 "Name": "Trainer",
				 "Parent": 3
				 }]`

	var roles []Roles
	assert.Nil(t, json.Unmarshal([]byte(jsonData), &roles))
	for _, role := range roles {
		assert.Nil(t, role.Validate())
	}
}

func TestInvalidID(t *testing.T) {
	jsonData = `[{
				 "Id": "1",
				 "Name": "System Administrator",
				 "Parent": 0
				 },
				 {
				 "Id": 2,
				 "Name": "Location Manager",
				 "Parent": 1
				 }]`

	var roles []Roles
	assert.Error(t, json.Unmarshal([]byte(jsonData), &roles))
}

func TestInvalidID2(t *testing.T) {
	jsonData = `[{
				 "Id": 1a,
				 "Name": "System Administrator",
				 "Parent": 0
				 },
				 {
				 "Id": 2,
				 "Name": "Location Manager",
				 "Parent": 1
				 }]`

	var roles []Roles
	assert.Error(t, json.Unmarshal([]byte(jsonData), &roles))
}

func TestEmptyID(t *testing.T) {
	jsonData = `[{
				 "Name": "System Administrator",
				 "Parent": 3
				 },
				 {
				 "Name": "Location Manager",
				 "Parent": 1
				 }]`

	var roles []Roles
	assert.Nil(t, json.Unmarshal([]byte(jsonData), &roles))
	totalErrors := 0
	for _, role := range roles {
		totalErrors += len(role.Validate())
	}
	// First Case - Empty ID, first input

	assert.Equal(t, 2, totalErrors)
}

// Invalid Name, should return 1 Error
func TestInvalidName(t *testing.T) {
	jsonData = `[{
				 "Id": 1,
				 "Name": 3,
				 "Parent": 0
				 },
				 {
				 "Id": 2,
				 "Name": "Location Manager",
				 "Parent": 1
				 }]`

	var roles []Roles
	assert.Error(t, json.Unmarshal([]byte(jsonData), &roles))
}

// Invalid Name, should return 1 Error
func TestInvalidNameEmpty(t *testing.T) {
	jsonData = `[{
				 "Id": 1,
				 "Parent": 0
				 },
				 {
				 "Id": 2,
				 "Name": "Location Manager",
				 "Parent": 1
				 }]`

	var roles []Roles
	assert.Nil(t, json.Unmarshal([]byte(jsonData), &roles))
	totalErrors := 0
	for _, role := range roles {
		totalErrors += len(role.Validate())
	}
	// First Case - Empty Name
	assert.Equal(t, 1, totalErrors)
}

func TestParentID(t *testing.T) {
	jsonData = `[{
				 "Id": 1,
				 "Name": "System Administrator",
				 "Parent": "2"
				 },
				 {
				 "Id": 2,
				 "Name": "Location Manager",
				 "Parent": 1
				 }]`

	var roles []Roles
	assert.Error(t, json.Unmarshal([]byte(jsonData), &roles))
}

func TestInvalidParentID2(t *testing.T) {
	jsonData = `[{
				 "Id": 1,
				 "Name": "System Administrator",
				 "Parent": 2a
				 },
				 {
				 "Id": 2,
				 "Name": "Location Manager",
				 "Parent": 1
				 }]`

	var roles []Roles
	assert.Error(t, json.Unmarshal([]byte(jsonData), &roles))
}

func TestEmptyParentID(t *testing.T) {
	jsonData = `[{
				 "Name": "System Administrator",
				 "Id": 3
				 },
				 {
				 "Name": "Location Manager",
				 "Id": 1
				 }]`

	var roles []Roles
	assert.Nil(t, json.Unmarshal([]byte(jsonData), &roles))

	for _, role := range roles {
		assert.Equal(t, uint64(0), role.Parent)
	}

}
