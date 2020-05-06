package wind

import (
	"fmt"
	"kx/commons"
)

type WINDPointLineQueueProcess struct {}

func (wind WINDPointLineQueueProcess) NoticeAdd(queue* commons.PointLineQueue,pl *commons.PointLine){
	if len(*queue) == 16 {
		queue.Pop()
	}
}

func (wind WINDPointLineQueueProcess) ExeCalcuate(idx int) bool {
	return (idx - 7) % 8 == 0
}

func (wind WINDPointLineQueueProcess) CalculateLatVal(queue* commons.PointLineQueue,idx int) string {
	var rs float64
	cnt:=0.0
	var plq = *queue
	if len(plq) == 8 {
		rs = plq[4].LatVals[idx] + plq[5].LatVals[idx] + plq[6].LatVals[idx]
		cnt = 3
	} else {
		rs = plq[7].LatVals[idx] + plq[12].LatVals[idx] + plq[13].LatVals[idx] + plq[14].LatVals[idx]
		cnt = 4
	}
	rs = rs / cnt
	return fmt.Sprintf("%.3f", rs)
}
