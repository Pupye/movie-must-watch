package store

import (
	"database/sql"
	_ "github.com/lib/pq" //drivers
)

//Store ...
type Store struct {
	config *Config
	db     *sql.DB
	movieRepository *MovieRepository
}

//New ...
func New(config *Config) *Store {
	return &Store{
		config: config,
	}
}

//Open ...
func (s *Store) Open() error {
	db, err := sql.Open("postgres", s.config.DatabaseURL)

	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}

	s.db = db
	return nil
}

//Close ...
func (s *Store) Close() {
	s.db.Close()
}

func (s *Store) Movie() *MovieRepository{
	if s.movieRepository != nil {
		return s.movieRepository
	}
	s.movieRepository = &MovieRepository{
		store: s,
	}

	return s.movieRepository
}