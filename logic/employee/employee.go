package employee

import (
	"employee/models"
	"employee/service/simpledb"
	"errors"
	"strconv"
)

type Employee struct {
	db *simpledb.Database[int, models.Employee]
}

// NewEmployee creates a new instance of the Employee struct and initializes its db field with a new instance of the simpledb.Database[int, models.Employee] struct.
//
// It returns a pointer to the newly created Employee struct.
func NewEmployee() *Employee {
	var d simpledb.Database[int, models.Employee]
	return &Employee{
		db: d.Init(),
	}
}

// CreateEmployee creates a new employee in the system.
//
// It takes in a models.Employee object as a parameter and returns an error.
func (eh *Employee) CreateEmployee(employee models.Employee) error {

	if employee.ID == 0 {
		return GetEmpError(InvalidID)
	}

	if employee.Name == "" || len(employee.Name) > 100 {
		return GetEmpError(NameInvalid)
	}

	if employee.Position == "" || len(employee.Position) > 100 {
		return GetEmpError(InvalidPosition)
	}

	if employee.Salary <= 0 {
		return GetEmpError(InvalidSalary)
	}

	if err := eh.db.SetItem(employee.ID, employee); err != nil {
		if errors.Is(err, simpledb.KeyAlreadyPresent) {
			return GetEmpError(EmpAlreadyExists)
		}
		return GetEmpError(ErrorAddingEmp)
	}

	return nil
}

func (eh *Employee) GetEmployee(empID, LastEvalKeyID, numRecords string) (models.GetEmployeeResponse, error) {
	var res models.GetEmployeeResponse

	// If empID is not empty, return the employee with the given ID
	if empID != "" {
		empIDInt, err := strconv.Atoi(empID)
		if err != nil {
			return res, GetEmpError(InvalidID)
		}
		emp, present := eh.db.GetItem(empIDInt)
		if !present {
			return res, GetEmpError(InvalidID)
		}
		res.Employees = append(res.Employees, emp)
		return res, nil
	}

	// Else return the next batch of employees
	var LastEvalKeyIDInt int
	var err error
	if LastEvalKeyID != "" {
		LastEvalKeyIDInt, err = strconv.Atoi(LastEvalKeyID)
		if err != nil {
			return res, GetEmpError(InvalidLastEvalKeyID)
		}
	}
	numRecordsInt, err := strconv.Atoi(numRecords)
	if err != nil {
		numRecordsInt = 10 // Default value if numRecords is not provided
	}

	// Get the next batch of employees
	res.Employees, res.LastEvalKeyID, err = eh.db.GetItems(LastEvalKeyIDInt, numRecordsInt)
	if err != nil {
		if errors.Is(err, simpledb.InvalidLastEvalKeyID) {
			return res, GetEmpError(InvalidLastEvalKeyID)
		}
		return res, GetEmpError(ErrorGettingEmp)
	}
	if res.Employees == nil {
		res.Employees = []models.Employee{}
	}
	return res, nil
}

func (eh *Employee) UpdateEmployee(empID string, empUpdateReq models.EmployeeUpdateRequest) error {

	// Validate the employee ID
	if empID == "" {
		return GetEmpError(InvalidID)
	}
	empInt, err := strconv.Atoi(empID)
	if err != nil {
		return GetEmpError(InvalidID)
	}

	// Validate the employee update request
	if empUpdateReq.Position == "" && empUpdateReq.Salary == 0 {
		return GetEmpError(InvalidEmpUpdate)
	}

	if empUpdateReq.Salary < 0 {
		return GetEmpError(InvalidSalary)
	}

	// Get the employee from the database
	currentEmp, ok := eh.db.GetItem(empInt)
	if !ok {
		return GetEmpError(InvalidID)
	}

	// Update the employee with the new values
	if empUpdateReq.Position != "" {
		currentEmp.Position = empUpdateReq.Position
	}

	if empUpdateReq.Salary != 0 {
		currentEmp.Salary = empUpdateReq.Salary
	}

	if err := eh.db.UpdateItem(empInt, currentEmp); err != nil {
		if errors.Is(err, simpledb.KeyAbsent) {
			return GetEmpError(InvalidID)
		}
		return GetEmpError(ErrorUpdateEmp)
	}

	return nil
}

func (eh *Employee) DeleteEmployee(empID string) error {

	if empID == "" {
		return GetEmpError(InvalidID)
	}
	empInt, err := strconv.Atoi(empID)
	if err != nil {
		return GetEmpError(InvalidID)
	}

	if err := eh.db.DeleteItem(empInt); err != nil {
		if errors.Is(err, simpledb.KeyAbsent) {
			return GetEmpError(InvalidID)
		}
		return GetEmpError(ErrorDeleteEmp)
	}

	return nil
}
