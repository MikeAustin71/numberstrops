package numberstr

import (
	"fmt"
	"strings"
	"sync"
)

// Do NOT access these maps without first getting
// the lock on 'lockBaseNumberSystemType'.

var mBaseNumberSystemTypeCodeToString = map[BaseNumberSystemType]string{
	BaseNumberSystemType(0):  "None",
	BaseNumberSystemType(2):  "Binary",
	BaseNumberSystemType(8):  "Octal",
	BaseNumberSystemType(10): "Decimal",
	BaseNumberSystemType(16): "Hexadecimal",
}

var mBaseNumberSystemTypeStringToCode = map[string]BaseNumberSystemType{
	"None":        BaseNumberSystemType(0),
	"Binary":      BaseNumberSystemType(2),
	"Octal":       BaseNumberSystemType(8),
	"Decimal":     BaseNumberSystemType(10),
	"Hexadecimal": BaseNumberSystemType(16),
}

var mBaseNumberSystemTypeLwrCaseStringToCode = map[string]BaseNumberSystemType{
	"none":        BaseNumberSystemType(0),
	"binary":      BaseNumberSystemType(2),
	"octal":       BaseNumberSystemType(8),
	"decimal":     BaseNumberSystemType(10),
	"hexadecimal": BaseNumberSystemType(16),
}

// BaseNumberSystemType - The 'Base Number System Type' is
// an enumeration base number systems such as Base-10, Base-2,
// Base-16 and Base-8. These are otherwise known as the Decimal,
// Binary, Hexadecimal and Octal Numbering Systems.
// Reference:
//  https://en.wikipedia.org/wiki/Decimal
//  https://en.wikipedia.org/wiki/Binary_number
//  https://en.wikipedia.org/wiki/Hexadecimal
//  https://en.wikipedia.org/wiki/Octal
//
// Since the Go Programming Language does not directly support
// enumerations, the 'BaseNumberSystemType' type has been adapted
// to function in a manner similar to classic enumerations.
//
// 'BaseNumberSystemType' is declared as a type 'int'. The method
// names effectively represent an enumeration of Base Numbering
// Systems. These methods are listed as follows:
//
// Method             Integer
// Name                Value
// ------             -------
//
// None                  (0)
//  - Signals that the Base Number System Type
//    (BaseNumberSystemType) is not initialized. This is an error
//    condition.
//
//
// Binary                (2)
//  - This designation specifies the base-2 or binary numbering
//    system.
//
//
// Octal                 (8)
//  - This designation specifies the base-8 or octal numbering
//    system.
//
//
// Decimal              (10)
//  - This designation specifies the base-10 or decimal numbering
//    system.
//
//
// Hexadecimal          (16)
//  - This designation specifies the base-16 or hexadecimal
//    numbering system.
//
//
// To use the intellisense listing of available methods, apply the
// following syntax.
//
//     BaseNumberSystemType(0).Decimal()
//
// An simpler alternative requires using the constant BaseNumSys
// and the following syntax.
//
//     BaseNumSys.Decimal()
//
// Depending on your editor, intellisense (a.k.a. intelligent code
// completion) may not list the BaseNumberSystemType methods in
// alphabetical order.
//
// Be advised that all 'BaseNumberSystemType' methods beginning
// with 'X', as well as the method 'String()', are utility methods
// and not part of the enumeration.
//
type BaseNumberSystemType int

var lockBaseNumberSystemType sync.Mutex

// None - Signals that the NumStrFormatTypeCode Type is uninitialized.
// This is an error condition.
//
// This method is part of the standard enumeration.
//
func (baseNumSysType BaseNumberSystemType) None() BaseNumberSystemType {

	lockBaseNumberSystemType.Lock()

	defer lockBaseNumberSystemType.Unlock()

	return BaseNumberSystemType(0)
}

// Binary - This designation specifies the base-2 or binary
// numbering system.
//
// Reference
//  https://en.wikipedia.org/wiki/Binary_number
//
// This method is part of the standard enumeration.
//
func (baseNumSysType BaseNumberSystemType) Binary() BaseNumberSystemType {

	lockBaseNumberSystemType.Lock()

	defer lockBaseNumberSystemType.Unlock()

	return BaseNumberSystemType(2)
}

// Octal - This designation specifies the base-8 or octal numbering
// system.
//
// Reference
//  https://en.wikipedia.org/wiki/Octal
//
//
// This method is part of the standard enumeration.
//
func (baseNumSysType BaseNumberSystemType) Octal() BaseNumberSystemType {

	lockBaseNumberSystemType.Lock()

	defer lockBaseNumberSystemType.Unlock()

	return BaseNumberSystemType(8)
}

// Decimal - This designation specifies the base-10 or decimal
// numbering system.
//
// Reference
//  https://en.wikipedia.org/wiki/Decimal
//
//
// This method is part of the standard enumeration.
//
func (baseNumSysType BaseNumberSystemType) Decimal() BaseNumberSystemType {

	lockBaseNumberSystemType.Lock()

	defer lockBaseNumberSystemType.Unlock()

	return BaseNumberSystemType(10)
}

// Hexadecimal - This designation specifies the base-16 or
// hexadecimal numbering system.
//
// Reference
//  https://en.wikipedia.org/wiki/Hexadecimal
//
//
// This method is part of the standard enumeration.
//
func (baseNumSysType BaseNumberSystemType) Hexadecimal() BaseNumberSystemType {

	lockBaseNumberSystemType.Lock()

	defer lockBaseNumberSystemType.Unlock()

	return BaseNumberSystemType(16)
}

// String - Returns a string with the name of the enumeration associated
// with this current instance of 'BaseNumberSystemType'.
//
// This is a standard utility method and is not part of the valid
// enumerations for this type.
//
// ------------------------------------------------------------------------
//
// Usage
//
// t:= BaseNumberSystemType(0).Decimal()
// str := t.String()
//     str is now equal to 'Decimal'
//
func (baseNumSysType BaseNumberSystemType) String() string {

	lockBaseNumberSystemType.Lock()

	defer lockBaseNumberSystemType.Unlock()

	result, ok :=
		mBaseNumberSystemTypeCodeToString[baseNumSysType]

	if !ok {
		return "Error: Base Number System Type Value UNKNOWN!"
	}

	return result
}

// XIsValid - Returns a boolean value signaling whether the current
// Number String Value Specification (NumStrFormatTypeCode) value is valid.
//
// This is a standard utility method and is not part of the valid enumerations
// for this type.
//
// ------------------------------------------------------------------------
//
// Usage
//
//  baseNumType := BaseNumberSystemType(0).Binary()
//
//  isValid := baseNumType.XIsValid()
//
//  In this case the boolean value of 'isValid' is 'true'.
//
//  Be advised, the value BaseNumberSystemType(0).None() is
//  classified as an 'invalid' value.
//
func (baseNumSysType BaseNumberSystemType) XIsValid() bool {

	lockBaseNumberSystemType.Lock()

	defer lockBaseNumberSystemType.Unlock()

	if baseNumSysType < 1 {
		return false
	}

	_, isValid :=
		mBaseNumberSystemTypeCodeToString[baseNumSysType]

	return isValid
}

// XParseString - Receives a string and attempts to match it with the
// string value of a supported enumeration. If successful, a new
// instance of BaseNumberSystemType is returned set to the value of the
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
//       sensitive and will require an exact match. Therefore,
//       'decimal' will NOT match the enumeration name, 'Decimal'.
//
//       A case sensitive search will match any of the following strings:
//           "None"
//           "Binary"
//           "Octal"
//           "Decimal"
//           "Hexadecimal"
//
//       If 'false', a case insensitive search is conducted for the
//       enumeration name. In this example, 'decimal' will match
//       enumeration name 'Decimal'.
//
//       A case insensitive search will match any of the following
//       lower case names:
//           "none"
//           "binary"
//           "octal"
//           "decimal"
//           "hexadecimal"
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  BaseNumberSystemType
//     - Upon successful completion, this method will return a new
//       instance of BaseNumberSystemType set to the value of the enumeration
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
//  t, err := BaseNumberSystemType(0).XParseString("Decimal", true)
//
//     t is now equal to BaseNumberSystemType(0).Decimal()
//
func (baseNumSysType BaseNumberSystemType) XParseString(
	valueString string,
	caseSensitive bool) (BaseNumberSystemType, error) {

	lockBaseNumberSystemType.Lock()

	defer lockBaseNumberSystemType.Unlock()

	ePrefix := "BaseNumberSystemType.XParseString() "

	if len(valueString) < 3 {
		return BaseNumberSystemType(0),
			fmt.Errorf(ePrefix+
				"\nInput parameter 'valueString' is INVALID!\n"+
				"String length is less than '3'.\n"+
				"valueString='%v'\n", valueString)
	}

	var ok bool
	var nStrBaseNumSys BaseNumberSystemType

	if caseSensitive {

		nStrBaseNumSys, ok =
			mBaseNumberSystemTypeStringToCode[valueString]

		if !ok {
			return BaseNumberSystemType(0),
				fmt.Errorf(ePrefix+
					"\n'valueString' did NOT MATCH a valid BaseNumberSystemType Value.\n"+
					"valueString='%v'\n", valueString)
		}

	} else {

		nStrBaseNumSys, ok = mBaseNumberSystemTypeLwrCaseStringToCode[strings.ToLower(valueString)]

		if !ok {
			return BaseNumberSystemType(0),
				fmt.Errorf(ePrefix+
					"\n'valueString' did NOT MATCH a valid BaseNumberSystemType Value.\n"+
					"valueString='%v'\n", valueString)
		}
	}

	return nStrBaseNumSys, nil
}

// XValue - This method returns the enumeration value of the current
// BaseNumberSystemType instance.
//
// This is a standard utility method and is not part of the valid
// enumerations for this type.
//
//
func (baseNumSysType BaseNumberSystemType) XValue() BaseNumberSystemType {

	lockBaseNumberSystemType.Lock()

	defer lockBaseNumberSystemType.Unlock()

	return baseNumSysType
}

// XValueInt - This method returns the integer value of the current
// BaseNumberSystemType value as an integer.
//
// This is a standard utility method and is not part of the valid
// enumerations for this type.
//
//
func (baseNumSysType BaseNumberSystemType) XValueInt() int {

	lockBaseNumberSystemType.Lock()

	defer lockBaseNumberSystemType.Unlock()

	return int(baseNumSysType)
}

// BaseNumSys - A constant value for BaseNumberSystemType(0).
//
// This constant serves as an easier, short hand technique for
// accessing BaseNumberSystemType values.
//
// For easy access to these enumeration values, use this
// constant 'BaseNumSys'.
//  Example: BaseNumSys.Decimal()
//
// Otherwise you will need to use the formal syntax.
//  Example: BaseNumberSystemType(0).Decimal()
//
// Usage:
// BaseNumSys.None(),
// BaseNumSys.Binary(),
// BaseNumSys.Octal(),
// BaseNumSys.Decimal(),
// BaseNumSys.Hexadecimal(),
//
const BaseNumSys = BaseNumberSystemType(0)
