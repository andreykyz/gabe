package tests

import (
	"testing"

	"github.com/juddus/gabe/src/cmd"
	"github.com/juddus/gabe/src/types"
)

func TestMultArgs(t *testing.T) {
	slice1, slice2 := []string{"echo", "one", "two", "three", "four"}, []string{"one", "two", "three", "four"}
	sh := types.Shell{
		true,
		"echo",
		slice1,
		slice2,
	}

	ret := cmd.Exec(&sh)
	if ret == false {
		t.Errorf("Command returned Error") // TODO: make more concise
	}
}
