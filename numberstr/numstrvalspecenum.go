package numberstr

import (
	"fmt"
	"strings"
	"sync"
)

var mNumStrValSpecCodeToString = map[NumStrValSpec]string{
	NumStrValSpec(0): "None",
	NumStrValSpec(1): "StdPlusOrMinus",
	NumStrValSpec(2): "AbsoluteValue",
}

var mNumStrValSpecStringToCode = map[string]NumStrValSpec{
	"None":           NumStrValSpec(0),
	"StdPlusOrMinus": NumStrValSpec(1),
	"PlusOrMinus":    NumStrValSpec(1),
	"Plus Or Minus":  NumStrValSpec(1),
	"AbsoluteValue":  NumStrValSpec(2),
	"Absolute Value": NumStrValSpec(2),
}

var mNumStrValSpecLwrCaseStringToCode = map[string]NumStrValSpec{
	"none":           NumStrValSpec(0),
	"stdplusorminus": NumStrValSpec(1),
	"plusorminus":    NumStrValSpec(1),
	"plus or minus":  NumStrValSpec(1),
	"absolutevalue":  NumStrValSpec(2),
	"absolute value": NumStrValSpec(2),
}

// NumStrValSpec - The 'Number String Value Specification' is an
// enumeration of specification or formatting codes for the display
// of numeric values in text number strings.
//
// Since the Go Programming Language does not directly support
// enumerations, the 'NumStrValSpec' type has been adapted to
// function in a manner similar to classic enumerations.
// NumStrValSpec is declared as a type 'int'. The method names
// effectively represent an enumeration of for numeric
// value display specifications. These methods are listed as
// follows:
//
//
// None               (0)
//  - Signals that the Number String Value Specification (NumStrValSpec)
//    Type is not initialized. This is an error condition.
//
// StdPlusOrMinus     (1)
//  - Signals that the numeric value will be displayed in text as a
//    standard positive or negative value depending on the number sign
//    associated with the numeric value. This is the default handling
//    for numeric values. It simply means that positive values will be
//    displayed as positive numbers and negative values will be displayed
//    as negative numbers.
//         Example:  -123 = -123 and +132 = 132
//
// AbsoluteValue      (2)
//  - This specification signals that a numeric value will be displayed as
//    a positive number regardless of whether the native value is positive
//    or negative. Effectively, this means that negative values will be
//    displayed as positive numbers and positive values will be displayed
//    as positive numbers.
//         Example:  -123 = 123  and 132 = 132
//
//
// For easy access to these enumeration values, use the global variable
// 'NStrValSpec'. Example: NStrValSpec.StdPlusOrMinus()
//
// Otherwise you will need to use the formal syntax.
//     Example: NumStrValSpec(0).StdPlusOrMinus()
//
//
// Depending on your editor, intellisense (a.k.a. intelligent code
// completion) may not list the NumStrValSpec methods in alphabetical
// order. Be advised that all 'NumStrValSpec' methods beginning with 'X', as well as the method 'String()',
// are utility methods and not part of the enumeration.
//
type NumStrValSpec int

var lockNumStrValSpec sync.Mutex

// None - Signals that the NumStrValSpec Type is uninitialized.
// This is an error condition.
//
// This method is part of the standard enumeration.
//
func (nStrValSpec NumStrValSpec) None() NumStrValSpec {

	lockNumStrValSpec.Lock()

	defer lockNumStrValSpec.Unlock()

	return NumStrValSpec(0)
}

// StdPlusOrMinus - The 'Standard Plus or Minus' specification
// signals that positive numeric values will be displayed in
// number strings as positive numbers and negative numeric
// values will be displayed as negative numbers.
//
//
// This method is part of the standard enumeration.
//
func (nStrValSpec NumStrValSpec) StdPlusOrMinus() NumStrValSpec {

	lockNumStrValSpec.Lock()

	defer lockNumStrValSpec.Unlock()

	return NumStrValSpec(1)
}

// AbsoluteValue - The 'Absolute Value' specification signals
// that both positive and negative numeric values will be
// displayed in number strings as positive numbers.
//
// This method is part of the standard enumeration.
//
func (nStrValSpec NumStrValSpec) AbsoluteValue() NumStrValSpec {

	lockNumStrValSpec.Lock()

	defer lockNumStrValSpec.Unlock()

	return NumStrValSpec(2)
}

// String - Returns a string with the name of the enumeration associated
// with this current instance of 'NumStrValSpec'.
//
// This is a standard utility method and is not part of the valid enumerations
// for this type.
//
// ------------------------------------------------------------------------
//
// Usage
//
// t:= NumStrValSpec(0).StdPlusOrMinus()
// str := t.String()
//     str is now equal to 'StdPlusOrMinus'
//
func (nStrValSpec NumStrValSpec) String() string {

	lockNumStrValSpec.Lock()

	defer lockNumStrValSpec.Unlock()

	result, ok := mNumStrValSpecCodeToString[nStrValSpec]

	if !ok {
		return "Error: Number String Value Specification UNKNOWN!"
	}

	return result
}

// XIsValid - Returns a boolean value signaling whether the current
// Number String Value Specification (NumStrValSpec) value is valid.
//
// This is a standard utility method and is not part of the valid enumerations
// for this type.
//
// ------------------------------------------------------------------------
//
// Usage
//
//  nStrValueSpec := NumStrValSpec(0).AbsoluteValue()
//
//  isValid := nStrValueSpec.XIsValid()
//
//  In this case the boolean value of 'isValid' is 'true'.
//
func (nStrValSpec NumStrValSpec) XIsValid() bool {

	lockNumStrValSpec.Lock()

	defer lockNumStrValSpec.Unlock()

	if nStrValSpec > 2 ||
		nStrValSpec < 1 {
		return false
	}

	return true
}

// XParseString - Receives a string and attempts to match it with the
// string value of a supported enumeration. If successful, a new
// instance of NumStrValSpec is returned set to the value of the
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
//           "StdPlusOrMinus"
//           "PlusOrMinus"
//           "Plus Or Minus"
//           "AbsoluteValue"
//           "Absolute Value"
//
//       If 'false' a case insensitive search is conducted
//       for the enumeration name. In this case, 'parentheses'
//       will match match enumeration name 'Parentheses'.
//
//       A case insensitive search will match any of the following
//       lower case names:
//           "none"
//           "stdplusorminus"
//           "plusorminus"
//           "plus or minus"
//           "absolutevalue"
//           "absolute value"
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
// NumStrValSpec
//     - Upon successful completion, this method will return a new
//       instance of NumStrValSpec set to the value of the enumeration
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
// t, err := NumStrValSpec(0).XParseString("StdPlusOrMinus", true)
//
//     t is now equal to NumStrValSpec(0).StdPlusOrMinus()
//
func (nStrValSpec NumStrValSpec) XParseString(
	valueString string,
	caseSensitive bool) (NumStrValSpec, error) {

	lockNumStrValSpec.Lock()

	defer lockNumStrValSpec.Unlock()

	ePrefix := "NumStrValSpec.XParseString() "

	if len(valueString) < 4 {
		return NumStrValSpec(0),
			fmt.Errorf(ePrefix+
				"\nInput parameter 'valueString' is INVALID!\n"+
				"String length is less than '4'.\n"+
				"valueString='%v'\n", valueString)
	}

	var ok bool
	var nStrValueSpec NumStrValSpec

	if caseSensitive {

		nStrValueSpec, ok = mNumStrValSpecStringToCode[valueString]

		if !ok {
			return NumStrValSpec(0),
				fmt.Errorf(ePrefix+
					"\n'valueString' did NOT MATCH a valid NumStrValSpec Value.\n"+
					"valueString='%v'\n", valueString)
		}

	} else {

		nStrValueSpec, ok = mNumStrValSpecLwrCaseStringToCode[strings.ToLower(valueString)]

		if !ok {
			return NumStrValSpec(0),
				fmt.Errorf(ePrefix+
					"\n'valueString' did NOT MATCH a valid NumStrValSpec Value.\n"+
					"valueString='%v'\n", valueString)
		}
	}

	return nStrValueSpec, nil
}

// XValue - This method returns the enumeration value of the current
// NumStrValSpec instance.
//
// This is a standard utility method and is not part of the valid enumerations
// for this type.
//
//
func (nStrValSpec NumStrValSpec) XValue() NumStrValSpec {

	lockNumStrValSpec.Lock()

	defer lockNumStrValSpec.Unlock()

	return nStrValSpec
}

// XValueInt - This method returns the integer value of the current
// NumStrValSpec value as an integer.
//
// This is a standard utility method and is not part of the valid enumerations
// for this type.
//
//
func (nStrValSpec NumStrValSpec) XValueInt() int {

	lockNumStrValSpec.Lock()

	defer lockNumStrValSpec.Unlock()

	return int(nStrValSpec)
}

// NStrValSpec - public global variable of
// type NumStrValSpec.
//
// This variable serves as an easier, short hand
// technique for accessing NumStrValSpec values.
//
// For easy access to these enumeration values, use the
// global variable 'NStrValSpec'.
//  Example: NStrValSpec.StdPlusOrMinus()
//
// Otherwise you will need to use the formal syntax.
//  Example: NumStrValSpec(0).StdPlusOrMinus()
//
// Usage:
// NStrValSpec.None(),
// NStrValSpec.StdPlusOrMinus(),
// NStrValSpec.AbsoluteValue(),
//
var NStrValSpec NumStrValSpec
