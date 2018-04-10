package SearchProvider

type EmployeeSearchProvider struct {
}

func (esp EmployeeSearchProvider) Query(term string) ([]string, error) {

}

// http://iacdn.onbase.net/images/hyland/employees/180/skelemen.jpg

type Employee struct {
	Name       string
	Title      string
	Department string
	Location   string
	Status     string
	Phones     []string
	Email      string
	Seat       string
	Photo      string
}
