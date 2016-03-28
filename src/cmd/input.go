package cmd

import (
	//"fmt"
	"strings"

	"github.com/juddus/gabe/src/types"
)

func HandleInput(sh *types.Shell, line string) {
	sh.C_all = strings.Split(strings.Trim(line, "\n"), " ")
	//fmt.Printf("%v (%T)", sh.all, sh.all) // debug
	sh.C_cmd = sh.C_all[:1][0]
	if len(sh.C_all) > 1 {
		sh.C_args = sh.C_all[1:]
	}
	Exec(sh)

	//sh.Running = false // debug
}
