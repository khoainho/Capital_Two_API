package main

import (
	"path/filepath"
	"os"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

// serve single page application
type spaHandler struct {
	staticPath string
	indexPath string
}

func (h spaHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path, err := filepath.Abs(r.URL.Path)
	if err != nil {
		// failed to get absolute path
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	
	path = filepath.Join(h.staticPath, path)
	
	// check if file exists at path
	_, err = os.Stat(path)
	if os.IsNotExist(err) {
		http.ServeFile(w,r,filepath.Join(h.staticPath, h.indexPath))
		return
	} else if err != nil {
		// return 500 error
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.FileServer(http.Dir(h.staticPath)).ServeHTTP(w,r)
}

func WriteJson(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

type APIFunc func(http.ResponseWriter, *http.Request) error 

type ApiError struct {
	Error string
}

func makeHTTPHandleFunc(f APIFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			WriteJson(w, http.StatusBadRequest, ApiError{Error: err.Error()})
		}
	}
}

type APIServer struct {
	listenAddr string
	store 	   Storage
}

func NewAPIServer(listenAddr string, store Storage) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
		store: 		store,
	}
}

func (s *APIServer) Run() {
	router := mux.NewRouter()

	router.HandleFunc("/account", makeHTTPHandleFunc(s.handleAccount))

	router.HandleFunc("/account/{id}", makeHTTPHandleFunc(s.handleGetAccountByID))

	spa := spaHandler{staticPath: "frontend", indexPath: "index.html"}
	router.PathPrefix("/").Handler(spa)

	log.Println("Capital Two API server running on port: ", s.listenAddr)

	log.Fatal(http.ListenAndServe(s.listenAddr, router))
}


func (s *APIServer) handleAccount (w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		return s.handleGetAccount(w, r)
	}
	if r.Method == "POST" {
		return s.handleCreateAccount(w, r)
	}
	if r.Method == "DELETE" {
		return s.handleDeleteAccount(w, r)
	}

	return fmt.Errorf("method not allowed %s", r.Method)
}

func (s *APIServer) handleGetAccount (w http.ResponseWriter, r *http.Request) error {
	accounts, err := s.store.GetAccounts()
	if err != nil {
		return err 
	}

	return WriteJson(w, http.StatusOK, accounts)
}


func (s *APIServer) handleGetAccountByID (w http.ResponseWriter, r *http.Request) error {
	id := mux.Vars(r)["id"]

	fmt.Println(id)
	
	return WriteJson(w, http.StatusOK, &Account{})
}

func (s *APIServer) handleCreateAccount (w http.ResponseWriter, r *http.Request) error {
	CreateAccountReq := new(CreateAccountRequest)
	if err := json.NewDecoder(r.Body).Decode(CreateAccountReq); err != nil {
		return err
	}
	
	account := NewAccount(CreateAccountReq.FirstName, CreateAccountReq.LastName)
	if err := s.store.CreateAccount(account); err != nil {
		return err 
	}

	return WriteJson(w, http.StatusOK, account)
}

func (s *APIServer) handleDeleteAccount (w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIServer) handleTransfer (w http.ResponseWriter, r *http.Request) error {
	return nil
}
