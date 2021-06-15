package utilities

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	c := m.Run()
	os.Exit(c)
}

func TestCurrentPath(t *testing.T) {
	path, _ := GetCurrentPath()
	assert.NotEqual(t, path, "")
}

func TestRootPath(t *testing.T) {
	path := GetRootProjectPath()
	assert.NotEqual(t, path, "")
}

func TestClearScreen(t *testing.T) {
	CallClear()
}
