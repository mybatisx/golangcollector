package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
)

type Util struct {
	values int
}

func (*Util) UploadImg(src string) error  {
	imgPath := "/Users/ming/mingtemp/demo/"
	imgUrl := "http://img2.bdstatic.com/img/image/166314e251f95cad1c8f496ad547d3e6709c93d5197.jpg"

	fileName := path.Base(imgUrl)


	res, err := http.Get(imgUrl)
	if err != nil {
		fmt.Println("A error occurred!")
		return err
	}
	defer res.Body.Close()
	// 获得get请求响应的reader对象

	body, _ := ioutil.ReadAll(res.Body)
	file, err := os.Create(imgPath + fileName)
	if err != nil {
		panic(err)
	}
	_,err=io.Copy(file, bytes.NewReader(body))


	return nil
}



