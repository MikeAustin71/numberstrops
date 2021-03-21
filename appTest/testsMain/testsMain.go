package testsMain

import (
	"fmt"
	numstr "github.com/MikeAustin71/numberstrops/numberstr"
)

type TestMain struct {
	input  string
	output string
}

func (tMain *TestMain) Test000010(
	ePrefix string) (
	err error) {

	errPrefixDto := numstr.ErrPrefixDto{}.NewEPrefOld(ePrefix)

	errPrefixDto.SetEPref("TestMain.TestAddNumStrs000010()")

	fmt.Printf("%v\n"+
		"Place Holder Test # 1\n",
		errPrefixDto.String())

	err = nil
	return err
}
