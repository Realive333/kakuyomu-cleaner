package file

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

	"github.com/Realive333/kakuyomu_cleaner/internal/works"
)

func ReadAllHTML(dir string, flag int) (work *works.Work, err error) {
	work = &works.Work{}
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return
	}

	for _, f := range files {
		if f.Name() == "abstract.html" {
			err = work.GetAbstract(filepath.Join(dir, f.Name()))
			if err != nil {
				return nil, err
			}
		} else if f.Name() == "record.json" {
			continue
		} else {
			err = work.GetContent(filepath.Join(dir, f.Name()), flag)
			if err != nil {
				return nil, err
			}
		}
		time.Sleep(1 * time.Millisecond)
	}
	return
}

func AppendJSON(path string, input []byte) (err error) {
	f, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return
	}
	defer f.Close()

	input = append(input, '\n')
	if _, err = f.Write(input); err != nil {
		return
	}
	return
}
