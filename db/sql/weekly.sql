-- Weekly Net Results (Last 30 Days)
SELECT 
    date_trunc('week', s.created_at)::date AS week_start,
    (date_trunc('week', s.created_at) + INTERVAL '6 days')::date AS week_end,
    SUM(s.withdraw_amount - s.amount) AS weekly_net_result
FROM 
    sessions s
JOIN 
    session_members sm ON s.id = sm.session_id
WHERE 
    sm.user_id = $1
    AND s.created_at >= CURRENT_DATE - INTERVAL '30 days'
GROUP BY 
    date_trunc('week', s.created_at),
    week_end
ORDER BY 
    week_start;