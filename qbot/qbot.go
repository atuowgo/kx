package qbot

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"kx/commons"
	"math"
	"os"
)

var rhQueue []float64

func add(rh float64) {
	if len(rhQueue) == 4 {
		rhQueue = rhQueue[1:]
	}
	rhQueue = append(rhQueue,rh)
}

func ProcessQbot(dirIn/*T,P,Q*/ [3]string,dirOut string,indexName string) {
	dirInT := dirIn[0]
	fileInfos,err := ioutil.ReadDir(dirInT)
	if err != nil {
		panic(err)
	}
	commons.MakeSureFileExists(dirOut)
	for _,fileInfo := range fileInfos {
		name := fileInfo.Name()
		pathOut := dirOut + string(os.PathSeparator) + name
		if name == indexName {
			src,err := os.Open(dirIn[0] + string(os.PathSeparator) + name)
			if err != nil {
				panic(err)
			}
			dst,err := os.Create(pathOut)
			if err != nil {
				panic(err)
			}
			io.Copy(dst,src)
			continue
		}
		var pathIn [3]string
		for i,dir := range dirIn {
			pathIn[i] = dir + string(os.PathSeparator) + name
		}
		processQbotFile(pathIn,pathOut)
	}
}

func processQbotFile(pathIn/*T,P,Q*/ [3]string,pathOut string) {
	filesIn := openFiles(pathIn)
	defer func() {
		for _,f := range filesIn {
			f.Close()
		}
	}()

	fileOut,err := os.Create(pathOut)
	if err != nil {
		panic(err)
	}
	defer fileOut.Close()

	var scanners [3]*bufio.Scanner
	for i,f := range filesIn {
		scanners[i] = bufio.NewScanner(f)
	}

	idx := 0
	for ;; {
		var items [3]float64
		needBreak := false
		for i, scanner := range scanners {
			hasNext := scanner.Scan()
			if !hasNext {
				needBreak = true
				continue
			}
			fmt.Sscanf(scanner.Text(),"%e",&items[i])
		}
		if needBreak {
			break
		}
		//fmt.Println("idx",idx,"val",items)
		if idx == 0 {
			fileOut.WriteString("19010101")
		} else {
			rh := calculateRh(items)
			add(rh)
			writeRh(idx,fileOut)
		}
		idx++
	}
}

func calculateRh(items/*T,P,Q*/ [3]float64) float64 {
	T := items[0]
	P := items[1] / 1000
	q := items[2]
	es := 6.1078 * math.Pow(math.E,17.2693882 * (T-273.16) / (T - 35.86))
	qs:=0.622*es / (P-0.378*es)
	rh := 100 * q /qs
	return rh
}

func writeRh(idx int,fileOut *os.File) {
	if (idx -2) % 4 != 0 {
		return
	}
	var rs float64
	cnt:=0.0
	for _,rh := range rhQueue {
		rs += rh
		cnt+=1
	}
	if cnt != 0 {
		rs = rs / cnt
	}
	val := fmt.Sprintf("%.3f", rs)
	fileOut.WriteString("\r\n")
	fileOut.WriteString(val)
}

func openFiles(pathIn [3]string) [3]*os.File {
	res := [3]*os.File{}
	for i,path := range pathIn {
		file,err := os.Open(path)
		if err != nil{
			panic(err)
		}
		res[i] = file
	}
	return res
}