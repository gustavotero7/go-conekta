package models

type Subscription struct {
	ID                string `json:"id,omitempty"`
	Object            string `json:"object,omitempty"`
	CreatedAt         int64  `json:"created_at,omitempty"`
	UpdatedAt         int64  `json:"updated_at,omitempty"`
	PausedAt          int64  `json:"paused_at,omitempty"`
	BillingCycleStart int64  `json:"billing_cycle_start,omitempty"`
	BillingCycleEnd   int64  `json:"billing_cycle_end,omitempty"`
	TrialStart        int64  `json:"trial_start,omitempty"`
	TrialEnd          int64  `json:"trial_end,omitempty"`
	PlanID            string `json:"plan_id,omitempty"`
	Status            string `json:"status,omitempty"`
}
