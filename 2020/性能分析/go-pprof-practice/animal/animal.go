package animal

import (
	"github.com/linzhenlong/my-go-code/2020/性能分析/go-pprof-practice/animal/canidae/dog"
)

var (
	Allanimals =[]Animal {
		&dog.Dog{},
		// &wolf.Wolf{},

		// &cat.Cat{},
		// &tiger.Tiger{},

		// &mouse.Mouse{},
	}
)

type Animal interface{
	Name() string
	Live()

	Eat()
	Drink()
	Shit()
	Pee()
}