package report

type ReportUsecase interface {
	Create(*Report) error
	Read(id int) (*Report, error)
	Update(*Report) error
	Delete(*Report) error
}
