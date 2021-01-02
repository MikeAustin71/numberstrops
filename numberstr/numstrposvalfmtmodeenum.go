package numberstr

import (
	"fmt"
	"strings"
	"sync"
)

var mNumStrPosValFmtModeCodeToString = map[NumStrPosValFmtMode]string{
	NumStrPosValFmtMode(0): "None",
	NumStrPosValFmtMode(1): "NoLeadingPlusSign",
	NumStrPosValFmtMode(2): "LeadingPlusSign",
}

var mNumStrPosValFmtModeStringToCode = map[string]NumStrPosValFmtMode{
	"None":              NumStrPosValFmtMode(0),
	"NoLeadingPlusSign": NumStrPosValFmtMode(1),
	"NoPlusSign":        NumStrPosValFmtMode(1),
	"LeadingPlusSign":   NumStrPosValFmtMode(2),
	"LeadingPlus":       NumStrPosValFmtMode(2),
	"PlusSign":          NumStrPosValFmtMode(2),
}

var mNumStrPosValFmtModeLwrCaseStringToCode = map[string]NumStrPosValFmtMode{
	"none":              NumStrPosValFmtMode(0),
	"noleadingplussign": NumStrPosValFmtMode(1),
	"noplussign":        NumStrPosValFmtMode(1),
	"leadingplussign":   NumStrPosValFmtMode(2),
	"leadingplus":       NumStrPosValFmtMode(2),
	"plussign":          NumStrPosValFmtMode(2),
}

// NumStrPosValFmtMode - An enumeration of Positive Value Format Modes
// used in formatting positive numeric values in number strings for display
// purposes.
//
// Since the Go Programming Language does not directly support
// enumerations, the 'NumStrPosValFmtMode' type has been adapted to
// function in a manner similar to classic enumerations. 'NumStrPosValFmtMode'
// is declared as a type 'int'. The method names effectively represents
// an enumeration of format modes for positive numeric values. These
// methods are listed as follows:
//
//
// None                (0)
//  - Signals that the Positive Value Format Mode (NumStrPosValFmtMode)
//    Type is not initialized. This is an error condition.
//
// NoLeadingPlusSign   (1)
//  - No Leading Plus Sign signals that positive numeric values will be
//    formatted in text without a leading plus sign. This is typically
//    the default for positive number formatting.
//         Example: 123
//
// LeadingPlusSign     (2)
//  - Leading Plus signals that the positive value will have a plus sign
//    ('+') in the first character position of the number string.
//         Example: +123
//
//
// For easy access to these enumeration values, use the global variable
// 'NStrPosValFmtMode'. Example: NStrPosValFmtMode.NoLeadingPlusSign()
//
// Otherwise you will need to use the formal syntax.
//     Example: NumStrPosValFmtMode(0).NoLeadingPlusSign()
//
//
// Depending on your editor, intellisense (a.k.a. intelligent code
// completion) may not list the NumStrPosValFmtMode methods in
// alphabetical order. Be advised that all 'NumStrPosValFmtMode'
// methods beginning with 'X', as well as the method 'String()',
// are utility methods and not part of the enumeration.
//
type NumStrPosValFmtMode int

var lockNumStrPosValFmtMode sync.Mutex

// None - Signals that the NumStrPosValFmtMode Type is uninitialized.
// This is an error condition.
//
// This method is part of the standard enumeration.
//
func (posFmtMode NumStrPosValFmtMode) None() NumStrPosValFmtMode {

	lockNumStrPosValFmtMode.Lock()

	defer lockNumStrPosValFmtMode.Unlock()

	return NumStrPosValFmtMode(0)
}

// NoLeadingPlusSign - Signals that positive numeric values will NOT
// have a leading plus sign. This is the default presentation for
// positive numeric values.
//
//      Example: '1234'
//
// This method is part of the standard enumeration.
//
func (posFmtMode NumStrPosValFmtMode) NoLeadingPlusSign() NumStrPosValFmtMode {

	lockNumStrPosValFmtMode.Lock()

	defer lockNumStrPosValFmtMode.Unlock()

	return NumStrPosValFmtMode(1)
}

// LeadingPlusSign - Signals that positive numeric values should
// be formatted with with a leading plus ('+') sign.
//
//      Example: '+1234'
//
// This method is part of the standard enumeration.
//
func (posFmtMode NumStrPosValFmtMode) LeadingPlusSign() NumStrPosValFmtMode {

	lockNumStrPosValFmtMode.Lock()

	defer lockNumStrPosValFmtMode.Unlock()

	return NumStrPosValFmtMode(2)
}

// String - Returns a string with the name of the enumeration associated
// with this current instance of 'NumStrPosValFmtMode'.
//
// This is a standard utility method and is not part of the valid enumerations
// for this type.
//
// ------------------------------------------------------------------------
//
// Usage
//
// t:= NumStrPosValFmtMode(0).LeadingPlusSign()
// str := t.String()
//     str is now equal to 'LeadingPlusSign'
//
func (posFmtMode NumStrPosValFmtMode) String() string {

	lockNumStrPosValFmtMode.Lock()

	defer lockNumStrPosValFmtMode.Unlock()

	result, ok := mNumStrPosValFmtModeCodeToString[posFmtMode]

	if !ok {
		return "Error: Number String Positive Value Format Mode UNKNOWN!"
	}

	return result
}

// XIsValid - Returns a boolean value signaling whether the current
// Number String Positive Value Format Mode (NumStrPosValFmtMode)
// value is valid.
//
// This is a standard utility method and is not part of the valid enumerations
// for this type.
//
// ------------------------------------------------------------------------
//
// Usage
//
//  posValueFmtMode := NumStrPosValFmtMode(0).NoLeadingPlusSign()
//
//  isValid := posValueFmtMode.XIsValid()
//
//  In this case the boolean value of 'isValid' is 'true'.
//
func (posFmtMode NumStrPosValFmtMode) XIsValid() bool {

	lockNumStrPosValFmtMode.Lock()

	defer lockNumStrPosValFmtMode.Unlock()

	if posFmtMode > 2 ||
		posFmtMode < 1 {
		return false
	}

	return true
}

// XParseString - Receives a string and attempts to match it with the
// string value of a supported enumeration. If successful, a new
// instance of NumStrPosValFmtMode is returned set to the value of the
// associated enumeration.
//
// This is a standard utility method and is not part of the valid
// enumerations for this type.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
// valueString   string
//     - A string which will be matched against the enumeration string
//       values. If 'valueString' is equal to one of the enumeration
//       names, this method will proceed to successful completion and
//       return the correct enumeration value.
//
// caseSensitive   bool
//     - If 'true' the search for enumeration names will be case
//       sensitive and will require an exact match. Therefore, 'leadingplusaign' will NOT
//       match the enumeration name, 'LeadingPlusSign'.
//
//       A case sensitive search will match any of the following strings:
//           "None"
//           "NoLeadingPlusSign"
//           "NoPlusSign"
//           "LeadingPlusSign"
//           "LeadingPlus"
//           "PlusSign"
//
//       If 'false' a case insensitive search is conducted
//       for the enumeration name. In this case, 'parentheses'
//       will match match enumeration name 'Parentheses'.
//
//       A case insensitive search will match any of the following
//       lower case names:
//           "none"
//           "noleadingplussign"
//           "noplussign"
//           "leadingplussign"
//           "leadingplus"
//           "plussign"
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
// NumStrPosValFmtMode
//     - Upon successful completion, this method will return a new
//       instance of NumStrPosValFmtMode set to the value of the enumeration
//       matched by the string search performed on input parameter,
//       'valueString'.
//
// error
//     - If this method completes successfully, the returned error
//       Type is set equal to 'nil'. If an error condition is encountered,
//       this method will return an error type which encapsulates an
//       appropriate error message.
//
// ------------------------------------------------------------------------
//
// Usage
//
// t, err := NumStrPosValFmtMode(0).XParseString("LeadingPlusSign", true)
//
//     t is now equal to NumStrPosValFmtMode(0).LeadingPlusSign()
//
func (posFmtMode NumStrPosValFmtMode) XParseString(
	valueString string,
	caseSensitive bool) (NumStrPosValFmtMode, error) {

	lockNumStrPosValFmtMode.Lock()

	defer lockNumStrPosValFmtMode.Unlock()

	ePrefix := "NumStrPosValFmtMode.XParseString() "

	if len(valueString) < 4 {
		return NumStrPosValFmtMode(0),
			fmt.Errorf(ePrefix+
				"\nInput parameter 'valueString' is INVALID!\n"+
				"String length is less than '4'.\n"+
				"valueString='%v'\n", valueString)
	}

	var ok bool
	var posValueFmtMode NumStrPosValFmtMode

	if caseSensitive {

		posValueFmtMode, ok = mNumStrPosValFmtModeStringToCode[valueString]

		if !ok {
			return NumStrPosValFmtMode(0),
				fmt.Errorf(ePrefix+
					"\n'valueString' did NOT MATCH a valid NumStrPosValFmtMode Value.\n"+
					"valueString='%v'\n", valueString)
		}

	} else {

		posValueFmtMode, ok = mNumStrPosValFmtModeLwrCaseStringToCode[strings.ToLower(valueString)]

		if !ok {
			return NumStrPosValFmtMode(0),
				fmt.Errorf(ePrefix+
					"\n'valueString' did NOT MATCH a valid NumStrPosValFmtMode Value.\n"+
					"valueString='%v'\n", valueString)
		}
	}

	return posValueFmtMode, nil
}

// XValue - This method returns the enumeration value of the current
// NumStrPosValFmtMode instance.
//
// This is a standard utility method and is not part of the valid enumerations
// for this type.
//
//
func (posFmtMode NumStrPosValFmtMode) XValue() NumStrPosValFmtMode {

	lockNumStrPosValFmtMode.Lock()

	defer lockNumStrPosValFmtMode.Unlock()

	return posFmtMode
}

// XValueInt - This method returns the integer value of the current
// NumStrPosValFmtMode value as an integer.
//
// This is a standard utility method and is not part of the valid enumerations
// for this type.
//
//
func (posFmtMode NumStrPosValFmtMode) XValueInt() int {

	lockNumStrPosValFmtMode.Lock()

	defer lockNumStrPosValFmtMode.Unlock()

	return int(posFmtMode)
}

// NStrPosValFmtMode - public global variable of
// type NumStrPosValFmtMode.
//
// This variable serves as an easier, short hand
// technique for accessing NumStrPosValFmtMode values.
//
// The formal syntax looks like this:  NumStrPosValFmtMode(0).LeadingPlusSign
//
// The abbreviated syntax looks this:  NStrPosValFmtMode.LeadingPlusSign()
//
// Usage:
// NStrPosValFmtMode.None(),
// NStrPosValFmtMode.NoLeadingPlusSign(),
// NStrPosValFmtMode.LeadingPlusSign(),
//
var NStrPosValFmtMode NumStrPosValFmtMode
