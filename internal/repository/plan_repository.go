package repository

import (
	"vatansoft-case/internal/model"

	"gorm.io/gorm"
)

type planRepository struct {
	db *gorm.DB
}

type PlanRepository interface {
	CreatePlan(plan *model.Plan) error
	GetPlansByUserId(id uint) ([]model.Plan, error)
	ChangeState(id uint, state string) error
	CheckPlanOverlap(plan *model.Plan) (bool, error)
}

func NewPlanRepository(db *gorm.DB) PlanRepository {
	return &planRepository{db: db}
}

func (pr *planRepository) ChangeState(id uint, state string) error {
	return pr.db.Model(&model.Plan{}).Where("id = ?", id).Update("state", state).Error
}

func (pr *planRepository) CreatePlan(plan *model.Plan) error {
	return pr.db.Create(plan).Error
}

func (pr *planRepository) GetPlansByUserId(id uint) ([]model.Plan, error) {
	var plans []model.Plan
	err := pr.db.Where("user_id = ?", id).Find(&plans).Error
	return plans, err
}

func (pr *planRepository) CheckPlanOverlap(plan *model.Plan) (bool, error) {
	var count int64
	err := pr.db.Model(&model.Plan{}).
		Where("user_id = ?", plan.UserID).
		Where("(start_time <= ? AND end_time >= ?) OR (start_time <= ? AND end_time >= ?) OR (start_time >= ? AND end_time <= ?)",
			plan.StartTime, plan.StartTime,
			plan.EndTime, plan.EndTime,
			plan.StartTime, plan.EndTime).
		Count(&count).
		Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
