package report

type Report struct {
	ReportID string     `json:"ReportID"`
	Content  ReportData `json:"content"`
	Store
}

type ReportData struct {
	Name string `json:"name"`
}

// Crud Interface
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

func (r *Report) Create(report *Report) (bool, error) {
	return r.Store.Create(report)
}

func (r *Report) Read(i int) (*Report, error) {
	return r.Store.Read(i)
}

func (r *Report) Update(report *Report) (bool, error) {
	return r.Store.Update(report)
}

func (r *Report) Delete(*Report) (bool, error) {
	return r.Store.Delete(r)
}
