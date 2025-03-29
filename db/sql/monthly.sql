-- Monthly Aggregated Data
SELECT 
    to_char(date_trunc('month', s.created_at), 'Mon YYYY') AS month,
    SUM(s.amount) AS deposits,
    SUM(s.withdraw_amount) AS withdrawals,
    SUM(s.withdraw_amount - s.amount) AS net_result
FROM 
    sessions s
JOIN 
    session_members sm ON s.id = sm.session_id
WHERE 
    sm.user_id = $1
GROUP BY 
    date_trunc('month', s.created_at)
ORDER BY 
    date_trunc('month', s.created_at) DESC;