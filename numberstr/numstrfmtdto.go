package numberstr

import (
	"sync"
)

type NumStrFormatDto struct {
	numStrFmtConfigs map[NumStrValSpec]NumStrFormatter
	numFieldDto      NumberFieldDto
	lock             *sync.Mutex
}

// copyIn - Receives pointers to an instance of NumStrFormatDto
// labeled as input parameter, 'nStrFmtDtoIn'. This method then
// proceeds to copy all of the data from 'nStrFmtDtoIn' into the
// current NumStrFormatDto instance.  When the copy process is
// completed, the data values in 'nStrFmtDtoIn' and the current
// NumStrFormatDto instance will be identical.
//
// IMPORTANT
//
// Be advised, this method will overwrite the data values contained
// in the current instance of NumStrFormatDto.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  nStrFmtDtoIn        *NumStrFormatDto
//     - A pointer to an instance of NumStrFormatDto. This method
//       will copy all of the data values contained in
//       'numStrFmtDtoIn' to the current instance of
//       NumStrFormatDto.
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
//  err                 error
//     - If this method completes successfully, the returned error Type
//       is set to 'nil'. If errors are encountered during processing,
//       the returned error Type will encapsulate an error message.
//       Note that this error message will incorporate the method
//       chain and text passed by input parameter, 'ePrefix'. This
//       error prefix, 'ePrefix' will be prefixed to the beginning
//       of the error message.
//
func (nStrFmtDto *NumStrFormatDto) CopyIn(
	nStrFmtDtoIn *NumStrFormatDto,
	ePrefix string) (err error) {

	if nStrFmtDto.lock == nil {
		nStrFmtDto.lock = new(sync.Mutex)
	}

	nStrFmtDto.lock.Lock()

	defer nStrFmtDto.lock.Unlock()

	ePrefix += "NumStrFormatDto.CopyIn() "

	nStrFmtDtoAtom := numStrFmtDtoAtom{}

	err =
		nStrFmtDtoAtom.copyIn(
			nStrFmtDto,
			nStrFmtDtoIn,
			ePrefix)

	return err
}

// CopyOut - Returns a deep copy of the current NumStrFormatDto
// instance.
//
func (nStrFmtDto *NumStrFormatDto) CopyOut() NumStrFormatDto {

	if nStrFmtDto.lock == nil {
		nStrFmtDto.lock = new(sync.Mutex)
	}

	nStrFmtDto.lock.Lock()

	defer nStrFmtDto.lock.Unlock()

	nStrFmtDtoAtom := numStrFmtDtoAtom{}

	newNumStrFmtDto,
		_ :=
		nStrFmtDtoAtom.copyOut(
			nStrFmtDto,
			"")

	return newNumStrFmtDto
}

// IsValidInstanceError - Performs a diagnostic review of the current
// NumStrFormatDto instance to determine whether the current instance
// is a valid in all respects.
//
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
//     - If this method completes successfully, the returned boolean
//       value will signal whether the current NumStrFormatDto is
//       valid, or not. If the current NumStrFormatDto contains
//       valid data, this method returns 'true'. If the data is
//       invalid, this method returns 'false'.
//
//
func (nStrFmtDto *NumStrFormatDto) IsValidInstance() (
	isValid bool) {

	if nStrFmtDto.lock == nil {
		nStrFmtDto.lock = new(sync.Mutex)
	}

	nStrFmtDto.lock.Lock()

	defer nStrFmtDto.lock.Unlock()

	nStrFmtDtoNanobot := numStrFmtDtoNanobot{}

	isValid,
		_ =
		nStrFmtDtoNanobot.testNumStrFormatDtoValidity(
			nStrFmtDto,
			"")

	return isValid
}

// IsValidInstanceError - Performs a diagnostic review of the current
// NumStrFormatDto instance to determine whether the current instance
// is a valid in all respects.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  ePrefix             string
//     - This is an error prefix which is included in all returned
//       error messages. Usually, it contains the names of the calling
//       method or methods. Note: Be sure to leave a space at the end
//       of 'ePrefix'.
//
//
// -----------------------------------------------------------------
//
// Return Values
//
//  err                 error
//     - If this method completes successfully, the returned error Type
//       is set equal to 'nil'. If errors are encountered during processing,
//       the returned error Type will encapsulate an error message. Note
//       that this error message will incorporate the method chain and text
//       passed by input parameter, 'ePrefix'. Said text will be prefixed
//       to the beginning of the error message.
//
//       If this instance of NumStrFormatDto contains invalid data, a
//       detailed error message will be returned identifying the invalid
//       data item.
//
func (nStrFmtDto *NumStrFormatDto) IsValidInstanceError(
	ePrefix string) (
	err error) {

	if nStrFmtDto.lock == nil {
		nStrFmtDto.lock = new(sync.Mutex)
	}

	nStrFmtDto.lock.Lock()

	defer nStrFmtDto.lock.Unlock()

	ePrefix += "NumStrFormatDto.IsValidInstanceError() "

	nStrFmtDtoNanobot := numStrFmtDtoNanobot{}

	_,
		err =
		nStrFmtDtoNanobot.testNumStrFormatDtoValidity(
			nStrFmtDto,
			ePrefix)

	return err
}

// New - Creates and returns a new instance of
// NumStrFormatDto.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  numStrFormatterCollection  map[NumStrValSpec]NumStrFormatter
//   - An instance of type 'map' which contains three NumStrFormatter
//     objects. Each of the objects is tagged with a NumStrValSpec
//     specification identifying that format for use in one of three
//     type of number string formatting:
//
//       NumStrValSpec(0).CurrencyValue() - Provides format specifications
//       for converting numeric values to currency value number strings.
//
//       NumStrValSpec(0).AbsoluteValue() - Provides format specifications
//       for converting numeric values to absolute value number strings.
//
//       NumStrValSpec(0).SignedNumberValue() - Provides format specifications
//       for converting numeric values to signed number strings.
//
//       While it is possible to create this format information manually
//       for a specific culture or nationality, a much easier approach is
//       to simply use type NumStrFormatCountry and extract this information
//       by country.
//         Example:
//           numStrFormatterCollection :=
//              NumStrFormatCountry{}.NewPtr().UnitedStates()
//           -- Use intellisense to pick your country or culture.
//
//
//  ePrefix                    string
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
//  newFmtDto                  NumStrFormatDto
//     - If this method completes successfully, a new instance of NumStrFormatDto
//       be returned to calling function.
//
//
//  err                        error
//     - If this method completes successfully, the returned error Type is set
//       equal to 'nil'. If errors are encountered during processing, the
//       returned error Type will encapsulate an error message. Note this
//       error message will incorporate the method chain and text passed by
//       input parameter, 'ePrefix'. The 'ePrefix' text will be prefixed to
//       the beginning of the error message.
//
func (nStrFmtDto NumStrFormatDto) New(
	numStrFormatterCollection map[NumStrValSpec]NumStrFormatter,
	numberFieldLength int,
	ePrefix string) (
	newFmtDto NumStrFormatDto,
	err error) {

	if nStrFmtDto.lock == nil {
		nStrFmtDto.lock = new(sync.Mutex)
	}

	nStrFmtDto.lock.Lock()

	defer nStrFmtDto.lock.Unlock()

	ePrefix += "NumStrFormatDto.New() "
	newFmtDto = NumStrFormatDto{}

	nStrFmtDtoMolecule := numStrFmtDtoMolecule{}

	_,
		err = nStrFmtDtoMolecule.testValidityOfNumStrFmtCollection(
		numStrFormatterCollection,
		ePrefix+
			"Parameter 'numStrFormatterCollection' is invalid! ")

	if err != nil {
		return newFmtDto, err
	}

	newFmtDto.numFieldDto = NumberFieldDto{
		requestedNumFieldLength: numberFieldLength,
		actualNumFieldLength:    -1,
		minimumNumFieldLength:   -1,
		lock:                    new(sync.Mutex),
	}

	newFmtDto.lock = new(sync.Mutex)

	return newFmtDto, err
}

// SetNumberStringFormatters - Configures the current NumStrFormatDto
// with a collection of Number String Formatters which will be used
// to format absolute numeric values, signed number values and currency
// values. The easiest way to generate this collection is to select
// from a list of countries and cultures provided by type,
// NumStrFormatCountry.
//
// Example:
//   numStrFormatterCollection :=
//      NumStrFormatCountry{}.NewPtr().UnitedStates()
//   -- Use intellisense to pick your country or culture.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  numStrFormatterCollection  map[NumStrValSpec]NumStrFormatter
//   - An instance of type 'map' which contains three NumStrFormatter
//     objects. Each of the objects is tagged with a NumStrValSpec
//     specification identifying that format for use in one of three
//     types of number string formatting:
//
//       NumStrValSpec(0).CurrencyValue() - Provides format specifications
//       for converting numeric values to currency value number strings.
//
//       NumStrValSpec(0).AbsoluteValue() - Provides format specifications
//       for converting numeric values to absolute value number strings.
//
//       NumStrValSpec(0).SignedNumberValue() - Provides format specifications
//       for converting numeric values to signed number strings.
//
//       While it is possible to create this format information manually
//       for a specific culture or nationality, a much easier approach is
//       to simply use type NumStrFormatCountry and extract this information
//       by country.
//         Example:
//           numStrFormatterCollection :=
//              NumStrFormatCountry{}.NewPtr().UnitedStates()
//           -- Use intellisense to pick your country or culture.
//
//
//  ePrefix                    string
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
//  error
//     - If this method completes successfully, the returned error Type is set
//       equal to 'nil'. If errors are encountered during processing, the
//       returned error Type will encapsulate an error message. Note this
//       error message will incorporate the method chain and text passed by
//       input parameter, 'ePrefix'. The 'ePrefix' text will be prefixed to
//       the beginning of the error message.
//
func (nStrFmtDto *NumStrFormatDto) SetNumberStringFormatters(
	numStrFormatterCollection map[NumStrValSpec]NumStrFormatter,
	ePrefix string) error {

	if nStrFmtDto.lock == nil {
		nStrFmtDto.lock = new(sync.Mutex)
	}

	nStrFmtDto.lock.Lock()

	defer nStrFmtDto.lock.Unlock()

	ePrefix += "NumStrFormatDto.SetNumberStringFormatters() "

	nStrFmtDtoNanobot := numStrFmtDtoNanobot{}

	return nStrFmtDtoNanobot.setNumberStringFormatters(
		nStrFmtDto,
		numStrFormatterCollection,
		ePrefix)
}

// SetNumberFieldLength - Sets the length of the number field in
// which the formatted number string will be displayed.
//
// If the number field length is wider than the actual length of
// the formatted number string, that number string will be right
// justified in the number string. The left margin of the number
// field will be padded with spaces.
//
// Example:
//   numberFieldLen = 9
//   numberString =          "123.456"
//   Output Formatted Text = "  123.456"
//
// If the number field length is equal to or less than the actual
// length of the formatted number string, the number string will
// will be displayed in it's actual length.
//
// Example:
//   numberFieldLen = -1
//   numberString =          "123.456"
//   Output Formatted Text = "123.456"
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  numberFieldLen      int
//     - This parameter sets the length of the number field in which the
//       number string will be displayed.
//
//       Example:
//         numberFieldLen = 9
//         numberString =          "123.456"
//         Output Formatted Text = "  123.456"
//
//       If the number field length is equal to or less than the actual
//       length of the formatted number string, the number string will
//       will be displayed in it's actual length.
//
//       Example:
//         numberFieldLen = -1
//         numberString =          "123.456"
//         Output Formatted Text = "123.456"
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  --- NONE ---
//
//
func (nStrFmtDto *NumStrFormatDto) SetNumberFieldLength(
	numberFieldLen int) {

	if nStrFmtDto.lock == nil {
		nStrFmtDto.lock = new(sync.Mutex)
	}

	nStrFmtDto.lock.Lock()

	defer nStrFmtDto.lock.Unlock()

	nStrFmtDtoNanobot := numStrFmtDtoNanobot{}

	_ = nStrFmtDtoNanobot.setNumberFieldLength(
		nStrFmtDto,
		numberFieldLen,
		"")

	return
}

// SetToDefaults - This method will set all internal
// member variables to their default values.
//
// The Number String Format Defaults represent formatting
// parameters used in the United States.
//
func (nStrFmtDto *NumStrFormatDto) SetToDefaults() {

	if nStrFmtDto.lock == nil {
		nStrFmtDto.lock = new(sync.Mutex)
	}

	nStrFmtDto.lock.Lock()

	defer nStrFmtDto.lock.Unlock()

	nStrFmtDtoMech := numStrFormatDtoMechanics{}

	_ = nStrFmtDtoMech.setToDefaults(
		nStrFmtDto,
		"")

}

// SetToDefaultsIfEmpty - If the current NumStrFormatDto instance
// is empty or invalid, this method will set all of the data fields
// to the default values.
//
// The Number String Format Defaults represent formatting
// parameters used in the United States.
//
func (nStrFmtDto *NumStrFormatDto) SetToDefaultsIfEmpty() {

	if nStrFmtDto.lock == nil {
		nStrFmtDto.lock = new(sync.Mutex)
	}

	nStrFmtDto.lock.Lock()

	defer nStrFmtDto.lock.Unlock()

	nStrFmtDtoUtil := numStrFormatDtoUtility{}

	_ = nStrFmtDtoUtil.setToDefaultsIfEmpty(
		nStrFmtDto,
		"")

}
