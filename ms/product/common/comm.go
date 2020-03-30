package common

import (
	"errors"
	"log"
	"reflect"
	"strconv"
	"time"
)

// DataToStructByTagSQL 通过tag sql标签转成对应的结构体.
func DataToStructByTagSQL(data map[string]string, obj interface{}) {
	objValue := reflect.ValueOf(obj).Elem()
	log.Printf("objValue:%v", objValue)
	for i := 0; i < objValue.NumField(); i++ {
		// 获取sql 对应的值
		value := data[objValue.Type().Field(i).Tag.Get("sql")]
		log.Printf("objValue.Type().Field(i)=%v", objValue.Type().Field(i))
		log.Printf("objValue.Type().Field(i).Tag=%v", objValue.Type().Field(i).Tag)
		log.Printf("objValue.Type().Field(i).Tag.Get(\"sql\")=%v", objValue.Type().Field(i).Tag.Get("sql"))
		// 获取对应字段的名称
		name := objValue.Type().Field(i).Name
		log.Printf("name=%v", name)

		// 获取对应字段的类型
		structFieldType := objValue.Field(i).Type()
		log.Printf("structFieldType=%v", structFieldType)
		//获取变量类型
		val := reflect.ValueOf(value)
		log.Printf("val:%v", val)
		var err error
		if structFieldType != val.Type() {
			// 类型转换
			log.Printf("structFieldType.Name=%v", structFieldType.Name())

			val, err = TypeConversion(value, structFieldType.Name())
			if err != nil {
				log.Printf("err=%v", err)
			}
		}
		objValue.FieldByName(name).Set(val)
	}
}

// TypeConversion 类型转换.
func TypeConversion(value string, ntype string) (reflect.Value, error) {
	switch ntype {
	case "uint8":
		i, err := strconv.ParseInt(value, 10, 64)
		return reflect.ValueOf(i), err
	case "string":
		return reflect.ValueOf(value), nil
	case "int64":
		i, err := strconv.ParseInt(value, 10, 64)
		return reflect.ValueOf(i), err
	case "Time":
		local, err := time.LoadLocation("Local")
		timeVal, err := time.ParseInLocation("2006-01-02 15:04:05", value, local)
		return reflect.ValueOf(timeVal), err

	}
	return reflect.ValueOf(value), errors.New("未知错误")
}
