package handler

import (
	"context"
	"fmt"
	"omo-msa-analytics/model"

	"github.com/micro/go-micro/v2/logger"
	proto "github.com/xtech-cloud/omo-msp-analytics/proto/analytics"
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

	dao := model.NewAgentDAO()

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

	return dao.Upset(&agent)
}
