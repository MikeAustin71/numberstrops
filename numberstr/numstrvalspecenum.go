package numberstr

import (
	"fmt"
	"strings"
	"sync"
)

var mNumStrValSpecCodeToString = map[NumStrValSpec]string{
	NumStrValSpec(0): "None",
	NumStrValSpec(1): "AbsoluteValue",
	NumStrValSpec(2): "CurrencyValue",
	NumStrValSpec(3): "SignedNumberValue",
	NumStrValSpec(4): "ScientificNotation",
}

var mNumStrValSpecStringToCode = map[string]NumStrValSpec{
	"None":                 NumStrValSpec(0),
	"AbsoluteValue":        NumStrValSpec(1),
	"Absolute Value":       NumStrValSpec(1),
	"Absolute":             NumStrValSpec(1),
	"CurrencyValue":        NumStrValSpec(2),
	"Currency Value":       NumStrValSpec(2),
	"Currency":             NumStrValSpec(2),
	"SignedNumberValue":    NumStrValSpec(3),
	"Signed Number Value":  NumStrValSpec(3),
	"SignedNumericValue":   NumStrValSpec(3),
	"Signed Numeric Value": NumStrValSpec(3),
	"SignedNumber":         NumStrValSpec(3),
	"Signed Number":        NumStrValSpec(3),
	"Signed":               NumStrValSpec(3),
	"ScientificNotation":   NumStrValSpec(4),
	"Scientific Notation":  NumStrValSpec(4),
	"SCI":                  NumStrValSpec(4),
	"Scientific Form":      NumStrValSpec(4),
	"Standard Index Form":  NumStrValSpec(4),
	"Standard Form":        NumStrValSpec(4),
}

var mNumStrValSpecLwrCaseStringToCode = map[string]NumStrValSpec{
	"none":                 NumStrValSpec(0),
	"absolutevalue":        NumStrValSpec(1),
	"absolute value":       NumStrValSpec(1),
	"absolute":             NumStrValSpec(1),
	"currencyvalue":        NumStrValSpec(2),
	"currency value":       NumStrValSpec(2),
	"currency":             NumStrValSpec(2),
	"signednumbervalue":    NumStrValSpec(3),
	"signed number value":  NumStrValSpec(3),
	"signednumericvalue":   NumStrValSpec(3),
	"signed numeric value": NumStrValSpec(3),
	"signednumber":         NumStrValSpec(3),
	"signed number":        NumStrValSpec(3),
	"signed":               NumStrValSpec(3),
	"scientificnotation":   NumStrValSpec(4),
	"scientific notation":  NumStrValSpec(4),
	"sci":                  NumStrValSpec(4),
	"scientific form":      NumStrValSpec(4),
	"standard index form":  NumStrValSpec(4),
	"standard form":        NumStrValSpec(4),
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
//
// AbsoluteValue      (1)
//  - This specification signals that a numeric value will be displayed
//    in text as a positive number regardless of whether the native
//    value is positive or negative. Effectively, this means that
//    both negative values and positive values will be displayed as
//    positive numbers.
//
//    Examples:
//
//         Positive Values          Negative Values
//          +132 = +132              -123 = +123
//
//
// CurrencyValue      (2)
//  - The 'Currency Value' specification signals that all numeric values
//    will be displayed in number strings as currency formatted with
//    appropriate currency characters.
//
//    Currency number strings are always displayed as signed numeric
//    values with currency symbols included in the text string. This
//    means that positive values are displayed in text as positive
//    numbers with currency symbols (like the dollar sign) included
//    in the text string. Likewise, negative values are displayed in
//    text as negative numbers with currency symbols (like the dollar
//    sign) included in the text string.
//
//    Examples:
//         Positive Values          Negative Values
//          +132 = $132               -123 = ($123)
//
//
// SignedNumberValue  (3)
//  - Signals that the numeric value will be displayed in text as a
//    standard positive or negative value contingent upon the number
//    sign associated with the numeric value. NO CURRENCY Symbols will
//    be display in the resulting text number strings.
//
//    This is the default handling for numeric values.
//
//    'SignedNumberValue' means that positive values will be displayed
//     as positive numbers and negative values will be displayed as
//     negative numbers.
//
//     Examples:
//         Positive Values          Negative Values
//          +132 = 132               -123 = -123
//
//
// ScientificNotation (4)
//  - Signals that the numeric value will be displayed in text as
//    Scientific Notation.
//
//    Examples: '2.652e+8'     '2.652e-8'
//
//
// For easy access to these enumeration values, use the global variable
// 'NStrValSpec'. Example: NStrValSpec.SignedNumberValue()
//
// Otherwise you will need to use the formal syntax.
//     Example: NumStrValSpec(0).SignedNumberValue()
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

// AbsoluteValue - The 'Absolute Value' specification signals
// that both positive and negative numeric values will be
// displayed in number strings as positive numbers.
//
//         Examples:
//         Positive Values          Negative Values
//          +132 = +132              -123 = +123
//
// Note: Placement and formatting of positive and negative
// value signs, such as  plus ('+'), minus ('-') and
// parentheses is controlled by type, 'NumberStrFormatDto'.
//
// This method is part of the standard enumeration.
//
func (nStrValSpec NumStrValSpec) AbsoluteValue() NumStrValSpec {

	lockNumStrValSpec.Lock()

	defer lockNumStrValSpec.Unlock()

	return NumStrValSpec(1)
}

// CurrencyValue - The 'Currency Value' specification
// signals that all numeric values will be displayed in
// number strings as currency formatted with appropriate
// currency characters.
//
// Currency number strings are always displayed as signed
// numeric values with currency symbols included in the
// text string.
//
// This means that positive values are displayed in text
// as positive numbers with currency symbols (like the
// dollar sign) included in the text string.
//
// Likewise, negative values are displayed in text as
// negative numbers with currency symbols (like the dollar
// sign) included in the text string.
//
// Examples:
//         Positive Values          Negative Values
//          +132 = $132              -123 = ($123)
//
// Note: Placement and formatting formatting characters
// such as  plus ('+'), minus ('-') parentheses ('()') and
// currency ('$') is controlled by type, 'NumberStrFormatDto'.
//
//
// This method is part of the standard enumeration.
//
func (nStrValSpec NumStrValSpec) CurrencyValue() NumStrValSpec {

	lockNumStrValSpec.Lock()

	defer lockNumStrValSpec.Unlock()

	return NumStrValSpec(2)
}

// SignedNumberValue - The 'Signed Number Value' specification
// signals that positive numeric values will be displayed in
// number strings as positive numbers and negative numeric
// values will be displayed as negative numbers.
//
// This is the default for processing and converting numeric
// values into text strings.
//
//
// This method is part of the standard enumeration.
//
func (nStrValSpec NumStrValSpec) SignedNumberValue() NumStrValSpec {

	lockNumStrValSpec.Lock()

	defer lockNumStrValSpec.Unlock()

	return NumStrValSpec(3)
}

// ScientificNotation - The 'Scientific Notation' specification
// signals that numeric values, both positive and negative, will
// be displayed in text number strings as Scientific Notation.
//
//    Example Text Display:
//        "2.652e+8"
//        "2.652e-8"
//
// Reference:
//  https://en.wikipedia.org/wiki/Scientific_notation
//
// This is the default for processing and converting numeric
// values into text strings.
//
//
// This method is part of the standard enumeration.
//
func (nStrValSpec NumStrValSpec) ScientificNotation() NumStrValSpec {

	lockNumStrValSpec.Lock()

	defer lockNumStrValSpec.Unlock()

	return NumStrValSpec(4)
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
// t:= NumStrValSpec(0).SignedNumberValue()
// str := t.String()
//     str is now equal to 'SignedNumberValue'
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
//  Be advised, the value NumStrValSpec(0).None() is
//  classified as an 'invalid' value.
//
func (nStrValSpec NumStrValSpec) XIsValid() bool {

	lockNumStrValSpec.Lock()

	defer lockNumStrValSpec.Unlock()

	if nStrValSpec > 4 ||
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
//       sensitive and will require an exact match. Therefore, 'ScientificNotation' will NOT
//       match the enumeration name, 'scientificnotation'.
//
//       A case sensitive search will match any of the following strings:
//           "None"
//           "AbsoluteValue"
//           "Absolute Value"
//           "Absolute"
//           "CurrencyValue"
//           "Currency Value"
//           "Currency"
//           "SignedNumberValue"
//           "Signed Number Value"
//           "SignedNumericValue"
//           "Signed Numeric Value"
//           "SignedNumber"
//           "Signed Number"
//           "ScientificNotation"
//           "Scientific Notation"
//           "SCI"
//           "Scientific Form"
//           "Standard Index Form"
//           "Standard Form"
//
//       If 'false', a case insensitive search is conducted
//       for the enumeration name. In this case, 'scientificnotation'
//       will match match enumeration name 'ScientificNotation'.
//
//       A case insensitive search will match any of the following
//       lower case names:
//             "none"
//             "absolutevalue"
//             "absolute value"
//             "absolute"
//             "currencyvalue"
//             "currency value"
//             "currency"
//             "signednumbervalue"
//             "signed number value"
//             "signednumericvalue"
//             "signed numeric value"
//             "signednumber"
//             "signed number"
//             "signed"
//             "scientificnotation"
//             "scientific notation"
//             "sci"
//             "scientific form"
//             "standard index form"
//             "standard form"
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  NumStrValSpec
//     - Upon successful completion, this method will return a new
//       instance of NumStrValSpec set to the value of the enumeration
//       matched by the string search performed on input parameter,
//       'valueString'.
//
//  error
//     - If this method completes successfully, the returned error
//       Type is set equal to 'nil'. If an error condition is encountered,
//       this method will return an error type which encapsulates an
//       appropriate error message.
//
// ------------------------------------------------------------------------
//
// Usage
//
//  t, err := NumStrValSpec(0).XParseString("SignedNumberValue", true)
//
//     t is now equal to NumStrValSpec(0).SignedNumberValue()
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
//  Example: NStrValSpec.SignedNumberValue()
//
// Otherwise you will need to use the formal syntax.
//  Example: NumStrValSpec(0).SignedNumberValue()
//
// Usage:
// NStrValSpec.None(),
// NStrValSpec.AbsoluteValue(),
// NStrValSpec.CurrencyValue(),
// NStrValSpec.SignedNumberValue(),
// NStrValSpec.ScientificNotation(),
//
var NStrValSpec NumStrValSpec
