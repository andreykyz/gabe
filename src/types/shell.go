package types

import "bytes"

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
	var buffer bytes.Buffer
	for _, value := range s.C_args {
		buffer.WriteString(value)
		buffer.WriteString(" ")
	}
	return buffer.String()
}
