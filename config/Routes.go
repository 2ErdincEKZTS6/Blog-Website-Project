package config

import (
	admin "goblog/admin/controllers"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Routes() *httprouter.Router {
	r := httprouter.New()

	// ? /ADMİN
	r.GET("/admin", admin.Dashboard{}.Index)

	//? SERVE FİLES Admin dashboard dosyalarını gösterebilmek için
	r.ServeFiles("/admin/assets/*filepath", http.Dir("admin/assets"))
	return r
}
