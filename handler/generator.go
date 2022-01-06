package handler

import (
	"context"
	"encoding/json"
	"ogm-analytics/model"

	"github.com/asim/go-micro/v3/logger"
	proto "github.com/xtech-cloud/ogm-msp-analytics/proto/analytics"
)

type Generator struct{}

func (this *Generator) Agent(_ctx context.Context, _req *proto.GeneratorAgentRequest, _rsp *proto.GeneratorAgentResponse) error {
	logger.Infof("Received Generator.Agent, request is %v", _req)
	_rsp.Status = &proto.Status{}

	offset := int64(0)
	if _req.Offset > 0 {
		offset = _req.Offset
	}

	count := int64(0)
	if _req.Count > 0 {
		count = _req.Count
	}

	dao := model.NewAgentDAO(nil)
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

func (this *Generator) Record(_ctx context.Context, _req *proto.GeneratorRecordRequest, _rsp *proto.GeneratorRecordResponse) error {
	logger.Infof("Received Generator.Record, request is %v", _req)
	_rsp.Status = &proto.Status{}

	dao := model.NewActivityDAO(nil)
	query := &model.ActivityQuery{
		Offset:         _req.Offset,
		Count:          _req.Count,
		StartTime:      _req.StartTime,
		EndTime:        _req.EndTime,
		AppID:          _req.AppID,
		DeviceID:       _req.DeviceID,
		UserID:         _req.UserID,
		EventID:        _req.EventID,
		EventParameter: _req.EventParameter,
	}
	activityAry, err := dao.List(query)
	if nil != err {
		return err
	}

	if len(activityAry) > 0 {
		// TODO use template
		content, err := json.Marshal(activityAry)
		if nil != err {
			_rsp.Status.Code = 2
			_rsp.Status.Message = err.Error()
			return nil
		}
		_rsp.Content = model.ToBase64(content)
	}
	return nil
}
