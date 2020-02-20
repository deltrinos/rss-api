package models

type ListParams struct {
	ItemsByPage uint   `form:"itemsByPage" binding:"required"`
	Page        uint   `form:"page"`
	Category    string `form:"category"`
	Provider    string `form:"provider"`
}
