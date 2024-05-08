package employee

import (
	"employee/pkg/apierror"
	"net/http"
)

func GetEmpError(c EmpError) *apierror.APIError {
	return EmpErrors[c]
}

type EmpError int

const (
	InvalidID EmpError = iota + 100010
	NameInvalid
	InvalidPosition
	InvalidSalary
	EmpAlreadyExists
	ErrorAddingEmp
	InvalidLastEvalKeyID
	ErrorGettingEmp
	ErrorDeleteEmp
	InvalidEmpUpdate
	ErrorUpdateEmp
)

var EmpErrors = map[EmpError]*apierror.APIError{
	InvalidID:            {HttpStatusCode: http.StatusBadRequest, ErrCode: int(InvalidID), ErrorMessage: "Provide a valid ID"},
	NameInvalid:          {HttpStatusCode: http.StatusBadRequest, ErrCode: int(NameInvalid), ErrorMessage: "Name cannot be empty or longer than 100 characters"},
	InvalidPosition:      {HttpStatusCode: http.StatusBadRequest, ErrCode: int(InvalidPosition), ErrorMessage: "Position cannot be empty or longer than 100 characters"},
	InvalidSalary:        {HttpStatusCode: http.StatusBadRequest, ErrCode: int(InvalidSalary), ErrorMessage: "Salary cannot be less than equal to 0"},
	EmpAlreadyExists:     {HttpStatusCode: http.StatusBadRequest, ErrCode: int(EmpAlreadyExists), ErrorMessage: "Employee already exists"},
	ErrorAddingEmp:       {HttpStatusCode: http.StatusInternalServerError, ErrCode: int(ErrorAddingEmp), ErrorMessage: "Error adding employee"},
	InvalidLastEvalKeyID: {HttpStatusCode: http.StatusBadRequest, ErrCode: int(InvalidLastEvalKeyID), ErrorMessage: "Invalid last evaluated ID"},
	ErrorGettingEmp:      {HttpStatusCode: http.StatusInternalServerError, ErrCode: int(ErrorGettingEmp), ErrorMessage: "Error getting employees"},
	ErrorDeleteEmp:       {HttpStatusCode: http.StatusInternalServerError, ErrCode: int(ErrorDeleteEmp), ErrorMessage: "Error deleting employee"},
	InvalidEmpUpdate:     {HttpStatusCode: http.StatusBadRequest, ErrCode: int(InvalidEmpUpdate), ErrorMessage: "Invalid employee update request"},
	ErrorUpdateEmp:       {HttpStatusCode: http.StatusInternalServerError, ErrCode: int(ErrorUpdateEmp), ErrorMessage: "Error updating employee"},
}
