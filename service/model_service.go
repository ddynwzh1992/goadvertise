package service

import (
	"arm-test/context"
	"arm-test/proto"
	"arm-test/strategy"
	stdContext "context"
	"errors"
	"fmt"
	"go.uber.org/zap"
)

type ModelService struct{}

func (s *ModelService) PredictV2(c stdContext.Context, req *proto.PredictionRequest) (*proto.PredictionResponse, error) {
	defer func() {
		if err := recover(); err != nil {
			zap.S().Errorw("panic", "err", err)
		}

	}()

	// 校验参数
	if req.Context == nil || len(req.PosInfos) == 0 {
		return nil, errors.New("invalid request")
	}

	// 构建上下文
	ctx := context.NewXContext(c, req)

	// 解析请求
	s.parseRequest(ctx)

	// 执行策略
	posteriorStrategy := strategy.Posterior{}
	err := posteriorStrategy.Process(ctx)
	if err != nil {
		zap.S().Errorw("posteriorStrategy.Process", "request_id", req.Context.RequestId, "err", err)
	}

	return ctx.Response, nil
}

func (s *ModelService) parseRequest(ctx *context.XContext) {
	s.parsePos(ctx)
}

func (s *ModelService) parsePos(ctx *context.XContext) {
	for _, simplePos := range ctx.Request.PosInfos {
		if simplePos == nil {
			continue
		}
		ctx.AllPoswrapsByPosid[simplePos.PosId] = &context.PosInfo{
			PosId:         simplePos.PosId,
			SimplePosInfo: simplePos,
			AdInfos:       make([]*context.AdInfo, 0, len(simplePos.Ads)),
		}
		s.parseAd(ctx, simplePos)
	}
}

func (s *ModelService) parseAd(ctx *context.XContext, simplePos *proto.SimplePosInfo) {
	for _, simpleAd := range simplePos.Ads {
		if simpleAd == nil {
			continue
		}
		adInfo := &context.AdInfo{
			PosId:        simplePos.PosId,
			AdId:         simpleAd.AdId,
			BillingType:  1,
			EventType:    1,
			SimpleAdInfo: simpleAd,
		}
		ctx.AllPoswrapsByPosid[simplePos.PosId].AdInfos = append(ctx.AllPoswrapsByPosid[simplePos.PosId].AdInfos, adInfo)
		s.parseCreative(ctx, adInfo)
	}
}

func (s *ModelService) parseCreative(ctx *context.XContext, adInfo *context.AdInfo) {
	for _, simpleCreative := range adInfo.SimpleAdInfo.Creatives {
		key := fmt.Sprintf("%d_%d_%d", adInfo.PosId, adInfo.AdId, simpleCreative.CreativeId)
		predictEcpmResult := &proto.PredictEcpmResult{
			PosId:      adInfo.PosId,
			AdId:       adInfo.AdId,
			CreativeId: simpleCreative.CreativeId,
			Ecpm:       1,
		}
		ctx.Response.ModelResult[key] = predictEcpmResult

		adInfo.CreativeInfos = append(adInfo.CreativeInfos, &context.MaCreativeInfo{
			CreativeId:         simpleCreative.CreativeId,
			SimpleCreativeInfo: simpleCreative,
			PredictEcpmResult:  predictEcpmResult,
			Attributions: map[string]interface{}{
				"bid_price":    1.23,
				"billing_type": adInfo.BillingType,
				"order_type":   1,
			},
		})
	}
}
