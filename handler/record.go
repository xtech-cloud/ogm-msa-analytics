package handler

import (
	"context"
	"fmt"
	"ogm-msa-analytics/model"

	"github.com/micro/go-micro/v2/logger"
	proto "github.com/xtech-cloud/ogm-msp-analytics/proto/analytics"
)

type Record struct{}

func (this *Record) Wake(_ctx context.Context, _req *proto.Agent, _rsp *proto.BlankResponse) error {
	logger.Infof("Received Record.Wake, request is %v", _req)
	_rsp.Status = &proto.Status{}
	if "" == _req.SerialNumber {
		_rsp.Status.Code = 1
		_rsp.Status.Message = "serialnumber is required"
		return nil
	}

	if "" == _req.SoftwareFamily {
		_rsp.Status.Code = 1
		_rsp.Status.Message = "softwarefamily is required"
		return nil
	}

	dao := model.NewAgentDAO(nil)

	uid := model.ToUUID(fmt.Sprintf("%s%s", _req.SoftwareFamily, _req.SerialNumber))

	agent := model.Agent{
		ID:              uid,
		SerialNumber:    _req.SerialNumber,
		SoftwareFamily:  _req.SoftwareFamily,
		SoftwareVersion: _req.SoftwareVersion,
		SystemFamily:    _req.SystemFamily,
		SystemVersion:   _req.SystemVersion,
		DeviceModel:     _req.DeviceModel,
		DeviceType:      _req.DeviceType,
		Profile:         _req.Profile,
	}

	return dao.Upsert(&agent)
}

func (this *Record) Activity(_ctx context.Context, _req *proto.RecordActivityRequest, _rsp *proto.BlankResponse) error {
	logger.Infof("Received Record.Activity, request is %v", _req)
	_rsp.Status = &proto.Status{}
	if "" == _req.AppID {
		_rsp.Status.Code = 1
		_rsp.Status.Message = "appID is required"
		return nil
	}

	if "" == _req.DeviceID {
		_rsp.Status.Code = 1
		_rsp.Status.Message = "deviceID is required"
		return nil
	}

	if "" == _req.EventID {
		_rsp.Status.Code = 1
		_rsp.Status.Message = "eventID is required"
		return nil
	}

	dao := model.NewActivityDAO(nil)

	uid := model.NewUUID()

	if len(_req.Parameter) > 0 {
		for k, v := range _req.Parameter {
			activity := &model.Activity{
				ID:         uid,
				AppID:      _req.AppID,
				DeviceID:   _req.DeviceID,
				UserID:     _req.UserID,
				EventID:    _req.EventID,
				EventKey:   k,
				EventValue: v,
			}
			dao.Insert(activity)
		}
	} else {
		activity := &model.Activity{
			ID:       uid,
			AppID:    _req.AppID,
			DeviceID: _req.DeviceID,
			UserID:   _req.UserID,
			EventID:  _req.EventID,
		}
		dao.Insert(activity)
	}
	return nil
}
