package numberstr

import (
	"fmt"
	"strings"
	"sync"
)

type strOpsAtom struct {
	lock *sync.Mutex
}

// breakTextAtLineLength - Breaks string text into lines. Takes a
// string and inserts a line delimiter character (a.k.a 'rune') at
// the specified line length ('lineLength').
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  targetStr           string
//     - The string which will be parsed into text lines. If
//       'targetStr' is an empty string an error will be returned.
//        If 'targetStr' consists entirely of white space, this
//        method will return a string consisting of a new-line
//        character and an error value of 'nil'.
//
//
//  lineLength          int
//     - The maximum length of each line.
//
//
//  lineDelimiter       rune
//     - The line delimiter character which will be inserted at the
//       end of a line break.
//
//
//  Note: If the caller specifies a line length of 50, the line delimiter
//  character may be placed in the 51st character position depending upon
//  the word breaks.
//
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  string
//     - If this method completes successfully, this string
//       parameter will contain the text with line breaks delimited
//       by the input parameter 'lineDelimiter'.
//
//  error
//     - If this method completes successfully, the returned error
//       Type is set equal to 'nil'. If errors are encountered during
//       processing, the returned error Type will encapsulate an error
//       message.
//
func (sOpsAtom *strOpsAtom) breakTextAtLineLength(
	targetStr string,
	lineLength int,
	lineDelimiter rune,
	ePrefix string) (
	string,
	error) {

	if sOpsAtom.lock == nil {
		sOpsAtom.lock = new(sync.Mutex)
	}

	sOpsAtom.lock.Lock()

	defer sOpsAtom.lock.Unlock()

	if len(ePrefix) > 0 {
		ePrefix += "\n"
	}

	ePrefix += "strOpsAtom.breakTextAtLineLength() "

	targetLen := len(targetStr)

	if targetLen == 0 {
		return "",
			fmt.Errorf("%v\n"+
				"Error: Input parameter 'targetStr' is a ZERO LENGTH STRING!\n",
				ePrefix)
	}

	if lineLength < 5 {
		return "",
			fmt.Errorf(ePrefix+"Error: Input parameter 'lineLength' is LESS THAN 5-CHARACTERS! "+
				"lineLength='%v' ", lineLength)
	}

	if lineDelimiter == 0 {
		return "",
			fmt.Errorf("%v\n"+
				"Error: Input parameter 'lineDelimiter' is ZERO VALUE!\n",
				ePrefix)
	}

	sOpsQuark := strOpsQuark{}

	if sOpsQuark.isEmptyOrWhiteSpace(targetStr) {
		return "\n", nil
	}

	var err error

	var b strings.Builder
	b.Grow(((targetLen / lineLength) * targetLen) + 50)

	var begIdx, endWrdIdx, actualLastIdx, beginWrdIdx int
	var isAllOneWord, isAllSpaces bool
	sOpsElectron := strOpsElectron{}

	for begIdx < targetLen && begIdx > -1 {

		// skip spaces at the beginning of the line
		begIdx, err = sOpsElectron.findFirstNonSpaceChar(
			targetStr,
			begIdx,
			targetLen-1,
			ePrefix)

		if err != nil {
			return "",
				fmt.Errorf(ePrefix+
					"Error returned by sops.FindFirstNonSpaceChar(targetStr, begIdx, actualLastIdx). "+
					"targetStr='%v' begIdx='%v' actualLastIdx='%v'   Error='%v' ",
					targetStr, begIdx, actualLastIdx, err.Error())
		}

		if begIdx == -1 {

			if b.Len() == 0 {
				b.WriteRune(lineDelimiter)
			}

			break // Exit loop
		}

		if begIdx == targetLen-1 {
			b.WriteByte(targetStr[begIdx])
			b.WriteRune(lineDelimiter)
			break
		}

		actualLastIdx = begIdx + lineLength - 1

		if actualLastIdx >= targetLen {
			actualLastIdx = targetLen - 1
		}

		// Find the last complete word in this string segment
		beginWrdIdx, endWrdIdx, isAllOneWord, isAllSpaces, err =
			sOpsQuark.findLastWord(
				targetStr,
				begIdx,
				actualLastIdx,
				ePrefix)

		if err != nil {
			return "",
				fmt.Errorf(ePrefix+
					"Error returned by sops.FindLastWord(targetStr,begIdx, actualLastIdx). "+
					"targetStr='%v' begIdx='%v' actualLastIdx='%v'   Error='%v' ",
					targetStr, begIdx, actualLastIdx, err.Error())
		}

		if isAllSpaces {
			// This string segment is all spaces
			// write a line delimiter and continue
			begIdx = actualLastIdx + 1

		} else if isAllOneWord {
			// This string segment is all one word
			// and contains NO spaces.

			if actualLastIdx+1 >= targetLen {
				// If this is end of the main string,
				// just write the remaining segment and
				// exit.
				//
				b.WriteString(targetStr[begIdx:])
				b.WriteRune(lineDelimiter)
				break

			} else if actualLastIdx-begIdx+1 <= lineLength {
				// If this string segment is less than the specified
				// line length, just write the entire line segment.
				// Be careful, we may be at the end of the main
				// string.

				if actualLastIdx+1 >= targetLen {
					// This is the end of the main string,
					// just exit.
					b.WriteString(targetStr[begIdx:])
					b.WriteRune(lineDelimiter)
					break

				} else {

					b.WriteString(targetStr[begIdx : actualLastIdx+1])
					begIdx = actualLastIdx + 1
				}

			} else {
				// Out of options. Nothing left to do but hyphenate
				// the word.
				b.WriteString(targetStr[begIdx : actualLastIdx-1])
				b.WriteRune('-')
				begIdx = actualLastIdx

			}

		} else {
			// The segment is NOT All spaces nor is it all one word.

			if endWrdIdx+1 >= targetLen {
				// Are we at the end of targetStr
				b.WriteString(targetStr[begIdx:])
				b.WriteRune(lineDelimiter)
				break

			} else if targetStr[endWrdIdx+1] != ' ' {
				// This word crosses a line break boundary. Try not to split the word.

				// Find  the end of the last word.
				idx, err :=
					sOpsQuark.findLastNonSpaceChar(
						targetStr,
						begIdx,
						beginWrdIdx-1,
						ePrefix)

				if err != nil {
					return "",
						fmt.Errorf(ePrefix+
							"Error returned by sOpsQuark.findLastNonSpaceChar(targetStr,begIdx, beginWrdIdx-1).\n"+
							"targetStr='%v'\n"+
							"begIdx='%v'\nactualLastIdx='%v'\n"+
							"Error='%v' ",
							targetStr, begIdx, actualLastIdx, err.Error())
				}

				if idx == -1 {
					begIdx = beginWrdIdx
					// Do not write end of line delimiter
					// Set bigIdx to beginning of word and
					// loop again
					continue

				} else {
					// Success we found the end of the last word.
					b.WriteString(targetStr[begIdx : idx+1])
					begIdx = idx + 1
				}

			} else {
				// The word does not cross a line break boundary.
				// The next character after the last word is a
				// space.

				b.WriteString(targetStr[begIdx : endWrdIdx+1])
				begIdx = endWrdIdx + 1
			}
		}

		b.WriteRune(lineDelimiter)

	}

	return b.String(), nil
}
