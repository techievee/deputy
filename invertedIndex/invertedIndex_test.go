package invertedIndex

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	roles "github.com/techievee/deputy/data/roles"
	users "github.com/techievee/deputy/data/users"
)

var (
	roleIndex, userIndex *InvertedIndex
)

func TestMain(m *testing.M) {

	roleIndex = NewInvertedIndex()
	userIndex = NewInvertedIndex()

	c := m.Run()
	os.Exit(c)
}

// Add roles index
func TestRolesIndex(t *testing.T) {

	roleData1 := []roles.Roles{
		{
			Id:     1,
			Name:   "System Administrator",
			Parent: 0,
		},
		{
			Id:     2,
			Name:   "Location Manager",
			Parent: 1,
		},
		{
			Id:     3,
			Name:   "Supervisor",
			Parent: 2,
		},
		{
			Id:     4,
			Name:   "Employee",
			Parent: 3,
		},
		{
			Id:     5,
			Name:   "Trainer",
			Parent: 3,
		},
		{
			Id:     6,
			Name:   "New Additional Supervisor",
			Parent: 2,
		},
	}

	for _, val := range roleData1 {
		roleIndex.AddItem(val.Parent, val.Id, &val)
	}

	// Test
	result := roleIndex.FindMaps(0)
	assert.Equal(t, 1, len(result))

	result = roleIndex.FindMaps(1)
	assert.Equal(t, 1, len(result))

	result = roleIndex.FindMaps(2)
	assert.Equal(t, 2, len(result))

	result = roleIndex.FindMaps(3)
	assert.Equal(t, 2, len(result))

}

func TestUserIndex(t *testing.T) {

	userData1 := []users.Users{
		{
			Id:   1,
			Name: "Adam Admin",
			Role: 1,
		},
		{
			Id:   2,
			Name: "Emily Employee",
			Role: 4,
		},
		{
			Id:   3,
			Name: "Sam Supervisor",
			Role: 3,
		},
		{
			Id:   31,
			Name: "Dandy Supervisor",
			Role: 3,
		},
		{
			Id:   4,
			Name: "Mary Manager",
			Role: 2,
		},
		{
			Id:   5,
			Name: "Steve Trainer",
			Role: 5,
		},
		{
			Id:   6,
			Name: "Jack Trainer",
			Role: 5,
		},
		{
			Id:   7,
			Name: "Ben Trainer",
			Role: 5,
		},
	}

	for _, val := range userData1 {
		userIndex.AddItem(val.Role, val.Id, &val)
	}

	// Test
	result := userIndex.FindMaps(1)
	assert.Equal(t, 1, len(result))

	result = userIndex.FindMaps(2)
	assert.Equal(t, 1, len(result))

	result = userIndex.FindMaps(3)
	assert.Equal(t, 2, len(result))

	result = userIndex.FindMaps(4)
	assert.Equal(t, 1, len(result))

	result = userIndex.FindMaps(5)
	assert.Equal(t, 3, len(result))

}
