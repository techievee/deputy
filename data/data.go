package data

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"

	"golang.org/x/sync/errgroup"

	roles "github.com/techievee/deputy/data/roles"
	users "github.com/techievee/deputy/data/users"
)

// Data in-memory representation of the Files in data structures
type Data struct {

	// Pointer to the structure arrays
	Roles *[]roles.Roles
	Users *[]users.Users

	// File pointer to files, used only till indexing
	// Can be used later for disk based indexing
	rolesFile *os.File
	usersFile *os.File
}

// NewDataStore opens the files in the read mode and initialized the Data object to be used later
func NewDataStore(rFile string, uFile string) (*Data, error) {

	var err error
	// Default to Read Access
	rolesFile, err := os.Open(rFile)
	if err != nil {
		return nil, err
	}
	usersFile, err := os.Open(uFile)
	if err != nil {
		return nil, err
	}

	d := &Data{
		Roles:     nil,
		Users:     nil,
		rolesFile: rolesFile,
		usersFile: usersFile,
	}

	return d, nil

}

// ReadFiles Read the file and decode them to the object asynchronously using the wait group
func (d *Data) ReadFiles() error {

	// Errorgroups for reading file asynchronously
	var errorGroup errgroup.Group

	// Read and decode the file, async
	errorGroup.Go(d.decodeRoles)
	errorGroup.Go(d.decodeUsers)

	if err := errorGroup.Wait(); err != nil {
		return err
	}

	return nil
}

// Private functions - Reads each files as a stream and loads them to memory
// Tested with large files ˜1GB, designed to use less memory footprint during load
// decodeRoles Decodes the roles file to memory
func (d *Data) decodeRoles() error {
	var err error

	dec := json.NewDecoder(d.rolesFile)
	var rolesArray []roles.Roles

	// read open bracket
	_, err = dec.Token()
	if err != nil {
		log.Fatal(err)
	}

	// while the array contains values
	for dec.More() {
		var role roles.Roles
		// decode an array value
		err = dec.Decode(&role)
		if err != nil {
			return err
		}
		// Add custom validator for the file as per the rule
		if errorList := role.Validate(); len(errorList) > 0 {
			var errString string
			for _, errVal := range errorList {
				errString = fmt.Sprintf("%s\n%s", errString, errVal)
			}
			return errors.New(errString)
		}

		rolesArray = append(rolesArray, role)
	}

	// read closing bracket
	_, err = dec.Token()
	if err != nil {
		return err
	}

	if len(rolesArray) > 0 {
		d.Roles = &rolesArray
		fmt.Printf("\nRead %d record from the Roles file", len(rolesArray))
	}
	return err
}

// decodeUsers Decodes the users file to memory
func (d *Data) decodeUsers() error {

	var err error

	dec := json.NewDecoder(d.usersFile)
	var usersArray []users.Users

	// read open bracket
	_, err = dec.Token()
	if err != nil {
		log.Fatal(err)
	}

	// while the array contains values
	for dec.More() {
		var user users.Users
		// decode an array value
		err = dec.Decode(&user)
		if err != nil {
			return err
		}

		// Add custom validator for the file as per the rule
		if errorList := user.Validate(); len(errorList) > 0 {
			var errString string
			for _, errVal := range errorList {
				errString = fmt.Sprintf("%s\n%s", errString, errVal)
			}
			return errors.New(errString)
		}

		usersArray = append(usersArray, user)
	}

	// read closing bracket
	_, err = dec.Token()
	if err != nil {
		return err
	}

	if len(usersArray) > 0 {
		d.Users = &usersArray
		fmt.Printf("\nRead %d record from the users file", len(usersArray))
	}
	return err
}
