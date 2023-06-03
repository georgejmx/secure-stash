package main

import (
	"example/secure-stash/cli"
	"example/secure-stash/manager"
	"os"
)

const APP_NAME = "secure-stash"

func main() {
	var err error
	password := cli.ScanPassword()
	if ok, err := manager.Init(password); !ok {
		cli.ShowLoginMessage(false, err.Error())
		return
	}
	cli.ShowLoginMessage(true, "")

	keys, err := manager.RetrieveEntries()
	if err != nil {
		panic("Fatal: Unable to retrieve entries from cache")
	}
	cli.ShowCurrentEntries(keys)

	for true {
		pivotRune := cli.DetermineAction()
		if pivotRune == "Z" || pivotRune == "z" {
			os.Exit(0)
		} else if pivotRune == "A" || pivotRune == "a" {
			key := cli.ParseKeyToView()
			readValue, err := manager.RetrieveEntry(key)
			if err == nil {
				cli.ShowRetrieveEntrySuccess(key, readValue)
			} else {
				cli.ShowInvalidKeyScreen(key)
			}
		} else if pivotRune == "B" || pivotRune == "b" {
			newKey, newVal := cli.ParseNewEntry()
			err = manager.InsertEntry(newKey, newVal)
			if err == nil {
				cli.ShowAddEntrySuccess(newKey)
			} else {
				cli.ShowInvalidEntryScreen()
			}
		} else {
			cli.ShowInvalidActionScreen()
		}
	}
}
