package numberstr

import (
	"fmt"
	"strings"
	"sync"
)

var mNumStrNegValFmtModeStringToCode = map[string]NumStrNegValFmtMode{
	"None":                NumStrNegValFmtMode(0),
	"LeadingMinusSign":    NumStrNegValFmtMode(1),
	"LeadingMinus":        NumStrNegValFmtMode(1),
	"MinusSign":           NumStrNegValFmtMode(1),
	"Minus":               NumStrNegValFmtMode(1),
	"ParenthesesSurround": NumStrNegValFmtMode(2),
	"Parentheses":         NumStrNegValFmtMode(2),
	"ParenthesisSurround": NumStrNegValFmtMode(2),
	"Parenthesis":         NumStrNegValFmtMode(2),
}

var mNumStrNegValFmtModeLwrCaseStringToCode = map[string]NumStrNegValFmtMode{
	"none":                NumStrNegValFmtMode(0),
	"leadingminussign":    NumStrNegValFmtMode(1),
	"leadingminus":        NumStrNegValFmtMode(1),
	"minussign":           NumStrNegValFmtMode(1),
	"minus":               NumStrNegValFmtMode(1),
	"parenthesessurround": NumStrNegValFmtMode(2),
	"parentheses":         NumStrNegValFmtMode(2),
	"parenthesissurround": NumStrNegValFmtMode(2),
	"parenthesis":         NumStrNegValFmtMode(2),
}

var mNumStrNegValFmtModeCodeToString = map[NumStrNegValFmtMode]string{
	NumStrNegValFmtMode(0): "None",
	NumStrNegValFmtMode(1): "LeadingMinusSign",
	NumStrNegValFmtMode(2): "Parentheses",
}

// NumStrNegValFmtMode - An enumeration of Negative Value Format Modes
// used to format negative numeric values in number strings for display
// purposes.
//
// Since the Go Programming Language does not directly support
// enumerations, the 'NumStrNegValFmtMode' type has been adapted to
// function in a manner similar to classic enumerations.
// 'NumStrNegValFmtMode' is declared as a type 'int'. The method names
// effectively represent an enumeration of for negative numeric values.
// These methods are listed as follows:
//
//
// None                (0)
//  - Signals that the Negative Value Format Mode (NumStrNegValFmtMode)
//    Type is not initialized. This is an error condition.
//
// LeadingMinusSign    (1)
//  - Leading Minus Sign signals that negative numeric values will be
//    formatted in text with a leading minus sign. A leading minus is
//    displayed as a minus character ('-') in the first character
//    position of the number string text.
//         Example:  -123
//
// Parentheses         (2)
//  - 'Parentheses' signals that the negative value will be surrounded by
//     parentheses.
//         Example:  (123)
//
//
// For easy access to these enumeration values, use the global variable
// 'NStrNegValFmtMode'. Example: NStrNegValFmtMode.Parentheses()
//
// Otherwise you will need to use the formal syntax.
//     Example: NumStrNegValFmtMode(0).Parentheses()
//
//
// Depending on your editor, intellisense (a.k.a. intelligent code
// completion) may not list the NumStrNegValFmtMode methods in
// alphabetical order. Be advised that all 'NumStrNegValFmtMode'
// methods beginning with 'X', as well as the method 'String()',
// are utility methods and not part of the enumeration.
//
type NumStrNegValFmtMode int

var lockNumStrNegValFmtMode sync.Mutex

// None - Signals that the NumStrNegValFmtMode Type is uninitialized.
// This is an error condition.
//
// This method is part of the standard enumeration.
//
func (negFmtMode NumStrNegValFmtMode) None() NumStrNegValFmtMode {

	lockNumStrNegValFmtMode.Lock()

	defer lockNumStrNegValFmtMode.Unlock()

	return NumStrNegValFmtMode(0)
}

// LeadingMinusSign - Signals that negative numeric values should
// be formatted with a leading minus sign.
//
//      Example: '-1234'
//
// This method is part of the standard enumeration.
//
func (negFmtMode NumStrNegValFmtMode) LeadingMinusSign() NumStrNegValFmtMode {

	lockNumStrNegValFmtMode.Lock()

	defer lockNumStrNegValFmtMode.Unlock()

	return NumStrNegValFmtMode(1)
}

// Parentheses - Signals that negative numeric values should
// be formatted with surrounding parentheses.
//
//      Example: '(1234)'
//
// This method is part of the standard enumeration.
//
func (negFmtMode NumStrNegValFmtMode) Parentheses() NumStrNegValFmtMode {

	lockNumStrNegValFmtMode.Lock()

	defer lockNumStrNegValFmtMode.Unlock()

	return NumStrNegValFmtMode(2)
}

// String - Returns a string with the name of the enumeration associated
// with this current instance of 'NumStrNegValFmtMode'.
//
// This is a standard utility method and is not part of the valid enumerations
// for this type.
//
// ------------------------------------------------------------------------
//
// Usage
//
// t:= NumStrNegValFmtMode(0).LeadingMinusSign()
// str := t.String()
//     str is now equal to 'LeadingMinusSign'
//
func (negFmtMode NumStrNegValFmtMode) String() string {

	lockNumStrNegValFmtMode.Lock()

	defer lockNumStrNegValFmtMode.Unlock()

	result, ok := mNumStrNegValFmtModeCodeToString[negFmtMode]

	if !ok {
		return "Error: Number String Negative Value Format Mode UNKNOWN!"
	}

	return result
}

// XIsValid - Returns a boolean value signaling whether the current
// Number String Negative Value Format Mode (NumStrNegValFmtMode)
// value is valid.
//
// This is a standard utility method and is not part of the valid enumerations
// for this type.
//
// ------------------------------------------------------------------------
//
// Usage
//
//  negValueFmtMode := NumStrNegValFmtMode(0).Parentheses()
//
//  isValid := negValueFmtMode.XIsValid()
//
//  In this case the boolean value of 'isValid' is 'true'.
//
func (negFmtMode NumStrNegValFmtMode) XIsValid() bool {

	lockNumStrNegValFmtMode.Lock()

	defer lockNumStrNegValFmtMode.Unlock()

	if negFmtMode > 2 ||
		negFmtMode < 1 {
		return false
	}

	return true
}

// XParseString - Receives a string and attempts to match it with the
// string value of a supported enumeration. If successful, a new
// instance of NumStrNegValFmtMode is returned set to the value of the
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
//       sensitive and will require an exact match. Therefore, 'parentheses' will NOT
//       match the enumeration name, 'Parentheses'.
//
//       A case sensitive search will match any of the following strings:
//           "None"
//           "LeadingMinusSign"
//           "LeadingMinus"
//           "MinusSign"
//           "Minus"
//           "ParenthesesSurround"
//           "Parentheses"
//           "ParenthesisSurround"
//           "Parenthesis"
//
//       If 'false' a case insensitive search is conducted
//       for the enumeration name. In this case, 'parentheses'
//       will match match enumeration name 'Parentheses'.
//
//       A case insensitive search will match any of the following
//       lower case names:
//           "none"
//           "leadingminussign"
//           "leadingminus"
//           "minussign"
//           "minus"
//           "parenthesessurround"
//           "parentheses"
//           "parenthesissurround"
//           "parenthesis"
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
// NumStrNegValFmtMode
//     - Upon successful completion, this method will return a new
//       instance of NumStrNegValFmtMode set to the value of the enumeration
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
// t, err := NumStrNegValFmtMode(0).XParseString("ParenthesesSurround", true)
//
//     t is now equal to NumStrNegValFmtMode(0).ParenthesesSurround()
//
func (negFmtMode NumStrNegValFmtMode) XParseString(
	valueString string,
	caseSensitive bool) (NumStrNegValFmtMode, error) {

	lockNumStrNegValFmtMode.Lock()

	defer lockNumStrNegValFmtMode.Unlock()

	ePrefix := "NumStrNegValFmtMode.XParseString() "

	if len(valueString) < 4 {
		return NumStrNegValFmtMode(0),
			fmt.Errorf(ePrefix+
				"\nInput parameter 'valueString' is INVALID!\n"+
				"String length is less than '4'.\n"+
				"valueString='%v'\n", valueString)
	}

	var ok bool
	var negValueFmtMode NumStrNegValFmtMode

	if caseSensitive {

		negValueFmtMode, ok = mNumStrNegValFmtModeStringToCode[valueString]

		if !ok {
			return NumStrNegValFmtMode(0),
				fmt.Errorf(ePrefix+
					"\n'valueString' did NOT MATCH a valid NumStrNegValFmtMode Value.\n"+
					"valueString='%v'\n", valueString)
		}

	} else {

		negValueFmtMode, ok = mNumStrNegValFmtModeLwrCaseStringToCode[strings.ToLower(valueString)]

		if !ok {
			return NumStrNegValFmtMode(0),
				fmt.Errorf(ePrefix+
					"\n'valueString' did NOT MATCH a valid NumStrNegValFmtMode Value.\n"+
					"valueString='%v'\n", valueString)
		}
	}

	return negValueFmtMode, nil
}

// XValue - This method returns the enumeration value of the current
// NumStrNegValFmtMode instance.
//
// This is a standard utility method and is not part of the valid enumerations
// for this type.
//
//
func (negFmtMode NumStrNegValFmtMode) XValue() NumStrNegValFmtMode {

	lockNumStrNegValFmtMode.Lock()

	defer lockNumStrNegValFmtMode.Unlock()

	return negFmtMode
}

// XValueInt - This method returns the integer value of the current
// NumStrNegValFmtMode value as an integer.
//
// This is a standard utility method and is not part of the valid enumerations
// for this type.
//
//
func (negFmtMode NumStrNegValFmtMode) XValueInt() int {

	lockNumStrNegValFmtMode.Lock()

	defer lockNumStrNegValFmtMode.Unlock()

	return int(negFmtMode)
}

// NStrNegValFmtMode - public global variable of
// type NumStrNegValFmtMode.
//
// This variable serves as an easier, short hand
// technique for accessing NumStrNegValFmtMode values.
//
// Usage:
// NStrNegValFmtMode.None(),
// NStrNegValFmtMode.LeadingMinusSign(),
// NStrNegValFmtMode.Parentheses(),
//
var NStrNegValFmtMode NumStrNegValFmtMode
