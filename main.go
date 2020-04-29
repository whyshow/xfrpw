package main

import (
	"./controllers"
	"./models"
	_ "./routers"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/time/rate"
	"strings"
	"sync"
)

func init() {
	url := beego.AppConfig.String("jdbc.url")
	port := beego.AppConfig.String("jdbc.port")
	username := beego.AppConfig.String("jdbc.username")
	password := beego.AppConfig.String("jdbc.password")
	orm.RegisterDataBase("default", "mysql", username+":"+password+"@tcp("+url+":"+port+")/xfrpw?charset=utf8")
	orm.RegisterModel(new(models.Ccit_admin))
	orm.RegisterModel(new(models.Ccit_classify))
	orm.RegisterModel(new(models.Ccit_course))
	orm.RegisterModel(new(models.Ccit_user))
	orm.RegisterModel(new(models.Ccit_course0))
	orm.RegisterModel(new(models.Ccit_section))
	orm.RegisterModel(new(models.Ccit_knobble))
}

func main() {
	beego.ErrorController(&controllers.ErrorController{})
	//注册过滤器  判断在用户访问admin后台时是否处于登录状态，如果没有登录，则跳转到 登录页面
	beego.InsertFilter("/admin/*", beego.BeforeRouter, FilterUser)
	//IP过滤
	beego.InsertFilter("/*", beego.BeforeRouter, IPFiltration)
	//打开session
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.Run()
}

//管理后台的登录过滤中间件
var FilterUser = func(ctx *context.Context) {
	_, ok := ctx.Input.Session("admin_id").(string)
	ok2 := strings.Contains(ctx.Request.RequestURI, "/login")
	if !ok && !ok2 {
		ctx.Redirect(302, "/admin/login")
	}
}

//IP过滤中间件
var limiter = NewIPRateLimiter(1, 10)

var IPFiltration = func(ctx *context.Context) {
	limiter := limiter.GetLimiter(ctx.Request.RemoteAddr)
	if !limiter.Allow() {
		fmt.Println("拦截请求")
		return
	}
}

//令牌桶中间件
// IP令牌桶结构体模型 .
type IPRateLimiter struct {
	ips map[string]*rate.Limiter //限制器
	mu  *sync.RWMutex            //读写锁
	r   rate.Limit
	b   int //数量
}

// 新建一个IP令牌桶结构体 .
func NewIPRateLimiter(r rate.Limit, b int) *IPRateLimiter {
	return &IPRateLimiter{
		ips: make(map[string]*rate.Limiter),
		mu:  &sync.RWMutex{},
		r:   r,
		b:   b,
	}
}

// AddIP 创建了一个新的速率限制器，并将其添加到 ips 映射中,
// 使用 IP地址作为密钥
func (i *IPRateLimiter) AddIP(ip string) *rate.Limiter {
	i.mu.Lock()
	defer i.mu.Unlock()
	limiter := rate.NewLimiter(i.r, i.b)
	i.ips[ip] = limiter
	return limiter
}

// GetLimiter 返回所提供的IP地址的速率限制器(如果存在的话).
// 否则调用 AddIP 将 IP 地址添加到映射中
func (i *IPRateLimiter) GetLimiter(ip string) *rate.Limiter {
	//结构体锁
	i.mu.Lock()
	limiter, exists := i.ips[ip]
	if !exists {
		i.mu.Unlock()
		return i.AddIP(ip)
	}
	i.mu.Unlock()
	return limiter
}
