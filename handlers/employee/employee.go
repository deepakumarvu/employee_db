package employee

import (
	"employee/logic/employee"
	"employee/models"
	"employee/pkg/apierror"
	"employee/pkg/logger"
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

	logger.Log.Info().Str("method", "CreateEmployee").Msg("Request received")

	var employee models.Employee

	if err := c.BindJSON(&employee); err != nil {
		logger.Log.Error().Err(err).Msg("Failed to parse the request body")
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if err := eh.emp.CreateEmployee(employee); err != nil {
		apiError := err.(*apierror.APIError)
		logger.Log.Error().Err(err).Msg("Failed to create employee")
		c.AbortWithStatusJSON(apiError.HttpStatusCode, apiError)
		return
	}

	logger.Log.Info().Str("method", "CreateEmployee").Msg("Request processed successfully")
	c.JSON(http.StatusOK, models.APIResponse{Message: "OK"})
}

func (eh *EmployeeHandler) GetEmployee(c *gin.Context) {
	logger.Log.Info().Str("method", "GetEmployee").Msg("Request received")

	empID := c.Query("id")
	LastEvalKeyID := c.Query("last_eval_id")
	numRecords := c.Query("num_records")
	var res models.GetEmployeeResponse
	var err error
	if res, err = eh.emp.GetEmployee(empID, LastEvalKeyID, numRecords); err != nil {
		apiError := err.(*apierror.APIError)
		logger.Log.Error().Err(err).Msg("Failed to get employee")
		c.AbortWithStatusJSON(apiError.HttpStatusCode, apiError)
		return
	}
	logger.Log.Info().Str("method", "GetEmployee").Msg("Request processed successfully")
	c.JSON(http.StatusOK, res)
}

func (eh *EmployeeHandler) UpdateEmployee(c *gin.Context) {

	logger.Log.Info().Str("method", "UpdateEmployee").Msg("Request received")

	empID := c.Query("id")

	var employee models.EmployeeUpdateRequest

	if err := c.BindJSON(&employee); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if err := eh.emp.UpdateEmployee(empID, employee); err != nil {
		apiError := err.(*apierror.APIError)
		logger.Log.Error().Err(err).Msg("Failed to update employee")
		c.AbortWithStatusJSON(apiError.HttpStatusCode, apiError)
		return
	}

	logger.Log.Info().Str("method", "UpdateEmployee").Msg("Request processed successfully")

	c.JSON(http.StatusOK, models.APIResponse{Message: "OK"})
}

func (eh *EmployeeHandler) DeleteEmployee(c *gin.Context) {

	logger.Log.Info().Str("method", "DeleteEmployee").Msg("Request received")

	empID := c.Query("id")

	if err := eh.emp.DeleteEmployee(empID); err != nil {
		logger.Log.Error().Err(err).Msg("Failed to delete employee")
		apiError := err.(*apierror.APIError)
		c.AbortWithStatusJSON(apiError.HttpStatusCode, apiError)
		return
	}

	logger.Log.Info().Str("method", "DeleteEmployee").Msg("Request processed successfully")
	c.JSON(http.StatusOK, models.APIResponse{Message: "OK"})
}
