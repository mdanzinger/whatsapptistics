package report

type ReportCache interface {
	Get(*Report) (*Report, error)
	Set(*Report) error
}
