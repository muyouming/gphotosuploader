package models

import (
	"database/sql/driver"
	"time"

	"github.com/jinzhu/gorm"
)

// FileStatus The file status
type FileStatus uint8

const (
	// FilePending Pending status
	FilePending FileStatus = 0
	// FileSuccess Success status
	FileSuccess FileStatus = 1
	// FileFailed Failed status
	FileFailed FileStatus = 2
)

// Scan Scan value
func (e *FileStatus) Scan(value interface{}) error {
	*e = FileStatus(value.(int64))
	return nil
}

// Value To the value
func (e FileStatus) Value() (driver.Value, error) {
	return int64(e), nil
}

// File modle
type File struct {
	gorm.Model
	Path   string     `gorm:"unique_index"`
	Status FileStatus `gorm:"type:INTEGER"`
	URL    string     `gorm:"default:NULL"`
	SortAt time.Time  `gorm:"default:NULL"`
}
