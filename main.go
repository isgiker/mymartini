package main

import (
	"net/http"

	"github.com/go-martini/martini"
	"github.com/martini-contrib/binding"
	"github.com/martini-contrib/render"
)

type UserInfo struct {
	Name  string `form:"name"`
	Email string `form:"email"`
}

func main() {
	m := martini.Classic()
	m.Use(render.Renderer(render.Options{
		Layout: "layout",
	}))
	m.Get("/", func(r render.Render) {
		// return "Hello, Welcome to the homepage!"
		r.HTML(http.StatusOK, "index", nil)
	})
	m.Post("/user", binding.Bind(UserInfo{}), func(r render.Render, user UserInfo) {
		var retData struct {
			User UserInfo
		}
		retData.User = user
		// r.JSON(http.StatusOK, user)
		//return user
		// r.HTML(http.StatusOK, "user_info", user)
		//return retData
		r.HTML(http.StatusOK, "user_info2", retData)
	})
	m.Get("/user/:userid", func(r render.Render, p martini.Params) {
		var retData struct {
			ID string
		}

		retData.ID = p["userid"]
		// r.HTML(http.StatusOK, "user", retData)
		r.JSON(http.StatusOK, retData)
	})
	m.Run()
}
