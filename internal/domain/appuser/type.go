package appuser

type User struct {
	ID string
	GoogleAuth
	Email string
	OnboardingStatus
}

type GoogleAuth struct {
	ID       string
	GoogleID string
}
