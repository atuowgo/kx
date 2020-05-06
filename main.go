package main

import (
	"flag"
	"fmt"
	"kx/commons"
	"kx/fsds"
	"kx/prect"
	"kx/qbot"
	"kx/tbot"
	"kx/wind"
	"os"
	"runtime"
	"sync"
)

type ProcessType int

const (
	FSDS = iota
	PRECT
	TBOT
	QBOT
	WIND
	ALL
)

func process(t ProcessType,dirIn string,dirOut string) {
	var process commons.PointLineQueueProcess
	var indexPrefix string
	var subDirInName string
	var subDirOutName string
	switch t {
	case FSDS:
		process = fsds.FSDSPointLineQueueProcess{}
		indexPrefix = "slr"
		subDirInName = "FSDS"
		subDirOutName = subDirInName
	case PRECT:
		process = prect.PRECTPointLineQueueProcess{}
		indexPrefix = "pcp"
		subDirInName = "PRECT"
		subDirOutName = subDirInName
	case TBOT:
		process = tbot.TBOTPointLineQueueProcess{}
		indexPrefix = "temp"
		subDirInName = "TBOT"
		subDirOutName = subDirInName
	case WIND:
		process = wind.WINDPointLineQueueProcess{}
		indexPrefix = "wind"
		subDirInName = "QBOT_WIND"
		subDirOutName = "WIND"
	}


	if t == ALL {
		processAll(dirIn,dirOut)
	} else if t == QBOT {
		processQBOT(dirIn,dirOut,true)
	}else {
		dirIn = dirIn + string(os.PathSeparator) + subDirInName
		dirOut = dirOut + string(os.PathSeparator) + subDirOutName
		commons.Process(dirIn,dirOut,indexPrefix,process)
	}
}

func processQBOT(dirIn string,dirOut string,setProcs bool) {
	if setProcs {
		runtime.GOMAXPROCS(3)
	}
	wg := sync.WaitGroup{}
	wg.Add(3)
	go func() {
		commons.Process(dirIn + string(os.PathSeparator) + "TBOT",dirOut + string(os.PathSeparator) + "TBOT_PER","rh",qbot.PerPointLineQueueProcess{})
		wg.Done()
	}()

	go func() {
		commons.Process(dirIn + string(os.PathSeparator) + "PSRF",dirOut + string(os.PathSeparator) + "PSRF_PER","rh",qbot.PerPointLineQueueProcess{})
		wg.Done()
	}()

	go func() {
		commons.Process(dirIn + string(os.PathSeparator) + "QBOT_WIND",dirOut + string(os.PathSeparator) + "QBOT_PER","rh",qbot.QBOTPerPointLineQueueProcess{})
		wg.Done()
	}()

	wg.Wait()

	var dirIns = [3]string{
		dirOut + string(os.PathSeparator) + "TBOT_PER",
		dirOut + string(os.PathSeparator) + "PSRF_PER",
		dirOut + string(os.PathSeparator) + "QBOT_PER",
	}
	qbot.ProcessQbot(dirIns,dirOut + string(os.PathSeparator) + "QBOT","rh.txt")

	for _,dir := range dirIns {
		fmt.Println("remove temp dir",dir)
		os.RemoveAll(dir)
	}
}

func processAll(dirIn string,dirOut string) {
	runtime.GOMAXPROCS(4)
	wg := sync.WaitGroup{}
	wg.Add(4)

	go func() {
		process(FSDS,dirIn,dirOut)
		wg.Done()
	}()

	go func() {
		process(PRECT,dirIn,dirOut)
		wg.Done()
	}()

	go func() {
		process(TBOT,dirIn,dirOut)
		wg.Done()
	}()

	go func() {
		process(WIND,dirIn,dirOut)
		wg.Done()
	}()

	wg.Wait()

	processQBOT(dirIn,dirOut,false)
}

var (
	h bool

	in string
	out string
	t int
)

func init() {
	flag.BoolVar(&h,"h",false,"help")
	flag.StringVar(&in,"in","","数据文件路径,该路径下包含要处理类型的子目录，如/xxxx/aa/FDS，则in为/xxx/aa")
	flag.StringVar(&out,"out","","输出路径")
	flag.IntVar(&t,"t",5,"0:FSDS;1:PRECT;2:TBOT;3:QBOT;4:WIND;5:ALL")
}

func main() {
	//dirIn := "/Users/adam/dev/temp/1901-2016"
	//dirOut := "/Users/adam/go/src/kx/out"
	//process(ALL,dirIn,dirOut)
	flag.Parse()
	if h {
		usage()
		return
	}

	pt := ProcessType(t)
	if len(in) == 0 {
		fmt.Println("输入路径不能为空")
		return
	}
	if len(out) == 0 {
		fmt.Println("输出路径不能为空")
		return
	}
	process(pt,in,out)
	fmt.Println("done ...")

}

func usage() {
	fmt.Fprintf(os.Stderr, `科学家自救指南
Usage: kx [-h] [-in dirIn -out dirOut -t type]

Options:
`)
	flag.PrintDefaults()
}
