package travel

import (
	"fmt"
	"os"
	"strings"

	"github.com/schollz/progressbar/v3"

	"io/ioutil"
	"path/filepath"

	"github.com/Realive333/kakuyomu_cleaner/pkg/file"
)

func Folder(dir string) (err error) {
	folders, err := ioutil.ReadDir(dir)
	if err != nil {
		return
	}

	rdirs := strings.Split(dir, `\`)
	dpath := fmt.Sprintf("./cleaned/%s", rdirs[len(rdirs)-1])
	path := fmt.Sprintf(`%s\file_test.jsonl`, dpath)

	if _, err := os.Stat(dpath); err == nil {
		err := os.RemoveAll(dpath)
		if err != nil {
			return err
		}
	}

	err = os.MkdirAll(dpath, os.ModePerm)
	if err != nil {
		return
	}

	bar := progressbar.Default(int64(len(folders)))
	for _, folder := range folders {
		work, err := file.ReadAllHTML(filepath.Join(dir, folder.Name()))
		if err != nil {
			return err
		}
		result, err := work.ToJson()
		if err != nil {
			return err
		}

		err = file.AppendJSON(path, result)
		if err != nil {
			return err
		}
		bar.Add(1)
	}
	return
}
