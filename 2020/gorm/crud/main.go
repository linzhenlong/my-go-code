package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"math/rand"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// User 用户结构提.
type User struct {
	gorm.Model
	Name sql.NullString `gorm:"type:varchar(255)"`
	Age  *int           `gorm:"default:'18';not null"` // 默认值18
}

// TableName 表名.
func (User) TableName() string {
	return "user_crud"
}

func main() {
	db, err := gorm.Open("mysql", "root:@(127.0.0.1:3306)/ms?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}

	rand.Seed(time.Now().UnixNano())
	//rand := rand.Intn(100)
	defer db.Close()
	// 2.把模型与数据库中的表对应起来
	db.AutoMigrate(&User{})

	// 3.创建
	/* u1 := User{
		//Name: "张三:" + strconv.Itoa(rand),
		Name: sql.NullString{
			String: "",
			Valid:  true,
		},
		Age: new(int), // 零值
	}
	fmt.Println(db.NewRecord(&u1)) // 判断主键是否为空
	db.Debug().Create(&u1)
	fmt.Println(u1.ID)
	fmt.Printf("%#v\n", u1)
	fmt.Println(db.NewRecord(&u1)) // 判断主键是否为空 */

	/* age := rand
	 u2 := User{
		Name: sql.NullString{
			String: "张三",
			Valid:  true,
		},
		Age: &age,
	}
	db.Debug().Create(&u2)

	u3 := User{
		Name: sql.NullString{
			String: "李四",
			Valid:  true,
		},
		Age: &age,
	}
	db.Debug().Create(&u3) */

	// 查询操作
	// 声明结构体类型变量
	user := User{}

	// 查询表中第一条结果.
	db.Debug().First(&user) //SELECT * FROM `user_crud`  WHERE `user_crud`.`deleted_at` IS NULL ORDER BY `user_crud`.`id` ASC LIMIT 1
	fmt.Printf("%#v\n", user)

	// 随机取一条
	user2 := new(User)
	db.Debug().Take(user2) // SELECT * FROM `user_crud`  WHERE `user_crud`.`deleted_at` IS NULL LIMIT 1
	fmt.Printf("%#v\n", user2)

	// 取最后一条
	lastUser := User{}
	db.Debug().Last(&lastUser) // SELECT * FROM `user_crud`  WHERE `user_crud`.`deleted_at` IS NULL ORDER BY `user_crud`.`id` DESC LIMIT 1
	fmt.Printf("%#v\n", lastUser)

	// 查询所有结果
	users := []User{}
	db.Debug().Find(&users) // SELECT * FROM `user_crud`  WHERE `user_crud`.`deleted_at` IS NULL
	fmt.Printf("%#v\n", users)

	usersJSON, _ := json.Marshal(users)
	fmt.Println(string(usersJSON))

	// 查询指定的某条记录(仅当主键是整形是可用)
	var user3 User
	db.Debug().First(&user3, 2) // SELECT * FROM `user_crud`  WHERE `user_crud`.`deleted_at` IS NULL AND ((`user_crud`.`id` = 2)) ORDER BY `user_crud`.`id` ASC LIMIT 1
	fmt.Printf("%#v\n", user3)

	// where 条件
	// 普通sql查询

	// 获取第一条满足条件的记录
	user4 := User{}
	//  SELECT * FROM `user_crud`  WHERE `user_crud`.`deleted_at` IS NULL AND ((name='张三')) ORDER BY `user_crud`.`id` ASC LIMIT 1
	db.Debug().Where("name=?", "张三").First(&user4)
	fmt.Printf("%#v\n", user4)

	// 获取符合条件的所有记录
	user5 := []User{}
	// SELECT * FROM `user_crud`  WHERE `user_crud`.`deleted_at` IS NULL AND ((age>=10))
	db.Debug().Where("age>=?", 10).Find(&user5)
	fmt.Printf("%#v\n", user5)
	usersJSON, _ = json.Marshal(user5)
	fmt.Println(string(usersJSON))

	// 不等于
	user6 := User{}
	// SELECT * FROM `user_crud`  WHERE `user_crud`.`deleted_at` IS NULL AND ((name <> '张三'))
	db.Debug().Where("name <> ?", "张三").Find(&user6)
	fmt.Printf("%#v\n", user6)

	// in 查询
	user7 := []User{}
	// SELECT * FROM `user_crud`  WHERE `user_crud`.`deleted_at` IS NULL AND ((name in ('张三','李四')))
	db.Debug().Where("name in (?)", []string{"张三", "李四"}).Find(&user7)
	fmt.Printf("%#v\n", user7)
	usersJSON, _ = json.Marshal(user7)
	fmt.Println(string(usersJSON))

	// like 查询
	var user8 []User
	// SELECT * FROM `user_crud`  WHERE `user_crud`.`deleted_at` IS NULL AND ((name like '%张%'))
	db.Debug().Where("name like ?", "%张%").Find(&user8)
	fmt.Printf("%#v\n", user8)
	usersJSON, _ = json.Marshal(user8)
	fmt.Println(string(usersJSON))

	// AND
	user9 := []User{}
	// SELECT * FROM `user_crud`  WHERE `user_crud`.`deleted_at` IS NULL AND ((name='张三' and age > 100))
	db.Debug().Where("name=? and age > ?", "张三", 100).Find(&user9)
	fmt.Printf("%#v\n", user9)
	usersJSON, _ = json.Marshal(user9)
	fmt.Println(string(usersJSON))

	// TIME
	user10 := []User{}
	now := time.Now().Unix()
	lastWeekUnix := now - 7*24*3600
	lastWeek := time.Unix(lastWeekUnix, 0)
	// SELECT * FROM `user_crud`  WHERE `user_crud`.`deleted_at` IS NULL AND ((updated_at > '2020-04-07 22:45:53'))
	db.Debug().Where("updated_at > ?", lastWeek).Find(&user10)
	fmt.Printf("%#v\n", user10)
	usersJSON, _ = json.Marshal(user10)
	fmt.Println(string(usersJSON))

	// BETWEEN
	user11 := []User{}
	currentDay := time.Now()
	lastWeek2 := currentDay.AddDate(0, 0, -7)
	fmt.Println(currentDay)
	fmt.Println(lastWeek2)
	// SELECT * FROM `user_crud`  WHERE `user_crud`.`deleted_at` IS NULL AND ((updated_at BETWEEN '2020-04-07 22:49:56' and '2020-04-14 22:49:56'))
	db.Debug().Where("updated_at BETWEEN ? and ?", lastWeek2, currentDay).Find(&user11)
	fmt.Printf("%#v\n", user11)
	usersJSON, _ = json.Marshal(user11)
	fmt.Println(string(usersJSON))

	// struct &map 查询

	// struct
	age := 52
	user12 := User{
		Name: sql.NullString{
			String: "张三",
			Valid:  true,
		},
		Age: &age,
	}
	// SELECT * FROM `user_crud`  WHERE `user_crud`.`deleted_at` IS NULL AND ((`user_crud`.`name` = '张三') AND (`user_crud`.`age` = 52)) ORDER BY `user_crud`.`id` ASC LIMIT 1
	db.Debug().Where(&user12).First(&user12)
	fmt.Printf("%#v\n", user12)

	// map
	user13 := []User{}
	userMap := make(map[string]interface{})
	userMap["name"] = "李四"
	userMap["age"] = 55
	// SELECT * FROM `user_crud`  WHERE `user_crud`.`deleted_at` IS NULL AND ((`user_crud`.`name` = '李四') AND (`user_crud`.`age` = 55))
	db.Debug().Where(userMap).Find(&user13)
	fmt.Printf("%#v\n", user13)

	// 主键的切片
	user14 := []User{}
	// SELECT * FROM `user_crud`  WHERE `user_crud`.`deleted_at` IS NULL AND ((`user_crud`.`id` IN (1,2)))
	db.Debug().Where([]int64{1, 2}).Find(&user14)
	fmt.Printf("%#v\n", user14)

	//not 条件
	user15 := User{}
	// SELECT * FROM `user_crud`  WHERE `user_crud`.`deleted_at` IS NULL AND ((`user_crud`.`name` NOT IN ('李四')))
	db.Debug().Not("name", "李四").Find(&user15)
	fmt.Printf("%#v\n", user15)

	// or查询
	users16 := []User{}
	// SELECT * FROM `user_crud`  WHERE `user_crud`.`deleted_at` IS NULL AND ((name='张三') OR (name='李四'))
	db.Debug().Where("name=?", "张三").Or("name=?", "李四").Find(&users16)
	fmt.Printf("%#v\n", users16)

	// 内联条件
	// 作用与where查询类似，当内联条件与多个立即执行方法一起使用时，内联添加不会传递给后面的立即执行方法

	user17 := User{}
	// SELECT * FROM `user_crud`  WHERE `user_crud`.`deleted_at` IS NULL AND ((`user_crud`.`id` = 2)) ORDER BY `user_crud`.`id` ASC LIMIT 1
	db.Debug().First(&user17, 2)
	fmt.Printf("%#v\n", user17)

	user18 := User{}
	// SELECT * FROM `user_crud`  WHERE `user_crud`.`deleted_at` IS NULL AND ((id=100)) ORDER BY `user_crud`.`id` ASC LIMIT 1
	db.Debug().First(&user18, "id=?", 100)
	fmt.Printf("%#v\n", user18)

	users19 := []User{}
	// SELECT * FROM `user_crud`  WHERE `user_crud`.`deleted_at` IS NULL AND ((name <> '张三' and age > 50))
	db.Debug().Find(&users19, "name <> ? and age > ?", "张三", 50)
	fmt.Printf("%#v\n", users19)

	// FirstOrInit
	user20 := User{}

	userStruct1 := User{
		Name: sql.NullString{
			String: "不存在",
			Valid:  true,
		},
	}
	// SELECT * FROM `user_crud`  WHERE `user_crud`.`deleted_at` IS NULL AND ((`user_crud`.`name` = '不存在')) ORDER BY `user_crud`.`id` ASC LIMIT 1
	db.Debug().FirstOrInit(&user20, userStruct1)
	// 查找名字是不存在的用户，如果不存在，就用name=不存在 初始化一个user 结构体
	fmt.Printf("%#v\n", user20)

	// SELECT * FROM `user_crud`  WHERE `user_crud`.`deleted_at` IS NULL AND ((name='张三')) ORDER BY `user_crud`.`id` ASC LIMIT 1
	user21 := User{}
	db.Debug().FirstOrInit(&user21, "name=?", "张三")
	fmt.Println(user21.Name, *user21.Age, user21.CreatedAt.Format("2006-01-02 15:04:05"))

	// Attrs
	// 如果记录未找到时，将使用参数初始化struct
	user22 := User{}
	user22Age := 19

	userStruct2 := User{
		Name: sql.NullString{
			String: "不存在",
			Valid:  true,
		},
	}
	createAt := time.Now()
	userStruct3 := User{
		Age: &user22Age,
	}
	userStruct3.CreatedAt = createAt
	// SELECT * FROM `user_crud`  WHERE `user_crud`.`deleted_at` IS NULL AND ((name='张三')) ORDER BY `user_crud`.`id` ASC LIMIT 1
	db.Attrs(userStruct3).FirstOrInit(&user22, userStruct2)
	// 查找名字是不存在的用户，如果不存在，就用name=不存在 初始化一个user 结构体,并将Attrs中的值赋值给这个结构体
	fmt.Printf("user:%s,age:%d,time:%v\n", user22.Name.String, *user22.Age, user22.CreatedAt.Format("2006-01-02 15:04:05"))

	// Assign
	// 不管记录是否找到，都将参数赋值给 struct.

	// 找到的情况
	user23 := User{}
	user23Age := 99

	userStruct23 := User{
		Name: sql.NullString{
			String: "张三",
			Valid:  true,
		},
	}
	createAt23 := time.Now()
	userStruct2323 := User{
		Age: &user23Age,
	}
	userStruct2323.CreatedAt = createAt23

	db.Assign(userStruct2323).Where(userStruct23).FirstOrInit(&user23)
	fmt.Printf("user:%s,age:%d,time:%v\n", user23.Name.String, *user23.Age, user23.CreatedAt.Format("2006-01-02 15:04:05"))

	// 未找到
	user24 := User{}
	user24Age := 88

	userStruct24 := User{
		Name: sql.NullString{
			String: "张三xx",
			Valid:  true,
		},
	}
	createAt24 := time.Now()
	userStruct2424 := User{
		Age: &user24Age,
	}
	userStruct2424.CreatedAt = createAt24

	db.Assign(userStruct2424).Where(userStruct24).FirstOrInit(&user24)
	fmt.Printf("user:%s,age:%d,time:%v\n", user24.Name.String, *user24.Age, user24.CreatedAt.Format("2006-01-02 15:04:05"))
}
