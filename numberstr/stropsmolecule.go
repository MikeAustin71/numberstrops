package numberstr

import (
	"fmt"
	"strings"
	"sync"
)

type strOpsMolecule struct {
	lock *sync.Mutex
}

// StrPadLeftToCenter - Returns a blank string which allows centering of the target
// string in a fixed length field.
//
// Example:
//
// Assume that total field length ('fieldlen') is 70. Assume that the string to Center
// ('strToCenter') is 10-characters. In order to center a 10-character string in a
// 70-character field, 30-space characters would need to be positioned on each side
// of the string to center. This method only returns the left margin, or a string
// consisting of 30-spaces.
//
func (sOpsMolecule *strOpsMolecule) strPadLeftToCenter(
	strToCenter string,
	fieldLen int,
	ePrefix string) (
	string,
	error) {

	if sOpsMolecule.lock == nil {
		sOpsMolecule.lock = new(sync.Mutex)
	}

	sOpsMolecule.lock.Lock()

	defer sOpsMolecule.lock.Unlock()

	if len(ePrefix) > 0 {
		ePrefix += "\n"
	}

	ePrefix += "strOpsMolecule.strPadLeftToCenter()\n"

	sOpsQuark := strOpsQuark{}

	if sOpsQuark.isEmptyOrWhiteSpace(strToCenter) {
		return strToCenter,
			fmt.Errorf("%v\n"+
				"Error: Input parameter 'strToCenter' is All White Space or an EMPTY String!\n",
				ePrefix)
	}

	sLen := sOpsQuark.getRuneCountInStr(strToCenter)

	if sLen > fieldLen {
		return "",
			fmt.Errorf("%v\n"+
				"Error: Input Parameter String To Center ('strToCenter')\n"+
				"is longer than Field Length\n",
				ePrefix)
	}

	if sLen == fieldLen {
		return "", nil
	}

	margin := (fieldLen - sLen) / 2

	return strings.Repeat(" ", margin), nil
}
