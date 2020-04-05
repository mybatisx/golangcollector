package collector

import (
	"bytes"
	db2 "cookbook/db"
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly/v2"
	"log"
	"net/http"
	"regexp"
	"strings"
	"time"

	_ "github.com/lib/pq"
)

type Box struct {
}

func (*Box) Run() {
	c := colly.NewCollector()

	// Find and visit all links
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {

		match, _ := regexp.MatchString(`/cookbook/(\d+).html`, strings.TrimSpace(e.Attr("href")))
		if match {

			e.Request.Visit(e.Attr("href"))
			time.Sleep(time.Duration(2) * time.Second)
		}

	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
		match, _ := regexp.MatchString(`/cookbook/(\d+).html`, strings.TrimSpace(r.URL.String()))
		if match {

			res, err := http.Get(r.URL.String())
			if err != nil {
				log.Fatal(err)
			}
			defer res.Body.Close()
			if res.StatusCode != 200 {
				log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
			}

			// Load the HTML document
			doc, err := goquery.NewDocumentFromReader(res.Body)
			if err != nil {
				log.Fatal(err)
			}

			// Find the review items
			cookInfo := &CookInfo{}
			img, _ := doc.Find("#banner .wb100").Attr("src")
			title, _ := doc.Find("#banner .wb100").Attr("alt")
			jianjie := strings.TrimSpace(doc.Find("#left .intro").Text())

			if strings.TrimSpace(img) == "" {
				return
			}
			if strings.TrimSpace(title) == "" {
				return
			}
			if strings.TrimSpace(jianjie) == "" {
				return
			}
			cookInfo.Img = UploadImg(strings.TrimSpace(img))
			cookInfo.Key = strings.TrimSpace(r.URL.String())
			cookInfo.Name = strings.TrimSpace(title)
			cookInfo.Brief = strings.TrimSpace(jianjie)
			sb := strings.Builder{}
			mater := make(map[string]string)
			doc.Find("#left .retamr  td").Each(func(i int, s *goquery.Selection) {
				// For each item found, get the band and title
				scname := s.Find(".scname").Text()
				scnum := s.Find(".scnum").Text()
				sb.WriteString("<b>" + strings.TrimSpace(scname) + "<b/>")
				sb.WriteString("<i>" + strings.TrimSpace(scnum) + "<i/>")
				if strings.TrimSpace(scname) != "" {
					mater[strings.TrimSpace(scname)] = strings.TrimSpace(scnum)
				}

			})
			cookInfo.Material = sb.String()
			sb2 := strings.Builder{}
			sb2.WriteString(`<div class="step">`)
			step := make(map[string]string)
			doc.Find("#left .step  .stepcont").Each(func(i int, s *goquery.Selection) {
				// For each item found, get the band and title
				sb2.WriteString(`<div class="stepcont clearfix">`)
				imgstep, _ := s.Find("img").Attr("src")
				img := UploadImg(strings.TrimSpace(imgstep))
				text, _ := s.Find(".stepinfo").Html()
				sb2.WriteString(`<img src="` + img + `"/>`)
				sb2.WriteString(`<div class="info">` + strings.TrimSpace(text) + `</div>`)
				sb2.WriteString(`</div>`)
				step[strings.TrimSpace(text)] = strings.TrimSpace(imgstep)
			})
			sb2.WriteString("</div>")
			cookInfo.Content = sb2.String()

			bf := bytes.NewBuffer([]byte{})
			jsonEncoder := json.NewEncoder(bf)
			jsonEncoder.SetEscapeHTML(false)
			jsonEncoder.Encode(cookInfo)
			fmt.Println(bf.String())

			fmt.Print("----------------------------------------------------------------------------")

			db := db2.GetDb().Conn

			var id int

			err = db.QueryRow(`select * from shipu where id=$1 limit 1`, id).Scan(&id)

			if id == 0 {
				err = db.QueryRow(`insert into shipu (name,brief,img,material,content,key)
                                         values ($1,$2,$3,$4,$5,$6) RETURNING id`,
					cookInfo.Name, cookInfo.Brief, cookInfo.Img, cookInfo.Material, cookInfo.Content, cookInfo.Key).Scan(&id)
				if err == nil {
					fmt.Printf("新增id %d", id)
				} else {
					fmt.Printf("新增出错 %s", err.Error())
				}

			} else {

				fmt.Printf("已经存在 %s", cookInfo.Key)
			}

			fmt.Print("----------------------------------------------------------------------------")
		}

	})

	c.Visit("https://www.douguo.com/ingredients/%E9%BB%84%E7%93%9C")
}
