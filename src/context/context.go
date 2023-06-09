package context

import (
	"arm-test/proto"
	stdContext "context"
	"math/rand"
)

type XContext struct {
	Ctx      stdContext.Context
	Request  *proto.PredictionRequest
	Response *proto.PredictionResponse

	AllPoswrapsByPosid map[uint64]*PosInfo
}

func NewXContext(stdCtx stdContext.Context, req *proto.PredictionRequest) *XContext {
	return &XContext{
		Ctx:     stdCtx,
		Request: req,
		Response: &proto.PredictionResponse{
			ModelResult: make(map[string]*proto.PredictEcpmResult),
		},
		AllPoswrapsByPosid: make(map[uint64]*PosInfo),
	}
}

type PosInfo struct {
	PosId         uint64
	SimplePosInfo *proto.SimplePosInfo
	AdInfos       []*AdInfo
}

type AdInfo struct {
	PosId         uint64
	AdId          uint64
	BillingType   int
	EventType     int
	SimpleAdInfo  *proto.SimpleAdInfo
	CreativeInfos []*MaCreativeInfo
}

type MaCreativeInfo struct {
	CreativeId         uint64
	SimpleCreativeInfo *proto.SimpleCreativeInfo
	PredictEcpmResult  *proto.PredictEcpmResult
	Attributions       map[string]interface{}
}

func (ad *AdInfo) IsDpa() bool {
	return rand.Int()%2 == 0
}
