package stu

import (
	"encoding/json"
	"io/ioutil"
)

// Student 学生结构体.
type Student struct {
	Name string
	Age  int
	Sex  string
}

// Save save student.
func (stu *Student) Save() error {
	data, err := json.Marshal(stu)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile("./test.txt", data, 0755)
	if err != nil {
		return err
	}
	return nil
}

// Load load student.
func (stu *Student) Load() (Student, error) {
	data, err := ioutil.ReadFile("./test.txt")
	student := Student{}
	if err != nil {
		return student, err
	}

	err = json.Unmarshal(data, &student)
	if err != nil {
		return student, err
	}
	return student, nil
}
