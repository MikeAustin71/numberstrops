package numberstr

import (
	"fmt"
	"sync"
)

type strOpsNanobot struct {
	lock *sync.Mutex
}

// strCenterInStrLeft - returns a string which includes a left pad blank string
// plus the original string. It does NOT include the Right pad blank string.
//
// Nevertheless, the complete string will effectively center the original string
// in a field of specified length.
//
// Example:
//   In this example the total field length is 15. The length of the test string,
//   "Hello", is 5. In order to center the test string in a field with length of 15,
//   there will be 5-spaces on the left and 5-spaces on the right. This method will
//   compute the left-pad spaces necessary to center the string in a field with length
//   of 15, but will only include the padded left margin of 5-spaces. It will NOT
//   include the trailing 5-spaces on the right.
//
//   In the following example, the final result string will substitute the'@' character
//   for the white space character (0x20) in order to illustrate the padding added by
//   this method.
//
//    strToCenter     = "Hello"
//    fieldLen        = 15
//    Returned String = "@@@@@Hello" or "     Hello"
//
func (sOpsNanobot *strOpsNanobot) strCenterInStrLeft(
	strToCenter string,
	fieldLen int,
	ePrefix string) (
	string,
	error) {

	if sOpsNanobot.lock == nil {
		sOpsNanobot.lock = new(sync.Mutex)
	}

	sOpsNanobot.lock.Lock()

	defer sOpsNanobot.lock.Unlock()

	if len(ePrefix) > 0 {
		ePrefix += "\n"
	}

	ePrefix += "strOpsNanobot.strCenterInStrLeft()\n"

	sOpsQuark := strOpsQuark{}

	if sOpsQuark.isEmptyOrWhiteSpace(strToCenter) {
		return "",
			fmt.Errorf("%v\n"+
				"Error: Input parameter 'strToCenter' is All White Space or an EMPTY String!\n",
				ePrefix)
	}

	if fieldLen < len(strToCenter) {
		return "",
			fmt.Errorf(ePrefix+
				"\n"+
				"Error: Input parameter 'fieldLen' is less than length of 'strToCenter'.\n"+
				"strToCenter length='%v'\n"+
				"fieldLen='%v'\n",
				len(strToCenter), fieldLen)
	}

	sOpsMolecule := strOpsMolecule{}

	pad, err := sOpsMolecule.strPadLeftToCenter(
		strToCenter,
		fieldLen,
		ePrefix)

	if err != nil {
		return "",
			fmt.Errorf(ePrefix+
				"\n"+
				"Error returned by sops.StrPadLeftToCenter(strToCenter, fieldLen).\n"+
				"Error='%v'\n", err.Error())
	}

	return pad + strToCenter, nil

}
