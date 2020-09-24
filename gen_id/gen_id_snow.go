package gen_id

import (
	"fmt"
	"log"
	"sync"
	"time"
)

const (
	workerIDBits = uint64(10) // 机器id占的位数
	maxWorkerID = int64(-1) ^ (int64(-1) << workerIDBits) // 最大机器id 1023
	timestampBits = uint64(32) // 时间戳占32位
	sequenceBits  = uint64(20) // 自增序列占20位
	sequenceMax = int64(-1) ^ (int64(-1) << sequenceBits) // 自增序列号最大值
	timestampLeftShift = sequenceBits // 时间戳左移20位
	workerIDLeftShift   =  timestampBits + sequenceBits // 机器id左移52

	// 粒度暂定为s级
	// 2020-09-24T00:00:00.000Z
	epochTime = int64(1600876800) // 起始日期
)

var (
	mu  sync.Mutex
	workerID int64      // worker id  0 <= workerID <= maxWorkerID
	lastTs   int64 = -1 // the last timestamp in milliseconds
	seq int64
)

func init() {
	if workerID < 0 || workerID > maxWorkerID {
		log.Fatalf("worker id must be between 0 and %d", maxWorkerID)
	}
}

func GetInt64ID() (int64, error) {
	mu.Lock()
	defer mu.Unlock()

	ts := time.Now().Unix() // 获取当前时间

	switch {
	case ts < lastTs:
		return 0, fmt.Errorf("time is moving backwards, waiting until %d", lastTs)
	case ts == lastTs:
		seq = (seq + 1) & sequenceMax
		if seq == 0 {
			for ts <= lastTs {
				time.Sleep(time.Second)
				ts = time.Now().Unix()
			}
		}
	default:
		seq = 0
	}

	lastTs = ts

	return (workerID << workerIDLeftShift) | ((ts - epochTime) << timestampLeftShift) | seq, nil
}
