CREATE FUNCTION create_profile_from_role()
RETURNS TRIGGER AS $$
BEGIN
    IF (NEW.raw_user_meta_data->>'role' IS DISTINCT FROM OLD.raw_user_meta_data->>'role' OR NEW.raw_user_meta_data IS DISTINCT FROM OLD.raw_user_meta_data) THEN
        DELETE FROM public.student_profiles WHERE id = NEW.id;
        DELETE FROM public.staff_profiles WHERE id = NEW.id;
    END IF;
    IF NEW.raw_user_meta_data->>'role' = 'student' THEN
        INSERT INTO public.student_profiles (id, created_at, updated_at, first_name, middle_name, last_name, passport, date_of_birth, gender, joined_class_id, current_class_id, joining_date)
        VALUES (NEW.id, NEW.created_at, NEW.updated_at, NEW.raw_user_meta_data->>'first_name', NEW.raw_user_meta_data->>'middle_name', NEW.raw_user_meta_data->>'last_name', NEW.raw_user_meta_data->>'passport', TO_DATE(NEW.raw_user_meta_data->>'date_of_birth', 'YYYY-MM-DD HH24:MI:SS.US TZH:TZM'),
        (NEW.raw_user_meta_data->>'gender')::gender,
        (NEW.raw_user_meta_data->>'joined_class_id')::uuid, (NEW.raw_user_meta_data->>'current_class_id')::uuid, TO_DATE(NEW.raw_user_meta_data->>'joining_date', 'YYYY-MM-DD HH24:MI:SS.US TZH:TZM'));
    ELSIF NEW.raw_user_meta_data->>'role' = 'teacher' THEN
        INSERT INTO public.staff_profiles (id, created_at, updated_at, first_name, middle_name, last_name, passport, date_of_birth, gender, role, term_id, joining_date)
        VALUES (NEW.id, NEW.created_at, NEW.updated_at, NEW.raw_user_meta_data->>'first_name', NEW.raw_user_meta_data->>'middle_name', NEW.raw_user_meta_data->>'last_name', NEW.raw_user_meta_data->>'passport', TO_DATE(NEW.raw_user_meta_data->>'date_of_birth', 'YYYY-MM-DD HH24:MI:SS.US TZH:TZM'),
        (NEW.raw_user_meta_data->>'gender')::gender,
        ('teacher')::staf_roles, (NEW.raw_user_meta_data->>'term_id')::uuid, TO_DATE(NEW.raw_user_meta_data->>'joining_date', 'YYYY-MM-DD HH24:MI:SS.US TZH:TZM'));
    ELSIF NEW.raw_user_meta_data->>'role' = 'accountant' THEN
        INSERT INTO public.staff_profiles (id, created_at, updated_at, first_name, middle_name, last_name, passport, date_of_birth, gender, role, term_id, joining_date)
        VALUES (NEW.id, NEW.created_at, NEW.updated_at, NEW.raw_user_meta_data->>'first_name', NEW.raw_user_meta_data->>'middle_name', NEW.raw_user_meta_data->>'last_name', NEW.raw_user_meta_data->>'passport', TO_DATE(NEW.raw_user_meta_data->>'date_of_birth', 'YYYY-MM-DD HH24:MI:SS.US TZH:TZM'),
        (NEW.raw_user_meta_data->>'gender')::gender,
        ('accountant')::staff_roles, (NEW.raw_user_meta_data->>'term_id')::uuid, TO_DATE(NEW.raw_user_meta_data->>'joining_date', 'YYYY-MM-DD HH24:MI:SS.US TZH:TZM'));
    ELSIF NEW.raw_user_meta_data->>'role' = 'admin' THEN
        INSERT INTO public.staff_profiles (id, created_at, updated_at, first_name, middle_name, last_name, passport, date_of_birth, gender, role, term_id, joining_date)
        VALUES (NEW.id, NEW.created_at, NEW.updated_at, NEW.raw_user_meta_data->>'first_name', NEW.raw_user_meta_data->>'middle_name', NEW.raw_user_meta_data->>'last_name', NEW.raw_user_meta_data->>'passport', TO_DATE(NEW.raw_user_meta_data->>'date_of_birth', 'YYYY-MM-DD HH24:MI:SS.US TZH:TZM'),
        (NEW.raw_user_meta_data->>'gender')::gender,
        ('admin')::staff_roles, (NEW.raw_user_meta_data->>'term_id')::uuid, TO_DATE(NEW.raw_user_meta_data->>'joining_date', 'YYYY-MM-DD HH24:MI:SS.US TZH:TZM'));
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER create_profile_trigger
AFTER INSERT OR UPDATE ON auth.users
FOR EACH ROW
EXECUTE FUNCTION create_profile_from_role();



CREATE OR REPLACE FUNCTION set_admission_no()
RETURNS TRIGGER AS $$
DECLARE
    max_admission_no INTEGER;
BEGIN
    -- Get the maximum admission_no in the current table
    EXECUTE format('SELECT COALESCE(MAX(admission_no), 0) FROM %I', TG_TABLE_NAME) INTO max_admission_no;

    -- Set the NEW.admission_no to the maximum admission_no + 1
    NEW.admission_no := max_admission_no + 1;

    -- Return the NEW row
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;


CREATE TRIGGER set_admission_no_trigger
BEFORE INSERT ON your_table_name
FOR EACH ROW
EXECUTE FUNCTION set_admission_no();


IF NEW.raw_user_meta_data->>'role' IS DISTINCT FROM OLD.raw_user_meta_data->>'role' AND OLD.raw_user_meta_data->>'role' = 'student' THEN
        DELETE FROM public.student_profiles WHERE id = NEW.id;
        INSERT INTO public.staff_profiles (id, created_at, updated_at, first_name, middle_name, last_name, passport, date_of_birth, gender, role, term_id, joining_date)
        VALUES (NEW.id, NEW.created_at, NEW.updated_at, NEW.raw_user_meta_data->>'first_name', NEW.raw_user_meta_data->>'middle_name', NEW.raw_user_meta_data->>'last_name', NEW.raw_user_meta_data->>'passport', TO_DATE(NEW.raw_user_meta_data->>'date_of_birth', 'YYYY-MM-DD'),
        CASE
            WHEN NEW.raw_user_meta_data->>'gender' = 'male' THEN 'male'::public.gender
            WHEN NEW.raw_user_meta_data->>'gender' = 'female' THEN 'female'::public.gender
            ELSE NULL
        END,
        CASE
            WHEN NEW.raw_user_meta_data->>'role' = 'teacher' THEN 'teacher'::public.staff_roles
            WHEN NEW.raw_user_meta_data->>'role' = 'admin' THEN 'admin'::public.staff_roles
            WHEN NEW.raw_user_meta_data->>'role' = 'accountant' THEN 'accountant'::public.staff_roles
            ELSE NULL
        END,
        (NEW.raw_user_meta_data->>'term_id')::uuid,
        TO_DATE(NEW.raw_user_meta_data->>'joining_date', 'YYYY-MM-DD HH24:MI:SS.US'));
    END IF;

    BEGIN
    IF NEW.raw_user_meta_data->>'role' IS DISTINCT FROM OLD.raw_user_meta_data->>'role' AND OLD.raw_user_meta_data->>'role' = 'student' THEN
        DELETE FROM public.student_profiles WHERE id = NEW.id;
        INSERT INTO public.staff_profiles (id, created_at, updated_at, first_name, middle_name, last_name, passport, date_of_birth, gender, role, term_id, joining_date)
        VALUES (NEW.id, NEW.created_at, NEW.updated_at, NEW.raw_user_meta_data->>'first_name', NEW.raw_user_meta_data->>'middle_name', NEW.raw_user_meta_data->>'last_name', NEW.raw_user_meta_data->>'passport', TO_DATE(NEW.raw_user_meta_data->>'date_of_birth', 'YYYY-MM-DD'),
        CASE
            WHEN NEW.raw_user_meta_data->>'gender' = 'male' THEN 'male'::public.gender
            WHEN NEW.raw_user_meta_data->>'gender' = 'female' THEN 'female'::public.gender
            ELSE NULL
        END,
        CASE
            WHEN NEW.raw_user_meta_data->>'role' = 'teacher' THEN 'teacher'::public.staff_roles
            WHEN NEW.raw_user_meta_data->>'role' = 'admin' THEN 'admin'::public.staff_roles
            WHEN NEW.raw_user_meta_data->>'role' = 'accountant' THEN 'accountant'::public.staff_roles
            ELSE NULL
        END,
        (NEW.raw_user_meta_data->>'term_id')::uuid,
        TO_DATE(NEW.raw_user_meta_data->>'joining_date', 'YYYY-MM-DD HH24:MI:SS.US'));
    END IF;
    IF NEW.raw_user_meta_data->>'role' = 'student' THEN
        UPDATE public.student_profiles
        SET
            id = NEW.id,
            created_at = NEW.created_at,
            updated_at = NEW.updated_at,
            first_name = NEW.raw_user_meta_data->>'first_name',
            middle_name = NEW.raw_user_meta_data->>'middle_name',
            last_name = NEW.raw_user_meta_data->>'last_name',
            passport = NEW.raw_user_meta_data->>'passport',
            date_of_birth = TO_DATE(NEW.raw_user_meta_data->>'date_of_birth', 'YYYY-MM-DD'),
            gender = CASE
            WHEN NEW.raw_user_meta_data->>'gender' = 'male' THEN 'male'::public.gender
            WHEN NEW.raw_user_meta_data->>'gender' = 'female' THEN 'female'::public.gender
            ELSE NULL
            END,
            joined_class_id = (NEW.raw_user_meta_data->>'joined_class_id')::uuid,
            current_class_id = (NEW.raw_user_meta_data->>'current_class_id')::uuid,
            joining_date = TO_DATE(NEW.raw_user_meta_data->>'joining_date', 'YYYY-MM-DD') 
        WHERE id = NEW.id;
    ELSIF NEW.raw_user_meta_data->>'role' = 'teacher' THEN
        UPDATE public.staff_profiles
        SET
            id = NEW.id,
            created_at = NEW.created_at,
            updated_at = NEW.updated_at,
            first_name = NEW.raw_user_meta_data->>'first_name',
            middle_name = NEW.raw_user_meta_data->>'middle_name',
            last_name = NEW.raw_user_meta_data->>'last_name',
            passport = NEW.raw_user_meta_data->>'passport',
            date_of_birth = TO_DATE(NEW.raw_user_meta_data->>'date_of_birth', 'YYYY-MM-DD'),
            gender = CASE
            WHEN NEW.raw_user_meta_data->>'gender' = 'male' THEN 'male'::public.gender
            WHEN NEW.raw_user_meta_data->>'gender' = 'female' THEN 'female'::public.gender
            ELSE NULL
            END,
            role = 'teacher'::public.staff_roles,
            term_id = (NEW.raw_user_meta_data->>'term_id')::uuid,
            joining_date = TO_DATE(NEW.raw_user_meta_data->>'joining_date', 'YYYY-MM-DD') 
        WHERE id = NEW.id;
    ELSIF NEW.raw_user_meta_data->>'role' = 'accountant' THEN
        UPDATE public.staff_profiles
        SET
            id = NEW.id,
            created_at = NEW.created_at,
            updated_at = NEW.updated_at,
            first_name = NEW.raw_user_meta_data->>'first_name',
            middle_name = NEW.raw_user_meta_data->>'middle_name',
            last_name = NEW.raw_user_meta_data->>'last_name',
            passport = NEW.raw_user_meta_data->>'passport',
            date_of_birth = TO_DATE(NEW.raw_user_meta_data->>'date_of_birth', 'YYYY-MM-DD'),
            gender = CASE
            WHEN NEW.raw_user_meta_data->>'gender' = 'male' THEN 'male'::public.gender
            WHEN NEW.raw_user_meta_data->>'gender' = 'female' THEN 'female'::public.gender
            ELSE NULL
            END,
            role = 'accountant'::public.staff_roles,
            term_id = (NEW.raw_user_meta_data->>'term_id')::uuid,
            joining_date = TO_DATE(NEW.raw_user_meta_data->>'joining_date', 'YYYY-MM-DD') 
        WHERE id = NEW.id;
    ELSIF NEW.raw_user_meta_data->>'role' = 'admin' THEN
        UPDATE public.staff_profiles
        SET
            id = NEW.id,
            created_at = NEW.created_at,
            updated_at = NEW.updated_at,
            first_name = NEW.raw_user_meta_data->>'first_name',
            middle_name = NEW.raw_user_meta_data->>'middle_name',
            last_name = NEW.raw_user_meta_data->>'last_name',
            passport = NEW.raw_user_meta_data->>'passport',
            date_of_birth = TO_DATE(NEW.raw_user_meta_data->>'date_of_birth', 'YYYY-MM-DD'),
            gender = CASE
            WHEN NEW.raw_user_meta_data->>'gender' = 'male' THEN 'male'::public.gender
            WHEN NEW.raw_user_meta_data->>'gender' = 'female' THEN 'female'::public.gender
            ELSE NULL
            END,
            role = 'admin'::public.staff_roles,
            term_id = (NEW.raw_user_meta_data->>'term_id')::uuid,
            joining_date = TO_DATE(NEW.raw_user_meta_data->>'joining_date', 'YYYY-MM-DD') 
        WHERE id = NEW.id;
    END IF;
    RETURN NEW;
END;
