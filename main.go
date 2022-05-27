package main

import (
	"fmt"
	"training/imgproc/task"
)

func main() {
	f := task.BuildFileList("./imgs")
	fmt.Println(f)
}
