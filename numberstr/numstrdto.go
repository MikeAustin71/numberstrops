package numberstr

import (
	"sync"
)

//
//
// Supporting File hierarchy
// numstrdto.go ->
//  numstrdtoutility.go ->
//    numstrdtohelper.go ->
//     numstrdtomechanics.go ->
//       numstrdtonanobot.go ->
//         numstrdtomolecule.go ->
//           numstrdtoatom.go ->
//             numstrdtoelectron.go ->
//               numstrdtoquark.go
//

// NumStrDto - This Type contains data fields and methods used
// to format, manage, store, and transport number strings.
//
// Data Fields
//
//  signVal             int
//     - The sign value designates the numeric sign of the number
//       string. This integer contains one of three values:
//         +1 - Signals that the number string is a positive value.
//          0 - Signals that the number string is a zero value.
//         -1 - Signals that the number string is a negative value.
//
//
//  intDigitRunes       []rune
//     - This slice of runes holds all of the integer digits in the
//       number string.
//
//
//  fracDigitRunes      []rune
//     - This slice of runes holds all fo the fractional digits in
//       the number string. Fractional digits are those digits in
//       a floating number which are located to the right of the
//       decimal separator. For example, in the United States the
//       decimal separator is a period ('.') or decimal point.
//
//
//  baseNumSys                uint
//     - Identifies base numbering system for the number string's
//       numeric value. Supported base numbering systems are base
//       2, 8, 10, 16. These are otherwise known as binary, octal,
//       decimal and hexadecimal numbering systems.
//
//
//  fmtSpec             NumStrFmtSpecDto
//     - This object contains all of the format specifications
//       necessary to produce text versions of a number string.
//       Supported formats include, Absolute Value, Signed Number,
//       Currency and Scientific Notation.
//
type NumStrDto struct {
	signVal int // An integer value indicating the numeric sign of this number string.
	//          //   Valid values are +1, 0, or -1.
	//          //     +1 signals a positive value.
	//          //      0 signals a zero value
	//          //     -1 signals a negative value.
	intDigitRunes []rune // A slice of runes containing the integer digits component
	//                   //   of this number string.
	fracDigitRunes []rune // A slice of runes containing the fractional digits component
	//                    //   of this number string
	baseNumSys BaseNumberSystemType // Identifies the Base Numbering System.
	//                              //   Supports Base 2, 8, 10, and 16
	fmtSpec NumStrFmtSpecDto // Format Specifications
	lock    *sync.Mutex
}
