package main

import (
	"fmt"
	"log"

	. "github.com/xwb1989/sqlparser"
)

func main() {
	sql := "select id,name from product as a"
	stmt, err := Parse(sql)
	if err != nil {
		log.Fatal(err)
	}
	// Otherwise do something with stmt
	switch stmt := stmt.(type) {
	case *Select:
		buf := NewTrackedBuffer(nil)
		stmt.SelectExprs.Format(buf)
		println(buf.String())
		// for _, node := range stmt.From {
		// 	//fmt.Printf("%#v, %T\n", node, node)
		// 	getTable := node.(*sqlparser.AliasedTableExpr)
		// 	fmt.Println(getTable.As.String())
		// 	fmt.Printf("%T\n", getTable.Expr)
		// 	tableExpr, ok := getTable.Expr.(sqlparser.TableName)
		// 	if !ok {
		// 		break
		// 	}
		// 	fmt.Println(tableExpr.Name)
		// }
	case *Insert:
	default:
		fmt.Println("default")
	}
}
