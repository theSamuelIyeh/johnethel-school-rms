-- +goose Up

CREATE TRIGGER set_admission_no_trigger
BEFORE INSERT ON public.staff_profiles
FOR EACH ROW
EXECUTE FUNCTION set_admission_no();

CREATE TRIGGER set_admission_no_trigger
BEFORE INSERT ON public.student_profiles
FOR EACH ROW
EXECUTE FUNCTION set_admission_no();

CREATE TRIGGER updated_at_trigger
AFTER UPDATE ON public.nursery_second_term_subject_scores
FOR EACH ROW
EXECUTE FUNCTION set_updated_at();

CREATE TRIGGER updated_at_trigger
AFTER UPDATE ON public.nursery_first_term_subject_scores
FOR EACH ROW
EXECUTE FUNCTION set_updated_at();

CREATE TRIGGER updated_at_trigger
AFTER UPDATE ON public.nursery_third_term_subject_scores
FOR EACH ROW
EXECUTE FUNCTION set_updated_at();

CREATE TRIGGER updated_at_trigger
AFTER UPDATE ON public.primary_first_term_subject_scores
FOR EACH ROW
EXECUTE FUNCTION set_updated_at();

CREATE TRIGGER updated_at_trigger
AFTER UPDATE ON public.primary_third_term_subject_scores
FOR EACH ROW
EXECUTE FUNCTION set_updated_at();

CREATE TRIGGER updated_at_trigger
AFTER UPDATE ON public.primary_second_term_subject_scores
FOR EACH ROW
EXECUTE FUNCTION set_updated_at();

CREATE TRIGGER updated_at_trigger
AFTER UPDATE ON public.nur_subjects_questions
FOR EACH ROW
EXECUTE FUNCTION set_updated_at();

CREATE TRIGGER updated_at_trigger
AFTER UPDATE ON public.primary_first_term_results
FOR EACH ROW
EXECUTE FUNCTION set_updated_at();

CREATE TRIGGER updated_at_trigger
AFTER UPDATE ON public.nursery_third_term_results
FOR EACH ROW
EXECUTE FUNCTION set_updated_at();

CREATE TRIGGER updated_at_trigger
AFTER UPDATE ON public.nursery_first_term_results
FOR EACH ROW
EXECUTE FUNCTION set_updated_at();

CREATE TRIGGER updated_at_trigger
AFTER UPDATE ON public.nursery_second_term_results
FOR EACH ROW
EXECUTE FUNCTION set_updated_at();

CREATE TRIGGER updated_at_trigger
AFTER UPDATE ON public.primary_third_term_results
FOR EACH ROW
EXECUTE FUNCTION set_updated_at();

CREATE TRIGGER updated_at_trigger
AFTER UPDATE ON public.primary_second_term_results
FOR EACH ROW
EXECUTE FUNCTION set_updated_at();

CREATE TRIGGER updated_at_trigger
AFTER UPDATE ON public.student_profiles
FOR EACH ROW
EXECUTE FUNCTION set_updated_at();

CREATE TRIGGER updated_at_trigger
AFTER UPDATE ON public.subjects
FOR EACH ROW
EXECUTE FUNCTION set_updated_at();

CREATE TRIGGER updated_at_trigger
AFTER UPDATE ON public.classes
FOR EACH ROW
EXECUTE FUNCTION set_updated_at();

CREATE TRIGGER updated_at_trigger
AFTER UPDATE ON public.staff_profiles
FOR EACH ROW
EXECUTE FUNCTION set_updated_at();

CREATE TRIGGER updated_at_trigger
AFTER UPDATE ON public.terms
FOR EACH ROW
EXECUTE FUNCTION set_updated_at();

CREATE TRIGGER updated_at_trigger
AFTER UPDATE ON public.class_list
FOR EACH ROW
EXECUTE FUNCTION set_updated_at();

CREATE TRIGGER updated_at_trigger
AFTER UPDATE ON public.subject_list
FOR EACH ROW
EXECUTE FUNCTION set_updated_at();

CREATE TRIGGER updated_at_trigger
AFTER UPDATE ON public.sessions
FOR EACH ROW
EXECUTE FUNCTION set_updated_at();

-- +goose Down

DROP TRIGGER set_admission_no_trigger ON public.staff_profiles;
DROP TRIGGER set_admission_no_trigger ON public.student_profiles;
DROP TRIGGER updated_at_trigger ON public.nursery_second_term_subject_scores;
DROP TRIGGER updated_at_trigger ON public.nursery_first_term_subject_scores;
DROP TRIGGER updated_at_trigger ON public.nursery_third_term_subject_scores;
DROP TRIGGER updated_at_trigger ON public.primary_first_term_subject_scores;
DROP TRIGGER updated_at_trigger ON public.primary_third_term_subject_scores;
DROP TRIGGER updated_at_trigger ON public.primary_second_term_subject_scores;
DROP TRIGGER updated_at_trigger ON public.nur_subjects_questions;
DROP TRIGGER updated_at_trigger ON public.primary_first_term_results;
DROP TRIGGER updated_at_trigger ON public.nursery_third_term_results;
DROP TRIGGER updated_at_trigger ON public.nursery_first_term_results;
DROP TRIGGER updated_at_trigger ON public.nursery_second_term_results;
DROP TRIGGER updated_at_trigger ON public.primary_third_term_results;
DROP TRIGGER updated_at_trigger ON public.primary_second_term_results;
DROP TRIGGER updated_at_trigger ON public.student_profiles;
DROP TRIGGER updated_at_trigger ON public.subjects;
DROP TRIGGER updated_at_trigger ON public.classes;
DROP TRIGGER updated_at_trigger ON public.staff_profiles;
DROP TRIGGER updated_at_trigger ON public.terms;
DROP TRIGGER updated_at_trigger ON public.class_list;
DROP TRIGGER updated_at_trigger ON public.subject_list;
DROP TRIGGER updated_at_trigger ON public.sessions;