package util

import (
	"sync"
	"time"
)

// SnowflakeID 雪花算法ID生成器（对标Java版）
type SnowflakeID struct {
	mu         sync.Mutex
	epoch      int64
	nodeID     int64
	sequence   int64
	lastTime   int64
}

var snowflake *SnowflakeID

func InitSnowflake(nodeID int64) {
	snowflake = &SnowflakeID{
		epoch:    1609459200000, // 2021-01-01
		nodeID:   nodeID,
		sequence: 0,
		lastTime: 0,
	}
}

func NextID() int64 {
	if snowflake == nil {
		InitSnowflake(1)
	}
	return snowflake.Next()
}

func (s *SnowflakeID) Next() int64 {
	s.mu.Lock()
	defer s.mu.Unlock()

	now := time.Now().UnixMilli() - s.epoch

	if now == s.lastTime {
		s.sequence = (s.sequence + 1) & 4095
		if s.sequence == 0 {
			for now <= s.lastTime {
				now = time.Now().UnixMilli() - s.epoch
			}
		}
	} else {
		s.sequence = 0
	}

	s.lastTime = now
	return (now << 22) | (s.nodeID << 12) | s.sequence
}
