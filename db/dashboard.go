package db

import (
	"embed"
	"fmt"
	"log"

	"github.com/deposinator/models"
)

//go:embed sql/*.sql
var sqlFiles embed.FS

func getQuery(name string) string {
	data, _ := sqlFiles.ReadFile(fmt.Sprintf("sql/%s.sql", name))
	return string(data)
}

func GetUserDashboard(userID int) (*models.DashboardData, error) {
	data := &models.DashboardData{}
	var err error

	// Get totals
	if data.Totals, err = getTotals(userID); err != nil {
		log.Printf("Error getting totals: %v", err)
		return nil, err
	}

	// Get weekly data
	if data.Weekly, err = getWeeklyResults(userID); err != nil {
		log.Printf("Error getting weekly results: %v", err)
		return nil, err
	}

	// Get sessions
	if data.Sessions, err = getSessions(userID); err != nil {
		log.Printf("Error getting sessions: %v", err)
		return nil, err
	}

	// Get monthly data
	if data.Monthly, err = getMonthlyResults(userID); err != nil {
		log.Printf("Error getting monthly results: %v", err)
		return nil, err
	}

	return data, nil
}

func getTotals(userID int) (models.DashboardTotals, error) {
	var totals models.DashboardTotals
	query := getQuery("totals")

	err := db.QueryRow(query, userID).Scan(
		&totals.TotalDeposits,
		&totals.TotalWithdrawals,
		&totals.NetResult,
		&totals.Last30DaysNet,
	)

	return totals, err
}

func getWeeklyResults(userID int) ([]models.WeeklyResult, error) {
	query := getQuery("weekly")
	rows, err := db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []models.WeeklyResult
	for rows.Next() {
		var wr models.WeeklyResult
		if err := rows.Scan(
			&wr.WeekStart,
			&wr.WeekEnd,
			&wr.NetResult,
		); err != nil {
			return nil, err
		}
		results = append(results, wr)
	}

	return results, rows.Err()
}

func getSessions(userID int) ([]models.SessionResult, error) {
	query := getQuery("session")
	rows, err := db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var sessions []models.SessionResult
	for rows.Next() {
		var s models.SessionResult
		if err := rows.Scan(
			&s.ID,
			&s.Date,
			&s.DepositAmt,
			&s.WithdrawAmt,
			&s.NetResult,
			&s.Description,
		); err != nil {
			return nil, err
		}
		sessions = append(sessions, s)
	}

	return sessions, rows.Err()
}

func getMonthlyResults(userID int) ([]models.MonthlyResult, error) {
	query := getQuery("monthly")
	rows, err := db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var monthly []models.MonthlyResult
	for rows.Next() {
		var m models.MonthlyResult
		if err := rows.Scan(
			&m.Month,
			&m.Deposits,
			&m.Withdrawals,
			&m.NetResult,
		); err != nil {
			return nil, err
		}
		monthly = append(monthly, m)
	}

	return monthly, rows.Err()
}
