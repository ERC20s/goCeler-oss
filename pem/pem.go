// Copyright 2018-2019 Celer Network

package pem

import (
	"time"

	log "github.com/celer-network/goCeler-oss/clog"
)

func NewPem(machine string) *PayEventMessage {
	return &PayEventMessage{
		StartTimeStamp: time.Now().UnixNano(),
		Machine:        machine,
		SeqNums:        &SimplexSeqNums{},
	}
}
func CommitPem(pem *PayEventMessage) {
	pem.EndTimeStamp = time.Now().UnixNano()
	pem.ExecutionTimeMs = (float32)(pem.EndTimeStamp-pem.StartTimeStamp) / 1000000
	if len(pem.Error) > 0 {
		log.Errorln("LOGPEM:", pem)
	} else if pem.Nack != nil {
		log.Warnln("LOGPEM:", pem)
	} else {
		log.Infoln("LOGPEM:", pem)
	}
}
