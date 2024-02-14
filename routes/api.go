package routes

import (
	"github.com/goravel/framework/facades"

	"page_builder/app/http/controllers/api"
)

func Api() {
	userController := api.NewUserController()
	pageController := api.NewPageController()
	facades.Route().Get("/api/users/{id}", userController.Show)
	facades.Route().Get("/api/pages/{id}/load-project", pageController.LoadProjectAPI)
	facades.Route().Patch("/api/pages/{id}/store-project", pageController.StoreProjectAPI)
}
