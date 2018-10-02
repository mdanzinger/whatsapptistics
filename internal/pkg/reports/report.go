package reports

type Report struct {
	ReportID int
	Content  []byte
	Store
}

// Interface Crud
type Store interface {
	Create(*Report) (bool, error)
	Read(id int) (*Report, error)
	Update(*Report) (bool, error)
	Delete(*Report) (bool, error)
}

func NewReport(s Store) *Report {
	return &Report{
		Store: s,
	}
}
