package collector

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly/v2"
	"log"
	"net/http"
	"regexp"
	"strings"
	"time"
)

type Box struct {

}

func (*Box) Run()  {
	c := colly.NewCollector()

	// Find and visit all links
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {


		match, _ := regexp.MatchString(`/cookbook/(\d+).html`, strings.TrimSpace(e.Attr("href")))
		if match {

			e.Request.Visit(e.Attr("href"))
			time.Sleep(time.Duration(2)*time.Second)
		}

	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
		match, _ := regexp.MatchString(`/cookbook/(\d+).html`, strings.TrimSpace(r.URL.String()))
		if match {

			res, err := http.Get("https://www.douguo.com/cookbook/2441273.html")
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

			img,_ := doc.Find("#banner .wb100").Attr("src")
			title,_ := doc.Find("#banner .wb100").Attr("alt")
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
			sb:= strings.Builder{}
			doc.Find("#left .retamr  td").Each(func(i int, s *goquery.Selection) {
				// For each item found, get the band and title
				scname := s.Find("a").Text()
				scnum := s.Find(".scnum").Text()
				sb.WriteString("<b>"+scname+"<b/>")
				sb.WriteString("<i>"+scnum+"<i/>")

			})
			fmt.Print(sb.String())
			sb2:= strings.Builder{}
			sb2.WriteString(`<div class="step">`)
			doc.Find("#left .step  .stepcont").Each(func(i int, s *goquery.Selection) {
				// For each item found, get the band and title
				sb2.WriteString(`<div class="stepcont clearfix">`)
				imgstep,_ := s.Find("img").Attr("src")
				text ,_:= s.Find(".stepinfo").Html()
				sb2.WriteString(`<img src=">`+strings.TrimSpace(imgstep)+`"/>`)
				sb2.WriteString(`<div class="info">`+strings.TrimSpace(text)+`</div>`)
				sb2.WriteString(`</div>`)
			})
			sb2.WriteString("</div>")
			fmt.Print(sb2.String())
			fmt.Print("5566")
		}

	})

	c.Visit("https://www.douguo.com/ingredients/%E9%BB%84%E7%93%9C")
}
