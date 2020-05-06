package qbot

import (
	"fmt"
	"kx/commons"
)

type QBOTPointLineQueueProcess struct {}

func (qbot QBOTPointLineQueueProcess) NoticeAdd(queue* commons.PointLineQueue,pl *commons.PointLine){
	if len(*queue) == 16 {
		queue.Pop()
	}
}

func (qbot QBOTPointLineQueueProcess) ExeCalcuate(idx int) bool {
	return (idx - 7) % 8 == 0
}

func (qbot QBOTPointLineQueueProcess) CalculateLatVal(queue* commons.PointLineQueue,idx int) string {
	var rs float64
	cnt:=0.0
	var plq = *queue
	if len(plq) == 8 {
		rs = plq[0].LatVals[idx] + plq[1].LatVals[idx] + plq[2].LatVals[idx]
		cnt = 3
	} else {
		rs = plq[3].LatVals[idx] + plq[8].LatVals[idx] + plq[9].LatVals[idx] + plq[10].LatVals[idx]
		cnt = 4
	}
	rs = rs / cnt
	return fmt.Sprintf("%.3f", rs)
}
