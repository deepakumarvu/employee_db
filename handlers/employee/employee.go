package employee

import (
	"employee/logic/employee"
	"employee/models"
	"employee/pkg/apierror"
	"net/http"

	"github.com/gin-gonic/gin"
)

type EmployeeHandler struct {
	emp *employee.Employee
}

func NewEmployeeHandler() *EmployeeHandler {
	return &EmployeeHandler{
		emp: employee.NewEmployee(),
	}
}

func (eh *EmployeeHandler) CreateEmployee(c *gin.Context) {

	var employee models.Employee

	if err := c.BindJSON(&employee); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if err := eh.emp.CreateEmployee(employee); err != nil {
		apiError := err.(*apierror.APIError)
		c.AbortWithStatusJSON(apiError.HttpStatusCode, apiError)
		return
	}

	c.JSON(http.StatusOK, models.APIResponse{Message: "OK"})
}

func (eh *EmployeeHandler) GetEmployee(c *gin.Context) {
	empID := c.Query("id")
	LastEvalKeyID := c.Query("last_eval_id")
	numRecords := c.Query("num_records")
	var res models.GetEmployeeResponse
	var err error
	if res, err = eh.emp.GetEmployee(empID, LastEvalKeyID, numRecords); err != nil {
		apiError := err.(*apierror.APIError)
		c.AbortWithStatusJSON(apiError.HttpStatusCode, apiError)
		return
	}
	c.JSON(http.StatusOK, res)
}

func (eh *EmployeeHandler) UpdateEmployee(c *gin.Context) {

	empID := c.Query("id")

	var employee models.EmployeeUpdateRequest

	if err := c.BindJSON(&employee); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if err := eh.emp.UpdateEmployee(empID, employee); err != nil {
		apiError := err.(*apierror.APIError)
		c.AbortWithStatusJSON(apiError.HttpStatusCode, apiError)
		return
	}

	c.JSON(http.StatusOK, models.APIResponse{Message: "OK"})
}

func (eh *EmployeeHandler) DeleteEmployee(c *gin.Context) {

	empID := c.Query("id")

	if err := eh.emp.DeleteEmployee(empID); err != nil {
		apiError := err.(*apierror.APIError)
		c.AbortWithStatusJSON(apiError.HttpStatusCode, apiError)
		return
	}

	c.JSON(http.StatusOK, models.APIResponse{Message: "OK"})
}
