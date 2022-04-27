package main

import (
	"fmt"
	"os"

	"github.com/Realive333/kakuyomu_cleaner/pkg/travel"
)

func main() {
	args := os.Args[1:]
	path := fmt.Sprintf(`D:\Programs\kakuyomu_analyzer\works\%s`, args[0])
	err := travel.Folder(path)
	if err != nil {
		panic(err)
	}
}
