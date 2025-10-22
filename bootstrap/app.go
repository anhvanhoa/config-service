package bootstrap

import (
	"config-service/infrastructure/repo"

	"github.com/anhvanhoa/service-core/bootstrap/db"
	"github.com/anhvanhoa/service-core/domain/cache"
	"github.com/anhvanhoa/service-core/domain/log"
	"github.com/anhvanhoa/service-core/utils"
	"github.com/go-pg/pg/v10"
	"go.uber.org/zap/zapcore"
)

type Application struct {
	Env    *Env
	DB     *pg.DB
	Log    *log.LogGRPCImpl
	Repo   repo.Repositories
	Helper utils.Helper
	Cacher cache.CacheI
}

func App() *Application {
	env := Env{}
	NewEnv(&env)
	logConfig := log.NewConfig()
	log := log.InitLogGRPC(logConfig, zapcore.DebugLevel, env.IsProduction())
	db := db.NewPostgresDB(db.ConfigDB{
		URL:  env.UrlDb,
		Mode: env.NodeEnv,
	})
	helper := utils.NewHelper()

	cacher := cache.NewCache(cache.ConfigCache{
		Addr:        env.DbCache.Addr,
		DB:          env.DbCache.Db,
		Password:    env.DbCache.Password,
		MaxIdle:     env.DbCache.MaxIdle,
		MaxActive:   env.DbCache.MaxActive,
		IdleTimeout: env.DbCache.IdleTimeout,
	})

	repo := repo.InitRepositories(db, helper)

	return &Application{
		Env:    &env,
		DB:     db,
		Log:    log,
		Repo:   repo,
		Helper: helper,
		Cacher: cacher,
	}
}
