package note

import (
	"errors"

	"github.com/Sn0wo2/QuickNote/pkg/database/orm"
	"gorm.io/gorm"
)

func InitNoteSchema() error {
	return orm.Instance.Get().AutoMigrate(&Note{})
}

func GetNote(nid string) (Note, error) {
	var note Note

	err := orm.Instance.Get().First(&note, "n_id = ?", nid).Error

	return note, err
}

func SetNote(note Note) error {
	return orm.Instance.Get().Transaction(func(tx *gorm.DB) error {
		var existing Note
		if err := tx.Where("n_id = ?", note.NID).First(&existing).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return tx.Create(&note).Error
			}

			return err
		}

		return tx.Model(&existing).Updates(note).Error
	})
}

func DeleteNote(nid string) error {
	return orm.Instance.Get().Transaction(func(tx *gorm.DB) error {
		if err := tx.Unscoped().Where("n_id = ?", nid).Delete(&Note{}).Error; err != nil {
			return err
		}

		return nil
	})
}
