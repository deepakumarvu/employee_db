package testhelpers

import (
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
)

// Helper function to process a request and test its response
func TestHTTPResponse(r *gin.Engine, req *http.Request, f func(w *httptest.ResponseRecorder) bool) bool {

	// Create a response recorder
	w := httptest.NewRecorder()

	// Create the service and process the above request.
	r.ServeHTTP(w, req)

	return f(w)
}
