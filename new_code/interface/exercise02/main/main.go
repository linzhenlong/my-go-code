package main

type AInterFace interface {
	Test01()
	Test02()
}
type BInterFace interface {
	Test02()
	Test03()
}
type CInterFace interface {
	AInterFace
	BInterFace
}

func main()  {


}
