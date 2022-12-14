package snowflake

import (
	"errors"
	"time"
)

/*
组成：0(1 bit) | timestamp in millisecond (41 bit) | machine id (10 bit) | index (12 bit)
每毫秒最多生成4096个id，集群机器最多1024台
*/
type Snowflake struct {
	lastTimestamp int64
	index         int16
	machineId     int16
}

func (s *Snowflake) Init(id int16) error {
	if id > 0xff {
		return errors.New("illegal machine id")
	}

	s.machineId = id
	s.lastTimestamp = time.Now().UnixNano() / 1e6
	s.index = 0
	return nil
}

func (s *Snowflake) GetId() (int64, error) {
	curTimestamp := time.Now().UnixNano() / 1e6
	if curTimestamp == s.lastTimestamp {
		s.index++
		if s.index > 0xfff {
			s.index = 0xfff
			return -1, errors.New("out of range")
		}
	} else {
		//fmt.Printf("id/ms:%d -- %d\n", s.lastTimestamp, s.index)
		s.index = 0
		s.lastTimestamp = curTimestamp
	}
	return (0x1ffffffffff&s.lastTimestamp)<<22 + int64(0xff<<10) + int64(0xfff&s.index), nil
}


