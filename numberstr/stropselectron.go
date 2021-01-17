package numberstr

import (
	"fmt"
	"io"
	"sync"
)

type strOpsElectron struct {
	lock *sync.Mutex
}

// FindFirstNonSpaceChar - Returns the string index of the first non-space character in
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
