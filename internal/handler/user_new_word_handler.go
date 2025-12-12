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
	var wordIds []uint
	for _, uw := range newWords {
		wordIds = append(wordIds, uw.WordID) // 修正：取关联的单词ID
	}

	var words []models.Word
	err = config.DB.Where("id IN ?", wordIds).Find(&words).Error
	if err != nil {
		tool.JSONResponse(w, nil, http.StatusInternalServerError)
		return
	}

	wordMap := make(map[uint]models.Word)
	for _, word := range words {
		wordMap[word.ID] = word
	}
	var userNewWordVOs []models.UserNewWordVO
	for _, uw := range newWords {
		// 获取关联的单词（处理单词不存在的边界情况）
		word, ok := wordMap[uw.WordID]
		if !ok {
			word = models.Word{} // 单词不存在时赋空结构体
		}

		// 赋值VO的所有字段（严格匹配你定义的UserNewWordVO）
		userNewWordVOs = append(userNewWordVOs, models.UserNewWordVO{
			ID:         uw.ID,         // UserNewWord的ID
			WordID:     uw.WordID,     // 关联的单词ID
			Nick:       uw.Nick,       // 用户昵称
			IsMastered: uw.IsMastered, // 掌握状态
			Img:        word.Img,      // 单词图片
			Word:       word.Word,     // 单词本身
			Phonetic:   word.Phonetic, // 音标
			Mean:       word.Mean,     // 释义
			Sound:      word.Sound,    // 发音文件
			CreatedAt:  uw.CreatedAt,  // 生词添加时间
		})
	}
	var total int64
	err = config.DB.Model(&models.UserNewWord{}).Where("nick = ?", nick).Count(&total).Error
	if err != nil {
		tool.JSONResponse(w, nil, http.StatusInternalServerError)
		return
	}

	respData := map[string]interface{}{
		"list": userNewWordVOs,
		"pagination": map[string]int64{
			"page":        int64(page),
			"page_size":   int64(pageSize),
			"total":       total,
			"total_pages": (total + int64(pageSize) - 1) / int64(pageSize),
		},
	}
	tool.JSONResponse(w, respData)
}
