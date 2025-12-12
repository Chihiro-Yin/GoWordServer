package tool

import (
	"encoding/json"
	"net/http"
)

func JSONResponse(w http.ResponseWriter, data any, code ...int) {
	httpCode := http.StatusOK // 用标准库常量更规范
	if len(code) > 0 && code[0] > 0 {
		httpCode = code[0]
	}
	resp := map[string]any{
		"code": httpCode,               // 业务码和HTTP状态码统一（更符合前端习惯）
		"msg":  getMsgByCode(httpCode), // 自动匹配提示语
		"data": data,                   // 业务数据（支持任意类型，nil则返回null）
	}
	// 1. 设置响应头为JSON格式（固定）
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	// 2. 设置HTTP状态码
	if code != nil {
		w.WriteHeader(code[0])
	} else {
		w.WriteHeader(200)
	}
	// 3. 编码为JSON并写入响应体
	err := json.NewEncoder(w).Encode(resp)
	if err != nil {
		http.Error(w, `{"code":500,"msg":"响应序列化失败"}`, http.StatusInternalServerError)
	}
}
func getMsgByCode(code int) string {
	switch code {
	case http.StatusOK:
		return "success"
	case http.StatusBadRequest:
		return "参数错误"
	case http.StatusUnauthorized:
		return "未授权"
	case http.StatusForbidden:
		return "禁止访问"
	case http.StatusNotFound:
		return "资源不存在"
	case http.StatusInternalServerError:
		return "服务器内部错误"
	default:
		return "success" // 未知状态码默认返回success
	}
}
