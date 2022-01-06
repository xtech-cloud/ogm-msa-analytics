package model

import (
	"gorm.io/gorm/clause"
	"time"
)

type Agent struct {
	ID              string `gorm:"column:id;type:char(32);not null;unique;primary_key"`
	SerialNumber    string `gorm:"column:serial_number;type:varchar(128);not null;unique"`
	SoftwareFamily  string `gorm:"column:software_family;type:varchar(128);not null"`
	SoftwareVersion string `gorm:"column:software_version;type:varchar(128)"`
	SystemFamily    string `gorm:"column:system_family;type:varchar(128)"`
	SystemVersion   string `gorm:"column:system_version;type:varchar(128)"`
	DeviceModel     string `gorm:"column:device_model;type:varchar(128)"`
	DeviceType      string `gorm:"column:device_type;type:varchar(128)"`
	Profile         string `gorm:"column:profile;type:TEXT"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

func (Agent) TableName() string {
	return "ogm_analytics_agent"
}

type AgentDAO struct {
	conn *Conn
}

func NewAgentDAO(_conn *Conn) *AgentDAO {
	conn := DefaultConn
	if nil != _conn {
		conn = _conn
	}
	return &AgentDAO{
		conn: conn,
	}
}

func (this *AgentDAO) Upsert(_agent *Agent) error {
	db := this.conn.DB
	return db.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(_agent).Error
}

func (this *AgentDAO) List(_offset int64, _count int64) ([]*Agent, error) {
	db := this.conn.DB
	var agents []*Agent
	res := db.Offset(int(_offset)).Limit(int(_count)).Order("created_at desc").Find(&agents)
	return agents, res.Error
}

func (this *AgentDAO) Count() (int64, error) {
	db := this.conn.DB
	count := int64(0)
	res := db.Model(&Agent{}).Count(&count)
	return count, res.Error
}
