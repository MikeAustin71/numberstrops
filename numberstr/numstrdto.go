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
//             numstrdtoelectron.go
//
type NumStrDto struct {
	signVal int // An integer value indicating the numeric sign of this number string.
	//                      //   Valid values are +1 or -1
	absAllNumRunes []rune // An array of runes containing all the numeric digits in a number with
	//                      //   no preceding plus or minus sign character. Example: 123.456 =
	//                      //   []rune{'1','2','3','4','5','6'}
	fmtSpec   NumStrFmtSpecDto // Format Specifications
	precision uint             // The number of digits to the right of the decimal point.
	lock      *sync.Mutex
}
