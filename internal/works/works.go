package works

import (
	"encoding/json"
	"github.com/antchfx/htmlquery"
	"strconv"
	"strings"
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
	Title string `json:"id"`
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
	author := htmlquery.FindOne(doc, "//span[@id='author-information-activityName']/a/text()")
	genre := htmlquery.FindOne(doc, "//dd[@itemprop='genre']/a/text()")
	date := htmlquery.FindOne(doc, "//time[@itemprop='datePublished']")
	star := htmlquery.FindOne(doc, "//span[@class='js-total-review-point-element']/text()")
	tags := htmlquery.Find(doc, "//span[@itemprop='keywords']/a/text()")

	var labels []string
	for _, item := range tags {
		labels = append(labels, htmlquery.InnerText(item))
	}

	star_i, err := strconv.Atoi(htmlquery.InnerText(star))
	if err != nil {
		return
	}

	w := *work
	meta := Metadata{
		Title:         htmlquery.InnerText(title),
		Author:        htmlquery.InnerText(author),
		Genre:         htmlquery.InnerText(genre),
		PublishedDate: htmlquery.InnerText(date),
		Star:          star_i,
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

func (work *Work) GetContent(path string) (err error) {
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
	content := Content{
		Id:    strings.Split(id, `.`)[0],
		Title: htmlquery.InnerText(title),
		Body:  strings.Join(bodies, " "),
	}

	w := *work
	w.Contents = append(w.Contents, content)
	*work = w
	return
}
