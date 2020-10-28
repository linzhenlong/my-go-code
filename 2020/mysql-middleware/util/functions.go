package util

import "regexp"

func replaceProductTable(sql string) string {
	r := regexp.MustCompile("product")
	ret := r.ReplaceAllString(sql, "(select * from product_1 union select * from product_2) a")
	return ret
}
