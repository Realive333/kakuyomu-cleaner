package test

import (
	"github.com/Realive333/kakuyomu_cleaner/internal/works"
	"github.com/Realive333/kakuyomu_cleaner/pkg/file"
	"testing"
)

func Test_ReadAllHTML(t *testing.T) {
	_, err := file.ReadAllHTML(`..\works\16816927619978032613`)
	if err != nil {
		t.Errorf("ReadAllHTML error: %v", err)
	}
	//t.Logf("%v", actual)
}

func Test_AppendJSON(t *testing.T) {
	test := works.Work{
		Id: "001",
		Metadata: works.Metadata{
			Title:         "a title",
			Author:        "an author",
			Genre:         "現代ファンタジー",
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
			"幼女",
			"洋ロリ　ロリ娘💛",
			"現代ファンタジー",
		},
	}

	input, err := test.ToJson()
	if err != nil {
		t.Errorf("AppendJSON error: %v", err)
	}

	err = file.AppendJSON("./file_test.jsonl", input)
	if err != nil {
		t.Errorf("AppendJSON error: %v", err)
	}
}
