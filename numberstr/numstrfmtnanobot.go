package numberstr

import (
	"fmt"
	"strings"
	"sync"
)

type numStrFormatNanobot struct {
	lock *sync.Mutex
}

// testNumStrFormatValidity - Tests the validity of a Number Format
// string. If the string is invalid, the return parameter 'isValid'
// is set to 'false' and a detailed error message is returned.
//
// If the string is valid, 'isValid' is set to true and the returned
// error value is set to 'nil'.
//
func (nStrFmtNanobot *numStrFormatNanobot) testNumStrFormatValidity(
	numFormatStr string,
	ePrefix string) (
	isValid bool,
	err error) {

	if nStrFmtNanobot.lock == nil {
		nStrFmtNanobot.lock = new(sync.Mutex)
	}

	nStrFmtNanobot.lock.Lock()

	defer nStrFmtNanobot.lock.Unlock()

	ePrefix += "numStrFormatNanobot.testNumStrFormatValidity() "

	isValid = false

	if len(numFormatStr) == 0 {

		err = fmt.Errorf("%v\n"+
			"Error: The Number String Format is an EMPTY String!\n",
			ePrefix)

		return isValid, err
	}

	nStrFmtQuark := numStrFormatQuark{}

	validFmtChars := nStrFmtQuark.getValidFormatChars()

	var lenValidFmtChars = len(validFmtChars)
	runesToTest := []rune(numFormatStr)
	var lenCurrFmt = len(runesToTest)
	var isRuneValid bool

	for i := 0; i < lenCurrFmt; i++ {

		isRuneValid = false

		for j := 0; j < lenValidFmtChars; j++ {

			if runesToTest[i] != validFmtChars[j] {
				continue
			} else {
				isRuneValid = true
				break
			}
		}

		if !isRuneValid {
			isValid = false
			err = fmt.Errorf(ePrefix+"\n"+
				"Error: The Number String Format contains an invalid character!\n"+
				"Complete Number String Format= '%v'\n"+
				"invalid char == '%v' at Index [%v] \n",
				numFormatStr,
				string(runesToTest[i]),
				i)

			return isValid, err
		}
	}

	if strings.Contains(numFormatStr, "(") &&
		!strings.Contains(numFormatStr, ")") {

		isValid = false
		err = fmt.Errorf("%v\n"+
			"Error: The Number String Format is missing a closing\n"+
			"parenthesis mark ')'. The format contains a '(', but no ')'.\n"+
			"Complete Number String Format= '%v'\n",
			ePrefix,
			numFormatStr)

		return isValid, err
	}

	if strings.Contains(numFormatStr, ")") &&
		!strings.Contains(numFormatStr, "(") {

		isValid = false
		err = fmt.Errorf("%v\n"+
			"Error: The Number String Format is missing an opening\n"+
			"parenthesis mark '('. The format contains a ')', but no '('.\n"+
			"Complete Number String Format= '%v'\n",
			ePrefix,
			numFormatStr)

		return isValid, err
	}

	if !strings.Contains(numFormatStr, "127.54") &&
		!strings.Contains(numFormatStr, "NUMFIELD") {
		isValid = false
		err = fmt.Errorf("%v\n"+
			"Error: The Number String Format is missing a place holder\n"+
			"for the numeric value. The format contain either '127.54'\n"+
			"or 'NUMFIELD' to designate a place holder for the numeric value.\n"+
			"This Number String Format has neither placeholder!\n"+
			"Complete Number String Format= '%v'\n",
			ePrefix,
			numFormatStr)

		return isValid, err
	}

	isValid = true
	return isValid, err
}
