package cli

import (
	"fmt"
	"strings"
)

func ScanPassword() string {
	var password string
	fmt.Printf(strings.Repeat("*", 40) + "\n")
	fmt.Printf("************* Secure Stash *************\n")
	fmt.Printf(strings.Repeat("*", 40) + "\n")
	fmt.Printf(strings.Repeat(" ", 40) + "\n")
	fmt.Printf("Type a password to unlock: ")
	fmt.Scanln(&password)
	fmt.Printf("\nUnlocking..\n")
	fmt.Printf("\n")
	return password
}

func ShowLogin(ok bool) {
	if ok {
		fmt.Printf("Successful login :)\n")
	} else {
		fmt.Printf("Unsuccessful login attempt:(\n")
	}
}