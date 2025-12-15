package routers

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/go-chi/chi/v5"
	"imxy.top/wordserver/internal/handler"
	"imxy.top/wordserver/internal/middleware"
)

func InitAllRouters() {
	r := chi.NewRouter()
	workDir, _ := os.Getwd()
	staticDir := filepath.Join(workDir, "static")
	r.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir(staticDir))))

	// 4. 配置 HTML 路由
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		// 设置响应头：HTML 格式 + 中文编码
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		// 读取并返回 HTML 文件
		htmlPath := filepath.Join(workDir, "templates", "index.html")
		http.ServeFile(w, r, htmlPath)
	})
	r.Route("/api/v1", func(r chi.Router) {
		// r.Use(middleware.CORS)
		r.Route("/words", func(r chi.Router) {
			r.Get("/", handler.ListWords)       // 查询单词列表
			r.Get("/{wordId}", handler.GetWord) // 查询单个单词
		})
		// 查询用户生词接口（无需鉴权）
		r.Route("/user-new-words", func(r chi.Router) {
			r.Get("/", handler.ListNewWords)
		})
		r.Route("/user", func(r chi.Router) {
			r.Post("/register", handler.Register) // 注册
			r.Post("/login", handler.Login)
		})
		r.Group(func(r chi.Router) {
			r.Use(middleware.JWTAuth) // 验证Token有效性
			//修改生词接口
			r.Route("/new-word", func(r chi.Router) {
				r.Get("/", handler.ListNewWords)                      //已登录用户查询生词接口
				r.Post("/", handler.AddNewWord)                       // 添加生词
				r.Put("/{wordId}/master", handler.MarkAsMastered)     // 标记为已掌握
				r.Put("/{wordId}/unmaster", handler.MarkAsUnmastered) // 标记为未掌握
				r.Delete("/{wordId}", handler.DeleteNewWord)          // 删除生词
			})
		})
	})
	http.ListenAndServe(":1145", r)
}
