package taskrunner

import (
	"log"
	"testing"
	"time"
)

/* func TestMain(m *testing.M) {
	m.Run()
} */

func TestRunner(t *testing.T) {
	d := func(dc dataChan) error {
		for i := 0; i < 30; i++ {
			dc <- i
			log.Printf("Dispatcher sent:%d", i)
		}
		return nil
	}
	e := func(dc dataChan) error {
	LABLE:
		for {
			select {
			case d := <-dc:
				log.Printf("excutor recevied:%v", d)
			default:
				break LABLE
			}
		}
		return nil
	}
	runner := NewRunner(30, false, d, e)
	go runner.StartAll()
	time.Sleep(3 * time.Second)
}
