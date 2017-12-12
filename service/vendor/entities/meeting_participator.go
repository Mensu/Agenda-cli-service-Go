package entities

import (
	"github.com/jinzhu/gorm"
)

// MeetingParticipator .
type MeetingParticipator struct {
	gorm.Model
	Title    string `json:"title" gorm:"column:title"`
	Username string `json:"username" gorm:"column:username"`
}

// TableName .
func (*MeetingParticipator) TableName() string {
	return "meeting_participator"
}

// MPService .
type MPService struct{}

// MPServ .
var MPServ = MPService{}

func init() {
	addServ(&MPServ)
}

func (*MPService) load() {
	m := &MeetingParticipator{}
	if !gormDb.HasTable(m) {
		gormDb.CreateTable(m)
	}
}

// Add .
func (*MPService) Add(m *Meeting) {
	tx := gormDb.Begin()
	checkErr(tx.Error)

	for _, username := range m.ParticipatorsUsername {
		mp := &MeetingParticipator{
			Title:    m.Title,
			Username: username,
		}
		if err := tx.Create(mp).Error; err != nil {
			tx.Rollback()
			checkErr(err)
		}
	}

	tx.Commit()
}

// Delete .
func (*MPService) Delete(m *MeetingParticipator) {
	tx := gormDb.Begin()
	checkErr(tx.Error)

	if err := tx.Delete(m).Error; err != nil {
		tx.Rollback()
		checkErr(err)
	}

	tx.Commit()
}
