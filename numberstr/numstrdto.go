package numberstr

import (
	"sync"
)

// NumStrDto - This Type contains data fields and methods used
// to manage, store and transport number strings.
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
	fmtSpec NumStrFmtSpecDto // Format Specifications
	lock    *sync.Mutex
}
