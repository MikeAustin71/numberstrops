package numberstr

import (
	"fmt"
	"sync"
)

type numStrBasicNanobot struct {
	lock *sync.Mutex
}

// convertNumStrToFloatingPointNumStr - Receives a number string
// and converts to a pure floating point number string.
//
// The returned floating point number string will contain a leading
// minus sign character ('-') if the value is negative. Positive
// values will never be prefixed with a leading plus ('+') sign.
// Otherwise, the returned floating point number string will
// consist solely of numeric digits and a decimal separator.
//
// If input parameter 'rawNumStr' is an integer value with no
// fractional digits to the right of the decimal separator, the
// returned floating point number string will reflect that integer
// value and consist solely of integer digits.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  rawNumStr           string
//     - A string of numeric digits which may or may not contain
//       fractional numeric digits to the right of a decimal
//       separator character. This number string may include a leading
//       numeric sign character (plus '+' or minus '-').
//
//
//  decimalSeparators    rune
//     - A unicode character inserted into a number string to
//       separate integer and fractional digits. In the United
//       States, the decimal separator is the period character
//       ('.') and it is known as the decimal point.
//
//       If input parameter 'rawNumStr' contains fractional digits,
//       they should be separated from integer digits by this
//       decimal separator character.
//
//
//  ePrefix             ErrPrefixDto
//     - This object encapsulates an error prefix string which is
//       included in all returned error messages. Usually, it
//       contains the names of the calling method or methods.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  floatingPointNumStr string
//     - A floating point number string parsed from input parameter
//       'rawNumStr'. If fractional digits exist in 'rawNumStr'
//       they will be included in this string separated by the
//       decimal separator character.
//
//        'floatingPointNumStr' is pure number string and will
//        never contain thousands separators.
//
//
//  err                 error
//     - If this method completes successfully, the returned error
//       Type is set equal to 'nil'.
//
//       If errors are encountered during processing, the returned
//       error Type will encapsulate an error message. This
//       returned error message will incorporate the method chain
//       and text passed by input parameter, 'ePrefix'. The
//       'ePrefix' text will be attached to the beginning of the
//       error message.
//
func (nStrBasicNanobot *numStrBasicNanobot) convertNumStrToFloatingPointNumStr(
	rawNumStr string,
	decimalSeparator rune,
	ePrefix *ErrPrefixDto) (
	floatingPointNumStr string,
	err error) {

	if nStrBasicNanobot.lock == nil {
		nStrBasicNanobot.lock = new(sync.Mutex)
	}

	nStrBasicNanobot.lock.Lock()

	defer nStrBasicNanobot.lock.Unlock()

	ePrefix.SetEPref("numStrBasicNanobot.convertNumStrToFloatingPointNumStr() ")

	var signChar rune
	var intNumRunes, fracNumRunes []rune
	var lenIntNumRunes, lenFracNumRunes int

	signChar,
		intNumRunes,
		lenIntNumRunes,
		fracNumRunes,
		lenFracNumRunes,
		err = numStrBasicAtom{}.ptr().
		parseIntFracRunesFromNumStr(
			rawNumStr,
			decimalSeparator,
			ePrefix.XCtx("rawNumStr"))

	if err != nil {
		return floatingPointNumStr, err
	}

	ePrefix.SetCtxEmpty()

	if lenIntNumRunes == 0 &&
		lenFracNumRunes == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'rawNumStr' is invalid!\n"+
			"'rawNumStr' contained ZERO integer or fractional digits.\n"+
			"rawNumStr='%v'\n",
			ePrefix.String(),
			rawNumStr)

		return floatingPointNumStr, err
	}

	if signChar == '-' {
		floatingPointNumStr = "-"
	}

	if lenIntNumRunes == 0 {

		floatingPointNumStr += "0"

	} else {

		floatingPointNumStr += string(intNumRunes)

	}

	if lenFracNumRunes > 0 {
		floatingPointNumStr += string(decimalSeparator)
		floatingPointNumStr += string(fracNumRunes)
	}

	return floatingPointNumStr, err
}

// ptr - Returns a pointer to a new instance of
// numStrBasicNanobot.
//
func (nStrBasicNanobot numStrBasicNanobot) ptr() *numStrBasicNanobot {

	if nStrBasicNanobot.lock == nil {
		nStrBasicNanobot.lock = new(sync.Mutex)
	}

	nStrBasicNanobot.lock.Lock()

	defer nStrBasicNanobot.lock.Unlock()

	newNumStrBasicNanobot := new(numStrBasicNanobot)

	newNumStrBasicNanobot.lock = new(sync.Mutex)

	return newNumStrBasicNanobot
}
