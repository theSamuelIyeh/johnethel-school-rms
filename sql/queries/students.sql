-- name: GetAllStudent :many
SELECT
    student_profiles.*,
    sessions.name || ' - ' || terms.term_id || ' - ' || class_list.name AS current_class
FROM student_profiles
LEFT JOIN classes ON student_profiles.current_class_id = classes.id
LEFT JOIN class_list ON classes.class_id = class_list.id
LEFT JOIN terms ON classes.term_id = terms.id
LEFT JOIN sessions ON terms.session_id = sessions.id
LIMIT $1 OFFSET $2
ORDER BY student_profiles.admission_no DESC;