package journal

import "time"

type Journal struct {
	Id       uint      `json:"id"`
	AuthorId string    `json:"author_id" `
	Date     time.Time `json:"date"`
	Title    string    `json:"tittle"`
	Level    int       `json:"level"`
	Points   []Point   `json:"point" gorm:"foreignKey:JournalId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type Point struct {
	Id        uint   `json:"id"`
	JournalId string `json:"journal_id"`
	Body      string `json:"body"`
}
