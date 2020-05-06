package qbot

import (
	"fmt"
	"kx/commons"
)

type QBOTPerPointLineQueueProcess struct {}

func (qbot QBOTPerPointLineQueueProcess) NoticeAdd(queue* commons.PointLineQueue,pl *commons.PointLine){
	if len(*queue) == 1 {
		queue.Pop()
	}
}

func (qbot QBOTPerPointLineQueueProcess) ExeCalcuate(idx int) bool {
	return (idx / 4) % 2 == 0
}

func (qbot QBOTPerPointLineQueueProcess) CalculateLatVal(queue* commons.PointLineQueue,idx int) string {
	rs := (*queue)[0].LatVals[idx]
	return fmt.Sprintf("%.3e",rs)
}
