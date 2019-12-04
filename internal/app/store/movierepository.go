package store

import (
	"fmt"
	"github.com/Pupye/movie-must-watch/model"
)

type MovieRepository struct {
	store *Store
}

func (r *MovieRepository) Create(m *model.Movie) (*model.Movie, error) {
	if m == nil {
		fmt.Println("problem here should not be nil")
		return nil, nil
	}
	if err := r.store.db.QueryRow(
		"INSERT INTO mydb.public.movies (title, releasedate) VALUES ($1, $2) RETURNING id",
		m.Title,
		m.ReleaseDate,
		).Scan(&m.ID); err != nil{

			return nil, err
	}
	return m, nil
}

func (r *MovieRepository) findByTitle(title string) (*model.Movie, error){
	m := &model.Movie{}
	err := r.store.db.QueryRow(
		"SELECT id, title, releasedate FROM mydb.public.movies WHERE title=$1",
		title,
		).Scan(&m)

	if err != nil {
		return nil, err
	}

	return m, nil
}