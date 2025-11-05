package main

import (
	"fmt"
	admin_models "goblog/admin/models"
	"goblog/config"
	"net/http"
)

func main() {
	admin_models.Post{}.Migrate()
	/*post := admin_models.Post{}.Get(2)  bu şekilde databaseden tek bir veri alabilirsin
	  post := admin_models.Post{}.Get("title","go ile web")
	fmt.Println(post.Title)*/

	//fmt.Println(admin_models.Post{}.GetAll("category_id",1))

	/*
		post := admin_models.Post{}.Get(4)
		post.Update("description", "selam bebeks")
		post.Updates(admin_models.Post{Title: "erdinc ve sila", CategoryID: 6})
	*/
	fmt.Println(admin_models.Post{}.GetAll("category_id", 1))
	http.ListenAndServe(":8080", config.Routes())

}
