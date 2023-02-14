package main

import (
	"example/secure-stash/cacher"
	"fmt"
)

func main() {
	c := cacher.Cacher{}
	c.InitCacher()

	sampleKey := "keyx"
	sampleVal := "valx"

	err := c.InsertKey(sampleKey, sampleVal)
	if err != nil {
		panic("Unable to insert basic key")
	}
	
	val, err := c.RetrieveKey(sampleKey)
	if err != nil {
		panic("Unable to get basic key")
	} else if val != sampleVal {
		panic("Unable to ensure consistent insert")
	} else {
		fmt.Println("Redis is working :)")
	}
}
