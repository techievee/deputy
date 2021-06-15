package indexer

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/techievee/deputy/data"
	"github.com/techievee/deputy/utilities"
)

var (
	d *data.Data
)

func TestMain(m *testing.M) {

	// Read from the test_sample path to bootstrap the files
	path := utilities.GetRootProjectPath()
	rFile := fmt.Sprintf("%s/test_samples/%s", path, "roles.json")
	uFile := fmt.Sprintf("%s/test_samples/%s", path, "users.json")

	// Read the data from the current test directory for running test
	d, _ = data.NewDataStore(rFile, uFile)
	err := d.ReadFiles()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	c := m.Run()
	os.Exit(c)
}

func TestGetChildForRoles(t *testing.T) {
	indexer := NewIndexer(d)
	indexer.GenerateIndex()

	result := indexer.GetChildForRoles(2)
	assert.Equal(t, []uint64{3}, result)

	result = indexer.GetChildForRoles(1)
	assert.Equal(t, []uint64{2}, result)

	result = indexer.GetChildForRoles(0)
	assert.Equal(t, []uint64{1}, result)

	result = indexer.GetChildForRoles(3)
	assert.Equal(t, 2, len(result))

	result = indexer.GetChildForRoles(4)
	assert.Equal(t, 0, len(result))

	result = indexer.GetChildForRoles(5)
	assert.Equal(t, 0, len(result))

}

func TestFindSubordinatesRoles(t *testing.T) {
	indexer := NewIndexer(d)
	indexer.GenerateIndex()

	result := indexer.FindSubordinatesRoles(1)
	assert.Equal(t, 4, result.Size())

	result = indexer.FindSubordinatesRoles(2)
	assert.Equal(t, 0, result.Size())

	result = indexer.FindSubordinatesRoles(3)
	assert.Equal(t, 2, result.Size())

	result = indexer.FindSubordinatesRoles(4)
	assert.Equal(t, 3, result.Size())

	result = indexer.FindSubordinatesRoles(5)
	assert.Equal(t, 0, result.Size())

}

func TestFindUsers(t *testing.T) {
	indexer := NewIndexer(d)
	indexer.GenerateIndex()

	subRole := indexer.FindSubordinatesRoles(1)
	result := indexer.FindUsers(subRole.Items())
	assert.Equal(t, 4, len(result))

	subRole = indexer.FindSubordinatesRoles(2)
	result = indexer.FindUsers(subRole.Items())
	assert.Equal(t, 0, len(result))

	subRole = indexer.FindSubordinatesRoles(3)
	result = indexer.FindUsers(subRole.Items())
	assert.Equal(t, 2, len(result))

	subRole = indexer.FindSubordinatesRoles(4)
	result = indexer.FindUsers(subRole.Items())
	assert.Equal(t, 3, len(result))

	subRole = indexer.FindSubordinatesRoles(5)
	result = indexer.FindUsers(subRole.Items())
	assert.Equal(t, 0, len(result))

}

func TestPrintSubordinates(t *testing.T) {
	indexer := NewIndexer(d)
	indexer.GenerateIndex()

	indexer.PrintSubordinates(1)
	indexer.PrintSubordinates(2)
	indexer.PrintSubordinates(3)
	indexer.PrintSubordinates(4)
	indexer.PrintSubordinates(5)

}
