-- Main Dashboard Metrics
SELECT 
    COALESCE(SUM(s.amount), 0) AS total_deposits,
    COALESCE(SUM(s.withdraw_amount), 0) AS total_withdrawals,
    COALESCE(SUM(s.withdraw_amount - s.amount), 0) AS net_result,
    COALESCE(SUM(s.withdraw_amount - s.amount) FILTER (
        WHERE s.created_at >= CURRENT_DATE - INTERVAL '30 days'
    ), 0) AS last_30days_net_result
FROM 
    sessions s
JOIN 
    session_members sm ON s.id = sm.session_id
WHERE 
    sm.user_id = $1;
