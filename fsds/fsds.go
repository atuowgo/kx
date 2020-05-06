package fsds

import (
	"fmt"
	"kx/commons"
)


type FSDSPointLineQueueProcess struct{}

func (fsds FSDSPointLineQueueProcess) NoticeAdd(queue* commons.PointLineQueue,pl *commons.PointLine){
	if len(*queue) == 4 {
		queue.Pop()
	}
}

func (fsds FSDSPointLineQueueProcess) ExeCalcuate(idx int) bool {
	return (idx - 2) % 4 == 0
}


func (fsds FSDSPointLineQueueProcess)CalculateLatVal(queue* commons.PointLineQueue,idx int) string {
	var rs float64
	cnt:=0.0
	for _,pl := range *queue {
		if pl != nil {
			cnt+=1
			rs += pl.LatVals[idx]
		}
	}
	if cnt != 0 {
		rs = rs / cnt
	}
	rs *= 0.0864
	return fmt.Sprintf("%.3f", rs)
}