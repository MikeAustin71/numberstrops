package numberstr

import (
	"sync"
)

// FormatterCollection - Encapsulates all of the format
// specifications required to format the numeric values contained
// in type NumStrDto. The primary purpose of this type is to
// manage a collection of INumStrFormatter objects which provide
// formatting specifications for numeric values displayed in text
// strings. Member data elements are defined as follows as follows:
//
//  type FormatterCollection struct {
//   fmtCollection []INumStrFormatter
//  }
//
//
//  fmtCollection       []INumStrFormatter
//     - A collection of INumStrFormatter objects. Current
//       supported INumStrFormatter types include:
//          1. FormatterAbsoluteValue
//          2. FormatterBinary
//          3. FormatterCountry
//          4. FormatterCurrency
//          5. FormatterHexadecimal
//          6. FormatterOctal
//          7. FormatterSciNotation // Scientific Notation
//          8. FormatterSignedNumber
//
// An instance of FormatterCollection is maintained as member
// data variable by type NumStrDto.
//
type FormatterCollection struct {
	description   string
	tag           string
	fmtCollection []INumStrFormatter
	lock          *sync.Mutex
}

// IsValidInstance - Performs a diagnostic review of the current
// FormatterCollection instance to determine whether that instance
// is valid in all respects.
//
// The diagnostic analysis is performed on each element in the
// INumStrFormatter collection.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  --- NONE ---
//
//
// -----------------------------------------------------------------
//
// Return Values
//
//  isValid             bool
//     - If this method completes successfully, the returned
//       boolean value will signal whether the current
//       FormatterCollection is valid, or not. If the current
//       FormatterCollection contains valid data, this method
//       returns 'true'. If the data is invalid, this method
//       returns 'false'.
//
func (fmtCollection *FormatterCollection) IsValidInstance() (
	isValid bool) {

	if fmtCollection.lock == nil {
		fmtCollection.lock = new(sync.Mutex)
	}

	fmtCollection.lock.Lock()

	defer fmtCollection.lock.Unlock()

	isValid,
		_ = formatterCollectionQuark{}.ptr().
		testValidityOfFormatterCollection(
			fmtCollection,
			nil)

	return isValid
}