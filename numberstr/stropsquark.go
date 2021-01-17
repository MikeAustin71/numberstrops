package numberstr

import (
	"fmt"
	"strings"
	"sync"
)

type strOpsQuark struct {
	lock *sync.Mutex
}

// ConvertNonPrintableChars - Receives a string containing non-printable characters
// and converts them to 'printable' characters returned in a string.
//
// If the input parameter 'convertSpace' is set to 'true' then all spaces are converted
// to "[SPACE]" in the returned string.
//
// Reference:
//    https://www.juniper.net/documentation/en_US/idp5.1/topics/reference/general/intrusion-detection-prevention-custom-attack-object-extended-ascii.html
//
func (sOpsQuark *strOpsQuark) convertNonPrintableChars(
	nonPrintableChars []rune,
	convertSpace bool,
	ePrefix string) (
	printableChars string) {

	if sOpsQuark.lock == nil {
		sOpsQuark.lock = new(sync.Mutex)
	}

	sOpsQuark.lock.Lock()

	defer sOpsQuark.lock.Unlock()

	if len(ePrefix) > 0 {
		ePrefix += "\n"
	}

	ePrefix += "strOpsQuark.convertNonPrintableChars() "

	lenNonPrintableChars := len(nonPrintableChars)

	if lenNonPrintableChars == 0 {
		return "[EMPTY]"
	}

	var b strings.Builder
	b.Grow(lenNonPrintableChars * 5)

	for i := 0; i < lenNonPrintableChars; i++ {
		cRune := nonPrintableChars[i]

		switch cRune {
		case 0:
			b.WriteString("[NULL]") // 0x00 NULL
		case 1:
			b.WriteString("[SOH]") // 0x01 State of Heading
		case 2:
			b.WriteString("[STX]") // 0x02 State of Text
		case 3:
			b.WriteString("[ETX]") // 0x03 End of Text
		case 4:
			b.WriteString("[EOT]") // 0X04 End of Transmission
		case 5:
			b.WriteString("[ENQ]") // 0x05 Enquiry
		case 6:
			b.WriteString("[ACK]") // 0x06 Acknowledge
		case '\a':
			b.WriteString("\\a") // U+0007 alert or bell
		case '\b':
			b.WriteString("\\b") // U+0008 backspace
		case '\t':
			b.WriteString("\\t") // U+0009 horizontal tab
		case '\n':
			b.WriteString("\\n") // U+000A line feed or newline
		case '\v':
			b.WriteString("\\v") // U+000B vertical tab
		case '\f':
			b.WriteString("\\f") // U+000C form feed
		case '\r':
			b.WriteString("\\r") // U+000D carriage return
		case 14:
			b.WriteString("[SO]") // U+000E Shift Out
		case 15:
			b.WriteString("[SI]") // U+000F Shift In
		case '\\':
			b.WriteString("\\") // U+005c backslash
		case ' ':
			// 0X20 Space character
			if convertSpace {
				b.WriteString("[SPACE]")
			} else {
				b.WriteRune(' ')
			}

		default:
			b.WriteRune(cRune)

		}

	}

	return b.String()

}

// DoesLastCharExist - returns true if the last character (rune) of
// input string 'testStr' is equal to input parameter 'lastChar' which
// is of type 'rune'.
func (sOpsQuark *strOpsQuark) doesLastCharExist(
	testStr string,
	lastChar rune) bool {

	if sOpsQuark.lock == nil {
		sOpsQuark.lock = new(sync.Mutex)
	}

	sOpsQuark.lock.Lock()

	defer sOpsQuark.lock.Unlock()
	testStrLen := len(testStr)

	if testStrLen == 0 {
		return false
	}

	strLastChar := rune(testStr[testStrLen-1])

	if strLastChar == lastChar {
		return true
	}

	return false

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

	if len(ePrefix) > 0 {
		ePrefix += "\n"
	}

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

// FindLastSpace - Returns a string index indicating the last space character (' ') in
// a string segment. The string segment is defined by input parameters, 'startIdx' and
// 'endIdx'.
//
// The string segment search proceeds backwards, from right to left. The search therefore
// starts at 'endIdx' and proceeds towards 'startIdx' until the last space character in
// the string segment is identified.
//
// If a valid index for the last space character is found in the string segment, that
// index value is returned. If a space character is NOT found in the specified string
// segment, a value of -1 is returned.
//
// if 'targetStr' is a zero length string, an error will be triggered. Likewise, if 'startIdx'
// of 'endIdx' are invalid, an error will be returned.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  targetStr           string
//     - The string segment to be searched for the last space
//       character resides in this string.
//
//  startIdx            int
//     - Defines the actual ending index for the string segment to
//       be searched. The search is a 'backwards' search and
//       proceeds from right to left. Therefore, the starting point
//       for the string segment is input parameter 'endIdx' while
//       the ending point for the string segment is this 'startIdx'.
//
//
//  endIdx              int
//     - Defines the actual beginning index of the string segment
//       to be searched. The search is a 'backwards' search and
//       proceeds from right to left. Therefore, the starting point
//       for the string segment is defined by this 'endIdx'
//       parameter while the ending point for the string segment is
//       marked by the input parameter, 'startIdx'.
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
//     - If this method completes successfully, this returned
//       integer value will constitute the string index of the last
//       space character in the string segment marked by input
//       parameters 'startIdx' and 'endIdx'.
//
//       If a space character is NOT found in the specified string
//       segment, a value of minus one (-1) is returned.
//
//
//  error
//     - If the method completes successfully and no errors are
//       encountered this return value is set to 'nil'. Otherwise,
//       if errors are encountered this return value will contain
//       an appropriate error message.
//
//       If an error message is returned, the input parameter
//       'ePrefix' will be inserted or prefixed at the beginning
//       of the error message.
//
func (sOpsQuark *strOpsQuark) findLastSpace(
	targetStr string,
	startIdx int,
	endIdx int,
	ePrefix string) (
	int,
	error) {

	if sOpsQuark.lock == nil {
		sOpsQuark.lock = new(sync.Mutex)
	}

	sOpsQuark.lock.Lock()

	defer sOpsQuark.lock.Unlock()

	if len(ePrefix) > 0 {
		ePrefix += "\n"
	}

	ePrefix += "strOpsQuark.findLastSpace() "

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
				"startIdx='%v'\n", startIdx)
	}

	if endIdx >= targetStrLen {
		return -1,
			fmt.Errorf(ePrefix+"\n"+
				"ERROR: Invalid input parameter. 'endIdx' is greater than target string length.\n"+
				"INDEX OUT OF RANGE!\n"+
				"endIdx='%v'\n"+
				"target string length='%v'\n",
				endIdx, targetStrLen)
	}

	if startIdx > endIdx {
		return -1,
			fmt.Errorf(ePrefix+"\n"+
				"ERROR: Invalid input parameter.\n"+
				"'startIdx' is GREATER THAN 'endIdx'\n"+
				"startIdx='%v' endIdx='%v'\n", startIdx, endIdx)
	}

	for endIdx >= startIdx {

		if targetStr[endIdx] == ' ' {
			return endIdx, nil
		}

		endIdx--
	}

	return -1, nil
}

// findLastWord - Returns the beginning and ending indexes of
// the last word in a target string segment. A 'word' is defined here
// as a contiguous set of non-space characters delimited by spaces or
// the beginning and ending indexes of the target string segment. Note,
// for purposes of this method, a 'word' my consist of a single non-space
// character such as an article 'a' or a punctuation mark '.'
//
// ------------------------------------------------------------------------
//
// Examples
//
//
//   Example (1)
//     In the text string segment:
//
//     "The cow jumped over the moon."
//
//     The last word would be defined as "moon."
//
//     Example (2)
//       In the text string segment:
//
//       "  somewhere over the rainbow  "
//
//       The last word would be defined as "rainbow"
//
// ------------------------------------------------------------------------
//
// The string to be searched is contained in input parameter, 'targetStr'.
// The string segment within 'targetStr' is defined by input parameters
// 'startIndex' and 'endIndex'.
//
// If the entire string segment is classified as a 'word', meaning that
// there are no space characters in the string segment, the returned
// values for 'beginWrdIdx' and 'endWrdIdx' will be equal to the input
// parameters 'startIndex' and 'endIndex'.
//
// If the string segment is consists entirely of space characters, the
// returned 'beginWrdIdx' and 'endWrdIdx' will be set equal to -1 and
// the returned value, 'isAllSpaces' will be set to 'true'.
//
// If 'targetStr' is an empty string, an error will be returned.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//
//  targetStr      string
//     - The string containing the string segment which
//       will be searched to identify the last word
//       in the string segment.
//
//
//  startIndex     int
//     - The index marking the beginning of the string
//       segment in 'targetStr'.
//
//
//  endIndex       int
//     - The index marking the end of the string segment
//       in 'targetStr'.
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
//  beginWrdIdx    int  - The index marking the beginning of the last word
//                        in the string segment identified by input parameters
//                        'startIndex' and 'endIndex'. If the string segment
//                        consists of all spaces or is empty, this value is
//                        set to -1.
//
//  endWrdIdx      int  - The index marking the end of the last word in the
//                        string segment identified by input parameters 'startIndex'
//                        and 'endIndex'. If the string segment consists of all
//                        spaces or is empty, this value is set to -1.
//
//  isAllOneWord   bool - If the string segment identified by input parameters
//                        'startIndex' and 'endIndex' consists entirely of non-space
//                        characters (characters other than ' '), this value is set
//                        to 'true'.
//
//  isAllSpaces    bool - If the string segment identified by input parameters
//                        'startIndex' and 'endIndex' consists entirely of space
//                        characters (character = ' '), this value is set to 'true'.
//
//  err           error - If targetStr is empty or if startIndex or endIndex is invalid,
//                        an error is returned. If the method completes successfully,
//                        err = nil.
//
func (sOpsQuark *strOpsQuark) findLastWord(
	targetStr string,
	startIndex int,
	endIndex int,
	ePrefix string) (
	beginWrdIdx int,
	endWrdIdx int,
	isAllOneWord bool,
	isAllSpaces bool,
	err error) {

	if sOpsQuark.lock == nil {
		sOpsQuark.lock = new(sync.Mutex)
	}

	sOpsQuark.lock.Lock()

	defer sOpsQuark.lock.Unlock()

	if len(ePrefix) > 0 {
		ePrefix += "\n"
	}

	ePrefix += "strOpsQuark.findLastNonSpaceChar() "

	beginWrdIdx = -1
	endWrdIdx = -1
	isAllOneWord = false
	isAllSpaces = false

	targetStrLen := len(targetStr)

	if targetStrLen == 0 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'targetStr' is an EMPTY STRING!\n",
			ePrefix)

		return beginWrdIdx,
			endWrdIdx,
			isAllOneWord,
			isAllSpaces,
			err
	}

	if startIndex < 0 {

		err = fmt.Errorf(ePrefix+"\n"+
			"ERROR: Invalid input parameter.\n"+
			"'startIndex' is LESS THAN ZERO!\n"+
			"startIndex='%v'\n", startIndex)

		return beginWrdIdx,
			endWrdIdx,
			isAllOneWord,
			isAllSpaces,
			err
	}

	if endIndex < 0 {
		err = fmt.Errorf(ePrefix+"\n"+
			"ERROR: Invalid input parameter.\n"+
			"'endIndex' is LESS THAN ZERO!\n"+
			"startIndex='%v'\n",
			startIndex)

		return beginWrdIdx,
			endWrdIdx,
			isAllOneWord,
			isAllSpaces,
			err
	}

	if endIndex >= targetStrLen {

		err = fmt.Errorf(ePrefix+"\n"+
			"ERROR: Invalid input parameter. 'endIndex' is greater than\n"+
			"target string length. INDEX OUT OF RANGE!\n"+
			"endIndex='%v'\n"+
			"target string length='%v'\n",
			endIndex,
			targetStrLen)

		return beginWrdIdx,
			endWrdIdx,
			isAllOneWord,
			isAllSpaces,
			err
	}

	if startIndex > endIndex {
		err = fmt.Errorf(ePrefix+"\n"+
			"ERROR: Invalid input parameter.\n"+
			"'startIndex' is GREATER THAN 'endIndex'.\n"+
			"startIndex='%v' endIndex='%v'\n",
			startIndex, endIndex)

		return beginWrdIdx,
			endWrdIdx,
			isAllOneWord,
			isAllSpaces,
			err
	}

	beginWrdIdx = startIndex
	endWrdIdx = endIndex

	idx := endIndex

	var endingIdxFound bool

	isAllSpaces = true
	isAllOneWord = true

	if startIndex == endIndex {

		beginWrdIdx = startIndex
		endWrdIdx = startIndex

		if targetStr[startIndex] == ' ' {
			isAllSpaces = true
			isAllOneWord = false
		} else {
			isAllSpaces = false
			isAllOneWord = true
		}

		err = nil

		return beginWrdIdx,
			endWrdIdx,
			isAllOneWord,
			isAllSpaces,
			err
	}

	for idx >= startIndex {

		if targetStr[idx] != ' ' {
			isAllSpaces = false
		} else {
			isAllOneWord = false
		}

		if !endingIdxFound &&
			targetStr[idx] != ' ' {

			endWrdIdx = idx
			endingIdxFound = true
			idx--
			continue
		}

		if endingIdxFound &&
			targetStr[idx] == ' ' {

			beginWrdIdx = idx + 1
			break
		}

		idx--
	}

	if isAllSpaces {
		isAllOneWord = false
		beginWrdIdx = -1
		endWrdIdx = -1
		err = nil
		return beginWrdIdx,
			endWrdIdx,
			isAllOneWord,
			isAllSpaces,
			err
	}

	if isAllOneWord {
		beginWrdIdx = startIndex
		endWrdIdx = endIndex
		isAllSpaces = false
		err = nil
		return beginWrdIdx,
			endWrdIdx,
			isAllOneWord,
			isAllSpaces,
			err
	}

	err = nil

	return beginWrdIdx,
		endWrdIdx,
		isAllOneWord,
		isAllSpaces,
		err
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
