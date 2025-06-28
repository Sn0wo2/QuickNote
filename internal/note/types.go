package note

import (
	"time"

	"gorm.io/gorm"
)

type Note struct {
	// gorm.Model
	ID        uint           `gorm:"primarykey" json:"-"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index"      json:"-"`

	NID  string `gorm:"index" json:"nid"`
	Data []byte `json:"-"` // Encoded Payload

	Title   []byte `gorm:"-" json:"-"` // Decoded Title
	Content []byte `gorm:"-" json:"-"` // Decoded Content

	// TODO: REFACTOR THIS
	DisplayTitle   string `json:"title,omitempty"`
	DisplayContent string `json:"content,omitempty"`
}
