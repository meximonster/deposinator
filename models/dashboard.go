package models

import "time"

type DashboardData struct {
	Totals   DashboardTotals `json:"totals"`
	Weekly   []WeeklyResult  `json:"weekly"`
	Sessions []SessionResult `json:"sessions"`
	Monthly  []MonthlyResult `json:"monthly"`
}

type DashboardTotals struct {
	TotalDeposits    float64 `json:"total_deposits"`
	TotalWithdrawals float64 `json:"total_withdrawals"`
	NetResult        float64 `json:"net_result"`
	Last30DaysNet    float64 `json:"last_30days_net_result"`
}

type WeeklyResult struct {
	WeekStart time.Time `json:"week_start"`
	WeekEnd   time.Time `json:"week_end"`
	NetResult float64   `json:"net_result"`
}

type SessionResult struct {
	ID          int       `json:"id"`
	Date        time.Time `json:"date"`
	DepositAmt  float64   `json:"deposit_amount"`
	WithdrawAmt float64   `json:"withdraw_amount"`
	NetResult   float64   `json:"net_result"`
	Description string    `json:"description"`
}

type MonthlyResult struct {
	Month       string  `json:"month"`
	Deposits    float64 `json:"deposits"`
	Withdrawals float64 `json:"withdrawals"`
	NetResult   float64 `json:"net_result"`
}
