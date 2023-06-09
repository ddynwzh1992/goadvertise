package proto

type PredictionRequest struct {
	PosInfos []*SimplePosInfo `protobuf:"bytes,1,rep,name=pos_infos,json=posInfos,proto3" json:"pos_infos,omitempty"`
	Context  *Context         `protobuf:"bytes,2,opt,name=context,proto3" json:"context,omitempty"`
}

type SimplePosInfo struct {
	PosId uint64          `protobuf:"varint,1,opt,name=pos_id,json=posId,proto3" json:"pos_id,omitempty"`
	Ads   []*SimpleAdInfo `protobuf:"bytes,2,rep,name=ads,proto3" json:"ads,omitempty"`
}

type SimpleAdInfo struct {
	AdId      uint64                `protobuf:"varint,1,opt,name=ad_id,json=adId,proto3" json:"ad_id,omitempty"`
	Creatives []*SimpleCreativeInfo `protobuf:"bytes,2,rep,name=creatives,proto3" json:"creatives,omitempty"`
}

type SimpleCreativeInfo struct {
	CreativeId uint64 `protobuf:"varint,1,opt,name=creative_id,json=creativeId,proto3" json:"creative_id,omitempty"`
}

type Context struct {
	RequestId string `json:"request_id,omitempty"`
	BelayId   string `json:"belay_id,omitempty"`
	GpType    int32  `protobuf:"varint,3,opt,name=gp_type,json=gpType,proto3" json:"gp_type,omitempty"`
	Isocode   string `protobuf:"bytes,4,opt,name=isocode,proto3" json:"isocode,omitempty"`
	AppName   string `protobuf:"bytes,5,opt,name=app_name,json=appName,proto3" json:"app_name,omitempty"`
}

type PredictionResponse struct {
	ModelResult map[string]*PredictEcpmResult `json:"model_result,omitempty"`
}

type PredictEcpmResult struct {
	PosId      uint64  `json:"pos_id,omitempty"`
	AdId       uint64  `json:"ad_id,omitempty"`
	CreativeId uint64  `json:"creative_id,omitempty"`
	Ecpm       float64 `json:"ecpm,omitempty"`
}
