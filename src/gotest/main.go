package main

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

const (
	_DB_NAME   = "mysql"
	_DB_DRIVER = "root:zhang1991@tcp(127.0.0.1:3306)/gotest?charset=utf8&parseTime=true&loc=Local"
)

type User struct {
	Id   int64
	Name string `orm:"size(100)"`
}

func AddUser(id int64, name string) {
	o := orm.NewOrm()

	user := &User{
		Id: id,
		Name: name,
	}
	_, err := o.Insert(user)
	if err != nil {
		// return err
	}
}


func init(){
	orm.RegisterDataBase("default", _DB_NAME, _DB_DRIVER, 30)

	orm.RegisterModel(new(User))

	orm.RunSyncdb("default", false, true)
}

func main(){
	fmt.Println("--------")
	users := map[string]*User{
		"01":{Id:01,Name:"tom",},
		"02":{Id:02,Name:"smith",},
		"03":{Id:03,Name:"tony",},
		"04":{Id:04,Name:"ben"},
		"05":{Id:05,Name:"daive"},
		"06":{Id:06,Name:"lilei"},
		"07":{Id:07,Name:"hanmeimei"},
	}

	for _, user := range users {
		fmt.Printf("%T, %T\n", user.Id, user.Name)
		AddUser(user.Id, user.Name)
	}
}