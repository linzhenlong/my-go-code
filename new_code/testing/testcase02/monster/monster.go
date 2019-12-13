package monster

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type monster struct {
	Name string `json:"name"`
	Age int `json:"age"`
	Skill []string `json:"skill"`
}

const FILENAME = "/Users/smzdm/webroot/golang/src/go_dev/new_code/testing/testcase02/main/test.log"

func (m *monster)Store() (bool, error){
	str , err := json.Marshal(m)
	if err != nil {
		return false, err
	}
	file, err := os.Create(FILENAME)
	defer file.Close()
	n, err := file.Write(str)
	fmt.Println(n)
	return n>0, err
}

func (m *monster)ReStore() (bool, error) {
	data ,err := ioutil.ReadFile(FILENAME)
	if err !=nil {
		return  false,err
	}

	err = json.Unmarshal(data,m)
	if err !=nil {
		return false,err
	}
	return true,err

}

func NewMonster(name string ,age int ,skill []string) *monster  {
	return &monster{Name:name,Age:age,Skill:skill}
}