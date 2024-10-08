// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package gql

import (
	"time"
)

type BudgetPeriodInput struct {
	StartDate *time.Time `json:"startDate"`
	EndDate   *time.Time `json:"endDate"`
}

type OnboardPayload struct {
	UserID string `json:"userID"`
}

type OnboardUserInput struct {
	UserID       string             `json:"userID"`
	BudgetPeriod *BudgetPeriodInput `json:"budgetPeriod"`
}
