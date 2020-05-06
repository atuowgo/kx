package commons

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strconv"
	"strings"
)

type PointLine struct {
	Year int
	Month int
	Day int
	Lon float64 //经度
	Lats [11]float64 //维度
	LatVals [11]float64 //维度值列表
}

var HEAD = [11]float64{35.25,35.75,36.25,36.75,37.25,37.75,38.25,38.75,39.25,39.75,40.25}

type PointLineQueue []*PointLine

func CreatePointLineQueue() PointLineQueue {
	return make(PointLineQueue,0)
}

func (queue* PointLineQueue) Add(pl *PointLine) {
	*queue = append(*queue,pl)
}

func (queue* PointLineQueue) Pop() *PointLine {
	rs := (*queue)[0]
	*queue = (*queue)[1:]
	return rs
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func MakeSureFileExists(path string) error {
	exits,err := PathExists(path)
	if err != nil {
		panic(err)
	}
	if !exits {
		os.MkdirAll(path,os.ModePerm)
	}
	return nil
}

type PointLineQueueProcess interface {
	NoticeAdd(queue* PointLineQueue,pl *PointLine)
	ExeCalcuate(idx int) bool
	CalculateLatVal(queue* PointLineQueue,idx int) string
}

func Process(dirIn string,dirOut string,indexPrefix string,plqp PointLineQueueProcess) {
	MakeSureFileExists(dirOut)
	indexFile := createIndexFile(dirOut,indexPrefix)
	defer indexFile.Close()
	files,err := ioutil.ReadDir(dirIn)
	if err != nil {
		panic(err)
	}
	for i,file := range files {
		path := dirIn + string(os.PathSeparator) + file.Name()
		fmt.Println("process",path)
		processFile(path,dirOut,indexFile,indexPrefix,plqp,i)
	}
}

func processFile(pathIn string,dirOut string,indexFile *os.File,indexPrefix string,plqp PointLineQueueProcess,fileIdx int) {
	file, err := os.Open(pathIn)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	plq := CreatePointLineQueue()
	fileNameWithSuffix := path.Base(file.Name())
	suffix := path.Ext(file.Name())
	outFiles,itemNames := initOutFileList(dirOut,strings.TrimSuffix(fileNameWithSuffix,suffix),indexPrefix)
	defer func() {
		for _,f := range outFiles {
			f.Close()
		}
	}()
	scanner := bufio.NewScanner(file)
	idx := 0
	for scanner.Scan() {
		lineText := scanner.Text()
		pl := readPointLine(lineText)
		//fmt.Println(pl)
		if idx == 0 {
			appendIndex(indexFile,itemNames,pl,fileIdx * 11)
		}
		plqp.NoticeAdd(&plq,pl)
		plq.Add(pl)
		if plqp.ExeCalcuate(idx) {
			for k,f := range outFiles {
				rs := plqp.CalculateLatVal(&plq,k)
				f.WriteString("\r\n")
				f.WriteString(rs)
			}
		}
		idx++
	}
}

func initOutFileList(dirOut string,name string,indexPrefix string) ([11]*os.File,[11]string) {
	res := [11]*os.File{}
	itemNames := [11]string{}
	for i:=1;i<=11;i++ {
		itemName := indexPrefix+name + fmt.Sprintf("%02v", i)
		path := dirOut + string(os.PathSeparator) + itemName +".txt"
		file,err := os.Create(path)
		if err != nil {
			panic(err)
		}
		file.WriteString("19790101")
		res[i-1] = file
		itemNames[i-1] = itemName
	}

	return res,itemNames
}

func createIndexFile(dirOut string,indexPrefix string) *os.File {
	path := dirOut + string(os.PathSeparator) + indexPrefix+".txt"
	file,err := os.Create(path)
	if err!=nil {
		panic(err)
	}
	file.WriteString("ID,NAME,LAT,LONG,ELEVATION")
	return file
}

func appendIndex(file *os.File,itemNames [11]string,pl *PointLine,startIdx int) {
	for i,itemName:= range itemNames {
		file.WriteString("\r\n")
		file.WriteString(fmt.Sprintf("%v,%v,%v,%v,NULL",startIdx+i+1,itemName,pl.Lon,pl.Lats[i]))
	}
}

func readPointLine(l string) *PointLine {
	pl := &PointLine{}
	cols := strings.Fields(l)
	pl.Year,_ = strconv.Atoi(cols[0])
	pl.Month,_ = strconv.Atoi(cols[1])
	pl.Day,_ = strconv.Atoi(cols[2])
	pl.Lats = HEAD
	fmt.Sscanf(cols[3], "%e", &pl.Lon)
	//fmt.Println(pl)
	for i:=0; i<len(pl.LatVals) ;i++  {
		fmt.Sscanf(cols[4+i], "%e", &pl.LatVals[i])
	}

	return pl
}
