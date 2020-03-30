package web

import (
	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris/v12"
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

	tmpl := iris.HTML("./web/assets/page", ".html") //.Layout("layout/layout.html")
	tmpl.Reload(true)
	app.RegisterView(tmpl)
	app.Get("/", func(ctx iris.Context) {
		ctx.ViewLayout(iris.NoLayout)
		//ctx.View("page/index.html")
		//ctx.Redirect("/page/index.html")

		if err := ctx.View("home.html"); err != nil {
			ctx.StatusCode(iris.StatusInternalServerError)
			ctx.Writef(err.Error())
		}
	})
	app.Get("/shipu/{pagename}", func(ctx iris.Context) {
		pagename := ctx.Params().Get("pagename")

		if err := ctx.View(pagename); err != nil {
			ctx.StatusCode(iris.StatusInternalServerError)
			ctx.Writef(err.Error())
		}
	})


	app.Get("/test", func(ctx iris.Context) {
		age := ctx.URLParamInt64Default("age",0)
		slice := make([]iris.Map, 0)
		slice = append(slice,iris.Map{"username":"nihao","id":1,"age":age})
		slice = append(slice,iris.Map{"username":"nihao2","id":2,"age":age})
		ctx.JSON(iris.Map{"code":0,"msg":"","count":1,"data":slice})
	})



	app.Listen("0.0.0.0:"+strconv.Itoa(s.Port))
	return nil
}
func (s *Server) Run() error {

	go func() {
		s.start()
	}()

	return nil
}