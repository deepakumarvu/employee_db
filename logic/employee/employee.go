package employee

import (
	"employee/models"
	"employee/pkg/logger"
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
	logger.Log.Debug().Str("id", strconv.Itoa(employee.ID)).
		Str("name", employee.Name).Str("position", employee.Position).Float64("salary", employee.Salary).
		Msg("Create Request received")

	if employee.ID == 0 {
		logger.Log.Error().Str("id", strconv.Itoa(employee.ID)).
			Msg("Invalid employee ID")
		return GetEmpError(InvalidID)
	}

	if employee.Name == "" || len(employee.Name) > 100 {
		logger.Log.Error().Str("name", employee.Name).
			Msg("Invalid employee name")
		return GetEmpError(NameInvalid)
	}

	if employee.Position == "" || len(employee.Position) > 100 {
		logger.Log.Error().Str("position", employee.Position).
			Msg("Invalid employee position")
		return GetEmpError(InvalidPosition)
	}

	if employee.Salary <= 0 {
		logger.Log.Error().Float64("salary", employee.Salary).
			Msg("Invalid employee salary")
		return GetEmpError(InvalidSalary)
	}

	if err := eh.db.SetItem(employee.ID, employee); err != nil {
		if errors.Is(err, simpledb.KeyAlreadyPresent) {
			logger.Log.Error().Int("id", employee.ID).
				Msg("Employee already exists")
			return GetEmpError(EmpAlreadyExists)
		}
		logger.Log.Error().Err(err).
			Msg("Error adding employee")
		return GetEmpError(ErrorAddingEmp)
	}

	logger.Log.Debug().Str("id", strconv.Itoa(employee.ID)).
		Msg("Request processed successfully")
	return nil
}

func (eh *Employee) GetEmployee(empID, LastEvalKeyID, numRecords string) (models.GetEmployeeResponse, error) {
	logger.Log.Debug().Str("empId", empID).Str("lastEvalKeyId", LastEvalKeyID).Str("numRecords", numRecords).
		Msg("Get Request received")

	var res models.GetEmployeeResponse

	// If empID is not empty, return the employee with the given ID
	if empID != "" {
		empIDInt, err := strconv.Atoi(empID)
		if err != nil {
			logger.Log.Error().Err(err).
				Str("empId", empID).
				Msg("Invalid employee ID")
			return res, GetEmpError(InvalidID)
		}
		emp, present := eh.db.GetItem(empIDInt)
		if !present {
			logger.Log.Error().Str("empId", empID).
				Msg("Employee not found")
			return res, GetEmpError(InvalidID)
		}
		res.Employees = append(res.Employees, emp)
		logger.Log.Debug().
			Str("empId", empID).Msg("Request processed successfully")
		return res, nil
	}

	// Else return the next batch of employees
	var LastEvalKeyIDInt int
	var err error
	if LastEvalKeyID != "" {
		LastEvalKeyIDInt, err = strconv.Atoi(LastEvalKeyID)
		if err != nil {
			logger.Log.Error().Err(err).
				Str("lastEvalKeyId", LastEvalKeyID).
				Msg("Invalid lastEvalKeyId")
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
			logger.Log.Error().Err(err).
				Str("lastEvalKeyId", LastEvalKeyID).
				Msg("Invalid lastEvalKeyId")
			return res, GetEmpError(InvalidLastEvalKeyID)
		}
		logger.Log.Error().Err(err).
			Msg("Error getting employees")
		return res, GetEmpError(ErrorGettingEmp)
	}
	if res.Employees == nil {
		res.Employees = []models.Employee{}
	}
	logger.Log.Debug().
		Msg("Request processed successfully")
	return res, nil
}

func (eh *Employee) UpdateEmployee(empID string, empUpdateReq models.EmployeeUpdateRequest) error {

	logger.Log.Debug().
		Str("empId", empID).
		Str("position", empUpdateReq.Position).
		Float64("salary", empUpdateReq.Salary).
		Msg("Update Request received")

	// Validate the employee ID
	if empID == "" {
		logger.Log.Error().Str("empId", empID).Msg("Invalid employee ID")
		return GetEmpError(InvalidID)
	}
	empInt, err := strconv.Atoi(empID)
	if err != nil {
		logger.Log.Error().Err(err).Str("empId", empID).Msg("Invalid employee ID")
		return GetEmpError(InvalidID)
	}

	// Validate the employee update request
	if empUpdateReq.Position == "" && empUpdateReq.Salary == 0 {
		logger.Log.Error().
			Str("position", empUpdateReq.Position).
			Float64("salary", empUpdateReq.Salary).
			Msg("Invalid employee update request")
		return GetEmpError(InvalidEmpUpdate)
	}

	if empUpdateReq.Salary < 0 {
		logger.Log.Error().Float64("salary", empUpdateReq.Salary).Msg("Invalid salary")
		return GetEmpError(InvalidSalary)
	}

	// Get the employee from the database
	currentEmp, ok := eh.db.GetItem(empInt)
	if !ok {
		logger.Log.Error().Int("empId", empInt).Msg("Employee not found")
		return GetEmpError(InvalidID)
	}

	// Update the employee with the new values
	if empUpdateReq.Position != "" {
		logger.Log.Debug().Str("position", empUpdateReq.Position).Msg("Updating position")
		currentEmp.Position = empUpdateReq.Position
	}

	if empUpdateReq.Salary != 0 {
		logger.Log.Debug().Float64("salary", empUpdateReq.Salary).Msg("Updating salary")
		currentEmp.Salary = empUpdateReq.Salary
	}

	if err := eh.db.UpdateItem(empInt, currentEmp); err != nil {
		if errors.Is(err, simpledb.KeyAbsent) {
			logger.Log.Error().Int("empId", empInt).Msg("Employee not found")
			return GetEmpError(InvalidID)
		}
		logger.Log.Error().Err(err).Msg("Error updating employee")
		return GetEmpError(ErrorUpdateEmp)
	}

	logger.Log.Debug().
		Str("empId", empID).
		Msg("Request processed successfully")
	return nil
}

func (eh *Employee) DeleteEmployee(empID string) error {

	logger.Log.Debug().Str("empId", empID).Msg("Delete Request received")

	if empID == "" {
		logger.Log.Error().Str("empId", empID).Msg("Invalid employee ID")
		return GetEmpError(InvalidID)
	}
	empInt, err := strconv.Atoi(empID)
	if err != nil {
		logger.Log.Error().Err(err).Str("empId", empID).Msg("Invalid employee ID")
		return GetEmpError(InvalidID)
	}

	if err := eh.db.DeleteItem(empInt); err != nil {
		if errors.Is(err, simpledb.KeyAbsent) {
			logger.Log.Error().Int("empId", empInt).Msg("Employee not found")
			return GetEmpError(InvalidID)
		}
		logger.Log.Error().Err(err).Msg("Error deleting employee")
		return GetEmpError(ErrorDeleteEmp)
	}
	logger.Log.Debug().Str("empId", empID).Msg("Request processed successfully")
	return nil
}
