package repo

import (
	"config-service/domain/repository"

	"github.com/go-pg/pg/v10"
)

type Repositories interface {
	SystemConfiguration() repository.SystemConfigurationRepository
}

type RepositoriesImpl struct {
	systemConfigurationRepo repository.SystemConfigurationRepository
}

func (r *RepositoriesImpl) SystemConfiguration() repository.SystemConfigurationRepository {
	return r.systemConfigurationRepo
}

func InitRepositories(db *pg.DB) Repositories {
	return &RepositoriesImpl{
		systemConfigurationRepo: NewSystemConfigurationRepository(db),
	}
}
