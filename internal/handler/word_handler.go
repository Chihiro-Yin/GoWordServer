package handler

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"imxy.top/wordserver/internal/config"
	"imxy.top/wordserver/internal/models"
	"imxy.top/wordserver/internal/tool"
)

func ListWords(w http.ResponseWriter, r *http.Request) {
	page, pageSize := tool.ParsePageParams(r)
	var words []models.Word
	err := config.DB.Limit(pageSize).Offset((page - 1) * pageSize).Find(&words).Error
	if err != nil {
		tool.JSONResponse(w, nil, http.StatusInternalServerError)
		return
	}
	var total int64
	_ = config.DB.Model(&models.Word{}).Count(&total).Error
	respData := map[string]interface{}{
		"list": words,
		"pagination": map[string]int64{
			"page":        int64(page),
			"page_size":   int64(pageSize),
			"total":       total,
			"total_pages": (total + int64(pageSize) - 1) / int64(pageSize), // 总页数
		},
	}
	tool.JSONResponse(w, respData)
}
func GetWord(w http.ResponseWriter, r *http.Request) {
	wordIdStr := chi.URLParam(r, "wordId")
	wordId, err := strconv.Atoi(wordIdStr)
	if err != nil || wordId < 1 {
		tool.JSONResponse(w, nil, http.StatusBadRequest)
		return
	}
	var word models.Word
	err = config.DB.First(&word, wordId).Error
	if err != nil {
		tool.JSONResponse(w, nil, http.StatusNotFound)
		return
	}
	tool.JSONResponse(w, word)
}
