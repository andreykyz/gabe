package util

import (
	"fmt"
	"os"
	"os/user"
)

var (
	USER     string
	HOSTNAME string
	CWD      string
)

func init() {
	USER = getUser()
	HOSTNAME = getHostname()
	CWD = getCwd()
}

func getUser() string {
	user, err := user.Current()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error (user):", err)
		return "" // temp
	}
	return user.Username
}

func getHostname() string {
	hn, err := os.Hostname()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error (hostname):", err)
		return "" // temp
	}
	return hn
}
func getCwd() string {
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error (cwd):", err)
		return "" // temp
	}
	return cwd
}
