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
		make([]string, 0),
		make([]string, 0),
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

func (s *Shell) Tick() {
	// clear current commands for next loop
	// TODO: implement history
	s.C_cmd = ""
	s.C_all, s.C_args = make([]string, 0), make([]string, 0)
}
