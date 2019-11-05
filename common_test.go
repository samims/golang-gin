package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var tmpArticleList [] article

// this function is used for setup before executing the test functions

func TestMain(m *testing.M) {
	// set gin to test mode
	gin.SetMode(gin.TestMode)

	// run the other test
	os.Exit(m.Run())
}

// helper function to create a router during testing

func getRouter(withTemplates bool) *gin.Engine {
	r := gin.Default()
	if withTemplates {
		r.LoadHTMLGlob("templates/*")
	}
	return r
}

// helper functions to process a request and test it's response

func testHTTPResponse(t *testing.T, r *gin.Engine, req *http.Request, f func(w *httptest.ResponseRecorder) bool) {
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	if !f(w) {
		t.Fail()
	}
}

func saveLists() {
	tmpArticleList = articleList
}

func restoreLists() {
	articleList = tmpArticleList

}
