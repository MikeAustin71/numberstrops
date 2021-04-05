package numberstr

import (
	"fmt"
	"sync"
)

type NumStrBasic struct {
	lock *sync.Mutex
}

// GetCountryFormatters - Receives a pointer to a Formatter
// Collection and a Country/Culture ID. Using these parameters the
// formatter collection is populated with a series of number string
// formatter specifications associated with the country or culture
// designated by the input parameter 'countryCulture'.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  fmtCollection       *FormatterCollection
//     - A pointer to an instance of FormatterCollection. This
//       object encapsulates multiple format specifications (a.k.a
//       'formatters) required to format the numeric values
//       contained in type NumStrDto.
//
//       This method will populate this formatter collection with
//       a series of format specifications associated with the
//       country or culture identified by input parameter,
//       'countryCulture'.
//
//       Currently, supported formatters include, binary, signed
//       number, absolute value, currency, hexadecimal and octal.
//
//
//  countryCulture      CountryCultureId
//     - An enumeration value specifying the country or culture for
//       which a series of number string formatter objects will be
//       created.
//
//  errorPrefix         interface{}
//     - This object encapsulates an error prefix string which is
//       included in all returned error messages. Usually, it
//       contains the names of the calling method or methods.
//
//       If no error prefix information is needed, set this parameter
//       to 'nil'.
//
//       This empty interface must be convertible to one of the
//       following types:
//
//
//       1. nil - A nil value is valid and generates an empty
//                collection of error prefix and error context
//                information.
//
//       2. string - A string containing error prefix information.
//
//       3. []string A one-dimensional slice of strings containing
//                   error prefix information
//
//       4. [][2]string A two-dimensional slice of strings containing
//                      error prefix and error context information.
//
//       5. ErrPrefixDto - An instance of ErrPrefixDto. The
//                         ErrorPrefixInfo from this object will be
//                         copied to 'errPrefDto'.
//
//       6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//                          ErrorPrefixInfo from this object will be
//                         copied to 'errPrefDto'.
//
//       7. IBasicErrorPrefix - An interface to a method generating
//                              a two-dimensional slice of strings
//                              containing error prefix and error
//                              context information.
//
//       If parameter 'errorPrefix' is NOT convertible to one of
//       the valid types listed above, it will be considered
//       invalid and trigger the return of an error.
//
//
// -----------------------------------------------------------------
//
// Return Values
//
//
//
//  error
//     - If this method completes successfully, the returned error
//       Type is set equal to 'nil'.
//
//       If errors are encountered during processing, the returned
//       error Type will encapsulate an error message. This
//       returned error message will incorporate the method chain
//       and text passed by input parameter, 'errorPrefix'. The
//       'errorPrefix' text will be attached to the beginning of
//       the error message.
//
func (numStrBasic *NumStrBasic) GetCountryFormatters(
	fmtCollection *FormatterCollection,
	countryCulture CountryCultureId,
	errorPrefix interface{}) error {

	if numStrBasic.lock == nil {
		numStrBasic.lock = new(sync.Mutex)
	}

	numStrBasic.lock.Lock()

	defer numStrBasic.lock.Unlock()

	methodName := "NumStrBasic." +
		"GetCountryFormatters()"

	var ePrefix *ErrPrefixDto
	var err error

	ePrefix,
		err = ErrPrefixDto{}.NewFromIEmpty(
		errorPrefix)

	if err != nil {
		return fmt.Errorf("%v\n"+
			"Error: Input parameter 'errorPrefix' is invalid!\n"+
			"%v\n",
			methodName,
			err.Error())
	}

	ePrefix.SetEPref(methodName)

	return numStrBasicQuark{}.ptr().getCountryFormatters(
		fmtCollection,
		countryCulture,
		ePrefix)
}
