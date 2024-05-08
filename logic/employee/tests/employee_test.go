package employee_test

import (
	"employee/logic/employee"
	"employee/models"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Employee", func() {
	var (
		eh *employee.Employee
	)

	BeforeEach(func() {
		eh = employee.NewEmployee()
	})

	Context("CreateEmployee function", func() {
		It("should return no error when given a valid employee", func() {
			// given
			emp := models.Employee{
				ID:       1,
				Name:     "John Doe",
				Position: "Software Engineer",
				Salary:   10000,
			}

			// when
			err := eh.CreateEmployee(emp)

			// then
			Expect(err).To(BeNil())
		})

		It("should return an error when given an invalid employee ID", func() {
			// given
			emp := models.Employee{
				ID:       0,
				Name:     "John Doe",
				Position: "Software Engineer",
				Salary:   10000,
			}

			// when
			err := eh.CreateEmployee(emp)

			// then
			Expect(err).To(Equal(employee.GetEmpError(employee.InvalidID)))
		})

		It("should return an error when given an empty employee name", func() {
			// given
			emp := models.Employee{
				ID:       1,
				Name:     "",
				Position: "Software Engineer",
				Salary:   10000,
			}

			// when
			err := eh.CreateEmployee(emp)

			// then
			Expect(err).To(Equal(employee.GetEmpError(employee.NameInvalid)))
		})

		It("should return an error when given an invalid employee name", func() {
			// given
			emp := models.Employee{
				ID:       1,
				Name:     "ThisNameIsWayTooLongAndInvalidThisNameIsWayTooLongAndInvalidThisNameIsWayTooLongAndInvalidThisNameIsWayTooLongAndInvalid",
				Position: "Software Engineer",
				Salary:   10000,
			}

			// when
			err := eh.CreateEmployee(emp)

			// then
			Expect(err).To(Equal(employee.GetEmpError(employee.NameInvalid)))
		})

		It("should return an error when given an empty employee position", func() {
			// given
			emp := models.Employee{
				ID:       1,
				Name:     "John Doe",
				Position: "",
				Salary:   10000,
			}

			// when
			err := eh.CreateEmployee(emp)

			// then
			Expect(err).To(Equal(employee.GetEmpError(employee.InvalidPosition)))
		})

		It("should return an error when given an invalid employee position", func() {
			// given
			emp := models.Employee{
				ID:       1,
				Name:     "John Doe",
				Position: "ThisNameIsWayTooLongAndInvalidThisNameIsWayTooLongAndInvalidThisNameIsWayTooLongAndInvalidThisNameIsWayTooLongAndInvalid",
				Salary:   10000,
			}

			// when
			err := eh.CreateEmployee(emp)

			// then
			Expect(err).To(Equal(employee.GetEmpError(employee.InvalidPosition)))
		})

		It("should return an error when given an invalid employee salary", func() {
			// given
			emp := models.Employee{
				ID:       1,
				Name:     "John Doe",
				Position: "Software Engineer",
				Salary:   0,
			}

			// when
			err := eh.CreateEmployee(emp)

			// then
			Expect(err).To(Equal(employee.GetEmpError(employee.InvalidSalary)))
		})

		It("should return an error when the employee already exists", func() {
			// Create a valid employee object
			emp := models.Employee{
				ID:       1,
				Name:     "John Doe",
				Position: "Software Engineer",
				Salary:   10000,
			}

			// Call the CreateEmployee function
			err := eh.CreateEmployee(emp)
			Expect(err).To(BeNil())
			err = eh.CreateEmployee(emp)
			// Check if the function returned the correct error
			Expect(err).To(Equal(employee.GetEmpError(employee.EmpAlreadyExists)))
		})

		It("should return an error when the employee has a negative salary", func() {
			// Create an employee object with a negative salary
			emp := models.Employee{
				ID:       1,
				Name:     "John Doe",
				Position: "Software Engineer",
				Salary:   -1000,
			}

			// Call the CreateEmployee function
			err := eh.CreateEmployee(emp)

			// Check if the function returned the correct error
			Expect(err).To(Equal(employee.GetEmpError(employee.InvalidSalary)))
		})

	})

	Context("GetEmployee function", func() {
		It("should return the employee with the given ID when empID is not empty", func() {
			// given
			empID := "1"
			LastEvalKeyID := ""
			numRecords := ""
			expected := models.GetEmployeeResponse{
				Employees: []models.Employee{
					{ID: 1, Name: "John Doe", Position: "Developer", Salary: 50000},
				},
			}
			eh.CreateEmployee(models.Employee{ID: 1, Name: "John Doe", Position: "Developer", Salary: 50000})

			// when
			res, err := eh.GetEmployee(empID, LastEvalKeyID, numRecords)

			// then
			Expect(err).To(BeNil())
			Expect(res).To(Equal(expected))
		})

		It("should return the employees starting from the specified LastEvalKeyID and with the specified number of records when empID is empty and numRecords is not empty", func() {
			// given
			empID := ""
			LastEvalKeyID := "2"
			numRecords := "3"
			emps := []models.Employee{
				{ID: 2, Name: "Jane Doe", Position: "Manager", Salary: 80000},
				{ID: 3, Name: "Bob Smith", Position: "Designer", Salary: 70000},
				{ID: 4, Name: "Alice Johnson", Position: "Engineer", Salary: 60000},
			}
			expected := models.GetEmployeeResponse{
				Employees: emps[1:],
			}
			for _, emp := range emps {
				eh.CreateEmployee(emp)
			}

			// when
			res, err := eh.GetEmployee(empID, LastEvalKeyID, numRecords)

			// then
			Expect(err).To(BeNil())
			Expect(res).To(Equal(expected))
		})

		It("should return the employees starting from the specified LastEvalKeyID and with the specified number of records when empID is empty and numRecords is not empty - no items left", func() {
			// given
			empID := ""
			LastEvalKeyID := "4"
			numRecords := "3"
			emps := []models.Employee{
				{ID: 4, Name: "Alice Johnson", Position: "Engineer", Salary: 60000},
			}
			expected := models.GetEmployeeResponse{
				Employees: []models.Employee{},
			}
			for _, emp := range emps {
				eh.CreateEmployee(emp)
			}

			// when
			res, err := eh.GetEmployee(empID, LastEvalKeyID, numRecords)

			// then
			Expect(err).To(BeNil())
			Expect(res).To(Equal(expected))
		})

		It("should return the employees starting from the oldest ID with a default number of records (10) when empID and numRecords are empty", func() {
			// given
			empID := ""
			LastEvalKeyID := ""
			numRecords := ""
			expected := models.GetEmployeeResponse{
				Employees: []models.Employee{
					{ID: 1, Name: "John Doe", Position: "Developer", Salary: 50000},
					{ID: 2, Name: "Jane Doe", Position: "Manager", Salary: 80000},
					{ID: 3, Name: "Bob Smith", Position: "Designer", Salary: 70000},
					{ID: 4, Name: "Alice Johnson", Position: "Engineer", Salary: 60000},
					{ID: 5, Name: "Mark Brown", Position: "Analyst", Salary: 55000},
					{ID: 6, Name: "Sarah Lee", Position: "Architect", Salary: 75000},
					{ID: 7, Name: "Chris Evans", Position: "Coordinator", Salary: 45000},
					{ID: 8, Name: "Emma Watson", Position: "Manager", Salary: 82000},
					{ID: 9, Name: "Tom Hardy", Position: "Developer", Salary: 48000},
					{ID: 10, Name: "Olivia Williams", Position: "Designer", Salary: 69000},
				},
				LastEvalKeyID: 10,
			}
			for _, emp := range expected.Employees {
				eh.CreateEmployee(emp)
			}

			// when
			res, err := eh.GetEmployee(empID, LastEvalKeyID, numRecords)

			// then
			Expect(err).To(BeNil())
			Expect(res.Employees).To(Equal(expected.Employees))
		})

		It("should return an error with InvalidID when empID is not a valid integer", func() {
			// given
			empID := "invalid"
			startID := ""
			numRecords := ""

			// when
			_, err := eh.GetEmployee(empID, startID, numRecords)

			// then
			Expect(err).To(Equal(employee.GetEmpError(employee.InvalidID)))
		})

		It("should return an error with InvalidID when empID is not a valid integer", func() {
			// given
			empID := "1"
			startID := ""
			numRecords := ""

			// when
			_, err := eh.GetEmployee(empID, startID, numRecords)

			// then
			Expect(err).To(Equal(employee.GetEmpError(employee.InvalidID)))
		})

		It("should return an error with InvalidLastEvalKeyID when startID is not a valid integer", func() {
			// given
			empID := ""
			startID := "invalid"
			numRecords := ""

			// when
			_, err := eh.GetEmployee(empID, startID, numRecords)

			// then
			Expect(err).To(Equal(employee.GetEmpError(employee.InvalidLastEvalKeyID)))
		})

		It("should return an error with InvalidLastEvalKeyID when startID is not valid", func() {
			// given
			empID := ""
			startID := "2"
			numRecords := ""

			// when
			_, err := eh.GetEmployee(empID, startID, numRecords)

			// then
			Expect(err).To(Equal(employee.GetEmpError(employee.InvalidLastEvalKeyID)))
		})

	})

	Context("Update employee", func() {
		When("the employee ID is invalid", func() {
			It("should return an error with InvalidID", func() {
				// given
				empID := ""
				empUpdateReq := models.EmployeeUpdateRequest{}

				// when
				err := eh.UpdateEmployee(empID, empUpdateReq)

				// then
				Expect(err).To(Equal(employee.GetEmpError(employee.InvalidID)))
			})
		})

		When("the employee update request is invalid", func() {
			It("should return an error with InvalidEmpUpdate", func() {
				// given
				empID := "1"
				empUpdateReq := models.EmployeeUpdateRequest{
					Position: "",
					Salary:   0,
				}

				// when
				err := eh.UpdateEmployee(empID, empUpdateReq)

				// then
				Expect(err).To(Equal(employee.GetEmpError(employee.InvalidEmpUpdate)))
			})
		})

		When("the salary is invalid", func() {
			It("should return an error with InvalidSalary", func() {
				// given
				empID := "1"
				empUpdateReq := models.EmployeeUpdateRequest{
					Position: "Manager",
					Salary:   -100,
				}

				// when
				err := eh.UpdateEmployee(empID, empUpdateReq)

				// then
				Expect(err).To(Equal(employee.GetEmpError(employee.InvalidSalary)))
			})
		})

		When("the employee is not found in the database", func() {
			It("should return an error with InvalidID", func() {
				// given
				empID := "999"
				empUpdateReq := models.EmployeeUpdateRequest{
					Position: "Manager",
					Salary:   50000,
				}

				// when
				err := eh.UpdateEmployee(empID, empUpdateReq)

				// then
				Expect(err).To(Equal(employee.GetEmpError(employee.InvalidID)))
			})
		})

		When("the employee is not found in the database", func() {
			It("should return an error with InvalidID", func() {
				// given
				empID := "invalid"
				empUpdateReq := models.EmployeeUpdateRequest{
					Position: "Manager",
					Salary:   50000,
				}

				// when
				err := eh.UpdateEmployee(empID, empUpdateReq)

				// then
				Expect(err).To(Equal(employee.GetEmpError(employee.InvalidID)))
			})
		})

		When("updating the employee position", func() {
			It("should update the employee position in the database", func() {
				// given
				empID := "1"
				empUpdateReq := models.EmployeeUpdateRequest{
					Position: "Manager",
				}
				eh.CreateEmployee(models.Employee{
					ID:       1,
					Name:     "John Doe",
					Position: "Software Engineer",
					Salary:   10000,
				})

				// when
				err := eh.UpdateEmployee(empID, empUpdateReq)

				// then
				Expect(err).To(BeNil())
				emp, err := eh.GetEmployee(empID, "", "")
				Expect(err).To(BeNil())
				Expect(emp.Employees[0].Position).To(Equal("Manager"))
			})
		})

		When("updating the employee Salary", func() {
			It("should update the employee salary in the database", func() {
				// given
				empID := "1"
				empUpdateReq := models.EmployeeUpdateRequest{
					Salary: 20000,
				}
				eh.CreateEmployee(models.Employee{
					ID:       1,
					Name:     "John Doe",
					Position: "Software Engineer",
					Salary:   10000,
				})

				// when
				err := eh.UpdateEmployee(empID, empUpdateReq)

				// then
				Expect(err).To(BeNil())
				emp, err := eh.GetEmployee(empID, "", "")
				Expect(err).To(BeNil())
				Expect(emp.Employees[0].Salary).To(Equal(float64(20000)))
			})
		})
	})

	Context("Delete Employee", func() {

		When("the employee ID is empty", func() {
			It("should return an error with InvalidID", func() {
				// given
				empID := ""

				// when
				err := eh.DeleteEmployee(empID)

				// then
				Expect(err).To(Equal(employee.GetEmpError(employee.InvalidID)))
			})
		})

		When("the employee ID is invalid", func() {
			It("should return an error with InvalidID", func() {
				// given
				empID := "invalid"

				// when
				err := eh.DeleteEmployee(empID)

				// then
				Expect(err).To(Equal(employee.GetEmpError(employee.InvalidID)))
			})
		})

		When("the employee is not found in the database", func() {
			It("should return an error with InvalidID", func() {
				// given
				empID := "999"

				// when
				err := eh.DeleteEmployee(empID)

				// then
				Expect(err).To(Equal(employee.GetEmpError(employee.InvalidID)))
			})
		})

		When("the employee is found in the database", func() {
			It("should delete the employee from the database", func() {
				// given
				empID := "1"
				eh.CreateEmployee(models.Employee{
					ID:       1,
					Name:     "John Doe",
					Position: "Software Engineer",
					Salary:   10000,
				})

				// when
				err := eh.DeleteEmployee(empID)

				// then
				Expect(err).To(BeNil())
			})
		})
	})

})
