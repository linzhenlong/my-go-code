package repositories

import "github.com/linzhenlong/my-go-code/ms/Iris/datamodels"

// MovieRepository 接口.
type MovieRepository interface {
	GetMovieName() string
}

// MovieManager 结构体.
type MovieManager struct {
}

// NewMovieManager 创建NewMovieManager 实例.
func NewMovieManager() MovieRepository {
	return &MovieManager{}
}

// GetMovieName 实现 MovieRepository接口.
func (m *MovieManager) GetMovieName() string {
	// 模拟操作数据库.
	movie := &datamodels.Movie{Name: "iris say hello world"}
	return movie.Name
}
