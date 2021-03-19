package numberstr

type INumStrFormatter interface {
	SetNumStrFormatTypeCode()

	GetNumStrFormatTypeCode() NumStrFormatTypeCode
}
