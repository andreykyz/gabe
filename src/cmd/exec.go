package cmd

import (
	"bytes"
	"fmt"
	"os/exec"

	"github.com/juddus/gabe/src/types"
)

func Exec(sh *types.Shell) bool {
	cmd := exec.Command(sh.C_cmd, sh.JoinArgs())
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		fmt.Printf("%s: Unknown command '%s'\n", "gabe", sh.C_cmd)
		return false
	} else {
		fmt.Print(out.String())
	}
	return true
}
