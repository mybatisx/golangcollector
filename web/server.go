package web

import (
	"cookbook/collector"
	db2 "cookbook/db"
	"encoding/base64"
	"fmt"
	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris/v12"
	"strings"

	"html/template"
	"io/ioutil"
	log2 "log"
	"strconv"
)

type Server struct {
	Port int
}

func (s *Server) start() error {
	app := iris.New()
	app.Logger().SetLevel("debug")
	app.UseGlobal(cors.AllowAll())
	app.HandleDir("/static", "./web/assets")
	//app.HandleDir("/page", "./assets/page")
	tmpl := iris.HTML("./web/assets/page", ".html").Layout("layout.html")
	tmpl.Reload(true)
	app.RegisterView(tmpl)
	app.Get("/", func(ctx iris.Context) {
		list := getRandomRows(20,``)
		ctx.ViewData("list", list)
		ctx.ViewData("title","豆饼网_菜谱_菜谱大全_食谱美食")
		ctx.ViewData("keywords","菜谱,菜谱大全,菜谱做法,家常菜,食谱,美食,豆饼网食谱")
		ctx.ViewData("desc","豆饼网是最优质的美食菜谱,提供各种菜谱大全,食谱大全,家常菜做法大全,丰富的菜谱大全可以让您轻松地学会怎么做美食,展现自己的高超厨艺,开启美好生活！")
		if err := ctx.View("index.html"); err != nil {
			ctx.StatusCode(iris.StatusInternalServerError)
			ctx.Writef(err.Error())
		}
	})
	app.Get("/cookbook/{id:string regexp(^[0-9]{1,10}\\.html)}", func(ctx iris.Context) {
		idStr:= ctx.Params().GetStringDefault("id","")
		idStr= strings.Replace(idStr,".html","",1)
		id,err:=strconv.Atoi(idStr)
		var name, brief, content, img, material string
		err = db2.GetDb().Conn.QueryRow(`select name,brief,content,img,material  from shipu where id= $1`, id).
			Scan(&name, &brief, &content, &img, &material)

		print(err)
		ctx.ViewData("title",name+"怎么做_好吃__详细_图文_视频_步骤_家庭版_懒人食谱")
		ctx.ViewData("keywords",name+"详细做法_图解视频_懒人食谱")
		ctx.ViewData("desc",name+"详细做法_图解视频_视频步骤_懒人食谱")

		list := getRandomRows(4,``)
		ctx.ViewData("list", list)
		ctx.ViewData("name", name)
		ctx.ViewData("brief", brief)
		ctx.ViewData("img", img)
		ctx.ViewData("material", template.HTML(material))
		ctx.ViewData("content", template.HTML(content))
		if err := ctx.View("one.html", ); err != nil {
			ctx.StatusCode(iris.StatusInternalServerError)
			ctx.Writef(err.Error())
		}
	})

	app.Get("/search/{name:string}", func(ctx iris.Context) {
		name0 := ctx.Params().GetStringDefault("name","")

		var name, brief, content, img, material string

		ctx.ViewData("title",name0+"怎么做_好吃__详细_图文_视频_步骤_家庭版_豆饼网")
		ctx.ViewData("keywords",name0+"详细做法_图解视频_豆饼网")
		ctx.ViewData("desc",name0+"详细做法_图解视频_视频步骤_豆饼网")

		list := getRandomRows(20,name0)
		ctx.ViewData("list", list)
		ctx.ViewData("name", name)
		ctx.ViewData("keyword",name0)
		ctx.ViewData("brief", brief)
		ctx.ViewData("img", img)
		ctx.ViewData("material", template.HTML(material))
		ctx.ViewData("content", template.HTML(content))
		if err := ctx.View("search.html", ); err != nil {
			ctx.StatusCode(iris.StatusInternalServerError)
			ctx.Writef(err.Error())
		}
	})



	app.Post("/upload", func(ctx iris.Context) {

		var upFile collector.UpFile
		err:=ctx.ReadJSON(&upFile)
		if err != nil {
			ctx.JSON(iris.Map{"code": 3, "msg":  err.Error()})
			return
		}
		ddd, err := base64.StdEncoding.DecodeString(upFile.Base64Str)
		if err != nil {
			ctx.JSON(iris.Map{"code": 5, "msg":  err.Error()})
			return
		}
		err = ioutil.WriteFile(fmt.Sprintf("/home/assets/images/%s",upFile.Name), ddd, 0666)
		if err != nil {
			ctx.JSON(iris.Map{"code": 6, "msg":  err.Error()})
			return
		}
		ctx.JSON(iris.Map{"code": 0, "msg": ""})
	})

	app.Listen("0.0.0.0:" + strconv.Itoa(s.Port))
	return nil
}
func getRandomRows(count int32,keyword string) []collector.CookInfo {
	list := make([]collector.CookInfo, 0)
	sql2 := fmt.Sprintf(`SELECT id, name,img,material,brief,content  from shipu order by random() LIMIT %d `, count)
	if keyword != ``{
		sql2 = fmt.Sprintf(`SELECT id, name,img,material,brief,content  from shipu where name like '%s' LIMIT %d `, "%" + keyword + "%", count)

	}

	rows, err := db2.GetDb().Conn.Query(sql2)
	if err != nil {
		log2.Print(err.Error())
		return list
	}
	var user collector.CookInfo
	for rows.Next() {
		rows.Scan(&user.Id,&user.Name, &user.Img, &user.Material, &user.Brief, &user.Content)
		list = append(list, user)
	}
	return list
}
func (s *Server) Run() error {

	go func() {
		s.start()
	}()

	return nil
}
