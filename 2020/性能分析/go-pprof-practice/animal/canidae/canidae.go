package canidae

import (
	"github.com/linzhenlong/my-go-code/2020/性能分析/go-pprof-practice/animal"
)


// canidae ...
type Canidae interface {
	animal.Animal
	Run()
	Howl()
}