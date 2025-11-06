package config

import (
	admin "goblog/admin/controllers"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Routes() *httprouter.Router {
	r := httprouter.New()

	// ? /admin
	//?  /admin/yeni-ekle
	r.GET("/admin", admin.Dashboard{}.Index)
	r.GET("/admin/yeni-ekle", admin.Dashboard{}.NewItem)

	//? SERVE FİLES Admin dashboard dosyalarını gösterebilmek için
	r.ServeFiles("/admin/assets/*filepath", http.Dir("admin/assets"))
	return r
}
