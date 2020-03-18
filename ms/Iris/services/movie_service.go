package services

import (
	"fmt"
	"github.com/linzhenlong/my-go-code/ms/Iris/repositories"
)

// MovieService 接口.
type MovieService interface {
	ShowMovieName() string
}

// MovieServiceManger 结构体.
type MovieServiceManger struct {
	repo repositories.MovieRepository
}

// NewMovieServiceManger 实例化.
func NewMovieServiceManger(repo repositories.MovieRepository) MovieService {
	return &MovieServiceManger{
		repo: repo,
	}
}

// ShowMovieName 展示电影名称.
func (m *MovieServiceManger) ShowMovieName() string {
	return fmt.Sprintf("展示电影名称:%s", m.repo.GetMovieName())
}
