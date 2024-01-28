package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func (s *Server) run() {

	router := mux.NewRouter()
	router.HandleFunc("/blog", makeHandlerFunc(s.handleBlogs))
	router.HandleFunc("/blog/{id}", makeHandlerFunc(s.handleBlog))

	http.ListenAndServe(s.address, router)

}
func (s *Server) handleBlogs(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "POST" {
		return s.createBlog(w, r)
	} else if r.Method == "GET" {
		return s.listBlogs(w, r)
	}
	return fmt.Errorf("method not allowed %s", r.Method)
}

func (s *Server) handleBlog(w http.ResponseWriter, r *http.Request) error {
	id := mux.Vars(r)["id"]
	fmt.Println("blog id ", id)

	if r.Method == "PUT" {
		return s.updateBlog(w, r)
	} else if r.Method == "DELETE" {
		return s.deleteBlog(w, r)
	} else if r.Method == "GET" {
		return s.getBlog(w, r)
	}
	return fmt.Errorf("method not allowed %s", r.Method)
}

func (s *Server) listBlogs(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *Server) createBlog(w http.ResponseWriter, r *http.Request) error {

	return WriteJson(w, http.StatusOK, &Blog{})
}

func (s *Server) getBlog(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *Server) deleteBlog(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *Server) updateBlog(w http.ResponseWriter, r *http.Request) error {
	return nil
}

type Server struct {
	address string
}

func WriteJson(w http.ResponseWriter, status int, v any) error {
	w.WriteHeader(status)
	w.Header().Add("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(v)
}

type ApiFunc func(http.ResponseWriter, *http.Request) error

type ApiError struct {
	Error string
}

func makeHandlerFunc(f ApiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			WriteJson(w, http.StatusBadRequest, ApiError{Error: err.Error()})
		}
	}
}

func makeServer(address string) *Server {
	return &Server{
		address: address,
	}
}
