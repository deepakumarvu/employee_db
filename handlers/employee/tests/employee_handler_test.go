package employee_test

import (
	"employee/pkg/testhelpers"
	"employee/service/router"
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Employee", func() {
	var r *gin.Engine = router.NewRouter()

	When("GET /employee", func() {
		It("returns 200 OK", func() {
			req, _ := http.NewRequest("GET", "/employee", nil)
			res := testhelpers.TestHTTPResponse(r, req, func(w *httptest.ResponseRecorder) bool {
				// Test that the http status code is 200
				statusOK := w.Code == http.StatusOK
				return statusOK
			})
			Expect(res).To(Equal(true))
		})
	})

	When("GET /employee?id=1", func() {
		It("returns 400", func() {
			req, _ := http.NewRequest("GET", "/employee?id=1", nil)
			res := testhelpers.TestHTTPResponse(r, req, func(w *httptest.ResponseRecorder) bool {
				// Test that the http status code is 200
				statusOK := w.Code == http.StatusOK
				return statusOK
			})
			Expect(res).To(Equal(false))
		})
	})
})
