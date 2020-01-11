package taskrunner

const (
	readyToDispatch = "d"
	readyToExcute   = "e"
	closeMsg        = "c"
	videoPath = "./videos/"
)

type controlChan chan string

type dataChan chan interface{}

type fn func(dc dataChan) error
