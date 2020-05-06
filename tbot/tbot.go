package tbot

import (
	"fmt"
	"kx/commons"
	"math"
)

type TBOTPointLineQueueProcess struct {}

func (tbot TBOTPointLineQueueProcess) NoticeAdd(queue* commons.PointLineQueue,pl *commons.PointLine){
	if len(*queue) == 4 {
		queue.Pop()
	}
}

func (tbot TBOTPointLineQueueProcess) ExeCalcuate(idx int) bool {
	return (idx - 2) % 4 == 0
}

func (tbot TBOTPointLineQueueProcess) CalculateLatVal(queue* commons.PointLineQueue,idx int) string {
	min := (*queue)[0].LatVals[idx]
	max := min
	for i:= 1;i<len(*queue);i++ {
		num := (*queue)[i].LatVals[idx]
		max = math.Max(max,num)
		min = math.Min(min,num)
	}
	min += -273.15
	max += -273.15
	return fmt.Sprintf("%.3f,%.3f",max,min)
}