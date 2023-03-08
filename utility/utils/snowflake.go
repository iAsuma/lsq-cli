package utils

import (
	"github.com/bwmarrin/snowflake"
	"time"
)

type snowFlake struct {
	node *snowflake.Node
}

type SnowFlakeFiller struct {
	MilliTimestamp int64
	Timestamp      int64
	Time           time.Time
	MachineId      int64
}

var insSnowFlake = snowFlake{}

func SnowFlake() *snowFlake {
	return &insSnowFlake
}

func (s *snowFlake) Init(startTime string) (err error) {
	var st time.Time
	st, err = time.Parse("2006-01-02", startTime)
	if err != nil {
		return
	}

	snowflake.Epoch = st.UnixNano() / 1e6
	return
}

func (s *snowFlake) Generate(machineID int64, startTime string) (snowflake.ID, error) {
	if s.node == nil {
		err := s.Init(startTime)

		s.node, err = snowflake.NewNode(machineID)
		if err != nil {
			return 0, err
		}
	}

	return s.node.Generate(), nil
}

func (s *snowFlake) Decrypt(id int64, startTime string) (t SnowFlakeFiller) {
	err := s.Init(startTime)
	if err != nil {
		return
	}

	var epoch = snowflake.Epoch
	var machineIDBits = snowflake.NodeBits
	var sequenceBits = snowflake.StepBits

	var timestampMilli = (id >> (machineIDBits + sequenceBits)) + epoch
	var machineID = (id >> sequenceBits) & ((1 << machineIDBits) - 1)

	t.MilliTimestamp = timestampMilli
	t.Timestamp = timestampMilli / 1000
	t.Time = time.Unix(timestampMilli/1000, (timestampMilli%1000)*int64(time.Millisecond))
	t.MachineId = machineID

	return t
}
