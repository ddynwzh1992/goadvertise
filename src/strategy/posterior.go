package strategy

import (
	"arm-test/context"
	"go.uber.org/zap"
)

type Posterior struct {
}

func (s *Posterior) Process(ctx *context.XContext) error {
	for _, pos := range ctx.AllPoswrapsByPosid {
		for _, ad := range pos.AdInfos {
			for _, c := range ad.CreativeInfos {
				defaultEcpm := s.getDefaultEcpm(ctx, ad)
				c.PredictEcpmResult.Ecpm = float64(defaultEcpm)
			}
		}
	}
	return nil
}

// 计算默认ecpm
func (s *Posterior) getDefaultEcpm(ctx *context.XContext, adInfo *context.AdInfo) (ecpm uint64) {
	var (
		p        uint64 = 1
		bidPrice uint64 = 123
	)
	switch adInfo.BillingType {
	case 1:
		ecpm = bidPrice * p
	case 2:
		ecpm = uint64(float64(bidPrice*p) * 0.01 * 1000)
	default:
		zap.S().Errorw("unknown_billing_type", "request_id", ctx.Request.Context.RequestId, "billing_type", adInfo.BillingType)
	}
	return
}
