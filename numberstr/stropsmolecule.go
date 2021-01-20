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

// StrLeftJustify - Creates a new string, left-justified, within a
// with a wider text field or output string. The text to be left
// justified is specified by input parameter 'strToJustify'. The
// length of the output string is defined by input parameter,
// 'fieldLen'.
//
// Input parameter 'strToJustify' is placed on the left side of the
// output string and spaces are padded to the right in order to
// create a string with total length of 'fieldLen'.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  strToJustify        string
//     - The text content which will be left-justified in the
//       output string returned by this method.
//
//  fieldLen            int
//     - Defines the length of the output string in which input
//       parameter 'strToJustify' will be left-justified.
//
//
//  ePrefix             string
//     - This is an error prefix which is included in all returned
//       error messages. Usually, it contains the names of the calling
//       method or methods. Be sure to leave a space at the end
//       of 'ePrefix'.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  string
//     - The output string resulting from the 'left-justify'
//       operation. Input parameter, 'strToJustify' will be
//       left-justified in this output string which will have a
//       total string length as defined by input parameter,
//       'fieldLen'.
//
//
//  error
//     - If the method completes successfully and no errors are
//       encountered this return value is set to 'nil'. Otherwise,
//       if errors are encountered this return value will contain
//       an appropriate error message.
//
//       If an error message is returned, the input parameter
//       'ePrefix' (error prefix) will be inserted or prefixed at
//       the beginning of the error message.
//
//
// ------------------------------------------------------------------------
//
// Example Usage
//
//
//  ePrefix := "TestStrOps_StrLeftJustify_01() "
//  fieldLen = 15
//  strToJustify    = "Hello World"
//  su := StrOps{}
//  justifiedStr, err := su.StrLeftJustify(
//                           strToJustify,
//                           fieldLen,
//                           ePrefix)
//
//  ------------------------------------------------
//  justifiedStr is now equal to "Hello World    "
//             String Length      123456789012345
//
//
func (sOpsMolecule *strOpsMolecule) strLeftJustify(
	strToJustify string,
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

	ePrefix += "strOpsMolecule.strLeftJustify() "

	sOpsQuark := strOpsQuark{}

	if sOpsQuark.isEmptyOrWhiteSpace(strToJustify) {
		return strToJustify,
			fmt.Errorf("%v\n"+
				"Error: Input parameter 'strToJustify' is All White Space or an EMPTY String!\n",
				ePrefix)
	}

	strLen := len(strToJustify)

	if fieldLen == strLen {
		return strToJustify, nil
	}

	if fieldLen < strLen {
		return strToJustify,
			fmt.Errorf(ePrefix+
				"\n"+
				"Error: Length of string to left justify is '%v'.\n"+
				"However, 'fieldLen' is less.\n"+
				"'fieldLen'= '%v'\n",
				strLen, fieldLen)
	}

	rightPadLen := fieldLen - strLen

	rightPadStr := strings.Repeat(" ", rightPadLen)

	return strToJustify + rightPadStr, nil
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

// StrRightJustify - Returns a string where input parameter
// 'strToJustify' is right justified. The length of the returned
// string is determined by input parameter 'fieldlen'.
//
// Example:
//
//  If the total field length ('fieldLen') is specified as 50-characters and the
//  length of string to justify ('strToJustify') is 20-characters, then this method
//  would return a string consisting of 30-space characters plus the 'strToJustify'.
//
func (sOpsMolecule *strOpsMolecule) strRightJustify(
	strToJustify string,
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

	ePrefix += "sOpsMolecule.strRightJustify() "

	sOpsQuark := strOpsQuark{}

	if sOpsQuark.isEmptyOrWhiteSpace(strToJustify) {
		return strToJustify,
			fmt.Errorf("%v\n"+
				"Error: Input parameter 'strToJustify' is "+
				"All White Space or an EMPTY String!\n",
				ePrefix)
	}

	strLen := len(strToJustify)

	if fieldLen == strLen {
		return strToJustify, nil
	}

	if fieldLen < strLen {
		return strToJustify,
			fmt.Errorf(ePrefix+
				"\n"+
				"Error: Length of string to right justify is '%v'.\n"+
				"However, 'fieldLen' is less.\n"+
				"'fieldLen'= '%v'\n",
				strLen, fieldLen)
	}

	// fieldLen must be greater than strLen
	lefPadCnt := fieldLen - strLen

	leftPadStr := strings.Repeat(" ", lefPadCnt)

	return leftPadStr + strToJustify, nil
}
