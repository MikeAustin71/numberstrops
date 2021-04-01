package numberstr

import (
	"sync"
)

/*
 NumStrBasicUtility
 ==================

 Type 'NumStrBasicUtility' provides a basic set of numeric
 conversion and management routines primarily focused on
 number strings. It is now deprecated and superseded by
 type 'NumStrDto' which provides many additional capabilities.

*/

type NumStrBasicUtility struct {
	NumStrFormatSpec FormatterCollection
	StrIn            string
	StrOut           string
	IsFractionalVal  bool
	IntegerStr       string
	FractionStr      string
	Int64Val         int64
	Float64Val       float64
	lock             *sync.Mutex
}
