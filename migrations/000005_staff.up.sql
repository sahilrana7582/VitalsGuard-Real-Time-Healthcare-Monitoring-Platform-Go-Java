CREATE TABLE IF NOT EXISTS staff (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    tenant_id UUID NOT NULL REFERENCES tenants(id) ON DELETE CASCADE,
    user_id UUID NOT NULL UNIQUE REFERENCES users(id) ON DELETE CASCADE,

    full_name TEXT NOT NULL,
    gender TEXT CHECK (gender IN ('Male', 'Female', 'Other')),
    dob DATE,
    contact_number TEXT,
    email TEXT,
    address TEXT,

    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);


CREATE TABLE IF NOT EXISTS doctors (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    tenant_id UUID NOT NULL REFERENCES tenants(id) ON DELETE CASCADE,
    staff_id UUID NOT NULL UNIQUE REFERENCES staff(id) ON DELETE CASCADE,

    specialization TEXT NOT NULL,
    license_number TEXT UNIQUE NOT NULL,
    years_of_experience INTEGER CHECK (years_of_experience >= 0),

    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS nurses (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    tenant_id UUID NOT NULL REFERENCES tenants(id) ON DELETE CASCADE,
    staff_id UUID NOT NULL UNIQUE REFERENCES staff(id) ON DELETE CASCADE,

    shift TEXT CHECK (shift IN ('Day', 'Night', 'Rotational')),
    floor_assigned TEXT,

    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);



CREATE OR REPLACE FUNCTION authorize_patient_write()
RETURNS TRIGGER AS $$
DECLARE
    user_role TEXT;
    user_tenant UUID;
BEGIN
    SELECT role, tenant_id INTO user_role, user_tenant
    FROM users
    WHERE id = current_setting('app.user_id')::UUID;

    IF NEW.tenant_id IS DISTINCT FROM user_tenant THEN
        RAISE EXCEPTION 'Unauthorized: tenant mismatch';
    END IF;

    IF user_role NOT IN ('doctor', 'nurse') THEN
        RAISE EXCEPTION 'Unauthorized: insufficient privileges';
    END IF;

    RETURN NEW;
END;
$$ LANGUAGE plpgsql;


CREATE TRIGGER patient_insert_trigger
BEFORE INSERT ON patients
FOR EACH ROW
EXECUTE FUNCTION authorize_patient_write();


CREATE TRIGGER patient_update_trigger
BEFORE UPDATE ON patients
FOR EACH ROW
EXECUTE FUNCTION authorize_patient_write();

CREATE TRIGGER patient_delete_trigger
BEFORE DELETE ON patients
FOR EACH ROW
EXECUTE FUNCTION authorize_patient_write();
