package indexer

import (
	"bytes"
	"encoding/json"
	"fmt"
	"sync"

	"github.com/techievee/deputy/data"
	roles "github.com/techievee/deputy/data/roles"
	users "github.com/techievee/deputy/data/users"
	"github.com/techievee/deputy/invertedIndex"
	"github.com/techievee/deputy/sets"
)

// Indexer IndexProcessor for the Program, used to store and process index
type Indexer struct {

	// Roles and Users inverted index which are maintained for data, all processing happens via the data
	// Can be maintained as index array or map
	rolesParentIndex       *invertedIndex.InvertedIndex
	usersRoleInvertedIndex *invertedIndex.InvertedIndex
	usersIndex             *invertedIndex.InvertedIndex

	// Reference to the data that is in memory
	data *data.Data
}

// NewIndexer Creates a new index processor
func NewIndexer(data *data.Data) *Indexer {

	// Creates a new empty inverted index for roles and users
	rolesInverterIndex := invertedIndex.NewInvertedIndex()
	usersInvertedIndex := invertedIndex.NewInvertedIndex()
	usersIndex := invertedIndex.NewInvertedIndex()

	indexer := &Indexer{
		rolesParentIndex:       rolesInverterIndex,
		usersRoleInvertedIndex: usersInvertedIndex,
		usersIndex:             usersIndex,
		data:                   data,
	}
	return indexer
}

// GenerateIndex Generates indexes for all the files asynchronously using wait groups
func (indexer *Indexer) GenerateIndex() {

	// Use waitgroups to generate the indexes asynchronously
	var wg sync.WaitGroup

	// Add 3 delta as we execute 3 go routines
	wg.Add(3)
	go indexer.GenerateRolesParentInvertedIndex(&wg)
	go indexer.GenerateUserRolesIndex(&wg)
	go indexer.GenerateUserIndex(&wg)

	// Wait for all the routine to complete
	wg.Wait()
	return

}

// GenerateRolesParentInvertedIndex Create inverted index for the roles data parentRole->childRoles
func (indexer *Indexer) GenerateRolesParentInvertedIndex(wg *sync.WaitGroup) {

	defer wg.Done()

	// return if no data present
	if indexer.data.Roles == nil {
		return
	}
	roleList := *indexer.data.Roles

	// Iterate through every roles record
	for _, record := range roleList {
		var role roles.Roles
		role = record
		indexer.rolesParentIndex.AddItem(record.Parent, record.Id, &role)
	}
	return
}

// GenerateUserRolesIndex Create inverted index for the users data, roles->users
func (indexer *Indexer) GenerateUserRolesIndex(wg *sync.WaitGroup) {

	defer wg.Done()
	// return if no data present
	if indexer.data.Users == nil {
		return
	}
	userList := *indexer.data.Users

	// Iterate through every users record
	for _, record := range userList {
		var user users.Users
		user = record
		indexer.usersRoleInvertedIndex.AddItem(record.Role, record.Id, &user)
	}
	return
}

// GenerateUserIndex Create index for the users data based on the user id, usersId-> users
func (indexer *Indexer) GenerateUserIndex(wg *sync.WaitGroup) {

	defer wg.Done()
	// return if no data present
	if indexer.data.Users == nil {
		return
	}
	userList := *indexer.data.Users

	// Iterate through every userList record
	for _, record := range userList {
		var user users.Users
		user = record
		indexer.usersIndex.AddItem(record.Id, 0, &user)
	}
	return
}

// FindUsers Find all the users with the input roles
func (indexer *Indexer) FindUsers(roleIds []uint64) []*users.Users {

	var userList []*users.Users
	// For all the users from input, get all the users
	for _, role := range roleIds {
		if userIndex := indexer.usersRoleInvertedIndex.FindMaps(role); userIndex != nil {
			for _, user := range userIndex {
				u := user.(*users.Users)
				userList = append(userList, u)
			}
		}

	}

	return userList

}

// FindSubordinatesRoles Find all the subordinates(child) role ids for the given users, return roles sets
func (indexer *Indexer) FindSubordinatesRoles(userId uint64) *sets.IntegerSet {

	// Generate a list of roles that is subordinate for the current user
	childRoleList := &sets.IntegerSet{}
	var processQueue []uint64

	// Add all the child for the current user
	if userDetails := indexer.usersIndex.FindMaps(userId); userDetails != nil {
		// Get the user details, to get the current user role
		if user := indexer.usersIndex.Find(userId, 0); user != nil {

			// Conver the interface to the user
			userStruct := user.(*users.Users)
			// Find the role of the current user
			currentUserRole := userStruct.Role

			// Get all the child role id for the current user role
			childRoles := indexer.GetChildForRoles(currentUserRole)

			// Add all those child role to the process list
			for _, childRoleId := range childRoles {

				// Check if the role was already present in the Sets
				if !childRoleList.Has(childRoleId) {
					// Not present, add to the queue and sets for processing
					processQueue = append(processQueue, childRoleId)
					childRoleList.Add(childRoleId)
				}
			}

			// For every item in the processQueue, process them till the queue is empty
			// Indexed based simple queue using dynamic array
			for queueIndex := 0; queueIndex < len(processQueue); queueIndex++ {

				// Get the child roles for the current processing items
				childRoles := indexer.GetChildForRoles(processQueue[queueIndex])

				// Add each child to the process queue for further processing and to the set to send
				for _, childRoleId := range childRoles {
					if !childRoleList.Has(childRoleId) {
						// Not processes, add to the queue and set
						processQueue = append(processQueue, childRoleId)
						childRoleList.Add(childRoleId)
					}
				}
			}
		}
	}
	// Key not found, level 1
	// user not found, level 2
	return childRoleList

}

// GetChildForRoles Get all the subordinate or child role id for the given role id
func (indexer *Indexer) GetChildForRoles(roleId uint64) []uint64 {

	var roleList []uint64
	// Search the index for the child roles for the given roles
	if roleMaps := indexer.rolesParentIndex.FindMaps(roleId); roleMaps != nil {
		for rId := range roleMaps {
			roleList = append(roleList, rId)
		}
	}
	return roleList
}

// PrintSubordinates Find all subordinates and prints the pretty json format
func (indexer *Indexer) PrintSubordinates(userId uint64) []byte {

	var buffer bytes.Buffer

	// Find all the subordinate roles
	subordinateRoleIds := indexer.FindSubordinatesRoles(userId)
	if subordinateRoleIds == nil {
		buffer.WriteString(fmt.Sprintf("\n No Subordinate roles found"))
		return buffer.Bytes()
	}
	// Find all the user for the found role ids
	userList := indexer.FindUsers(subordinateRoleIds.Items())
	if len(userList) <= 0 {
		buffer.WriteString(fmt.Sprintf("\n No Subordinate user found for the subordinate roles"))
		return buffer.Bytes()
	}
	// Use in case of network traffic, data to the app or other frontends
	//e, err := json.Marshal(userList)

	// Use this for printing in console, pretty printing
	e, err := json.MarshalIndent(userList, "", "  ")
	if err != nil {
		fmt.Println(err)
		return buffer.Bytes()
	}
	return e
}
