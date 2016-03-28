package cmd

import (
	"bytes"
	"fmt"
	"os/exec"

	"github.com/juddus/gabe/src/types"
)

func Exec(sh *types.Shell) bool {
	switch sh.C_cmd {
	case "exit", "gabe":
		handleBuiltin(sh)
	default:
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
	}
	return true
}

func handleBuiltin(sh *types.Shell) {
	switch sh.C_cmd {
	case "exit":
		sh.Running = false
	case "gabe":
		if len(sh.C_args) >= 1 {
			fmt.Printf("Gabe says %s\n", sh.JoinArgs())
		} else {
			fmt.Println("Gabe says woof")
		}
	}
}
