package author

import "gratitude_journal/derpl-del/journal"

type Author struct {
	Id      uint              `json:"id"`
	Name    string            `json:"name"`
	Auth    string            `json:"auth"`
	Journal []journal.Journal `json:"journal" gorm:"foreignKey:AuthorId;"`
}
