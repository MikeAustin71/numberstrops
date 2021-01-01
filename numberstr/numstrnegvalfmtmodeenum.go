package numberstr

import "sync"

var mNumStrNegValFmtModeStringToCode = map[string]NumStrNegValFmtMode{
	"None":                NumStrNegValFmtMode(0),
	"LeadingMinusSign":    NumStrNegValFmtMode(1),
	"ParenthesisSurround": NumStrNegValFmtMode(2),
	"AbsoluteValue":       NumStrNegValFmtMode(3),
}

var mNumStrNegValFmtModeCodeToString = map[NumStrNegValFmtMode]string{
	NumStrNegValFmtMode(0): "None",
	NumStrNegValFmtMode(1): "LeadingMinusSign",
	NumStrNegValFmtMode(2): "ParenthesisSurround",
	NumStrNegValFmtMode(3): "AbsoluteValue",
}

// NumStrNegValFmtMode - An enumeration Negative Value Format Modes
// used in formatting number strings for display purposes.
//
// Since the Go Programming Language does not directly support
// enumerations, the 'NumStrNegValFmtMode' type has been adapted to
// function in a manner similar to classic enumerations. 'NumStrNegValFmtMode'
// is declared as a type 'int'. The method names effectively represent
// an enumeration of Negative Format Modes. These methods are listed
// as follows:
//
//
// None                (0) - None - Signals that the Negative Value
//                           Format Mode (NumStrNegValFmtMode) Type
//                           is not initialized. This is an error
//                           condition.
//
// LeadingMinusSign    (1) - Leading Minus Sign signals that negative
//                           numeric values will be formatted in text
//                           with a leading minus sign. A leading minus
//                           is displayed as a minus character ('-') in
//                           first character position of the number string
//                           text.
//
// ParenthesisSurround (2) - Parenthesis signals that the negative value
//                           will be surrounded by parentheses. Example:
//                           (123)
//
// AbsoluteValue       (3) - Absolute Value signals that negative values
//                           will be displayed as positive or absolute values
//                           with no lead minus sign or surrounding parentheses.
//
//
// For easy access to these enumeration values, use the global variable
// 'NStrNegFmtMode'. Example: NStrNegFmtMode.ParenthesisSurround()
//
// Otherwise you will need to use the formal syntax.
//     Example: NumStrNegValFmtMode(0).AbsoluteValue()
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

// ParenthesisSurround - Signals that negative numeric values should
// be formatted with surrounding parentheses.
//
//      Example: '(1234)'
//
// This method is part of the standard enumeration.
//
func (negFmtMode NumStrNegValFmtMode) ParenthesisSurround() NumStrNegValFmtMode {

	lockNumStrNegValFmtMode.Lock()

	defer lockNumStrNegValFmtMode.Unlock()

	return NumStrNegValFmtMode(2)
}

// AbsoluteValue - Signals that negative numeric values should
// be formatted as positive or absolute values. Absolute values when
// displayed in number strings have no characters denoting negative
// value. The number string has no leading minus sign and no surrounding
// parentheses.
//
//      Example: '1234'
//
// This method is part of the standard enumeration.
//
func (negFmtMode NumStrNegValFmtMode) AbsoluteValue() NumStrNegValFmtMode {

	lockNumStrNegValFmtMode.Lock()

	defer lockNumStrNegValFmtMode.Unlock()

	return NumStrNegValFmtMode(3)
}

// String - Returns a string with the name of the enumeration associated
// with this instance of 'CalendarSpec'.
//
// This is a standard utility method and is not part of the valid enumerations
// for this type.
//
// ------------------------------------------------------------------------
//
// Usage
//
// t:= CalendarSpec(0).Gregorian()
// str := t.String()
//     str is now equal to 'Gregorian'
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
// Number String Negative Value Format Mode value is valid.
//
// This is a standard utility method and is not part of the valid enumerations
// for this type.
//
// ------------------------------------------------------------------------
//
// Usage
//
//  CalendarSpecification := CalendarSpec(0).Gregorian()
//
//  isValid := CalendarSpecification.XIsValid()
//
func (negFmtMode NumStrNegValFmtMode) XIsValid() bool {

	lockNumStrNegValFmtMode.Lock()

	defer lockNumStrNegValFmtMode.Unlock()

	if negFmtMode > 3 ||
		negFmtMode < 1 {
		return false
	}

	return true
}
