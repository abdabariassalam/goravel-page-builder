package api

import (
	"encoding/json"
	"page_builder/app/models"

	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
)

type PageController struct {
	//Dependent services
}

func NewPageController() *PageController {
	return &PageController{
		//Inject services
	}
}

func (r *PageController) LoadProjectAPI(ctx http.Context) http.Response {
	var page models.Page
	facades.Orm().Query().Where("id", ctx.Request().Route("id")).FirstOrFail(&page)
	var jsonMap map[string]interface{}
	json.Unmarshal([]byte(page.Content), &jsonMap)
	return ctx.Response().Json(
		200,
		map[string]any{
			"success": true,
			"data":    jsonMap,
			"messge":  "Success load page content",
		})
}
func (r *PageController) StoreProjectAPI(ctx http.Context) http.Response {
	var (
		request map[string]any
		page    models.Page
	)
	ctx.Request().Bind(&request)
	jsonStr, _ := json.Marshal(request["data"])
	facades.Orm().Query().Where("id", ctx.Request().Route("id")).FirstOrFail(&page)
	page.Content = string(jsonStr)
	facades.Orm().Query().Save(&page)
	return ctx.Response().Json(
		200,
		map[string]any{
			"success": true,
			"data":    page.Content,
			"messge":  "Project stored successfully",
		})
}
