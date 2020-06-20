package employee

type employee struct {
	department string
	name       string
	age        int
	salary     float64
}

func New(d string, name string, age int, salary float64) employee {
	var e = employee{d, name, age, salary}
	return e
}

func (e *employee) Zgz(m float64) float64 {
	e.salary += m
	return e.salary
}
