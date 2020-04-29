package routers

import (
	"../controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.WebController{}, "GET:Home")             // web/ 门户首页页面
	beego.Router("/classify", &controllers.WebController{}, "GET:Classify") // web/ 专业方向页面
	beego.Router("/learn", &controllers.WebController{}, "GET:Learn")       // web/ 课程页面
	beego.Router("/play", &controllers.WebController{}, "GET:Play")         // web/ 播放页面
	beego.Router("/login", &controllers.WebController{}, "GET:Login")       // web/ 用户登录页面

	beego.Router("/admin/login", &controllers.AdminUserController{}, "GET:LoginPage")       // admin/登录页面
	beego.Router("/admin/*", &controllers.AdminUserController{}, "GET:Home")                // admin 管理页面首页
	beego.Router("/admin/login_request", &controllers.AdminUserController{}, "POST:Login")  // admin/登录请求
	beego.Router("/admin/exit", &controllers.AdminUserController{}, "GET:Exit")             // admin/退出登录
	beego.Router("/admin/statistics", &controllers.AdminUserController{}, "GET:Statistics") // admin/数据统计
	beego.Router("/admin/queryusers", &controllers.UserController{}, "GET:QueryUsers")
	beego.Router("/admin/queryuser", &controllers.UserController{}, "GET:QueryUser")
	beego.Router("/admin/adduserpage", &controllers.UserController{}, "GET:AddUserPage")
	beego.Router("/admin/addusers", &controllers.TjUserController{})

	beego.Router("/admin/removesection", &controllers.AdminRemoveController{}, "POST:RemoveSection") //删除章节

	beego.Router("/admin/direction", &controllers.DController{})            // admin/专业方向管理页面
	beego.Router("/admin/tjzyfx", &controllers.TjzyfxController{})          // admin/添加专业方向
	beego.Router("/admin/getDirection", &controllers.DirectionController{}) // admin/获得单个方向的所有课程

	beego.Router("/admin/removefx", &controllers.RemoveFController{})  // admin/删除方向
	beego.Router("/admin/removekc", &controllers.RemoveKCController{}) // admin/删除方向

	beego.Router("/admin/alterfxpage", &controllers.AlterFxPageController{}) // admin/修改方向页面
	beego.Router("/admin/alterfx", &controllers.AlterFxController{})         // admin/修改方向
	beego.Router("/admin/alterkcpage", &controllers.AlterKcPageController{})
	beego.Router("/admin/alterk", &controllers.AlterKcController{})

	beego.Router("/admin/course/*", &controllers.CourseController{})   // admin/课程管理页面
	beego.Router("/admin/tjkc", &controllers.TjkcController{})         // admin/添加课程页面
	beego.Router("/admin/tjkcdata", &controllers.TjkcDataController{}) // admin/向数据库添加课程

	beego.Router("/admin/section/", &controllers.SectionController{})      // admin/章节管理页面
	beego.Router("/admin/getCourse", &controllers.GetCourseController{})   // admin/根据条件查询课程列表
	beego.Router("/admin/tjzj", &controllers.TjzjPageController{})         // admin/添加章节
	beego.Router("/admin/getSection", &controllers.GetSectionController{}) // admin/添加章节

	beego.Router("/updata", &controllers.TestController{})          //
	beego.Router("/admin/tjzjdata", &controllers.UpZipController{}) // 测试
	beego.Router("/admin/logs", &controllers.LogsController{})      // 测试

	// android设备请求
	beego.Router("/api/classify", &controllers.AndroidController{}, "GET:Classify") // 测试
	beego.Router("/api/course", &controllers.AndroidController{}, "POST:Course")
	beego.Router("/api/learn", &controllers.AndroidController{}, "POST:Learn")
	beego.Router("/api/play", &controllers.AndroidController{}, "POST:Play")
	beego.Router("/api/login", &controllers.AndroidController{}, "POST:Login")
	beego.Router("/api/findpassword", &controllers.AndroidController{}, "POST:Findpassword")
}
