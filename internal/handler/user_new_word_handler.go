package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"imxy.top/wordserver/internal/config"
	"imxy.top/wordserver/internal/models"
	"imxy.top/wordserver/internal/tool"
)

func ListNewWords(w http.ResponseWriter, r *http.Request) {
	page, pageSize := tool.ParsePageParams(r)
	nick := chi.URLParam(r, "nick")
	var newWords []models.UserNewWord
	err := config.DB.Where("nick = ?", nick).Limit(pageSize).Offset((page - 1) * pageSize).Find(&newWords).Error
	if err != nil {
		tool.JSONResponse(w, nil, http.StatusInternalServerError)
		return
	}
	if newWords == nil {
		tool.JSONResponse(w, nil, http.StatusNotFound)
		return
	}
	var total int64
	_ = config.DB.Where("nick = ?", nick).Model(&models.UserNewWord{}).Count(&total).Error
	respData := map[string]interface{}{
		"list": newWords,
		"pagination": map[string]int64{
			"page":        int64(page),
			"page_size":   int64(pageSize),
			"total":       total,
			"total_pages": (total + int64(pageSize) - 1) / int64(pageSize), // 总页数
		},
	}
	tool.JSONResponse(w, respData)
}
