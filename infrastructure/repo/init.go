package repo

import (
	"github.com/go-pg/pg/v10"
)

type Repositories interface {
}

type RepositoriesImpl struct {
}

func InitRepositories(db *pg.DB) Repositories {
	return &RepositoriesImpl{}
}
