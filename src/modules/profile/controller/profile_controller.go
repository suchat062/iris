package controller

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/sessions"
	"iris/src/modules/profile/usecase"
)

type ProfileController struct {
	Ctx iris.Context
	Session *sessions.Session
	ProfileUsecase usecase.ProfileUsecase
}