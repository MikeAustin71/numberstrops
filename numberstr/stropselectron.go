package numberstr

import (
	"fmt"
	"io"
	"sync"
)

type strOpsElectron struct {
	lock *sync.Mutex
}

// findFirstNonSpaceChar - Returns the string index of the first non-space character in
// a string segment. The string to be searched is input parameter 'targetStr'. The string
// segment which will be searched from left to right in 'targetStr' is defined by the
// starting index ('startIndex') and the ending index ('endIndex').
//
// Searching from left to right, this method identifies the first non-space character
// (any character that is NOT a space ' ') in the target string segment and returns
// the index associated with that non-space character.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  targetStr           string
//     - The string to be searched for the first non-space character.
//
//
//  startIdx            int
//     - Since the search is forwards from left to right, this is
//       the starting index for the search.
//
//
//  endIdx              int
//     - Since the search is forwards from left to right, this is
//       the ending index for the search.
//
//
//  ePrefix             string
//     - This is an error prefix which is included in all returned
//       error messages. Usually, it contains the names of the calling
//       method or methods. Note: Be sure to leave a space at the end
//       of 'ePrefix'.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  int
//     - This method returns the index of the first non-space
//       character in the target string segment using a left
//       to right search. If the target string is an empty string
//       or consists of entirely of space characters, this method
//       returns a value of minus one (-1).
//
//
//  error
//     - If the method completes successfully this value is
//       'nil'. If an error is encountered this value will
//       contain the error message.
//
func (sOpsElectron *strOpsElectron) findFirstNonSpaceChar(
	targetStr string,
	startIndex,
	endIndex int,
	ePrefix string) (
	int,
	error) {

	if sOpsElectron.lock == nil {
		sOpsElectron.lock = new(sync.Mutex)
	}

	sOpsElectron.lock.Lock()

	defer sOpsElectron.lock.Unlock()

	if len(ePrefix) > 0 {
		ePrefix += "\n"
	}

	ePrefix += "strOpsElectron.findFirstNonSpaceChar()\n"

	sOpsQuark := strOpsQuark{}

	if sOpsQuark.isEmptyOrWhiteSpace(targetStr) {
		return -1, nil
	}
	targetStrLen := len(targetStr)

	if startIndex < 0 {
		return -1, fmt.Errorf(ePrefix+"ERROR: Invalid input parameter. 'startIndex' is LESS THAN ZERO! "+
			"startIndex='%v' ", startIndex)
	}

	if endIndex < 0 {
		return -1, fmt.Errorf(ePrefix+"ERROR: Invalid input parameter. 'endIndex' is LESS THAN ZERO! "+
			"startIndex='%v' ", startIndex)
	}

	if endIndex >= targetStrLen {
		return -1, fmt.Errorf(ePrefix+"ERROR: Invalid input parameter. 'endIndex' is greater than "+
			"target string length. INDEX OUT OF RANGE! endIndex='%v' target string length='%v' ",
			endIndex, targetStrLen)
	}

	if startIndex > endIndex {
		return -1, fmt.Errorf(ePrefix+"ERROR: Invalid input parameter. 'startIndex' is GREATER THAN 'endIndex' "+
			"startIndex='%v' endIndex='%v' ", startIndex, endIndex)
	}

	idx := startIndex

	for idx <= endIndex {

		if targetStr[idx] != ' ' {
			return idx, nil
		}

		idx++
	}

	return -1, nil

}

// GetValidString - Validates the individual characters in input parameter string,
// 'targetStr'. To identify valid characters, the characters in 'targetStr' are
// compared against input parameter 'validRunes', an array of type rune. If a character
// exists in both 'targetStr' and 'validRunes' it is deemed valid and returned in
// an output string.
//
// ------------------------------------------------------------------------
//
// Input Parameter
//
//  targetStr           string
//     - The string which will be screened for valid characters.
//
//
//  validRunes          []rune
//     - An array of type rune containing valid characters. Characters
//       which exist in both 'targetStr' and 'validRunes' will be
//       returned as a new string. Invalid characters are discarded.
//
//
//  ePrefix             string
//     - This is an error prefix which is included in all returned
//       error messages. Usually, it contains the names of the calling
//       method or methods. Note: Be sure to leave a space at the end
//       of 'ePrefix'.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  string - This string will be returned containing valid characters extracted
//           from 'targetStr'. A character is considered valid if it exists in
//           both 'targetStr' and 'validRunes'. Invalid characters are discarded.
//           This means that if no valid characters are identified, a zero length
//           string will be returned.
//
//  error  - If the method completes successfully this value is 'nil'. If an error is
//           encountered this value will contain the error message. Examples of possible
//           errors include a zero length 'targetStr' (string) or a zero length
//           'validRunes' array.
//
func (sOpsElectron *strOpsElectron) getValidString(
	targetStr string,
	validRunes []rune,
	ePrefix string) (
	string,
	error) {

	if sOpsElectron.lock == nil {
		sOpsElectron.lock = new(sync.Mutex)
	}

	sOpsElectron.lock.Lock()

	defer sOpsElectron.lock.Unlock()

	if len(ePrefix) > 0 {
		ePrefix += "\n"
	}

	ePrefix += "strOpsElectron.getValidString() "

	if len(targetStr) == 0 {
		return "",
			fmt.Errorf("%v\n"+
				"Error: Input parameter 'targetStr' is a ZERO LENGTH STRING!\n",
				ePrefix)
	}

	if len(validRunes) == 0 {
		return "",
			fmt.Errorf("%v\n"+
				"Error: Input parameter 'validRunes' is a ZERO LENGTH ARRAY!\n",
				ePrefix)
	}

	sOpsQuark := strOpsQuark{}

	actualValidRunes, err :=
		sOpsQuark.getValidRunes(
			[]rune(targetStr),
			validRunes,
			ePrefix)

	if err != nil {
		return "", err
	}

	return string(actualValidRunes), err
}

// readBytes - Implements io.Reader interface for type StrOps.
// 'readBytes' reads up to len(p) bytes into 'p'. This
// method supports buffered 'read' operations.
//
// The internal member string variable, 'StrOps.stringData'
// is written into 'p'. When the end of 'StrOps.stringData'
// is written to 'p', the method returns error = 'io.EOF'.
//
// 'StrOps.stringData' can be accessed through Getter an Setter methods,
// StrOps.GetStringData() and StrOps.SetStringData()
//
func (sOpsElectron *strOpsElectron) readBytes(
	strOpsInstance *StrOps,
	p []byte,
	ePrefix string) (
	n int,
	err error) {

	if sOpsElectron.lock == nil {
		sOpsElectron.lock = new(sync.Mutex)
	}

	sOpsElectron.lock.Lock()

	defer sOpsElectron.lock.Unlock()

	if len(ePrefix) > 0 {
		ePrefix += "\n"
	}

	ePrefix += "strOpsElectron.readBytes() "

	n = 0
	err = nil

	if strOpsInstance == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input Parameter 'strOpsInstance' is a 'nil' pointer!\n",
			ePrefix)

		return n, err
	}

	n = len(p)
	err = io.EOF

	if n == 0 {
		strOpsInstance.cntBytesRead = 0
		err = fmt.Errorf("%v\n"+
			"Error: Input byte array 'p' is zero length!\n",
			ePrefix)

		return 0, err
	}

	strData := strOpsInstance.stringData

	w := []byte(strData)

	lenW := uint64(len(w))

	cntBytesRead := strOpsInstance.cntBytesRead

	if lenW == 0 ||
		cntBytesRead >= lenW {
		strOpsInstance.cntBytesRead = 0
		n = 0
		return n, err
	}

	startReadIdx := cntBytesRead

	remainingBytesToRead := lenW - cntBytesRead

	if uint64(n) < remainingBytesToRead {
		remainingBytesToRead = startReadIdx + uint64(n)
		err = nil
	} else {
		remainingBytesToRead += startReadIdx
		err = io.EOF
	}

	n = 0

	for i := startReadIdx; i < remainingBytesToRead; i++ {
		p[n] = w[i]
		n++
	}

	cntBytesRead += uint64(n)

	if cntBytesRead >= lenW {
		strOpsInstance.cntBytesRead = 0
	} else {
		strOpsInstance.cntBytesRead = cntBytesRead
	}

	return n, err
}

// readStringFromBytes - Receives a byte array and retrieves a string. The beginning of
// the string is designated by input parameter 'startIdx'. The end of the string is determined
// when a carriage return ('\r'), vertical tab ('\v') or a new line character ('\n') is encountered.
//
// The parsed string is returned to the caller along with 'nextStartIdx', which is the byte
// array index of the beginning of the next string.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  bytes          []byte - An array of bytes from which a string will be extracted
//                          and returned.
//
//  startIdx          int - The starting index in input parameter 'bytes' where the string
//                          extraction will begin. The string extraction will cease when
//                          a carriage return ('\r'), a vertical tab ('\v') or a new line
//                          character ('\n') is encountered.
//
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  extractedStr   string - The string extracted from input parameter 'bytes' beginning
//                          at the index in 'bytes' indicated by input parameter 'startIdx'.
//
//  nextStartIdx      int - The index of the beginning of the next string in the byte array
//                          'bytes' after 'extractedString'. If no more strings exist in the
//                          the byte array, 'nextStartIdx' will be set to -1.
//
func (sOpsElectron *strOpsElectron) readStringFromBytes(
	bytes []byte,
	startIdx int) (
	extractedStr string,
	nextStartIdx int) {

	if sOpsElectron.lock == nil {
		sOpsElectron.lock = new(sync.Mutex)
	}

	sOpsElectron.lock.Lock()

	defer sOpsElectron.lock.Unlock()

	extractedStr = ""
	nextStartIdx = -1

	bLen := len(bytes)

	if bLen == 0 {
		return extractedStr, nextStartIdx
	}

	if startIdx >= bLen || startIdx < 0 {
		return extractedStr, nextStartIdx
	}

	nextStartIdx = -1

	runeAry := make([]rune, 0, bLen+5)

	for i := startIdx; i < bLen; i++ {

		if bytes[i] == '\r' ||
			bytes[i] == '\n' ||
			bytes[i] == '\v' {

			if i+1 < bLen {

				j := 0

				for j = i + 1; j < bLen; j++ {
					if bytes[j] == '\r' ||
						bytes[j] == '\v' ||
						bytes[j] == '\n' {
						continue
					} else {
						break
					}
				}

				if j >= bLen {
					nextStartIdx = -1
				} else {
					nextStartIdx = j
				}

				break

			} else {
				nextStartIdx = -1
			}

			break
		}

		runeAry = append(runeAry, rune(bytes[i]))
	}

	extractedStr = string(runeAry)

	return extractedStr, nextStartIdx
}
