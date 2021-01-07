package main

import (
	"fmt"
	"github.com/MikeAustin71/numberstrops/appTest/testsMain"
)

func main() {

	ePrefix := "main() "

	tMain := testsMain.TestMain{}

	err := tMain.Test001000NewInt(ePrefix)

	if err != nil {
		fmt.Printf("%v\n", err.Error())
		return
	}

	fmt.Printf(ePrefix + "\n" +
		"Successful Completion!")
	return
}
