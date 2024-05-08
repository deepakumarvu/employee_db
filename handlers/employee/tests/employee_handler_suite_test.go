package employee_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestEmployeeHandler(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Employee Handler Suite")
}
