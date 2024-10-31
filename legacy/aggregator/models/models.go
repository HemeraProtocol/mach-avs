package models

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Model struct {
	DB *gorm.DB
}

func NewModel(db *gorm.DB) *Model {
	db.AutoMigrate(&Task{}, &TaskSignature{})
	return &Model{DB: db}
}

func (m *Model) CreateTask(task *Task) error {
	err := m.DB.Create(task).Error
	if err == gorm.ErrDuplicatedKey {
		return nil
	}
	return err
}

func (m *Model) SetTaskFinished(task *Task) error {
	return m.DB.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "alert_hash"}},
		DoUpdates: clause.AssignmentColumns([]string{"tx_hash", "block_hash", "block_number", "transaction_index"}),
	}).Create(task).Error
}

func (m *Model) CreateTaskSignature(taskSignature *TaskSignature) error {
	return m.DB.Create(taskSignature).Error
}
