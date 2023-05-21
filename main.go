package main

import (
	"example/secure-stash/manager"
	"fmt"
)

const IV = "12345678901234567890123456789012"

func main() {
	manager.Init(IV)
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
