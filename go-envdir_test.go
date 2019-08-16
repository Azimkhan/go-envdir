package main

import (
	"gotest.tools/assert"
	"os"
	"testing"
)

func TestSetEnvFromDir(t *testing.T) {
	err := SetEnvFromDir("testenv")
	assert.NilError(t, err)
	e1, _ := os.LookupEnv("A_ENV")
	e2, _ := os.LookupEnv("B_VAR")
	_, found := os.LookupEnv("C_VAR")

	assert.Equal(t, "123", e1)
	assert.Equal(t, "another_val", e2)
	assert.Equal(t, false, found)
}
