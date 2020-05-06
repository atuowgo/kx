package qbot

import (
	"fmt"
	"kx/commons"
)


type PerPointLineQueueProcess struct {}

func (tbot PerPointLineQueueProcess) NoticeAdd(queue* commons.PointLineQueue,pl *commons.PointLine){
	if len(*queue) == 1 {
		queue.Pop()
	}
}

func (tbot PerPointLineQueueProcess) ExeCalcuate(idx int) bool {
	return true
}

func (tbot PerPointLineQueueProcess) CalculateLatVal(queue* commons.PointLineQueue,idx int) string {
	rs := (*queue)[0].LatVals[idx]
	return fmt.Sprintf("%.3e",rs)
}