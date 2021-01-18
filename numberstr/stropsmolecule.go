package numberstr

import (
	"fmt"
	"strings"
	"sync"
)

type strOpsMolecule struct {
	lock *sync.Mutex
}

// StrCenterInStr - returns a string which includes a left pad blank
// string plus the original string ('strToCenter'), plus a right pad
// blank string.
//
// The returned string will effectively center the original string
// ('strToCenter') in a field of specified length ('fieldLen').
//
func (sOpsMolecule *strOpsMolecule) strCenterInStr(
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
				"Error: Input parameter 'strToCenter' is All White "+
				"Space or an EMPTY String!\n",
				ePrefix)
	}

	sLen := len(strToCenter)

	if sLen > fieldLen {
		return strToCenter,
			fmt.Errorf(ePrefix+
				"\n"+
				"Error: 'fieldLen' = '%v' strToCenter Length= '%v'.\n"+
				"'fieldLen' is shorter than 'strToCenter' Length!\n",
				fieldLen, sLen)
	}

	if sLen == fieldLen {
		return strToCenter, nil
	}

	leftPadCnt := (fieldLen - sLen) / 2

	leftPadStr := strings.Repeat(" ", leftPadCnt)

	rightPadCnt := fieldLen - sLen - leftPadCnt

	rightPadStr := ""

	if rightPadCnt > 0 {
		rightPadStr = strings.Repeat(" ", rightPadCnt)
	}

	return leftPadStr + strToCenter + rightPadStr, nil
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
