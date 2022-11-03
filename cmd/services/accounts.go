package services

import (
	accountshandler "metroanno-api/app/accounts/http"
	accountsrepo "metroanno-api/app/accounts/repo"
	accountsusecase "metroanno-api/app/accounts/usecase"
	"metroanno-api/infrastructure/config"
	"metroanno-api/infrastructure/database"
)

func RegisterServiceAccounts(db database.Connection, cfg config.Config) accountshandler.AccountHandler {
	accountsDB := accountsrepo.New(accountsrepo.AccountsRepo{
		MySQL: db,
		Cfg:   cfg,
	})

	accountsApp := accountsusecase.New(accountsusecase.AccountsApp{
		AccountsRepo: accountsDB,
		Cfg:          cfg,
	})

	accountHandler := accountshandler.AccountHandler{
		AccountsApp: accountsApp,
		Cfg:         cfg,
	}

	return accountHandler
}
