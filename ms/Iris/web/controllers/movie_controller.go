package controllers

import (
	"github.com/kataras/iris/mvc"
	"github.com/linzhenlong/my-go-code/ms/Iris/repositories"
	"github.com/linzhenlong/my-go-code/ms/Iris/services"
)

// MovieController 电影控制器.
type MovieController struct {
}

// Get 方法.
func (c *MovieController) Get() mvc.View {
	movieRepo := repositories.NewMovieManager()
	movieService := services.NewMovieServiceManger(movieRepo)
	movieName := movieService.ShowMovieName()
	return mvc.View{
		Name: "movie/index.html",
		Data: movieName,
	}
}
