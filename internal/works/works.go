package works

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/antchfx/htmlquery"
)

type Work struct {
	Id       string    `json:"id"`
	Metadata Metadata  `json:"metadata"`
	Contents []Content `json:"content"`
	Labels   []string  `json:"labels"`
}

type Metadata struct {
	Title         string `json:"title"`
	Author        string `json:"author"`
	Genre         string `json:"genre"`
	Star          int    `json:"star"`
	PublishedDate string `json:"published_date"`
}

type Content struct {
	Id    string `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

func (work *Work) ToJson() (result []byte, err error) {
	result, err = json.Marshal(&work)
	if err != nil {
		return
	}
	return
}

func (work *Work) GetAbstract(path string) (err error) {
	doc, err := htmlquery.LoadDoc(path)
	if err != nil {
		return
	}

	ids := strings.Split(path, `\`)
	id := ids[len(ids)-2]

	title := htmlquery.FindOne(doc, "//h4[@class='heading-level5']/a/text()")
	if title == nil {
		return fmt.Errorf("title did not exist")
	}
	author := htmlquery.FindOne(doc, "//span[@id='author-information-activityName']/a/text()")
	if author == nil {
		return fmt.Errorf("author did not exist")
	}
	genre := htmlquery.FindOne(doc, "//li[@id='workGenre']/a/text()")
	if genre == nil {
		return fmt.Errorf("genre did not exist")
	}
	date := htmlquery.FindOne(doc, "//time[@itemprop='datePublished']")
	if date == nil {
		return fmt.Errorf("date did not exist")
	}
	star := htmlquery.FindOne(doc, "//span[@class='js-total-review-point-element']/text()")
	if star == nil {
		return fmt.Errorf("star did not exist")
	}
	tags := htmlquery.Find(doc, "//span[@itemprop='keywords']/a/text()")

	var labels []string
	for _, item := range tags {
		labels = append(labels, htmlquery.InnerText(item))
	}

	starS := htmlquery.InnerText(star)
	starS = strings.Replace(starS, ",", "", -1)
	starI, err := strconv.Atoi(starS)
	if err != nil {
		return
	}

	w := *work
	meta := Metadata{
		Title:         htmlquery.InnerText(title),
		Author:        htmlquery.InnerText(author),
		Genre:         htmlquery.InnerText(genre),
		PublishedDate: htmlquery.InnerText(date),
		Star:          starI,
	}

	w = Work{
		Id:       id,
		Metadata: meta,
		Contents: w.Contents,
		Labels:   labels,
	}

	*work = w
	return
}

func (work *Work) GetContent(path string, flag int) (err error) {
	doc, err := htmlquery.LoadDoc(path)
	if err != nil {
		return
	}

	ids := strings.Split(path, `\`)
	id := ids[len(ids)-1]

	title := htmlquery.FindOne(doc, "//p[@class='widget-episodeTitle js-vertical-composition-item']/text()")

	b := htmlquery.Find(doc, "//div[@class='widget-episodeBody js-episode-body']/p/text()")
	var bodies []string
	for _, item := range b {
		bodies = append(bodies, htmlquery.InnerText(item))
	}
	strBody := strings.Join(bodies, "")

	if flag > 1 && len(strBody) > flag {
		strBody = TruncateString(strBody, flag)
	}
	content := Content{
		Id:    strings.Split(id, `.`)[0],
		Title: htmlquery.InnerText(title),
		Body:  strBody,
	}

	w := *work
	w.Contents = append(w.Contents, content)
	*work = w
	return
}

func TruncateString(str string, length int) string {
	if length <= 0 {
		return ""
	}
	truncated := ""
	count := 0
	for _, char := range str {
		truncated += string(char)
		count++
		if count >= length {
			break
		}
	}
	return truncated
}
