package routers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"imxy.top/wordserver/internal/handler"
	"imxy.top/wordserver/internal/middleware"
)

func InitAllRouters() {
	r := chi.NewRouter()
	r.Route("/api/v1", func(r chi.Router) {
		r.Route("/words", func(r chi.Router) {
			r.Get("/", handler.ListWords)       // 查询生词列表
			r.Get("/{wordId}", handler.GetWord) // 查询单个生词
		})
		// 用户生词接口
		r.Route("/new-word", func(r chi.Router) {
			r.Get("/{nick}", handler.ListNewWords)
		})
		r.Route("/user", func(r chi.Router) {
			r.Post("/register", handler.Register) // 注册
			r.Post("/login", handler.Login)
		})
		r.Group(func(r chi.Router) {
			r.Use(middleware.JWTAuth) // 验证Token有效性
			//修改生词接口
			// r.Route("/new-word", func(r chi.Router) {
			// 	r.Post("/", handler.AddNewWord)              // 添加生词
			// 	r.Delete("/{wordId}", handler.DeleteNewWord) // 删除生词
			// })
		})
	})
	http.ListenAndServe(":3000", r)
}
