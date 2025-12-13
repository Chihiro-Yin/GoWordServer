package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
	"imxy.top/wordserver/internal/config"
	"imxy.top/wordserver/internal/models"
	"imxy.top/wordserver/internal/tool"
)

func ListNewWords(w http.ResponseWriter, r *http.Request) {
	page, pageSize := tool.ParsePageParams(r)
	var nick string
	nick = r.URL.Query().Get("nick")
	if nick == "" {
		nick = r.Context().Value("nick").(string)
	}
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
func isDuplicateWordError(err error) bool {
	// MySQL 1062错误码：唯一键冲突
	return err != nil && err.Error() != "" && (err.Error() == "Error 1062: Duplicate entry '"+"' for key 'user_new_words.uk_user_word_id'" ||
		strings.Contains(err.Error(), "Duplicate entry") && strings.Contains(err.Error(), "uk_user_word_id"))
}
func AddNewWord(w http.ResponseWriter, r *http.Request) {
	var req models.AddNewWordRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		tool.JSONResponse(w, nil, http.StatusBadRequest)
		return
	}
	nick := r.Context().Value("nick").(string)
	// 检查是否已存在（包括软删除的）
	var existing models.UserNewWord
	err := config.DB.Unscoped().Where("nick = ? AND word_id = ?", nick, req.WordID).First(&existing).Error
	if err == nil {
		// 如果存在且被软删除，恢复它
		if existing.DeletedAt.Valid {
			existing.DeletedAt = gorm.DeletedAt{}
			err = config.DB.Save(&existing).Error
			if err != nil {
				tool.JSONResponse(w, nil, http.StatusInternalServerError)
				return
			}
			tool.JSONResponse(w, existing, http.StatusOK)
			return
		} else {
			// 已存在且未删除
			tool.JSONResponse(w, map[string]string{"msg": "生词已存在"}, http.StatusBadRequest)
			return
		}
	} else if err != gorm.ErrRecordNotFound {
		tool.JSONResponse(w, nil, http.StatusInternalServerError)
		return
	}
	// 不存在，插入新记录
	newWord := models.UserNewWord{
		Nick:       nick,
		WordID:     req.WordID,
		IsMastered: false, // 默认未掌握
	}
	err = config.DB.Create(&newWord).Error
	if err != nil {
		tool.JSONResponse(w, nil, http.StatusInternalServerError)
		return
	}
	tool.JSONResponse(w, newWord, http.StatusCreated)
}
func DeleteNewWord(w http.ResponseWriter, r *http.Request) {
	nick := r.Context().Value("nick").(string)
	wordIdStr := chi.URLParam(r, "wordId")
	wordId, err := strconv.Atoi(wordIdStr)
	if err != nil {
		tool.JSONResponse(w, nil, http.StatusBadRequest)
		return
	}
	err = config.DB.Where(&models.UserNewWord{Nick: nick, WordID: uint(wordId)}).Delete(&models.UserNewWord{}).Error
	if err != nil {
		tool.JSONResponse(w, nil, http.StatusInternalServerError)
		return
	}
	tool.JSONResponse(w, map[string]string{"msg": "删除成功"}, http.StatusOK)
}
func checkUserHasWord(nick string, wordId int) (*models.UserNewWord, error) {
	var existing models.UserNewWord
	err := config.DB.Where("nick = ? AND word_id = ?", nick, wordId).First(&existing).Error
	if err != nil {
		return nil, err
	}
	return &existing, nil
}
func MarkAsMastered(w http.ResponseWriter, r *http.Request) {
	wordIdStr := chi.URLParam(r, "wordId")
	wordId, err := strconv.Atoi(wordIdStr)
	if err != nil {
		tool.JSONResponse(w, nil, http.StatusBadRequest)
		return
	}
	nick := r.Context().Value("nick").(string)
	existing, err := checkUserHasWord(nick, wordId)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			tool.JSONResponse(w, map[string]string{"msg": "生词不存在"}, http.StatusNotFound)
			return
		}
		tool.JSONResponse(w, nil, http.StatusInternalServerError)
		return
	}
	err = config.DB.Model(existing).Update("is_mastered", true).Error
	if err != nil {
		tool.JSONResponse(w, nil, http.StatusInternalServerError)
		return
	}
	tool.JSONResponse(w, map[string]string{"msg": "标记为已掌握"}, http.StatusOK)
}
func MarkAsUnmastered(w http.ResponseWriter, r *http.Request) {
	wordIdStr := chi.URLParam(r, "wordId")
	wordId, err := strconv.Atoi(wordIdStr)
	if err != nil {
		tool.JSONResponse(w, nil, http.StatusBadRequest)
		return
	}
	nick := r.Context().Value("nick").(string)
	existing, err := checkUserHasWord(nick, wordId)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			tool.JSONResponse(w, map[string]string{"msg": "生词不存在"}, http.StatusNotFound)
			return
		}
		tool.JSONResponse(w, nil, http.StatusInternalServerError)
		return
	}
	err = config.DB.Model(existing).Update("is_mastered", false).Error
	if err != nil {
		tool.JSONResponse(w, nil, http.StatusInternalServerError)
		return
	}
	tool.JSONResponse(w, map[string]string{"msg": "标记为未掌握"}, http.StatusOK)
}
