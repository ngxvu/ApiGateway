package models

import (
	"github.com/jinzhu/gorm/dialects/postgres"
	"gorm.io/gorm"
	"time"
)

type BaseInt struct {
	ID        int             `gorm:"primary_key;autoIncrement" json:"id"`
	CreatedAt *time.Time      `gorm:"column:created_at;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt *time.Time      `gorm:"column:updated_at;default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt *gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`
}

type Log struct {
	BaseInt
	Endpoint  string         `json:"endpoint" gorm:"type:varchar(255);not null"`
	Method    string         `json:"method" gorm:"type:varchar(10);not null"`
	Params    postgres.Jsonb `json:"params" gorm:"type:jsonb"`
	Results   postgres.Jsonb `json:"results" gorm:"type:jsonb"`
	IPAddress string         `json:"ip_address" gorm:"type:varchar(50);not null"`
	Status    int            `json:"status" gorm:"type:int;not null"`
	Elapsed   float64        `json:"elapsed" gorm:"type:decimal(10,2);not null"`
}

func (Log) TableName() string {
	return "logs"
}
