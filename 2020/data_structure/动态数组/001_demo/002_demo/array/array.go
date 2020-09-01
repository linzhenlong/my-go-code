package array

// https://segmentfault.com/a/1190000015680429

// IArray ...
type IArray interface {
	Add(int, interface{}) // 插入元素
	GetCap() int
	AddLast(interface{})
	AddFirst(interface{})
	GetSize() int
	IsEmpty() bool
}

type array struct {
	data []interface{}
	size int
}

// NewArray ...
func NewArray(cap int) IArray {
	arr := &array{}
	arr.data = make([]interface{}, cap)
	arr.size = 0
	return arr
}

func (a *array) GetSize() int {
	return a.size
}

func (a *array) IsEmpty() bool {
	return a.size == 0
}

func (a *array) AddFirst(e interface{}) {
	a.Add(0, e)
}

// GetCap 数组容量
func (a *array) GetCap() int {
	return len(a.data)
}

// 数组扩容
func (a *array) resize(newCap int) {
	newArr := make([]interface{}, newCap)
	for i := 0; i < a.size; i++ {
		newArr[i] = a.data[i]
	}
	a.data = newArr
}
func (a *array) AddLast(element interface{}) {
	a.Add(a.size, element)
}

func (a *array) Add(index int, element interface{}) {
	if index < 0 || index > a.GetCap() {
		panic("Add failed, require index >=0 and index <= cap")
	}
	// 判断数组是否需求扩容
	if a.size == len(a.data) {
		a.resize(2 * a.size)
	}
	// 将插入的索引位置之后的原始后移，腾出插入位置
	// arr[0]=>1 arr[1]=>2  arr[2] => 3
	// arr[0]=>1 arr[1] => 10, arr[2]=>2 ,arr[3]=>3
	for i := a.size - 1; i >= index; i-- {
		a.data[i+1] = a.data[i]
	}
	a.data[index] = element
	a.size++
}
