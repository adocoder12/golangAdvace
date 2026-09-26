package dbmock

import (
	"time"

	"github.com/adocoder12/golangAdvace/model"
)

type Database struct {
	name      string
	movies    []model.Movie
	createdAt time.Time
}
