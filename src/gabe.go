package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/juddus/gabe/src/cmd"
	"github.com/juddus/gabe/src/types"
	"github.com/juddus/gabe/src/util"
)

func main() {
	sh := types.CreateShell()
	//fmt.Printf("%v (%T)", sh, sh) // debug
	for sh.Running == true {
		fmt.Printf("%s%s%s%s%s%s%s%s %s%s%s%s",
			util.GREEN_TEXT, util.USER, util.RESET,
			"@",
			util.YELLOW_TEXT, util.HOSTNAME, util.RESET,
			":",
			util.CYAN_TEXT, util.CWD, util.RESET,
			" $ ")

		args := bufio.NewReader(os.Stdin)
		if line, err := args.ReadString('\n'); err != nil {
			switch err.Error() {
			case "EOF":
				fmt.Println("\nEOF Detected... Exiting")
				sh.Running = false
			default:
				fmt.Fprintln(os.Stderr, "Error (stdin):", err)
			}
		} else {
			cmd.HandleInput(sh, line)
		}
		sh.Tick()
		//sh.Running = false // debug
	}
}
