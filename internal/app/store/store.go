package store

import (
	"database/sql"
	_ "github.com/lib/pq" //drivers
)

//Store ...
type Store struct {
	config         *Config
	db             *sql.DB
	userRepository *UserRepository
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

//User method which returns user repository to work with users table
func (s *Store) User() *UserRepository {
	//with the help of this method now we can deal with users by s.User().Create(<model>)
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &UserRepository{
		store: s,
	}

	return s.userRepository
}
