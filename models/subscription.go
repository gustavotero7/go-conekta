package models

type Subscription struct {
	ConektaBase
	ID                string `json:"id,omitempty"`
	PausedAt          int64  `json:"paused_at,omitempty"`
	CanceledAt        int64  `json:"canceled_at,omitempty"`
	BillingCycleStart int64  `json:"billing_cycle_start,omitempty"`
	BillingCycleEnd   int64  `json:"billing_cycle_end,omitempty"`
	TrialStart        int64  `json:"trial_start,omitempty"`
	TrialEnd          int64  `json:"trial_end,omitempty"`
	PlanID            string `json:"plan_id,omitempty"`
	// Status of the subscription. Allowed values are: n_trial, active, past_due, paused, and canceled.
	Status string `json:"status,omitempty"`
}
