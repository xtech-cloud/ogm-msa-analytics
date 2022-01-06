package model

import "time"

type Activity struct {
	ID             string `gorm:"column:id;type:char(32);not null;unique;primary_key"`
	AppID          string `gorm:"column:app_id;type:char(32)"`
	DeviceID       string `gorm:"column:device_id;type:char(32)"`
	UserID         string `gorm:"column:user_id;type:char(32)"`
	EventID        string `gorm:"column:event_id;type:char(32)"`
	EventParameter string `gorm:"column:event_parameter;type:varchar(256)"`
	CreatedAt      time.Time
}

func (Activity) TableName() string {
	return "ogm_analytics_activity"
}

type ActivityDAO struct {
	conn *Conn
}

type ActivityQuery struct {
	StartTime      int64
	EndTime        int64
	Offset         int64
	Count          int64
	AppID          string
	DeviceID       string
	UserID         string
	EventID        string
	EventParameter string
}

func NewActivityDAO(_conn *Conn) *ActivityDAO {
	conn := DefaultConn
	if nil != _conn {
		conn = _conn
	}
	return &ActivityDAO{
		conn: conn,
	}
}

func (this *ActivityDAO) Insert(_activity *Activity) error {
	return this.conn.DB.Create(_activity).Error
}

func (this *ActivityDAO) List(_query *ActivityQuery) ([]*Activity, error) {
	db := this.conn.DB
	if "" != _query.AppID {
		db = db.Where("app_id = ?", _query.AppID)
	}
	if "" != _query.DeviceID {
		db = db.Where("device_id = ?", _query.DeviceID)
	}
	if "" != _query.UserID {
		db = db.Where("user_id = ?", _query.UserID)
	}
	if "" != _query.EventID {
		db = db.Where("event_id = ?", _query.EventID)
	}
	if "" != _query.EventParameter {
		db = db.Where("event_parameter LIKE ?", "%"+_query.EventParameter+"%")
	}
	if _query.Offset > 0 {
		db = db.Offset(int(_query.Offset))
	}
	if _query.Count > 0 {
		db = db.Limit(int(_query.Count))
	}
	if _query.StartTime > 0 && _query.EndTime > 0 {
		db = db.Where("created_at BETWEEN ? AND ?", time.Unix(_query.StartTime, 0), time.Unix(_query.EndTime, 0))
	}
	var activityAry []*Activity
	res := db.Order("created_at desc").Find(&activityAry)
	return activityAry, res.Error
}
