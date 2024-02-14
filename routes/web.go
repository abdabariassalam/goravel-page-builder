package routes

import (
	"page_builder/app/http/controllers"

	"github.com/goravel/framework/facades"
)

func Web() {
	pageController := controllers.NewPageController()
	facades.Route().Static("assets", "./public")
	facades.Route().Get("/", pageController.Index)
	facades.Route().Get("/pages/create", pageController.Create)
	facades.Route().Get("/pages/{id}/edit", pageController.Edit)
	facades.Route().Resource("/pages", pageController)
}
