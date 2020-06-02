package handler

import (
	"context"
	"omo-msa-analytics/model"

	"github.com/micro/go-micro/v2/logger"
	proto "github.com/xtech-cloud/omo-msp-analytics/proto/analytics"
)

type Query struct{}

func (this *Query) Agent(_ctx context.Context, _req *proto.QueryAgentRequest, _rsp *proto.QueryAgentResponse) error {
	logger.Infof("Received Query.Agent, request is %v", _req)
	_rsp.Status = &proto.Status{}

	offset := int64(0)
	if _req.Offset > 0 {
		offset = _req.Offset
	}

	count := int64(0)
	if _req.Count > 0 {
		count = _req.Count
	}

	dao := model.NewAgentDAO()
	agents, err := dao.List(offset, count)
	if nil != err {
		return err
	}

	total, err := dao.Count()
	if nil != err {
		return err
	}
	_rsp.Total = total

	_rsp.Agent = make([]*proto.Agent, len(agents))
	for i := 0; i < len(agents); i++ {
		_rsp.Agent[i] = &proto.Agent{
			SerialNumber:    agents[i].SerialNumber,
			SoftwareFamily:  agents[i].SoftwareFamily,
			SoftwareVersion: agents[i].SoftwareVersion,
			SystemFamily:    agents[i].SystemFamily,
			SystemVersion:   agents[i].SystemVersion,
			DeviceModel:     agents[i].DeviceModel,
			DeviceType:      agents[i].DeviceType,
			Profile:         agents[i].Profile,
		}
	}
	return nil
}
