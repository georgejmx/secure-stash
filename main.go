package main

import (
	"example/secure-stash/cli"
	"example/secure-stash/manager"
	"fmt"
)

func main() {
	password := cli.ScanPassword()
	if ok, err := manager.Init(password); !ok {
		cli.ShowLoginMessage(false, err.Error())
		return
	}
	cli.ShowLoginMessage(true, "")
	
	sampleKey := "keyw"
	sampleVal := "valx"
	err := manager.InsertEntry(sampleKey, sampleVal)
	if err != nil {
		panic("Unable to insert basic key")
	}
	
	val, err := manager.RetrieveEntry(sampleKey)
	if err != nil {
		panic("Unable to get basic key")
	} else if val != sampleVal {
		panic("Unable to ensure consistent insert")
	} else {
		fmt.Println("Redis and GCM are working :)")
	}
}
