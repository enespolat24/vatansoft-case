package model

import (
	"time"

	"gorm.io/gorm"
)

type Plan struct {
	ID        uint      `gorm:"primaryKey"`
	UserID    uint      `gorm:"column:user_id"`
	StartTime time.Time `gorm:"column:start_time"`
	EndTime   time.Time `gorm:"column:end_time"`
	Title     string    `gorm:"column:title"`
	State     PlanState `gorm:"column:state"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
	DeletedAt gorm.DeletedAt
}

type PlanState string

const (
	StatePending    PlanState = "pending"
	StateInProgress PlanState = "in_progress"
	StateCompleted  PlanState = "completed"
)
