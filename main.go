package main

import (
	"example/secure-stash/cli"
	"example/secure-stash/manager"
	"fmt"
)

const APP_NAME = "secure-stash"

func main() {
	password := cli.ScanPassword()
	if ok, err := manager.Init(password); !ok {
		cli.ShowLoginMessage(false, err.Error())
		return
	}
	cli.ShowLoginMessage(true, "")

	// pivotNumber := cli.DetermineAction() // switch on this number below
	
	// INSERT ENTRY

	key := "coinbase"; val := "valy44444m!" // TOGO
	// key, val := cli.ReadInputAfterDisplaying(key, val string) {}
	err := manager.InsertEntry(key, val)
	if err != nil {
		str := fmt.Sprintf("Unable to insert '%s' into '%s' ", key, APP_NAME)
		panic(str)
	}
	
	// GET ENTRY

	// key = cli.ReadKeyFromDisplay()
	readValue, err := manager.RetrieveEntry(key)
	if err != nil {
		panic("Error when retrieving value")
	}
	// cli.WriteValueToDisplay(readValue)
	fmt.Printf("%s has been retrieved from database\n", readValue)

	// GET KEYS
	keys, err := manager.RetrieveEntries()
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(keys)
}
