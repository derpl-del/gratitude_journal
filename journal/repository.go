package journal

import (
	"errors"
	"gratitude_journal/derpl-del/infrastructure"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (j *Journal) Create() error {
	err := infrastructure.PGdatabase.Table.Create(j).Error
	if err != nil {
		return err
	}
	return nil
}

func (j *Journal) Get() error {
	err := infrastructure.PGdatabase.Table.Preload("Points").Find(&j).Error
	if errors.Is(err, gorm.ErrRecordNotFound); err != nil {
		return err
	}
	if len(j.AuthorId) == 0 {
		return errors.New("INVALID JOURNAL ID")
	}
	return nil
}

func (j *Journal) Delete() error {
	err := infrastructure.PGdatabase.Table.Select(clause.Associations).Delete(&j).Error
	if err != nil {
		return err
	}
	return nil
}

func (j *Journal) Update() error {
	updates := *j
	infrastructure.PGdatabase.Table.Find(&j)
	if len(j.AuthorId) == 0 {
		return errors.New("INVALID JOURNAL ID")
	}
	j.Level = updates.Level
	j.Title = updates.Title
	j.Points = updates.Points
	infrastructure.PGdatabase.Table.Where("journal_id = ?", j.Id).Delete(&j.Points)
	err := infrastructure.PGdatabase.Table.Save(&j).Error
	if err != nil {
		return err
	}
	return nil
}
