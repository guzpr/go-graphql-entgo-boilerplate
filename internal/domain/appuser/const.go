package appuser

type OnboardingStatus string

const (
	OnboardingStatusCreateBudgetPeriod            OnboardingStatus = "create_budget_period"
	OnboardingStatusCreateInitialFinancialAccount OnboardingStatus = "create_initial_financial_account"
	OnboardingStatusFinished                      OnboardingStatus = "finished"
)
