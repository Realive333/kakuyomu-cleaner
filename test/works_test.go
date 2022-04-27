package test

import (
	"github.com/Realive333/kakuyomu_cleaner/internal/works"
	"testing"
)

func Test_Content(t *testing.T) {
	assert := works.Content{
		Id:    "0123456",
		Title: "this is an example title",
		Body:  "this is an example body",
	}

	if assert.Id != "0123456" {
		t.Errorf(`Content value error, expect "0123456", is "%v"`, assert.Id)
	}

	if assert.Title != "this is an example title" {
		t.Errorf(`Content value error, expect "this is an example title", is "%v"`, assert.Title)
	}

	if assert.Body != "this is an example body" {
		t.Errorf(`Content value error, expect "this is an example body", is "%v"`, assert.Body)
	}
}

func Test_ToJson(t *testing.T) {
	test := works.Work{
		Id: "001",
		Metadata: works.Metadata{
			Title:         "a title",
			Author:        "an author",
			Genre:         "SF",
			Star:          999,
			PublishedDate: "2021-01-01",
		},
		Contents: []works.Content{
			works.Content{
				Id:    "a01",
				Title: "a01 title",
				Body:  "a01 body",
			},
			works.Content{
				Id:    "a02",
				Title: "a02 title",
				Body:  "a02 body",
			},
		},
		Labels: []string{
			"lable 1",
			"lable 2",
			"lable 3",
		},
	}

	assert, err := test.ToJson()
	if err != nil {
		t.Errorf("ToJson error: %v", err)
	}
	t.Log(string(assert))
}
