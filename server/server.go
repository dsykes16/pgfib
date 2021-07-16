package server

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/dsykes16/pgfib/fibonacci"
	"gotest.tools/gotestsum/log"
)

type FibServer struct {
	fib fibonacci.Fibonacci
	http.Handler
}

func New(fib fibonacci.Fibonacci) *FibServer {
	s := new(FibServer)
	s.fib = fib

	router := http.NewServeMux()
	router.Handle("/fib/", http.HandlerFunc(s.fibHandler))
	router.Handle("/cache", http.HandlerFunc(s.resetHandler))
	router.Handle("/size/", http.HandlerFunc(s.sizeRangeHandler))
	router.Handle("/size", http.HandlerFunc(s.sizeHandler))
	s.Handler = router

	return s
}

func (s *FibServer) fibHandler(w http.ResponseWriter, r *http.Request) {
	index, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/fib/"))
	if err != nil {
		w.WriteHeader(500)
		errmsg := fmt.Sprintf("could not parse request: %s", err)
		log.Error(errmsg)
		fmt.Fprint(w, errmsg)
		return
	}

	val, err := s.fib.GetFib(index)
	if err != nil {
		w.WriteHeader(500)
		errmsg := fmt.Sprintf("could not calculate fibonacci: %s", err)
		log.Error(errmsg)
		fmt.Fprint(w, errmsg)
		return
	}
	fmt.Fprintf(w, "%s", val)
}

func (s *FibServer) resetHandler(w http.ResponseWriter, r *http.Request) {
	err := s.fib.ClearCache()
	if err != nil {
		w.WriteHeader(500)
		errmsg := fmt.Sprintf("could not clear cache: %s", err)
		log.Error(errmsg)
		fmt.Fprint(w, errmsg)
		return
	}
	w.WriteHeader(200)
}

func (s *FibServer) sizeHandler(w http.ResponseWriter, r *http.Request) {
	size, err := s.fib.GetCacheSize()
	if err != nil {
		w.WriteHeader(500)
		errmsg := fmt.Sprintf("could not get cache size: %s", err)
		log.Error(errmsg)
		fmt.Fprint(w, errmsg)
		return
	}
	w.WriteHeader(200)
	fmt.Fprintf(w, "%d", size)
}

func (s *FibServer) sizeRangeHandler(w http.ResponseWriter, r *http.Request) {
	end, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/size/"))
	if err != nil {
		w.WriteHeader(500)
		errmsg := fmt.Sprintf("could not parse request: %s", err)
		log.Error(errmsg)
		fmt.Fprint(w, errmsg)
		return
	}

	size, err := s.fib.GetCacheSizeLT(end)
	if err != nil {
		w.WriteHeader(500)
		errmsg := fmt.Sprintf("could not get cache size for given range(0, %d): %s", end, err)
		log.Error(errmsg)
		fmt.Fprint(w, errmsg)
		return
	}
	fmt.Fprintf(w, "%d", size)
}
