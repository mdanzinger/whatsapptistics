package report

type ReportRepository interface {
	Create(*Report) error
	Read(id int) (*Report, error)
	Update(*Report) error
	Delete(*Report) error
}
