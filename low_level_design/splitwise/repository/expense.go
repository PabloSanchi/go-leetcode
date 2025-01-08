package repository

import (
	"gorm.io/gorm"
	models "splitwise/domain/db"
	"splitwise/domain/dto"
)

type ExpenseRepository interface {
	CreateExpense(expense *dto.Expense) error
}

type ExpenseRepositoryImpl struct {
	db *gorm.DB
}

func NewExpenseRepositoryImpl(db *gorm.DB) ExpenseRepository {
	return &ExpenseRepositoryImpl{db: db}
}

func (er *ExpenseRepositoryImpl) CreateExpense(expense *dto.Expense) error {
	tx := er.db.Begin()

	expenseShareEntities := make([]*models.ExpenseShare, len(expense.Shares))
	for i, share := range expense.Shares {
		expenseShareEntities[i] = &models.ExpenseShare{
			UserId:    share.UserId,
			Amount:    share.Amount,
			IsSettled: false,
		}
	}

	expenseEntity := &models.Expense{
		Description: expense.Description,
		Amount:      expense.Amount,
		GroupId:     expense.GroupId,
		PayerId:     expense.PayerId,
		Shares:      expenseShareEntities,
	}

	if err := er.db.Create(expenseEntity).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}
