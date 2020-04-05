package global

import (
	"cookbook/collector"
	"cookbook/web"
	"sync"
)

type Global struct {
	Server *web.Server
	//Util *Util
	Box *collector.Box

	//DBhelper *db.DBhelper
}

var once sync.Once
var global *Global
func MyGlobal() *Global {
	once.Do(func() {
		global = new(Global)
		global.Server = &web.Server{ Port :8080}
		//global.Util=new(Util)
		global.Box=new(collector.Box)
		//global.DBhelper = db.GetDb()
	})
	return global
}

