package main

import (
	"beeblog/models"
	_ "beeblog/routers"

	"github.com/astaxie/beego"
)

func init() {
	models.RegisterDB()

}

func main() {

	/*o := orm.NewOrm()

	user := models.User{Name: "bbb"}

	id, err := o.Insert(&user)
	fmt.Printf("ID: %d, ERR: %v\n", id, err)

	user.Name = "astaxie"
	num, err := o.Update(&user)
	fmt.Printf("NUM: %d, ERR: %v\n", num, err)

	u := models.User{Id: user.Id}
	err = o.Read(&u)
	fmt.Printf("ERR: %v\n", err)

	num, err = o.Delete(&u)
	fmt.Printf("NUM: %d, ERR: %v\n", num, err)

	fmt.Println("----------1-----------")
	name, err := o.QueryTable("student").Filter("name", "zsddd").Update(orm.Params{
		"name": "CDFSDFE",
	})
	fmt.Printf("Affected Num: %s, %s", name, err)
	*/
	// orm.Debug = true
	beego.Run()
}
