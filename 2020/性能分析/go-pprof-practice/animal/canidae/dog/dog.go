package dog

import "log"

import "github.com/wolfogre/go-pprof-practice/constant"

// Dog ...
type Dog struct {

}
// Name ...
func (d *Dog)Name()string {
	return "dog"
}

func (d *Dog)Live() {
	d.Eat()
	d.Drink()
	d.Shit()
	d.Pee()
	d.Run()
	d.Howl()
}

func (d *Dog)Drink() {
	log.Println(d.Name(),"Drink")
}

func (d *Dog)Eat() {
	log.Println(d.Name(),"Eat")
}

func (d *Dog)Pee() {
	log.Println(d.Name(),"Pee")
}

func (d *Dog)Shit() {
	log.Println(d.Name(),"Shit")
}
func (d *Dog)Run() {
	log.Println(d.Name(),"run")
	_ = make([]byte,constant.Mi * 16)
}

func (d *Dog)Howl() {
	log.Println(d.Name(), "howl")
}