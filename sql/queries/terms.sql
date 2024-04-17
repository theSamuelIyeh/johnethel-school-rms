-- name: GetAllUnlockedTerms :many
SELECT *, sessions.name || ' - ' || term_id AS display_name
FROM terms 
LEFT JOIN sessions on terms.session_id = sessions.id
WHERE lock = false
ORDER BY terms.created_at DESC;

-- name: GetAllTerms :many
SELECT *, sessions.name || ' - ' || term_id AS display_name
FROM terms 
LEFT JOIN sessions on terms.session_id = sessions.id
ORDER BY terms.created_at DESC;