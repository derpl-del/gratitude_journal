package author

import (
	"errors"
	"gratitude_journal/derpl-del/infrastructure"

	"gorm.io/gorm"
)

func (a *Author) Create() error {
	err := infrastructure.PGdatabase.Table.Create(a).Error
	if err != nil {
		return err
	}
	return nil
}

func (a *Author) Get() error {
	err := infrastructure.PGdatabase.Table.Preload("Journal").Preload("Journal.Points").Find(&a).Error
	if errors.Is(err, gorm.ErrRecordNotFound); err != nil {
		return err
	}
	return nil
}

func (a *Author) Delete() error {
	infrastructure.PGdatabase.Table.Preload("Journal").Preload("Journal.Points").Find(&a)
	infrastructure.PGdatabase.Table.Select("Points").Delete(&a.Journal)
	err := infrastructure.PGdatabase.Table.Select("Journal").Delete(&a).Error
	if errors.Is(err, gorm.ErrRecordNotFound); err != nil {
		return err
	}
	return nil
}

func (a *Author) Update() error {
	updates := *a
	infrastructure.PGdatabase.Table.Find(&a)
	if len(a.Name) == 0 {
		return errors.New("INVALID JOURNAL ID")
	}
	a.Auth = updates.Auth
	err := infrastructure.PGdatabase.Table.Save(&a).Error
	if err != nil {
		return err
	}
	return nil
}
