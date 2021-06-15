package console

import (
	"fmt"
	"os"
	"testing"

	"github.com/techievee/deputy/utilities"

	"github.com/stretchr/testify/assert"

	"github.com/techievee/deputy/data"
	"github.com/techievee/deputy/indexer"
)

var (
	d       *data.Data
	i       *indexer.Indexer
	console *Console
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

	i = indexer.NewIndexer(d)
	i.GenerateIndex()
	console = NewConsole(i)

	c := m.Run()
	os.Exit(c)
}

func TestRunInteractive(t *testing.T) {
	console.RunInteractive()
}

func TestRunNonInteractive(t *testing.T) {
	console.RunNonInteractive("2")
}

func TestOperationPrompt(t *testing.T) {
	_, _, _ = console.operationPrompt()
}

func TestPrompt(t *testing.T) {
	_, _ = console.Prompt()
}

func TestContinueTextPrompt(t *testing.T) {
	_, _ = console.continueTextPrompt("test")
}

func TestValidate(t *testing.T) {
	assert.NoError(t, console.validate("1"))
	assert.NoError(t, console.validate(" 1"))
	assert.NoError(t, console.validate("1"))
	assert.NoError(t, console.validate("1 "))
	assert.Error(t, console.validate("asd"))
	assert.Error(t, console.validate("'1'"))
	assert.Error(t, console.validate(",1"))
	assert.Error(t, console.validate("1s"))

}

func TestValidateInputString(t *testing.T) {
	assert.Equal(t, true, ValidateInputString("1"))
	assert.Equal(t, true, ValidateInputString(" 1"))
	assert.Equal(t, true, ValidateInputString("1"))
	assert.Equal(t, true, ValidateInputString("1 "))
	assert.Equal(t, false, ValidateInputString("asd"))
	assert.Equal(t, false, ValidateInputString("'1'"))
	assert.Equal(t, false, ValidateInputString(",1"))
	assert.Equal(t, false, ValidateInputString("1s"))
	assert.Equal(t, false, ValidateInputString(""))

}
