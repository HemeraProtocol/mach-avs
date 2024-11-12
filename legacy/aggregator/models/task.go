package models

import (
	"time"

	"gorm.io/gorm"
)

type AggregatorTask struct {
	gorm.Model
	ID                   uint    `gorm:"primary_key"`
	AlertHash            []byte  `gorm:"column:alert_hash;unique"`
	QuorumNumbers        []uint8 `gorm:"column:quorum_numbers"`
	TaskIndex            uint32  `gorm:"column:task_index"`
	ReferenceBlockNumber uint64  `gorm:"column:reference_block_number"`
	TxHash               string  `gorm:"column:tx_hash"`
	BlockHash            string  `gorm:"column:block_hash"`
	BlockNumber          uint64  `gorm:"column:block_number"`
	TransactionIndex     uint    `json:"transactionIndex"`
	CreatedAt            time.Time
	UpdatedAt            time.Time
}

type AggregatorTaskSignature struct {
	gorm.Model
	ID         uint   `gorm:"primary_key"`
	AlertHash  []byte `gorm:"column:alert_hash"`
	OperatorId []byte `gorm:"column:operator_id"`
	SignResult bool   `gorm:"column:sign_result"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
