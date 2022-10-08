package dto

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"log"
	"time"
)

const (
	ISO8661 = "2006-01-02T15:04:05Z"
)

type SystemItemType string

type SystemItem struct {
	Id       string         `binding:"required"`
	Url      string         `binding:"-"`
	Date     time.Time      `binding:"required"`
	ParentId string         `binding:"-"`
	Type     SystemItemType `binding:"required"`
	Size     int            `binding:"-"`
	Children []SystemItem   `binding:"-"`
}

type SystemItemImport struct {
	Id       string `json:"id" binding:"required"`
	Url      string `json:"url" binding:"-"`
	ParentId string `json:"parent-id" binding:"-"`
	Type     string `json:"type" binding:"required"`
	Size     int    `json:"size" binding:"-"`
}

type SystemItemImportRequest struct {
	Items      []SystemItemImport `json:"items" binding:"required,systemItemImport"`
	UpdateDate string             `json:"update-date" binding:"required,iso8601"`
}

type SystemItemHistoryUnit struct {
	Id       string         `binding:"required"`
	Url      string         `binding:"-"`
	ParentId string         `binding:"-"`
	Type     SystemItemType `binding:"required"`
	Size     int            `binding:"-"`
	Date     time.Time      `binding:"required"`
}

type SystemItemHistoryResponse struct {
	Items []SystemItemHistoryUnit `json:"items"`
}

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

var iso8601 validator.Func = func(fl validator.FieldLevel) bool {
	_, err := time.Parse(ISO8661, fl.Field().String())
	return err == nil
}

var systemItemImport validator.Func = func(fl validator.FieldLevel) bool {
	systemItemImports, ok := fl.Field().Interface().([]SystemItemImport)
	if !ok {
		return false
	}
	for i := 0; i < len(systemItemImports); i++ {
		if systemItemImports[i].Type != "FILE" && systemItemImports[i].Type != "FOLDER" {
			return false
		}
	}
	return true
}

func init() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		if err := v.RegisterValidation("iso8601", iso8601); err != nil {
			log.Fatalf("cannot register validation function:\n%s", err.Error())
		}
		if err := v.RegisterValidation("systemItemImport", systemItemImport); err != nil {
			log.Fatalf("cannot register validation function:\n%s", err.Error())
		}
	}
}
