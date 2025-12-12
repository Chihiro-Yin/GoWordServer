package handler

import (
	"encoding/json"
	"net/http"
	"strings"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"imxy.top/wordserver/internal/config"
	"imxy.top/wordserver/internal/middleware"
	"imxy.top/wordserver/internal/models"
	"imxy.top/wordserver/internal/tool"
)

func Register(w http.ResponseWriter, r *http.Request) {
	// 1. 解析请求体
	var req models.RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		tool.JSONResponse(w, nil, http.StatusBadRequest)
		return
	}

	// 2. 密码加密（bcrypt）
	hashPwd, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		tool.JSONResponse(w, nil, http.StatusInternalServerError)
		return
	}

	// 3. 构建用户模型
	user := models.User{
		Nick:     req.Nick,
		Email:    req.Email,
		Password: string(hashPwd),
		Token:    "", // 注册时token为空，登录时生成
	}

	// 4. 插入数据库（利用nick唯一索引，避免重复注册）
	err = config.DB.Create(&user).Error
	if err != nil {
		// 昵称重复：返回400
		if isDuplicateNickError(err) {
			tool.JSONResponse(w, map[string]string{"msg": "昵称已存在"}, http.StatusBadRequest)
			return
		}
		tool.JSONResponse(w, nil, http.StatusInternalServerError)
		return
	}
	token, err := middleware.GenerateToken(user.ID, user.Nick)
	if err != nil {
		tool.JSONResponse(w, nil, http.StatusInternalServerError)
		return
	}

	// 6. 更新用户Token到数据库
	config.DB.Model(&user).Update("token", token)

	// 7. 构建响应（包含Token）
	respData := map[string]interface{}{
		"id":    user.ID,
		"nick":  user.Nick,
		"email": user.Email,
		"token": token, // 注册成功直接返回Token
	}
	tool.JSONResponse(w, respData, http.StatusCreated)
}

// 辅助函数：判断是否是昵称重复错误（适配MySQL）
func isDuplicateNickError(err error) bool {
	// MySQL 1062错误码：唯一键冲突
	return err != nil && err.Error() != "" && (err.Error() == "Error 1062: Duplicate entry '"+"' for key 'users.uk_nick'" ||
		strings.Contains(err.Error(), "Duplicate entry") && strings.Contains(err.Error(), "uk_nick"))
}
func Login(w http.ResponseWriter, r *http.Request) {
	// 1. 解析并校验请求体
	var req models.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		tool.JSONResponse(w, nil, http.StatusBadRequest)
		return
	}

	// 2. 根据昵称查询用户（走nick唯一索引，性能最优）
	var user models.User
	err := config.DB.Where("nick = ?", req.Nick).First(&user).Error
	if err != nil {
		// 用户不存在 → 返回401（统一提示，避免泄露用户信息）
		if err == gorm.ErrRecordNotFound {
			tool.JSONResponse(w, map[string]string{"msg": "昵称或密码错误"}, http.StatusUnauthorized)
			return
		}
		// 数据库异常 → 返回500
		tool.JSONResponse(w, nil, http.StatusInternalServerError)
		return
	}

	// 3. 验证密码（bcrypt比对，不可逆加密）
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		tool.JSONResponse(w, map[string]string{"msg": "昵称或密码错误"}, http.StatusUnauthorized)
		return
	}

	token, err := middleware.GenerateToken(user.ID, user.Nick)
	if err != nil {
		tool.JSONResponse(w, nil, http.StatusInternalServerError)
		return
	}

	// 5. 更新用户最新Token到数据库
	config.DB.Model(&user).Update("token", token)

	// 6. 构建响应（与注册接口返回格式完全统一）
	respData := map[string]interface{}{
		"id":    user.ID,
		"nick":  user.Nick,
		"email": user.Email,
		"token": token,
	}
	tool.JSONResponse(w, respData)
}
