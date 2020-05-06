package prect

import (
	"fmt"
	"kx/commons"
)

type PRECTPointLineQueueProcess struct {}

func (prect PRECTPointLineQueueProcess) NoticeAdd(queue* commons.PointLineQueue,pl *commons.PointLine){
	if len(*queue) == 5 {
		queue.Pop()
	}
}

func (prect PRECTPointLineQueueProcess) ExeCalcuate(idx int) bool {
	return (idx - 2) % 4 == 0
}

func (prect PRECTPointLineQueueProcess) CalculateLatVal(queue* commons.PointLineQueue,idx int) string {
	var rs float64
	var plq = *queue
	if len(plq) == 3 {
		rs = plq[0].LatVals[idx] * 6 + plq[1].LatVals[idx] * 6 + plq[2].LatVals[idx] * 4
	} else {
		rs = plq[0].LatVals[idx] * 2+ plq[1].LatVals[idx] * 6 + plq[2].LatVals[idx] * 6 + plq[3].LatVals[idx] * 6 + plq[4].LatVals[idx] * 4
	}

	rs *= 3600

	return fmt.Sprintf("%.3f", rs)
}