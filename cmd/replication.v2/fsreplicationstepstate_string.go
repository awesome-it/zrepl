// Code generated by "stringer -type=FSReplicationStepState"; DO NOT EDIT.

package replication

import "strconv"

const (
	_FSReplicationStepState_name_0 = "StepReadyStepRetry"
	_FSReplicationStepState_name_1 = "StepPermanentError"
	_FSReplicationStepState_name_2 = "StepCompleted"
)

var (
	_FSReplicationStepState_index_0 = [...]uint8{0, 9, 18}
)

func (i FSReplicationStepState) String() string {
	switch {
	case 1 <= i && i <= 2:
		i -= 1
		return _FSReplicationStepState_name_0[_FSReplicationStepState_index_0[i]:_FSReplicationStepState_index_0[i+1]]
	case i == 4:
		return _FSReplicationStepState_name_1
	case i == 8:
		return _FSReplicationStepState_name_2
	default:
		return "FSReplicationStepState(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}
