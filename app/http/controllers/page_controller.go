package controllers

import (
	"os"
	"page_builder/app/models"

	"github.com/google/uuid"
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

type Page struct {
	ID               uuid.UUID      `json:"id"`
	Title            string         `json:"title" form:"title" validate:"required"`
	ShortDescription string         `json:"short_description" form:"short_description" validate:"required"`
	Content          map[string]any `json:"content" form:"content" validate:"required"`
	Template         *string        `json:"template" form:"template"`
}

// GET /
func (r *PageController) Show(ctx http.Context) http.Response {
	var page models.Page
	if err := facades.Orm().Query().Where("id", ctx.Request().Route("id")).FirstOrFail(&page); err != nil {
		return ctx.Response().Json(http.StatusBadRequest, map[string]any{
			"message": err.Error(),
		})
	}
	return ctx.Response().View().Make("preview.html", map[string]any{
		"page": page,
	})
}

// GET /pages/create
func (r *PageController) Create(ctx http.Context) http.Response {
	return ctx.Response().View().Make("create.html", map[string]any{})
}

// POST /pages
func (r *PageController) Store(ctx http.Context) http.Response {
	var (
		req  = Page{}
		file = []byte{}
	)
	ctx.Request().Bind(&req)
	switch *req.Template {
	case "web":
		file, _ = os.ReadFile("./storage/app/public/template2.json")
	case "newspaper":
		file, _ = os.ReadFile("./storage/app/public/template3.json")
	default:
		file, _ = os.ReadFile("./storage/app/public/template1.json")
	}
	if err := facades.Orm().Query().Create(&models.Page{
		ID:               uuid.New(),
		Title:            req.Title,
		ShortDescription: req.ShortDescription,
		TemplateType:     *req.Template,
		Content:          string(file),
	}); err != nil {
		return ctx.Response().Json(http.StatusBadRequest, "Failed Store Data")
	}
	return ctx.Response().Redirect(http.StatusFound, "/")
}

// PUT /
func (r *PageController) Update(ctx http.Context) http.Response {
	var (
		req  = Page{}
		file = []byte{}
		page models.Page
	)
	ctx.Request().Bind(&req)
	if err := facades.Orm().Query().Where("id", ctx.Request().Route("id")).FirstOrFail(&page); err != nil {
		return ctx.Response().Json(http.StatusBadRequest, map[string]any{
			"message": err.Error(),
		})
	}
	if page.Title == req.Title && page.ShortDescription == req.ShortDescription && page.TemplateType == *req.Template {
		return ctx.Response().Success().Json("tidak ada data yang diganti")
	}
	page.Title = req.Title
	page.ShortDescription = req.ShortDescription
	if req.Template != &page.TemplateType {
		switch *req.Template {
		case "web":
			file, _ = os.ReadFile("./storage/app/public/template2.json")
		case "newspaper":
			file, _ = os.ReadFile("./storage/app/public/template3.json")
		default:
			file, _ = os.ReadFile("./storage/app/public/template1.json")
		}
		page.TemplateType = *req.Template
		page.Content = string(file)
	}
	if _, err := facades.Orm().Query().Model(&models.Page{}).Where("id", ctx.Request().Route("id")).Update(&page); err != nil {
		return ctx.Response().Json(http.StatusBadRequest, "Failed to update page")
	}
	return ctx.Response().Success().Json("Page update successfully")
}

// DELETE /
func (r *PageController) Destroy(ctx http.Context) http.Response {
	if _, err := facades.Orm().Query().Delete(&models.Page{}, uuid.MustParse(ctx.Request().Route("id"))); err != nil {
		return ctx.Response().Json(http.StatusBadRequest, "Failed to delete page")
	}
	return ctx.Response().Success().Json("Page deleted successfully")
}

// GET /
func (r *PageController) Index(ctx http.Context) http.Response {
	var pages []models.Page
	if err := facades.Orm().Query().Order("updated_at desc").Get(&pages); err != nil {
		return ctx.Response().Json(http.StatusBadRequest, map[string]any{
			"message": err.Error(),
		})
	}
	return ctx.Response().View().Make("index.html", map[string]any{
		"pages": pages,
	})
}

// GET /pages/{id}/edit
func (r *PageController) Edit(ctx http.Context) http.Response {
	var page models.Page
	if err := facades.Orm().Query().Where("id", ctx.Request().Route("id")).FirstOrFail(&page); err != nil {
		return ctx.Response().Json(http.StatusBadRequest, map[string]any{
			"message": err.Error(),
		})
	}
	return ctx.Response().View().Make("edit.html", map[string]any{
		"page": page,
	})
}
