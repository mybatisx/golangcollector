package collector

import (
	"bytes"
	"database/sql"
	"fmt"
	"github.com/disintegration/imaging"
	"golang.org/x/tools/go/ssa/interp/testdata/src/errors"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"strings"
)

type Util struct {
	values int
}

func ResizeImg(path5 string) string {

	connStr := "postgres://postgres:123456@123.57.206.19:10001/lanren?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	var id int

	err = db.QueryRow(`insert into ider (name)   values ($1) RETURNING id`, "a").Scan(&id)
	if err != nil {
		log.Panic(errors.New(err.Error()))
	}

	// Open a test image.
	path1 := path.Base(path5)
	extA := strings.Split(path1, ".")
	ext := extA[len(extA)-1]
	src, err := imaging.Open(path5)
	if err != nil {
		log.Fatalf("failed to open image: %v", err)
	}
	destinationFile := fmt.Sprintf("/Users/ming/mingtemp/images/%d.%s", id, ext)
	if src.Bounds().Max.X > 220 {
		src = imaging.Resize(src, 220, 0, imaging.Lanczos)
		imaging.Save(src, destinationFile)
	} else {
		input, err := ioutil.ReadFile(path5)
		if err != nil {
			log.Panic(err.Error())
		}

		err = ioutil.WriteFile(destinationFile, input, 0644)
		if err != nil {
			log.Panic(err.Error())
		}

	}
	err = os.Remove(path5)

	if err != nil {
		log.Panic(err.Error())
	}
	return fmt.Sprintf("//asset.lanrenshipu.cn/images/%d.%s", id, ext)
}
func UploadImg(src string) string {
	imgPath := "/Users/ming/mingtemp/pk/"
	imgUrl := src

	fileName := path.Base(imgUrl)

	res, err := http.Get(imgUrl)
	if err != nil {
		log.Panic(err.Error())
	}
	defer res.Body.Close()
	// 获得get请求响应的reader对象

	body, _ := ioutil.ReadAll(res.Body)
	file, err := os.Create(imgPath + fileName)
	if err != nil {
		log.Panic(err.Error())
	}
	_, err = io.Copy(file, bytes.NewReader(body))
if err != nil{
	log.Panic(err.Error())
}
	fullPath := imgPath + fileName
	return ResizeImg(fullPath)
}
