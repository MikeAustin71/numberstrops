package numberstr

import (
	"fmt"
	"sync"
)

type strOpsQuark struct {
	lock *sync.Mutex
}

// findLastNonSpaceChar - Returns the string index of the last non-space character in a
// string segment.  The string to be searched is input parameter, 'targetStr'. The
// string segment is further defined by input parameters 'startIdx' and  'endIdx'. These
// indexes define a segment within 'targetStr' which will be searched to identify the last
// non-space character.
//
// The search is a backwards search, from right to left, conducted within the defined
// 'targetStr' segment. The search therefore starts at 'endIdx' and proceeds towards
// 'startIdx' until the last non-space character in the string segment is identified.
//
//
// If the last non-space character is found, that string index is returned. If the string
// segment consists entirely of space characters, the return value is -1.
//
// if 'targetStr' is a zero length string, an error will be triggered. Likewise, if 'startIdx'
// of 'endIdx' are invalid, an error will be returned.
//
func (sOpsQuark *strOpsQuark) findLastNonSpaceChar(
	targetStr string,
	startIdx,
	endIdx int,
	ePrefix string) (
	int,
	error) {

	if sOpsQuark.lock == nil {
		sOpsQuark.lock = new(sync.Mutex)
	}

	sOpsQuark.lock.Lock()

	defer sOpsQuark.lock.Unlock()

	ePrefix += "strOpsQuark.findLastNonSpaceChar() "

	targetStrLen := len(targetStr)

	if targetStrLen == 0 {
		return -1,
			fmt.Errorf("%v\n"+
				"ERROR: Invalid input parameter. 'targetStr' is a ZERO LENGTH STRING!\n",
				ePrefix)
	}

	if startIdx < 0 {
		return -1,
			fmt.Errorf(ePrefix+"\n"+
				"ERROR: Invalid input parameter. 'startIdx' is LESS THAN ZERO!\n"+
				"startIdx='%v'\n", startIdx)
	}

	if endIdx < 0 {
		return -1,
			fmt.Errorf(ePrefix+"\n"+
				"ERROR: Invalid input parameter. 'endIdx' is LESS THAN ZERO!\n"+
				"startIdx='%v' ", startIdx)
	}

	if endIdx >= targetStrLen {
		return -1,
			fmt.Errorf(ePrefix+"\n"+
				"ERROR: Invalid input parameter. 'endIdx' is greater than target string length.\n"+
				"INDEX OUT OF RANGE!\n"+
				"endIdx='%v' target string length='%v'\n",
				endIdx, targetStrLen)
	}

	if startIdx > endIdx {
		return -1, fmt.Errorf(ePrefix+"\n"+
			"ERROR: Invalid input parameter.\n"+
			"'startIdx' is GREATER THAN 'endIdx'\n"+
			"startIdx='%v' endIdx='%v'\n",
			startIdx, endIdx)
	}

	floor := startIdx - 1

	for endIdx > floor {

		if targetStr[endIdx] != ' ' {
			return endIdx, nil
		}

		endIdx--
	}

	return -1, nil
}

// getValidBytes - Receives an array of 'targetBytes' which will be examined to determine
// the validity of individual bytes or characters. Each character (byte) in input array
// 'targetBytes' will be compared to input parameter 'validBytes', another array of bytes.
// If a character in 'targetBytes' also exists in 'validBytes' it will be considered valid
// and included in the returned array of bytes.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  targetBytes  [] byte  - An array of characters (bytes) which will be examined
//                          for valid characters. The list of valid characters is
//                          found in input parameter 'validBytes'. Valid characters
//                          in targetBytes will be returned by this method as an
//                          array of bytes. Invalid characters will be discarded.
//
//
//  validBytes  [] byte   - An array of bytes containing valid characters. If a character
//                          (byte) in 'targetBytes' is also present in 'validBytes' it will
//                          be classified as 'valid' and included in the returned array of
//                          bytes. Invalid characters will be discarded.
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
//  [] byte
//     - An array of bytes which contains bytes that are
//       present in both 'targetBytes' and 'validBytes'.
//       Note: If all characters in 'targetBytes' are classified
//       as 'invalid', the returned array of bytes will be a zero
//       length array.
//
//
//  error
//     - If the method completes successfully this value is
//       'nil'. If an error is encountered this value will
//       contain the error message.
//
func (sOpsQuark *strOpsQuark) getValidBytes(
	targetBytes []byte,
	validBytes []byte,
	ePrefix string) (
	[]byte,
	error) {

	if sOpsQuark.lock == nil {
		sOpsQuark.lock = new(sync.Mutex)
	}

	sOpsQuark.lock.Lock()

	defer sOpsQuark.lock.Unlock()

	ePrefix += "strOpsQuark.getValidBytes() "

	lenTargetBytes := len(targetBytes)

	output := make([]byte, 0, lenTargetBytes+10)

	if lenTargetBytes == 0 {
		return output,
			fmt.Errorf("%v\n"+
				"Error: Input parameter 'targetBytes' is a ZERO LENGTH ARRAY!\n",
				ePrefix)
	}

	lenValidBytes := len(validBytes)

	if lenValidBytes == 0 {
		return output,
			fmt.Errorf("%v\n"+
				"Error: Input parameter 'validBytes' is a ZERO LENGTH ARRAY!\n",
				ePrefix)
	}

	for i := 0; i < lenTargetBytes; i++ {

		for j := 0; j < lenValidBytes; j++ {
			if targetBytes[i] == validBytes[j] {
				output = append(output, targetBytes[i])
				break
			}
		}

	}

	return output, nil
}

// getValidRunes - Receives an array of 'targetRunes' which will be examined to determine
// the validity of individual runes or characters. Each character (rune) in input array
// 'targetRunes' will be compared to input parameter 'validRunes', another array of runes.
// If a character in 'targetRunes' also exists in 'validRunes', that character will be considered
// valid and included in the returned array of runes.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  targetRunes    [] rune  - An array of characters (runes) which will be examined
//                            for valid characters. The list of valid characters is
//                            found in input parameter 'validRunes'. Valid characters
//                            in targetRunes will be returned by this method as an
//                            array of runes. Invalid characters will be discarded.
//
//  validRunes    [] rune  - An array of runes containing valid characters. If a character
//                           (rune) in targetRunes is also present in 'validRunes' it will
//                           be classified as 'valid' and included in the returned array of
//                           runes. Invalid characters will be discarded.
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
//  [] rune  - An array of runes which contains runes that are present in 'targetRunes' and
//             'validRunes'. Note: If all characters in 'targetRunes' are classified as
//             'invalid', the returned array of runes will be a zero length array.
//
//  error    - If the method completes successfully this value is 'nil'. If an error is
//             encountered this value will contain the error message. Examples of possible
//             errors include a zero length 'targetRunes array or 'validRunes' array.
//
func (sOpsQuark *strOpsQuark) getValidRunes(
	targetRunes []rune,
	validRunes []rune,
	ePrefix string) (
	[]rune,
	error) {

	if sOpsQuark.lock == nil {
		sOpsQuark.lock = new(sync.Mutex)
	}

	sOpsQuark.lock.Lock()

	defer sOpsQuark.lock.Unlock()

	ePrefix += "strOpsQuark.getValidRunes() "

	lenTargetRunes := len(targetRunes)

	output := make([]rune, 0, lenTargetRunes+10)

	if lenTargetRunes == 0 {
		return output,
			fmt.Errorf("%v\n"+
				"Error: Input parameter 'targetRunes' is a ZERO LENGTH ARRAY!\n",
				ePrefix)
	}

	lenValidRunes := len(validRunes)

	if lenValidRunes == 0 {
		return output,
			fmt.Errorf("%v\n"+
				"Error: Input parameter 'validRunes' is a ZERO LENGTH ARRAY!\n",
				ePrefix)
	}

	for i := 0; i < lenTargetRunes; i++ {

		for j := 0; j < lenValidRunes; j++ {
			if targetRunes[i] == validRunes[j] {
				output = append(output, targetRunes[i])
				break
			}
		}
	}

	return output, nil
}

// isEmptyOrWhiteSpace - Analyzes the incoming string and returns
// 'true' if the strings is empty or consists of all white space.
//
func (sOpsQuark *strOpsQuark) isEmptyOrWhiteSpace(
	targetStr string) bool {

	if sOpsQuark.lock == nil {
		sOpsQuark.lock = new(sync.Mutex)
	}

	sOpsQuark.lock.Lock()

	defer sOpsQuark.lock.Unlock()

	targetLen := len(targetStr)

	for i := 0; i < targetLen; i++ {
		if targetStr[i] != ' ' {
			return false
		}
	}

	return true
}
