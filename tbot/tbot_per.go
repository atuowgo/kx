package tbot

import (
	"fmt"
	"kx/commons"
)

type TBOTPerPointLineQueueProcess struct {}

func (tbot TBOTPerPointLineQueueProcess) NoticeAdd(queue* commons.PointLineQueue,pl *commons.PointLine){
	if len(*queue) == 1 {
		queue.Pop()
	}
}

func (tbot TBOTPerPointLineQueueProcess) ExeCalcuate(idx int) bool {
	return true
}

func (tbot TBOTPerPointLineQueueProcess) CalculateLatVal(queue* commons.PointLineQueue,idx int) string {
	rs := (*queue)[0].LatVals[idx] - 273.15
	return fmt.Sprintf("%.3f",rs)
}