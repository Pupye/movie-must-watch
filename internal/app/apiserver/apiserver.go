package apiserver

import (
	"encoding/json"
	"fmt"
	"github.com/Pupye/movie-must-watch/internal/app/store"
	"github.com/Pupye/movie-must-watch/model"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
	"net/http"
)

//APIServer ...
type APIServer struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
	store  *store.Store
}

//New ...
func New(config *Config) *APIServer {
	return &APIServer{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

//Start ...
func (s *APIServer) Start() error {
	if err := s.configureLogger(); err != nil {
		return err
	}

	s.configureRouter()

	if err := s.configureStore(); err != nil {
		return err
	}

	s.logger.Info("starting api server")
	return http.ListenAndServe(s.config.BindAddr, s.router)
}

func (s *APIServer) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)

	if err != nil {
		return err
	}

	s.logger.SetLevel(level)
	return nil
}

func (s *APIServer) configureRouter() {
	s.router.HandleFunc("/hello", s.handleHello())
	s.router.HandleFunc("/new-movie", s.handleCreateMovie()).Methods("POST")
}

func (s *APIServer) configureStore() error {
	st := store.New(s.config.Store)

	if err := st.Open(); err != nil {
		return err
	}

	s.store = st

	return nil
}

func (s *APIServer) handleHello() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "fuck yea")
	}
}

func (s *APIServer) handleCreateMovie() http.HandlerFunc {

	return func(writer http.ResponseWriter, request *http.Request) {
		newMovie := &model.Movie{}

		reqBody, err := ioutil.ReadAll(request.Body)
		if err != nil {
			fmt.Println("you should enter proper json")
			writer.WriteHeader(http.StatusBadRequest)
			return
		}

		json.Unmarshal(reqBody, newMovie)
		fmt.Println(newMovie)
		_, err = s.store.Movie().Create(newMovie)

		if err != nil {
			fmt.Println(err)
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}
		writer.WriteHeader(http.StatusCreated)
		json.NewEncoder(writer).Encode(newMovie)
	}
}

