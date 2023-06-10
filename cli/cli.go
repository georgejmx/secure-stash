package cli

import (
	"fmt"
	"strings"
)

func ScanPassword(appName string) string {
	var password string
	fmt.Printf(strings.Repeat("*", 40) + "\n")
	fmt.Printf("************* ")
	fmt.Printf(appName)
	fmt.Printf(" *************\n")
	fmt.Printf(strings.Repeat("*", 40) + "\n")
	fmt.Printf(strings.Repeat(" ", 40) + "\n")
	fmt.Printf("Type a password to initialise or unlock: ")
	fmt.Scanln(&password)
	fmt.Printf("\nUnlocking..\n")
	fmt.Printf("\n")
	return password
}

func ShowLoginMessage(ok bool, errMsg string) {
	if ok {
		fmt.Printf("Successful login :)\n\n")
	} else {
		if errMsg == "cipher: message authentication failed" {
			fmt.Printf("Invalid password :(\n")
		} else {
			fmt.Printf("Error logging in: %s\n", errMsg)
		}
	}
}

func ShowCurrentEntries(keys []string) {
	if len(keys) == 0 {
		fmt.Printf("You have no entries. ")
		return
	}

	fmt.Printf("|-----------------|\n")
	fmt.Printf("| Current Entries |\n")
	fmt.Printf("|-----------------|\n")
	for key, val := range keys {
		fmt.Printf(" (%d) %s\n", key, val)
	}
	fmt.Print("\n")
}

func DetermineAction() string {
	var action string
	fmt.Printf("Enter (A) to retrieve an entry, (B) to add an entry or (Z) to quit: ")
	fmt.Scanln(&action)
	fmt.Print("\n")
	return action
}

func ParseKeyToView() string {
	var key string
	fmt.Printf("Type a key to view its password: ")
	fmt.Scanln(&key)
	fmt.Print("\n")
	return key
}

func ParseNewEntry() (string, string) {
	var key string
	var val string
	fmt.Printf("Enter an entry name: ")
	fmt.Scanln(&key)
	fmt.Printf("Enter a secure password: ")
	fmt.Scanln(&val)
	fmt.Print("\n")
	return key, val
}

func ShowRetrieveEntrySuccess(key, val string) {
	fmt.Printf("%s : %s\n", key, val)
}

func ShowAddEntrySuccess(key string) {
	fmt.Printf("%s successfully added to cache :)\n", key)
} 

func ShowInvalidActionScreen() {
	fmt.Print("Please enter either (A), (B) or (Z)\n")
}

func ShowInvalidKeyScreen(key string) {
	fmt.Printf("%s not found\n", key)
}

func ShowInvalidEntryScreen() {
	fmt.Printf("Error when adding new entry\n")
}