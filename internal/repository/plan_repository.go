package repository

import (
	"time"
	"vatansoft-case/internal/model"

	"gorm.io/gorm"
)

type planRepository struct {
	db *gorm.DB
}

type PlanRepository interface {
	CreatePlan(plan *model.Plan) error
	GetPlansByUserId(id uint) ([]model.Plan, error)
	GetPlanByID(id uint) (*model.Plan, error)
	ChangeState(id uint, state string) error
	UpdatePlan(plan *model.Plan) error
	CheckPlanOverlap(plan *model.Plan) (bool, error)
	GetWeeklyPlansByUserID(userID uint, weekStartDate time.Time) ([]model.Plan, error)
	GetMonthlyPlansByUserID(userID uint, monthStartDate time.Time) ([]model.Plan, error)
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

func (pr *planRepository) UpdatePlan(plan *model.Plan) error {
	return pr.db.Model(&model.Plan{}).Where("id = ?", plan.ID).Updates(plan).Error
}

func (pr *planRepository) GetPlanByID(id uint) (*model.Plan, error) {
	var plan model.Plan
	err := pr.db.Where("id = ?", id).First(&plan).Error
	if err != nil {
		return nil, err
	}
	return &plan, nil
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

func (pr *planRepository) GetWeeklyPlansByUserID(userID uint, weekStartDate time.Time) ([]model.Plan, error) {
	weekStartDate = weekStartDate.AddDate(0, 0, -int(weekStartDate.Weekday())+1)

	weekEndDate := weekStartDate.AddDate(0, 0, 6)

	var plans []model.Plan
	err := pr.db.Where("user_id = ?", userID).
		Where("start_time >= ? AND start_time <= ?", weekStartDate, weekEndDate).
		Find(&plans).Error
	if err != nil {
		return nil, err
	}
	return plans, nil
}

func (pr *planRepository) GetMonthlyPlansByUserID(userID uint, monthStartDate time.Time) ([]model.Plan, error) {
	var plans []model.Plan
	monthEndDate := monthStartDate.AddDate(0, 1, 0)
	err := pr.db.Where("user_id = ? AND start_time >= ? AND start_time < ?", userID, monthStartDate, monthEndDate).
		Find(&plans).Error
	if err != nil {
		return nil, err
	}
	return plans, nil
}
