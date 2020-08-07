package model

import "time"

type Agent struct {
	ID              string `gorm:"primary_key"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	SerialNumber    string `gorm:"column:serial_number;type:varchar(128);not null;unique"`
	SoftwareFamily  string `gorm:"column:software_family;type:varchar(128);not null"`
	SoftwareVersion string `gorm:"column:software_version;type:varchar(128)"`
	SystemFamily    string `gorm:"column:system_family;type:varchar(128)"`
	SystemVersion   string `gorm:"column:system_version;type:varchar(128)"`
	DeviceModel     string `gorm:"column:device_model;type:varchar(128)"`
	DeviceType      string `gorm:"column:device_type;type:varchar(128)"`
	Profile         string `gorm:"column:profile;type:TEXT"`
}

func (Agent) TableName() string {
	return "msa_analytics_agent"
}

type AgentDAO struct {
}

func NewAgentDAO() *AgentDAO {
	return &AgentDAO{}
}

func (AgentDAO) Upset(_agent *Agent) error {
	db, err := openSqlDB()
	if nil != err {
		return err
	}
	defer closeSqlDB(db)

	if db.NewRecord(_agent) {
		// 插入
		return db.Create(_agent).Error
	} else {
		// 更新
		return db.Save(_agent).Error
	}
}

func (AgentDAO) List(_offset int64, _count int64) ([]*Agent, error) {
	db, err := openSqlDB()
	if nil != err {
		return nil, err
	}
	defer closeSqlDB(db)

	var agents []*Agent
	res := db.Offset(_offset).Limit(_count).Order("created_at desc").Find(&agents)
	return agents, res.Error
}

func (AgentDAO) Count() (int64, error) {
	db, err := openSqlDB()
	if nil != err {
		return 0, err
	}
	defer closeSqlDB(db)
	count := int64(0)
	res := db.Model(&Agent{}).Count(&count)
	return count, res.Error
}
