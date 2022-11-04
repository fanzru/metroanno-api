package services

import (
	accountshandler "metroanno-api/app/annotation/http"
	accountsrepo "metroanno-api/app/annotation/repo"
	accountsusecase "metroanno-api/app/annotation/usecase"
	"metroanno-api/infrastructure/config"
	"metroanno-api/infrastructure/database"
)

func RegisterServiceAnnotations(db database.Connection, cfg config.Config) accountshandler.AnnotationHandler {
	DB := accountsrepo.New(accountsrepo.AnnotationsRepo{
		MySQL: db,
		Cfg:   cfg,
	})

	App := accountsusecase.New(accountsusecase.AnnotationsApp{
		AnnotationsRepo: DB,
		Cfg:             cfg,
	})

	Handler := accountshandler.AnnotationHandler{
		App: App,
		Cfg: cfg,
	}

	return Handler
}
