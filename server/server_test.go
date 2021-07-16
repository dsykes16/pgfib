package server_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strconv"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/dsykes16/pgfib/server"
)

type MockFibonacci struct {
	size int
	err  error
}

func (f *MockFibonacci) GetFib(n int) (result string, err error) {
	return strconv.Itoa(n), f.err
}
func (f *MockFibonacci) GetCacheSize() (size int, err error) {
	return f.size, f.err
}
func (f *MockFibonacci) GetCacheSizeLT(end int) (size int, err error) {
	return end, f.err
}
func (f *MockFibonacci) ClearCache() error {
	return f.err
}

var _ = Describe("Server", func() {
	Describe("GET /fib/[n]", func() {
		It("returns code 200 and a fibonacci number given no backend errors", func() {
			server := server.New(&MockFibonacci{})
			request, _ := http.NewRequest(http.MethodGet, "/fib/5", nil)

			response := httptest.NewRecorder()

			server.ServeHTTP(response, request)

			Expect(response.Result().StatusCode).To(Equal(200))
			Expect(response.Body.String()).To(Equal("5"))
		})
	})

	Describe("GET /size", func() {
		It("returns code 200 and the cache size", func() {
			server := server.New(&MockFibonacci{size: 42})
			request, _ := http.NewRequest(http.MethodGet, "/size", nil)

			response := httptest.NewRecorder()

			server.ServeHTTP(response, request)

			Expect(response.Result().StatusCode).To(Equal(200))
			Expect(response.Body.String()).To(Equal("42"))
		})
	})

	Describe("GET /size/[n]", func() {
		It("returns code 200 and the number of cache entries for all indicies less than N", func() {
			server := server.New(&MockFibonacci{size: 42})
			request, _ := http.NewRequest(http.MethodGet, "/size/7", nil)

			response := httptest.NewRecorder()

			server.ServeHTTP(response, request)

			Expect(response.Result().StatusCode).To(Equal(200))
			Expect(response.Body.String()).To(Equal("7"))
		})
	})

	Describe("DELETE /cache", func() {
		It("return code 200 given no errors from backend", func() {
			server := server.New(&MockFibonacci{err: nil})
			request, _ := http.NewRequest(http.MethodDelete, "/cache", nil)

			response := httptest.NewRecorder()

			server.ServeHTTP(response, request)

			Expect(response.Result().StatusCode).To(Equal(200))
		})
		It("returns code 500 given a backend error", func() {
			server := server.New(&MockFibonacci{err: errors.New("fake backend error")})
			request, _ := http.NewRequest(http.MethodDelete, "/cache", nil)

			response := httptest.NewRecorder()

			server.ServeHTTP(response, request)

			Expect(response.Result().StatusCode).To(Equal(500))
		})
	})

})
