package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/Realive333/kakuyomu_cleaner/pkg/travel"
)

func main() {
	args := os.Args[1:]
	path := fmt.Sprintf(`D:\Programs\kakuyomu_analyzer\works\%s`, args[0])
	flag := -1

	i, err := strconv.Atoi(args[1])
	//fmt.Println(i)
	if err == nil {
		flag = i
	}
	//fmt.Println(flag)
	err = travel.Folder(path, flag)
	if err != nil {
		panic(err)
	}
}
