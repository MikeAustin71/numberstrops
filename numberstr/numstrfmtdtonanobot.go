package numberstr

import (
	"fmt"
	"sync"
)

type numStrFmtDtoNanobot struct {
	lock *sync.Mutex
}

// setNumberStringFormatters - Configures the NumStrFormatDto instance
// passed as an input parameter (nStrFmtDto) with new Number String
// Formatters. This collection of Number String Formatters which will be used
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
// IMPORTANT
// The old collection of Number String Formatters contained within
// input parameter 'nStrFmtDto' will be overwritten.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  nStrFmtDto                 *NumStrFormatDto
//     - This instance of NumStrFormatDto will receive new Number String
//       Formatters copied in from input parameter,
//       'numStrFormatterCollection'. The old collection will be deleted.
//
//
//  numStrFormatterCollection  map[NumStrValSpec]NumStrFormatter
//     - An instance of type 'map' which contains three NumStrFormatter
//       objects. Each of the objects is tagged with a NumStrValSpec
//       specification identifying that format for use in one of three
//       types of number string formatting:
//
//         NumStrValSpec(0).CurrencyValue() - Provides format specifications
//         for converting numeric values to currency value number strings.
//
//         NumStrValSpec(0).AbsoluteValue() - Provides format specifications
//         for converting numeric values to absolute value number strings.
//
//         NumStrValSpec(0).SignedNumberValue() - Provides format specifications
//         for converting numeric values to signed number strings.
//
//         While it is possible to create this format information manually
//         for a specific culture or nationality, a much easier approach is
//         to simply use type NumStrFormatCountry and extract this information
//         by country.
//           Example:
//             numStrFormatterCollection :=
//                NumStrFormatCountry{}.NewPtr().UnitedStates()
//             -- Use intellisense to pick your country or culture.
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
//  err                        error
//     - If this method completes successfully, the returned error Type is set
//       equal to 'nil'. If errors are encountered during processing, the
//       returned error Type will encapsulate an error message. Note this
//       error message will incorporate the method chain and text passed by
//       input parameter, 'ePrefix'. The 'ePrefix' text will be prefixed to
//       the beginning of the error message.
//
func (nStrFmtDtoNanobot *numStrFmtDtoNanobot) setNumberStringFormatters(
	nStrFmtDto *NumStrFormatDto,
	numStrFormatterCollection map[NumStrValSpec]NumStrFormatter,
	ePrefix string) (
	err error) {

	if nStrFmtDtoNanobot.lock == nil {
		nStrFmtDtoNanobot.lock = new(sync.Mutex)
	}

	nStrFmtDtoNanobot.lock.Lock()

	defer nStrFmtDtoNanobot.lock.Unlock()

	ePrefix += "numStrFmtDtoNanobot.setNumberStringFormatters() "

	if nStrFmtDto == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'nStrFmtDto' is a 'nil' pointer!\n",
			ePrefix)
		return err
	}

	nStrFmtDtoMolecule := numStrFmtDtoMolecule{}

	_,
		err = nStrFmtDtoMolecule.testValidityOfNumStrFmtCollection(
		numStrFormatterCollection,
		ePrefix+
			"Parameter 'numStrFormatterCollection' is invalid! ")

	if err != nil {
		return err
	}

	nStrFmtDto.numStrFmtConfigs =
		make(map[NumStrValSpec]NumStrFormatter, 0)

	nStrFmtDto.numStrFmtConfigs = make(map[NumStrValSpec]NumStrFormatter,
		len(numStrFormatterCollection))

	for idx, nStrFormtrs := range numStrFormatterCollection {
		nStrFmtDto.numStrFmtConfigs[idx] = nStrFormtrs.CopyOut()
	}

	return err
}

// setNumberFieldLength - Sets the number field length for the
// instance of NumStrFormatDto passed as an input parameter.
//
// This effectively sets the length of the number field in
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
//  nStrFmtDto          *NumStrFormatDto
//     - This instance of NumStrFormatDto will receive new Number String
//       Formatters copied in from input parameter,
//       'numStrFormatterCollection'. The old collection will be deleted.
//
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
//  err                        error
//     - If this method completes successfully, the returned error Type is set
//       equal to 'nil'. If errors are encountered during processing, the
//       returned error Type will encapsulate an error message. Note this
//       error message will incorporate the method chain and text passed by
//       input parameter, 'ePrefix'. The 'ePrefix' text will be prefixed to
//       the beginning of the error message.
//
func (nStrFmtDtoNanobot *numStrFmtDtoNanobot) setNumberFieldLength(
	nStrFmtDto *NumStrFormatDto,
	numberFieldLen int,
	ePrefix string) (
	err error) {

	if nStrFmtDtoNanobot.lock == nil {
		nStrFmtDtoNanobot.lock = new(sync.Mutex)
	}

	nStrFmtDtoNanobot.lock.Lock()

	defer nStrFmtDtoNanobot.lock.Unlock()

	ePrefix += "numStrFmtDtoNanobot.setNumberStringFormatters() "

	if nStrFmtDto == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'nStrFmtDto' is a 'nil' pointer!\n",
			ePrefix)
		return err
	}

	nStrFmtDto.numFieldDto.SetRequestedNumFieldLength(
		numberFieldLen)

	return err
}

// testNumStrFormatDtoValidity - Receives an instance of
// NumStrFormatDto and proceeds to test the validity of the
// member data fields.
//
// If one or more data elements are found to be invalid, an error
// is returned and the return boolean parameter, 'isValid', is set
// to 'false'.
//
func (nStrFmtDtoNanobot *numStrFmtDtoNanobot) testNumStrFormatDtoValidity(
	numStrFmtDto *NumStrFormatDto,
	ePrefix string) (
	isValid bool,
	err error) {

	if nStrFmtDtoNanobot.lock == nil {
		nStrFmtDtoNanobot.lock = new(sync.Mutex)
	}

	nStrFmtDtoNanobot.lock.Lock()

	defer nStrFmtDtoNanobot.lock.Unlock()

	ePrefix += "NumStrFormatDto.testNumStrFormatDtoValidity() "

	isValid = false

	if numStrFmtDto == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'numStrFmtDto' is a 'nil' pointer!\n",
			ePrefix)

		return isValid, err
	}

	nStrFmtDtoMolecule := numStrFmtDtoMolecule{}

	isValid,
		err = nStrFmtDtoMolecule.testValidityOfNumStrFmtCollection(
		numStrFmtDto.numStrFmtConfigs,
		ePrefix)

	return isValid, err
}
