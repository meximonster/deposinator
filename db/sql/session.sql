-- Session Data
SELECT 
    s.id,
    s.created_at AS date,
    s.amount AS deposit_amount,
    s.withdraw_amount,
    (s.withdraw_amount - s.amount) AS net_result,
    s.description
FROM 
    sessions s
JOIN 
    session_members sm ON s.id = sm.session_id
WHERE 
    sm.user_id = $1
ORDER BY 
    s.created_at DESC
LIMIT 50;