package types

import "strings"

type Shell struct {
	Running       bool
	C_cmd         string
	C_all, C_args []string
}

func CreateShell() *Shell {
	sh := &Shell{
		true,
		"",
		make([]string, 1),
		make([]string, 1),
	}
	return sh
}

func (s Shell) JoinArgs() string {
	return strings.Join(s.C_args, " ")
}
