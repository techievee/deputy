package data

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/techievee/deputy/utilities"
)

var data *Data

func TestMain(m *testing.M) {

	var rFile, uFile string
	var err error

	// Read from the test_sample path to bootstrap the files
	path := utilities.GetRootProjectPath()
	rFile = fmt.Sprintf("%s/test_samples/%s", path, "roles.json")
	uFile = fmt.Sprintf("%s/test_samples/%s", path, "users.json")

	// Read the data from the current test directory for running test
	data, err = NewDataStore(rFile, uFile)
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

	c := m.Run()
	os.Exit(c)
}

func TestDecodeRoles(t *testing.T) {

	err := data.decodeRoles()
	assert.NoError(t, err)
}

func TestDecodeUsers(t *testing.T) {

	err := data.decodeUsers()
	assert.NoError(t, err)
}
