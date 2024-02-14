package models

import (
	"github.com/google/uuid"
	"github.com/goravel/framework/database/orm"
)

type Page struct {
	ID               uuid.UUID `gorm:"primaryKey"`
	Title            string
	ShortDescription string
	TemplateType     string
	Content          string
	orm.SoftDeletes
	orm.Timestamps
}
