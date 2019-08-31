package routers

import (
	"beeblog/controllers"
	"os"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
)

func init() {

	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"*"},
		AllowCredentials: true,
	}))
	beego.Router("/", &controllers.MainController{})
	beego.Router("/category", &controllers.CategoryController{})
	beego.Router("/topic", &controllers.TopicController{})
	beego.Router("/reply", &controllers.ReplyController{})
	beego.Router("/reply/add", &controllers.ReplyController{}, "post:Add")
	beego.Router("/reply/delete", &controllers.ReplyController{}, "get:Delete")
	beego.AutoRouter(&controllers.TopicController{})
	beego.Router("login", &controllers.LoginController{})

	//创建附件目录
	os.Mkdir("attachment", os.ModePerm)
	/*
		//作为静态文件
		beego.SetStaticPath("/attachment", "attachment")
	*/

	//作为单独一个控制器来处理
	beego.Router("/attachment/:all", &controllers.AttachController{})

	//tech
	beego.Router("/tech", &controllers.TechController{})

	//detail
	beego.Router("/detail", &controllers.DetailController{}, "post:Detail")

	//homeat
	beego.Router("/homeat", &controllers.HomeController{})

	//categoryAt
	beego.Router("/categoryAt", &controllers.GetCategoryAtController{}, "post:GetCategoryAt")

	//topicat
	beego.Router("/topicat", &controllers.TopicAtController{})
	beego.Router("/topicat/add", &controllers.TopicAtController{})
	beego.Router("/topicat/edit", &controllers.TopicAtController{})
	beego.Router("/topicat/view", &controllers.TopicAtController{}, "get:View")
	beego.Router("/topicat/delete", &controllers.TopicAtController{}, "post:Delete")
	beego.AutoRouter(&controllers.TopicAtController{})

	//replyat
	beego.Router("/replyat", &controllers.ReplyatController{})
	beego.Router("/replyat/add", &controllers.ReplyatController{}, "post:Add")
	beego.Router("/replyat/delete", &controllers.ReplyatController{}, "get:Delete")

	//register
	beego.Router("/register", &controllers.RegisterController{})

	//login
	beego.Router("/loginat", &controllers.LoginatController{})

	//rist
	beego.Router("/rist", &controllers.RistController{}, "post:Rist")

	//upload
	beego.Router("/upload", &controllers.UploadController{})
}
