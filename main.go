package main

import (
	"cookbook/collector"
	"cookbook/global"
	"os"
)

func main() {


 if os.Getenv("WHOAMI") == "ALIYUN"{

	 ch := make(chan struct{})
	 global.MyGlobal().Server.Run()
	 ch <- struct{}{}
 } else{
	 box:= collector.Box{}
	 box.Run()

 }


}