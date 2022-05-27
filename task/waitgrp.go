package task

import "training/imgproc/filter"

type WaitGrpTask struct {
	dirCtx
	Filter filter.Filter
}

func NewWaitGrpTask(srcDir, dstDir string, filter filter.Filter) Tasker {
	return &WaitGrpTask{
		Filter: filter,
		dirCtx: dirCtx{
			SrcDir: srcDir,
			DstDir, dstDir,
			files: BuildFileList(srcDir),
		},
	}
}
