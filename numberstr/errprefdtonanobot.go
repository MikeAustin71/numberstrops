package numberstr

import (
	"fmt"
	"strings"
	"sync"
)

type errPrefixDtoNanobot struct {
	lock *sync.Mutex
}

// Ptr - Returns a pointer to a new instance of errPrefixDtoNanobot.
//
func (ePrefDtoNanobot errPrefixDtoNanobot) ptr() *errPrefixDtoNanobot {

	if ePrefDtoNanobot.lock == nil {
		ePrefDtoNanobot.lock = new(sync.Mutex)
	}

	ePrefDtoNanobot.lock.Lock()

	defer ePrefDtoNanobot.lock.Unlock()

	return &errPrefixDtoNanobot{
		lock: new(sync.Mutex),
	}
}

// setFromIBasicErrorPrefix - Receives an ErrPrefixDto object and
// then proceeds to overwrite and reset the ErrorPrefixInfo
// collection containing error prefix and error context
// information. New error information will be extracted from input
// parameter, 'iEPref' which implements the IBasicErrorPrefix
// interface.
//
// The IBasicErrorPrefix interface contains the method:
//  GetEPrefStrings() [][2]string
//
// This method returns a two-dimensional slice of error prefix and
// error context strings which can be used to create new
// ErrorPrefixInfo collection. This new ErrorPrefixInfo collection
// is assigned to the ErrPrefixDto object.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  errPrefDto          *ErrPrefixDto
//     - A pointer to an instance of ErrPrefixDto. Data values in
//       this ErrPrefixDto object will be overwritten an set to new
//       values extracted from input parameter, 'iEPref'.
//
//
//  iEPref              IBasicErrorPrefix
//     - An object which implements the IBasicErrorPrefix
//       interface.
//
//       This interface contains one method, 'GetEPrefStrings()'.
//       This single method returns a two dimensional array of
//       strings. Each 2-Dimensional array element holds a separate
//       string for error prefix and error context.
//
//       Information extracted from this object will be used to
//       generate error prefix information in a new instance of
//       ErrPrefixDto.
//
//
//  errorPrefStr        string
//     - This parameter contains an error prefix string which is
//       included in all returned error messages. Usually, it
//       contains the names of the calling method or methods.
//
//       If no error prefix information is needed, set this
//       parameter to an empty string: "".
//
//
// -----------------------------------------------------------------
//
// Return Values
//
//  error
//     - If this method completes successfully, the returned error
//       Type is set equal to 'nil'.
//
//       If errors are encountered during processing, the returned
//       error Type will encapsulate an error message. This
//       returned error message will incorporate the method chain
//       and text passed by input parameter, 'errorPrefStr'. The
//       'errorPrefStr' text will be attached to the beginning of
//       the error message.
//
func (ePrefDtoNanobot *errPrefixDtoNanobot) setFromIBasicErrorPrefix(
	errPrefDto *ErrPrefixDto,
	iEPref IBasicErrorPrefix,
	errorPrefStr string) error {

	if ePrefDtoNanobot.lock == nil {
		ePrefDtoNanobot.lock = new(sync.Mutex)
	}

	ePrefDtoNanobot.lock.Lock()

	defer ePrefDtoNanobot.lock.Unlock()

	methodName := errorPrefStr + "\nerrPrefixDtoNanobot." +
		"setFromIBasicErrorPrefix()"

	if errPrefDto == nil {
		return fmt.Errorf("%v\n"+
			"Error: Input parameter 'errPrefDto' is invalid!\n"+
			"'errPrefDto' is a 'nil' pointer.\n",
			methodName)
	}

	if iEPref == nil {
		return fmt.Errorf("%v\n"+
			"Error: Input parameter 'iEPref' is invalid!\n"+
			"'iEPref' is a 'nil' pointer.\n",
			methodName)
	}

	twoDSlice := iEPref.GetEPrefStrings()

	lenTwoDSlice := len(twoDSlice)

	if lenTwoDSlice == 0 {
		errPrefDto.ePrefCol = make([]ErrorPrefixInfo, 0)
		return nil
	}

	errPrefDto.ePrefCol = make([]ErrorPrefixInfo, lenTwoDSlice)

	var ePrefInfo ErrorPrefixInfo
	var lastIdx = lenTwoDSlice - 1
	var isFirstIdx, isLastIdx bool

	for i := 0; i < lenTwoDSlice; i++ {

		if i == 0 {
			isFirstIdx = true
		} else {
			isFirstIdx = false
		}

		if i == lastIdx {
			isLastIdx = true
		} else {
			isLastIdx = false
		}

		ePrefInfo = ErrorPrefixInfo{
			isFirstIdx:             isFirstIdx,
			isLastIdx:              isLastIdx,
			isPopulated:            true,
			errorPrefixStr:         twoDSlice[i][0],
			lenErrorPrefixStr:      uint(len(twoDSlice[i][0])),
			errPrefixHasContextStr: len(twoDSlice[i][1]) > 0,
			errorContextStr:        twoDSlice[i][1],
			lenErrorContextStr:     uint(len(twoDSlice[i][1])),
			lock:                   nil,
		}

		errPrefDto.ePrefCol[i] = ePrefInfo
	}

	return nil
}

// setFromString - Receives an ErrPrefixDto object and then
// proceeds to overwrite and reset the ErrorPrefixInfo collection
// containing error prefix and error context information. New
// error prefix information will be extracted from input
// parameter, 'iEPref', a string which contains error prefix
// information.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  errPrefDto          *ErrPrefixDto
//     - A pointer to an instance of ErrPrefixDto. Data values in
//       this ErrPrefixDto object will be overwritten an set to new
//       values extracted from input parameter, 'iEPref'.
//
//
//  iEPref              string
//     - A string containing error prefix information which will be
//       extracted and used to overwrite the existing error prefix
//       information contained in input parameter, 'errPrefDto'.
//
//       If this string is delimited by new line characters ("\n")
//       an array of strings will be created and used to create
//       multiple error prefix entries in the 'errPrefDto'
//       ErrorPrefixInfo collection.
//
//
//  errorPrefStr        string
//     - This parameter contains an error prefix string which is
//       included in all returned error messages. Usually, it
//       contains the names of the calling method or methods.
//
//       If no error prefix information is needed, set this
//       parameter to an empty string: "".
//
//
// -----------------------------------------------------------------
//
// Return Values
//
//  error
//     - If this method completes successfully, the returned error
//       Type is set equal to 'nil'.
//
//       If errors are encountered during processing, the returned
//       error Type will encapsulate an error message. This
//       returned error message will incorporate the method chain
//       and text passed by input parameter, 'errorPrefStr'. The
//       'errorPrefStr' text will be attached to the beginning of
//       the error message.
//
func (ePrefDtoNanobot *errPrefixDtoNanobot) setFromString(
	errPrefDto *ErrPrefixDto,
	iEPref string,
	errorPrefStr string) error {

	if ePrefDtoNanobot.lock == nil {
		ePrefDtoNanobot.lock = new(sync.Mutex)
	}

	ePrefDtoNanobot.lock.Lock()

	defer ePrefDtoNanobot.lock.Unlock()

	methodName := errorPrefStr + "\nerrPrefixDtoNanobot." +
		"setFromIBasicErrorPrefix()"

	if errPrefDto == nil {
		return fmt.Errorf("%v\n"+
			"Error: Input parameter 'errPrefDto' is invalid!\n"+
			"'errPrefDto' is a 'nil' pointer.\n",
			methodName)
	}

	if len(iEPref) == 0 {

		errPrefDto.ePrefCol = nil

		return nil
	}

	ePrefStrs := strings.Split(iEPref, "\n")

	lenEPrefStrs := len(ePrefStrs)

	errPrefDto.ePrefCol = make([]ErrorPrefixInfo, lenEPrefStrs)

	var isFirstIdx, isLastIdx bool
	var lastIdx = lenEPrefStrs - 1

	for i := 0; i < lenEPrefStrs; i++ {

		if i == 0 {
			isFirstIdx = true
		} else {
			isFirstIdx = false
		}

		if i == lastIdx {
			isLastIdx = true
		} else {
			isLastIdx = false
		}

		errPrefDto.ePrefCol[i] =
			ErrorPrefixInfo{
				isFirstIdx:             isFirstIdx,
				isLastIdx:              isLastIdx,
				isPopulated:            true,
				errorPrefixStr:         ePrefStrs[i],
				lenErrorPrefixStr:      uint(len(ePrefStrs[i])),
				errPrefixHasContextStr: false,
				errorContextStr:        "",
				lenErrorContextStr:     0,
				lock:                   nil,
			}
	}

	return nil
}

// setFromStringArray - Receives an ErrPrefixDto object and then
// proceeds to overwrite and reset the ErrorPrefixInfo collection
// containing error prefix and error context information. New
// error prefix information will be extracted from input
// parameter, 'iEPref', a string array which contains error prefix
// information.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  errPrefDto          *ErrPrefixDto
//     - A pointer to an instance of ErrPrefixDto. Data values in
//       this ErrPrefixDto object will be overwritten an set to new
//       values extracted from input parameter, 'iEPref'.
//
//
//  iEPref              []string
//     - A string containing error prefix information which will be
//       extracted and used to overwrite the existing error prefix
//       information contained in input parameter, 'errPrefDto'.
//
//       If this string is delimited by new line characters ("\n")
//       an array of strings will be created and used to create
//       multiple error prefix entries in the 'errPrefDto'
//       ErrorPrefixInfo collection.
//
//
//  errorPrefStr        string
//     - This parameter contains an error prefix string which is
//       included in all returned error messages. Usually, it
//       contains the names of the calling method or methods.
//
//       If no error prefix information is needed, set this
//       parameter to an empty string: "".
//
//
// -----------------------------------------------------------------
//
// Return Values
//
//  error
//     - If this method completes successfully, the returned error
//       Type is set equal to 'nil'.
//
//       If errors are encountered during processing, the returned
//       error Type will encapsulate an error message. This
//       returned error message will incorporate the method chain
//       and text passed by input parameter, 'errorPrefStr'. The
//       'errorPrefStr' text will be attached to the beginning of
//       the error message.
//
func (ePrefDtoNanobot *errPrefixDtoNanobot) setFromStringArray(
	errPrefDto *ErrPrefixDto,
	iEPref []string,
	errorPrefStr string) error {

	if ePrefDtoNanobot.lock == nil {
		ePrefDtoNanobot.lock = new(sync.Mutex)
	}

	ePrefDtoNanobot.lock.Lock()

	defer ePrefDtoNanobot.lock.Unlock()

	methodName := errorPrefStr + "\nerrPrefixDtoNanobot." +
		"setFromIBasicErrorPrefix()"

	if errPrefDto == nil {
		return fmt.Errorf("%v\n"+
			"Error: Input parameter 'errPrefDto' is invalid!\n"+
			"'errPrefDto' is a 'nil' pointer.\n",
			methodName)
	}

	lenEPrefStrs := len(iEPref)

	if lenEPrefStrs == 0 {

		errPrefDto.ePrefCol = nil

		return nil
	}

	errPrefDto.ePrefCol = make([]ErrorPrefixInfo, lenEPrefStrs)

	var isFirstIdx, isLastIdx bool
	var lastIdx = lenEPrefStrs - 1

	for i := 0; i < lenEPrefStrs; i++ {

		if i == 0 {
			isFirstIdx = true
		} else {
			isFirstIdx = false
		}

		if i == lastIdx {
			isLastIdx = true
		} else {
			isLastIdx = false
		}

		errPrefDto.ePrefCol[i] =
			ErrorPrefixInfo{
				isFirstIdx:             isFirstIdx,
				isLastIdx:              isLastIdx,
				isPopulated:            true,
				errorPrefixStr:         iEPref[i],
				lenErrorPrefixStr:      uint(len(iEPref[i])),
				errPrefixHasContextStr: false,
				errorContextStr:        "",
				lenErrorContextStr:     0,
				lock:                   nil,
			}
	}

	return nil
}

// setFromTwoDStrArray - Receives an ErrPrefixDto object and
// then proceeds to overwrite and reset the ErrorPrefixInfo
// collection containing error prefix and error context
// information. New error information will be extracted from the
// two two dimensional string array contained in input parameter,
// 'iEPref'.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  errPrefDto          *ErrPrefixDto
//     - A pointer to an instance of ErrPrefixDto. Data values in
//       this ErrPrefixDto object will be overwritten an set to new
//       values extracted from input parameter, 'iEPref'.
//
//
//  iEPref              [][2]string
//     - A two dimensional string array containing error prefix and
//       error context strings in pairs.
//
//       iEPref[x][0] - contains an error prefix string.
//       iEPref[x][0] - contains an error context string if one
//                      exists. Otherwise, it is populated with an
//                      empty string.
//
//       Information extracted from this array will be used to
//       generate and transfer error prefix and error context
//       information to the ErrPrefixDto input parameter,
//       'errPrefDto'.
//
//
//  errorPrefStr        string
//     - This parameter contains an error prefix string which is
//       included in all returned error messages. Usually, it
//       contains the names of the calling method or methods.
//
//       If no error prefix information is needed, set this
//       parameter to an empty string: "".
//
//
// -----------------------------------------------------------------
//
// Return Values
//
//  error
//     - If this method completes successfully, the returned error
//       Type is set equal to 'nil'.
//
//       If errors are encountered during processing, the returned
//       error Type will encapsulate an error message. This
//       returned error message will incorporate the method chain
//       and text passed by input parameter, 'errorPrefStr'. The
//       'errorPrefStr' text will be attached to the beginning of
//       the error message.
//
func (ePrefDtoNanobot *errPrefixDtoNanobot) setFromTwoDStrArray(
	errPrefDto *ErrPrefixDto,
	iEPref [][2]string,
	errorPrefStr string) error {

	if ePrefDtoNanobot.lock == nil {
		ePrefDtoNanobot.lock = new(sync.Mutex)
	}

	ePrefDtoNanobot.lock.Lock()

	defer ePrefDtoNanobot.lock.Unlock()

	methodName := errorPrefStr + "\nerrPrefixDtoNanobot." +
		"setFromTwoDStrArray()"

	if errPrefDto == nil {
		return fmt.Errorf("%v\n"+
			"Error: Input parameter 'errPrefDto' is invalid!\n"+
			"'errPrefDto' is a 'nil' pointer.\n",
			methodName)
	}

	if iEPref == nil {
		errPrefDto.ePrefCol = nil
		return nil
	}

	lenTwoDSlice := len(iEPref)

	if lenTwoDSlice == 0 {
		errPrefDto.ePrefCol = nil
		return nil
	}

	errPrefDto.ePrefCol = make([]ErrorPrefixInfo, lenTwoDSlice)

	var ePrefInfo ErrorPrefixInfo
	var lastIdx = lenTwoDSlice - 1
	var isFirstIdx, isLastIdx bool

	for i := 0; i < lenTwoDSlice; i++ {

		if i == 0 {
			isFirstIdx = true
		} else {
			isFirstIdx = false
		}

		if i == lastIdx {
			isLastIdx = true
		} else {
			isLastIdx = false
		}

		ePrefInfo = ErrorPrefixInfo{
			isFirstIdx:             isFirstIdx,
			isLastIdx:              isLastIdx,
			isPopulated:            true,
			errorPrefixStr:         iEPref[i][0],
			lenErrorPrefixStr:      uint(len(iEPref[i][0])),
			errPrefixHasContextStr: len(iEPref[i][1]) > 0,
			errorContextStr:        iEPref[i][1],
			lenErrorContextStr:     uint(len(iEPref[i][1])),
			lock:                   nil,
		}

		errPrefDto.ePrefCol[i] = ePrefInfo
	}

	return nil
}
