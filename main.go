package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
)

type item map[string]string

var collect []item

func main() {
	now := time.Now()
	c := colly.NewCollector()
	extensions.RandomUserAgent(c)
	extensions.Referer(c)
	url := `https://suumo.jp/jj/bukken/ichiran/JJ012FC001/?ar=030&bs=011&cn=9999999&cnb=0&ekTjCd=&ekTjNm=&kb=1&kt=9999999&mb=0&mt=9999999&ta=13&tj=0&po=0&pj=1&pc=100&pn=%d`
	c.OnHTML(`.dottable.dottable--cassette`, func(e *colly.HTMLElement) {
		tmp := make(item)
		e.ForEach(`.dottable-line`, func(_ int, e *colly.HTMLElement) {
			e.ForEach(`dl`, func(_ int, e *colly.HTMLElement) {
				tmp[e.ChildText(`dt`)] = e.ChildText(`dd`)
			})
		})
		collect = append(collect, tmp)
	})
	for i := 1; i < 230; i++ {
		c.Visit(fmt.Sprintf(url, i))
	}
	fName := "data.csv"
	file, _ := os.Create(fName)
	defer file.Close()
	json, _ := json.Marshal(collect)
	io.WriteString(file, string(json))
	fmt.Printf("spent time => %f s\n", time.Since(now).Minutes())
	fmt.Printf("number of data=> %d \n", len(collect))
}
