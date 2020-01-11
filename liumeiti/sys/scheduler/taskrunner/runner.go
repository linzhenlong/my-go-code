package taskrunner

type Runner struct {
	Controller controlChan
	Error      controlChan
	Data       dataChan
	dataSize   int
	longLived  bool
	Dispatcher fn
	Excutor    fn
}

func NewRunner(size int, longLived bool, dispatcher fn, excutor fn) *Runner {
	return &Runner{
		Controller: make(chan string, 1), // 使用代buffer的channel,非阻塞的
		Error:      make(chan string, 1),
		Data:       make(chan interface{}, size),
		longLived:  longLived,
		Dispatcher: dispatcher,
		Excutor:    excutor,
	}

}

// 生产者，消费者模型.
func (r *Runner) startDispatcher() {
	defer func() {
		// 如果不是常驻的,关闭管道.
		if !r.longLived {
			close(r.Controller)
			close(r.Data)
			close(r.Error)
		}
	}()
	for {
		select {
		case c := <-r.Controller:
			if c == readyToDispatch {
				err := r.Dispatcher(r.Data)
				if err != nil {
					r.Error <- closeMsg
				} else {
					r.Controller <- readyToExcute
				}
			}
			if c == readyToExcute {
				err := r.Excutor(r.Data)
				if err != nil {
					r.Error <- closeMsg
				} else {
					r.Controller <- readyToDispatch
				}
			}
		case e := <-r.Error:
			if e == closeMsg {
				return
			}
		default:
		}
	}
}

func (r *Runner) StartAll() {
	r.Controller <- readyToDispatch
	r.startDispatcher()
}

// 该5-4
