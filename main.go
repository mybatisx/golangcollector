package main

import (
	"cookbook/collector"
	"os"
)

func main() {



	//ch := make(chan struct{})
	//global.MyGlobal().Server.Run()
	//ch <- struct{}{}
	//imgUrl:=collector.UploadImg("http://img2.bdstatic.com/img/image/166314e251f95cad1c8f496ad547d3e6709c93d5197.jpg")
	//log.Printf(imgUrl)
 if os.Getenv("WHOAMI") != "ALIYUN"{
	 box:= collector.Box{}
	 box.Run()
 }


}