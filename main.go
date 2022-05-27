package main

import (
	"training/imgproc/filter"
	"training/imgproc/task"
)

func main() {
	var f filter.Filter = filter.Grayscale{}
	t := task.NewWaitGrpTask("./imgs", "output", f)
	t.Process()

}
