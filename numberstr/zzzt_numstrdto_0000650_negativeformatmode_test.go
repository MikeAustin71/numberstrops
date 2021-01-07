package numberstr

import "testing"

func TestNumStrNegValFmtMode_10(t *testing.T) {

	negFmtMode := NumStrNegValFmtMode(0).None()

	if negFmtMode.XIsValid() {
		t.Error("Error: Expected NumStrNegValFmtMode(0).None()\n" +
			"XIsValid()='false'. Instead, XIsValid()='true'.")
	}

}

func TestNumStrNegValFmtMode_20(t *testing.T) {

	negFmtMode := NumStrNegValFmtMode(0).LeadingMinusSign()

	if !negFmtMode.XIsValid() {
		t.Error("Error: Expected NumStrNegValFmtMode(0).LeadingMinusSign()\n" +
			"XIsValid()='true'. Instead, XIsValid()='false'.")
	}

}

func TestNumStrNegValFmtMode_30(t *testing.T) {

	negFmtModeLeadMinus := NumStrNegValFmtMode(0).LeadingMinusSign()

	str := negFmtModeLeadMinus.String()

	if str != "LeadingMinusSign" {
		t.Errorf("Error: String Value of NumStrNegValFmtMode(0).LeadingMinusSign()\n"+
			"should be 'LeadingMinusSign'\n"+
			"Instead, Int Value='%v'\n", int(negFmtModeLeadMinus))
	}

}

func TestNumStrNegValFmtMode_40(t *testing.T) {

	negFmtMode := NStrNegValFmtMode.Parentheses()

	str := negFmtMode.String()

	if str != "Parentheses" {
		t.Errorf("Error: String Value of NumStrNegValFmtMode(0).Parentheses()\n"+
			"should be 'Parentheses'.\n"+
			"Instead, Int Value='%v'\n", int(negFmtMode))
	}

}

func TestNumStrNegValFmtMode_50(t *testing.T) {

	negFmtMode, err :=
		NumStrNegValFmtMode(0).XParseString("Parentheses", true)

	if err != nil {
		t.Errorf("XParseString(\"Parentheses\") returned an error!\n"+
			"Error='%v'\n", err.Error())
		return
	}

	if negFmtMode != NStrNegValFmtMode.Parentheses() {
		t.Errorf("Error: XParseString() returned a bad value for \"Parentheses\".\n"+
			"Int Value of NumStrNegValFmtMode=%v\n",
			int(negFmtMode))
	}

}

func TestNumStrNegValFmtMode_60(t *testing.T) {

	negFmtMode, err :=
		NumStrNegValFmtMode(0).XParseString("LeadingMinusSign", true)

	if err != nil {
		t.Errorf("XParseString(\"LeadingMinusSign\") returned an error!\n"+
			"Error='%v'\n", err.Error())
		return
	}

	if negFmtMode != NStrNegValFmtMode.LeadingMinusSign() {
		t.Errorf("Error: XParseString() returned a bad value for \"LeadingMinusSign\".\n"+
			"Int Value of NumStrNegValFmtMode=%v\n",
			int(negFmtMode))
	}

}
