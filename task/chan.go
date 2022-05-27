package task

import (
	"fmt"
	"path"
	"path/filepath"
	"training/imgproc/filter"
)

type ChanTask struct {
	dirCtx
	Filter   filter.Filter
	PoolSize int
}

func NewChanTask(srcDir, dstDir string, filter filter.Filter, poolSize int) Tasker {
	return &ChanTask{
		Filter: filter,
		dirCtx: dirCtx{
			SrcDir: srcDir,
			DstDir: dstDir,
			files:  BuildFileList(srcDir),
		},
		PoolSize: poolSize,
	}
}

type jobReq struct {
	src string
	dst string
}

func worker(id int, chanTask *ChanTask, jobs <-chan jobReq, results chan<- string) {
	for j := range jobs {
		fmt.Printf("worker %d, started job %v\n", id, j)
		chanTask.Filter.Process(j.src, j.dst)
		fmt.Printf("worker %d, finished job %v\n", id, j)
		results <- j.dst
	}
}

func (c *ChanTask) Process() error {
	size := len(c.files)
	jobs := make(chan jobReq, size)
	results := make(chan string, size)

	// init workers
	for w := 1; w <= c.PoolSize; w++ {
		go worker(w, c, jobs, results)
	}

	for _, f := range c.files {
		filename := filepath.Base(f)
		dst := path.Join(c.DstDir, filename)
		jobs <- jobReq{
			src: f,
			dst: dst,
		}
	}
	close(jobs)

	for range c.files {
		fmt.Println(<-results)
	}

	return nil
}
