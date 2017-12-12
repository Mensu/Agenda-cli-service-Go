package entities

// Meeting .
type Meeting struct {
	Title                 string                `json:"title" gorm:"primary_key;column:title"`
	Speecher              User                  `gorm:"column:speecher;ForeignKey:Username"`
	SpeecherUsername      string                `json:"speecher" gorm:"-"`
	ParticipatorsUsername []string              `json:"participators" gorm:"-"`
	Participators         []MeetingParticipator `gorm:"ForeignKey:Username"`
	StartTime             string                `json:"startTime" gorm:"column:startTime"`
	EndTime               string                `json:"endTime" gorm:"column:endTime"`
}

// TableName .
func (*Meeting) TableName() string {
	return "meeting"
}

// MeetingService .
type MeetingService struct{}

// MeetingServ .
var MeetingServ = MeetingService{}

func init() {
	addServ(&MeetingServ)
}

func (*MeetingService) load() {
	m := &Meeting{}
	if !gormDb.HasTable(m) {
		gormDb.CreateTable(m)
	}
}

// Add .
func (*MeetingService) Add(m *Meeting) {
	tx := gormDb.Begin()
	checkErr(tx.Error)

	if err := tx.Create(m).Error; err != nil {
		tx.Rollback()
		checkErr(err)
	}

	tx.Commit()
}

// Delete .
func (*MeetingService) Delete(m *Meeting) {
	tx := gormDb.Begin()
	checkErr(tx.Error)

	if err := tx.Delete(m).Error; err != nil {
		tx.Rollback()
		checkErr(err)
	}

	tx.Commit()
}

// FindAll .
func (*MeetingService) FindAll() []Meeting {
	mlist := make([]Meeting, 0, 0)
	checkErr(gormDb.Find(&mlist).Error)
	return mlist
}

// FindByTitle .
func (*MeetingService) FindByTitle(title string) *Meeting {
	meetings := make([]Meeting, 0, 0)
	checkErr(gormDb.Where([]string{title}).Find(&meetings).Error)
	if len(meetings) == 0 {
		return nil
	}
	return &meetings[0]
}
